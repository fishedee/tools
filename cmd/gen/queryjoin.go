package main

import (
	"go/constant"
	"go/types"
	"html/template"
	"strings"
)

func analyseJoin(line string, joinType string) (string, string) {
	joinTypeArray := strings.Split(joinType, "=")
	if len(joinTypeArray) != 2 {
		Throw(1, "%v:join type should be two argument with equal operator", line)
	}
	leftJoinType := strings.Trim(joinTypeArray[0], " ")
	rightJoinType := strings.Trim(joinTypeArray[1], " ")
	return leftJoinType, rightJoinType
}

// QueryJoinGen QueryJoinGen
func QueryJoinGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgElem := firstArgSlice.Elem()

	//解析第二个参数
	secondArgSlice := getSliceType(line, args[1].Type)
	secondArgElem := secondArgSlice.Elem()

	//解析第三个参数
	thirdArgJoinPlace := getContantStringValue(line, args[2].Value)
	joinPlace := strings.Trim(strings.ToLower(thirdArgJoinPlace), " ")
	if joinPlace != "left" && joinPlace != "right" &&
		joinPlace != "inner" && joinPlace != "outer" {
		Throw(1, "%v:invalid join place %v", line, joinPlace)
	}

	//解析第四个参数
	forthArgJoinType := getContantStringValue(line, args[3].Value)
	leftJoinColumn, rightJoinColumn := analyseJoin(line, forthArgJoinType)
	leftFieldExtract, leftFieldType := getExtendFieldType(line, firstArgElem, leftJoinColumn)
	rightFieldExtract, rightFieldType := getExtendFieldType(line, secondArgElem, rightJoinColumn)
	if leftFieldType.String() != rightFieldType.String() {
		Throw(1, "%v:left join type should be equal to right join type %v!=%v", line, leftFieldType.String(), rightFieldType.String())
	}

	//解析第五个参数
	fifthArgFunc := getFunctionType(line, args[4].Type)
	fifthArgFuncArgument := getArgumentType(line, fifthArgFunc)
	fifthArgFuncReturn := getReturnType(line, fifthArgFunc)
	if len(fifthArgFuncArgument) != 2 {
		Throw(1, "%v:should be two argument", line)
	}
	if len(fifthArgFuncReturn) != 1 {
		Throw(1, "%v:should be one return", line)
	}
	if fifthArgFuncArgument[0].String() != firstArgElem.String() {
		Throw(1, "%v:joinFuctor first argument should be equal with first argument %v!=%v", line, fifthArgFuncArgument[0], firstArgElem)
	}
	if fifthArgFuncArgument[1].String() != secondArgElem.String() {
		Throw(1, "%v:joinFuctor second argument should be equal with second argument %v!=%v", line, fifthArgFuncArgument[1], secondArgElem)
	}

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, false, true, true, false})
	if hasQueryJoinGenerate[signature] == true {
		return nil
	}
	hasQueryJoinGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgElem, importPackage)
	setImportPackage(line, secondArgElem, importPackage)
	setImportPackage(line, fifthArgFuncReturn[0], importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, false, true, true, false})
	funcBody := excuteTemplate(queryJoinFuncTmpl, map[string]string{
		"signature":               signature,
		"firstArgElemType":        getTypeDeclareCode(line, firstArgElem),
		"secondArgElemType":       getTypeDeclareCode(line, secondArgElem),
		"fifthArgType":            getTypeDeclareCode(line, fifthArgFunc),
		"fifthArgReturnType":      getTypeDeclareCode(line, fifthArgFuncReturn[0]),
		"firstArgElemTypeDefine":  getTypeDefineCode(line, firstArgElem),
		"secondArgElemTypeDefine": getTypeDefineCode(line, secondArgElem),
		"joinPlace":               joinPlace,
		"joinFieldType":           getTypeDeclareCode(line, leftFieldType),
		"leftDataExtract":         leftFieldExtract,
		"rightDataExtract":        rightFieldExtract,
	})
	initBody := excuteTemplate(queryJoinInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      joinFuncPrefix + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

const (
	joinFuncPrefix = "queryJoin"
)

var (
	queryJoinFuncTmpl    *template.Template
	queryJoinInitTmpl    *template.Template
	hasQueryJoinGenerate map[string]bool
)

func init() {
	var err error
	queryJoinFuncTmpl, err = template.New("name").Parse(`
	func ` + joinFuncPrefix + `{{ .signature }}(leftData interface{},rightData interface{},joinPlace string,joinType string,joinFunctor interface{})interface{}{
		leftDataIn := leftData.([]{{ .firstArgElemType }})
		rightDataIn := rightData.([]{{ .secondArgElemType }})
		joinFunctorIn := joinFunctor.({{ .fifthArgType }})
		result := make([]{{ .fifthArgReturnType }},0,len(leftDataIn))

		emptyLeftData := {{ .firstArgElemTypeDefine }}
		emptyRightData := {{ .secondArgElemTypeDefine }}
		joinPlace = "{{ .joinPlace }}"

		nextData := make([]int,len(rightDataIn),len(rightDataIn))
		mapDataNext := make(map[{{ .joinFieldType }}]int,len(rightDataIn))
		mapDataFirst := make(map[{{ .joinFieldType }}]int,len(rightDataIn))

		for i := 0; i != len(rightDataIn); i++ {
			fieldValue := rightDataIn[i]{{ .rightDataExtract }}
			lastIndex,isExist := mapDataNext[fieldValue]
			if isExist {
				nextData[lastIndex] = i
			} else {
				mapDataFirst[fieldValue] = i
			}
			nextData[i] = -1
			mapDataNext[fieldValue] = i
		}

		rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
		for i := 0; i != len(leftDataIn); i++ {
			leftValue := leftDataIn[i]
			fieldValue := leftValue{{ .leftDataExtract }}
			rightIndex,isExist := mapDataFirst[fieldValue]
			if isExist {
				//找到右值
				j := rightIndex
				for nextData[j] != -1 {
					singleResult := joinFunctorIn(leftValue,rightDataIn[j])
					result = append(result,singleResult)
					rightHaveJoin[j] = true
					j = nextData[j]
				}
				singleResult := joinFunctorIn(leftValue,rightDataIn[j])
				result = append(result,singleResult)
				rightHaveJoin[j] = true
			} else {
				//找不到右值
				if joinPlace == "left" || joinPlace == "outer" {
					singleResult := joinFunctorIn(leftValue,emptyRightData)
					result = append(result,singleResult)
				}
			}
		}
		if joinPlace == "right" || joinPlace == "outer" {
			for j := 0; j != len(rightDataIn); j++ {
				if rightHaveJoin[j] {
					continue
				}
				singleResult := joinFunctorIn(emptyLeftData,rightDataIn[j])
				result = append(result,singleResult)
			}
		}
		return result
	}
	`)
	if err != nil {
		panic(err)
	}
	queryJoinInitTmpl, err = template.New("name").Parse(`
		query.JoinMacroRegister({{.argumentDefine}},` + joinFuncPrefix + `{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.Join", QueryJoinGen)
	registerQueryGen("github.com/fishedee/tools/query.LeftJoin", func(request QueryGenRequest) *QueryGenResponse {
		thridParty := types.TypeAndValue{
			Type:  nil,
			Value: constant.MakeString("left"),
		}
		newArgs := []types.TypeAndValue{}
		newArgs = append(newArgs, request.args[0:2]...)
		newArgs = append(newArgs, thridParty)
		newArgs = append(newArgs, request.args[2:]...)
		request.args = newArgs
		return QueryJoinGen(request)
	})
	registerQueryGen("github.com/fishedee/tools/query.RightJoin", func(request QueryGenRequest) *QueryGenResponse {
		thridParty := types.TypeAndValue{
			Type:  nil,
			Value: constant.MakeString("right"),
		}
		newArgs := []types.TypeAndValue{}
		newArgs = append(newArgs, request.args[0:2]...)
		newArgs = append(newArgs, thridParty)
		newArgs = append(newArgs, request.args[2:]...)
		request.args = newArgs
		return QueryJoinGen(request)
	})
	registerQueryGen("github.com/fishedee/tools/query.InnerJoin", func(request QueryGenRequest) *QueryGenResponse {
		thridParty := types.TypeAndValue{
			Type:  nil,
			Value: constant.MakeString("inner"),
		}
		newArgs := []types.TypeAndValue{}
		newArgs = append(newArgs, request.args[0:2]...)
		newArgs = append(newArgs, thridParty)
		newArgs = append(newArgs, request.args[2:]...)
		request.args = newArgs
		return QueryJoinGen(request)
	})
	registerQueryGen("github.com/fishedee/tools/query.OuterJoin", func(request QueryGenRequest) *QueryGenResponse {
		thridParty := types.TypeAndValue{
			Type:  nil,
			Value: constant.MakeString("outer"),
		}
		newArgs := []types.TypeAndValue{}
		newArgs = append(newArgs, request.args[0:2]...)
		newArgs = append(newArgs, thridParty)
		newArgs = append(newArgs, request.args[2:]...)
		request.args = newArgs
		return QueryJoinGen(request)
	})
	hasQueryJoinGenerate = map[string]bool{}
}
