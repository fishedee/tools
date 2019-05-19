package main

import (
	"html/template"
	"strings"
)

// QueryGroupGen QueryGroupGen
func QueryGroupGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgSliceElem := firstArgSlice.Elem()

	//解析第二个参数
	secondArgGroupType := getContantStringValue(line, args[1].Value)
	groupType := strings.Trim(secondArgGroupType, " ")
	groupFieldExtract, groupFieldType := getExtendFieldType(line, firstArgSliceElem, groupType)

	//解析第三个参数
	thirdArgFunc := getFunctionType(line, args[2].Type)
	thirdArgFuncArgument := getArgumentType(line, thirdArgFunc)
	thirdArgFuncReturn := getReturnType(line, thirdArgFunc)
	if len(thirdArgFuncArgument) != 1 {
		Throw(1, "%v:should be one argument", line)
	}
	if len(thirdArgFuncReturn) != 1 {
		Throw(1, "%v:should be one return", line)
	}
	if thirdArgFuncArgument[0].String() != firstArgSlice.String() {
		Throw(1, "%v:groupFunctor argument should be equal with first argument %v!=%v", line, thirdArgFuncArgument[0], firstArgSliceElem)
	}

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, true, false})
	if hasQueryGroupGenerate[signature] == true {
		return nil
	}
	hasQueryGroupGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgSliceElem, importPackage)
	setImportPackage(line, thirdArgFuncReturn[0], importPackage)
	setImportPackage(line, groupFieldType, importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, true, false})
	funcBody := excuteTemplate(queryGroupFuncTmpl, map[string]string{
		"signature":          signature,
		"firstArgElemType":   getTypeDeclareCode(line, firstArgSliceElem),
		"thirdArgType":       getTypeDeclareCode(line, thirdArgFunc),
		"thirdArgReturnType": getTypeDeclareCode(line, thirdArgFuncReturn[0]),
		"columnType":         getTypeDeclareCode(line, groupFieldType),
		"columnExtract":      groupFieldExtract,
	})
	initBody := excuteTemplate(queryGroupInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      groupFuncPrefix + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

const (
	groupFuncPrefix = "queryGroup"
)

var (
	queryGroupFuncTmpl    *template.Template
	queryGroupInitTmpl    *template.Template
	hasQueryGroupGenerate map[string]bool
)

func init() {
	var err error
	queryGroupFuncTmpl, err = template.New("name").Parse(`
	func ` + groupFuncPrefix + `{{ .signature }}(data interface{},groupType string,groupFunctor interface{})interface{}{
		dataIn := data.([]{{ .firstArgElemType }})
		groupFunctorIn := groupFunctor.({{ .thirdArgType }})
		bufferData := make([]{{ .firstArgElemType }},len(dataIn),len(dataIn))
		mapData := make(map[{{ .columnType }}]int,len(dataIn))
		result := make([]{{ .thirdArgReturnType}},0,len(dataIn))

		length := len(dataIn)
		nextData := make([]int, length, length)
		for i := 0; i != length; i++ {
			single := dataIn[i]{{ .columnExtract }}
			lastIndex,isExist := mapData[single]
			if isExist == true {
				nextData[lastIndex] = i
			}
			nextData[i] = -1
			mapData[single] = i
		}
		k := 0
		for i := 0; i != length; i++ {
			j := i
			if nextData[j] == 0 {
				continue
			}
			kbegin := k
			for nextData[j] != -1 {
				nextJ := nextData[j]
				bufferData[k] = dataIn[j]
				nextData[j] = 0
				j = nextJ
				k++
			}
			bufferData[k] = dataIn[j]
			k++
			nextData[j] = 0
			single := groupFunctorIn(bufferData[kbegin:k])
			result = append(result,single)
		}

		return result
	}
	`)
	if err != nil {
		panic(err)
	}
	queryGroupInitTmpl, err = template.New("name").Parse(`
		query.GroupMacroRegister({{.argumentDefine}},` + groupFuncPrefix + `{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.Group", QueryGroupGen)
	hasQueryGroupGenerate = map[string]bool{}
}
