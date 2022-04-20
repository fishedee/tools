package main

import (
	"html/template"

	"github.com/fishedee/tools/exception"
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
		exception.Throw(1, "%v:should be two argument", line)
	}
	if len(thirdArgFuncReturn) != 1 {
		exception.Throw(1, "%v:should be one return", line)
	}
	if thirdArgFuncArgument[0].String() != firstArgElem.String() {
		exception.Throw(1, "%v:groupFunctor first argument should be equal with first argument %v!=%v", line, thirdArgFuncArgument[0], firstArgElem)
	}
	if thirdArgFuncArgument[1].String() != secondArgElem.String() {
		exception.Throw(1, "%v:groupFunctor second argument should be equal with second argument %v!=%v", line, thirdArgFuncArgument[1], secondArgElem)
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
		funcName:      combineFuncPrefix + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

const (
	combineFuncPrefix = "queryCombineV"
)

var (
	queryCombineFuncTmpl    *template.Template
	queryCombineInitTmpl    *template.Template
	hasQueryCombineGenerate map[string]bool
)

func init() {
	var err error
	queryCombineFuncTmpl, err = template.New("name").Parse(`
	func ` + combineFuncPrefix + `{{ .signature }}(leftData []{{ .firstArgElemType }},rightData []{{ .secondArgElemType }},combineFunctor {{ .thirdArgType }})[]{{ .thirdArgReturnType }}{
		leftDataIn := leftData
		rightDataIn := rightData
		combineFunctorIn := combineFunctor
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
		query.CombineMacroRegister({{.argumentDefine}},` + combineFuncPrefix + `{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.Combine", QueryCombineGen)
	hasQueryCombineGenerate = map[string]bool{}
}
