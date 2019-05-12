package main

import (
	"html/template"
)

// QueryCombineGen QueryCombineGen
func QueryCombineGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgElem := firstArgSlice.Elem()

	//解析第二个参数
	secondArgSlice := getSliceType(line, args[1].Type)
	secondArgElem := secondArgSlice.Elem()

	//解析第三个参数
	thirdArgFunc := getFunctionType(line, args[2].Type)
	thirdArgFuncArgument := getArgumentType(line, thirdArgFunc)
	thirdArgFuncReturn := getReturnType(line, thirdArgFunc)
	if len(thirdArgFuncArgument) != 2 {
		Throw(1, "%v:should be two argument", line)
	}
	if len(thirdArgFuncReturn) != 1 {
		Throw(1, "%v:should be one return", line)
	}
	if thirdArgFuncArgument[0].String() != firstArgElem.String() {
		Throw(1, "%v:groupFunctor first argument should be equal with first argument %v!=%v", line, thirdArgFuncArgument[0], firstArgElem)
	}
	if thirdArgFuncArgument[1].String() != secondArgElem.String() {
		Throw(1, "%v:groupFunctor second argument should be equal with second argument %v!=%v", line, thirdArgFuncArgument[1], secondArgElem)
	}

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, false, false})
	if hasQueryCombineGenerate[signature] == true {
		return nil
	}
	hasQueryCombineGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgElem, importPackage)
	setImportPackage(line, secondArgElem, importPackage)
	setImportPackage(line, thirdArgFuncReturn[0], importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, false, false})
	funcBody := excuteTemplate(queryCombineFuncTmpl, map[string]string{
		"signature":          signature,
		"firstArgElemType":   getTypeDeclareCode(line, firstArgElem),
		"secondArgElemType":  getTypeDeclareCode(line, secondArgElem),
		"thirdArgType":       getTypeDeclareCode(line, thirdArgFunc),
		"thirdArgReturnType": getTypeDeclareCode(line, thirdArgFuncReturn[0]),
	})
	initBody := excuteTemplate(queryCombineInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      "queryCombine_" + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

var (
	queryCombineFuncTmpl    *template.Template
	queryCombineInitTmpl    *template.Template
	hasQueryCombineGenerate map[string]bool
)

func init() {
	var err error
	queryCombineFuncTmpl, err = template.New("name").Parse(`
	func queryCombine_{{ .signature }}(leftData interface{},rightData interface{},combineFunctor interface{})interface{}{
		leftDataIn := leftData.([]{{ .firstArgElemType }})
		rightDataIn := rightData.([]{{ .secondArgElemType }})
		combineFunctorIn := combineFunctor.({{ .thirdArgType }})
		newData := make([]{{ .thirdArgReturnType }},len(leftDataIn),len(leftDataIn))

		for i := 0 ;i != len(leftDataIn);i++{
			newData[i] = combineFunctorIn(leftDataIn[i],rightDataIn[i])
		}
		return newData
	}
	`)
	if err != nil {
		panic(err)
	}
	queryCombineInitTmpl, err = template.New("name").Parse(`
		query.CombineMacroRegister({{.argumentDefine}},queryCombine_{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/donnol/tools/query.Combine", QueryCombineGen)
	hasQueryCombineGenerate = map[string]bool{}
}
