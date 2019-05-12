package main

import (
	"html/template"
)

// QuerySelectGen QuerySelectGen
func QuerySelectGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgElem := firstArgSlice.Elem()

	//解析第二个参数
	secondArgFunc := getFunctionType(line, args[1].Type)
	secondArgFuncArguments := getArgumentType(line, secondArgFunc)
	secondArgFuncResults := getReturnType(line, secondArgFunc)
	if len(secondArgFuncArguments) != 1 {
		Throw(1, "%v:selector should be single argument")
	}
	if len(secondArgFuncResults) != 1 {
		Throw(1, "%v:selector should be single return")
	}
	secondArgFuncArgument := secondArgFuncArguments[0]
	secondArgFuncResult := secondArgFuncResults[0]
	if firstArgElem.String() != secondArgFuncArgument.String() {
		Throw(1, "%v:second selector argument should be equal with first argument")
	}

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, false})
	if hasQuerySelectGenerate[signature] == true {
		return nil
	}
	hasQuerySelectGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, secondArgFuncArgument, importPackage)
	setImportPackage(line, secondArgFuncResult, importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, false})
	funcBody := excuteTemplate(querySelectFuncTmpl, map[string]string{
		"signature":           signature,
		"firstArgElemType":    getTypeDeclareCode(line, firstArgElem),
		"secondArgType":       getTypeDeclareCode(line, secondArgFunc),
		"secondArgReturnType": getTypeDeclareCode(line, secondArgFuncResult),
	})
	initBody := excuteTemplate(querySelectInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      "querySelect_" + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

var (
	querySelectFuncTmpl    *template.Template
	querySelectInitTmpl    *template.Template
	hasQuerySelectGenerate map[string]bool
)

func init() {
	var err error
	querySelectFuncTmpl, err = template.New("name").Parse(`
	func querySelect_{{ .signature }}(data interface{},selectFunctor interface{})interface{}{
		dataIn := data.([]{{ .firstArgElemType }})
		selectFunctorIn := selectFunctor.({{ .secondArgType }})
		result := make([]{{ .secondArgReturnType }},len(dataIn),len(dataIn))

		for i,single := range dataIn{
			result[i] = selectFunctorIn(single)
		}
		return result
	}
	`)
	if err != nil {
		panic(err)
	}
	querySelectInitTmpl, err = template.New("name").Parse(`
		query.SelectMacroRegister({{.argumentDefine}},querySelect_{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/donnol/tools/query.Select", QuerySelectGen)
	hasQuerySelectGenerate = map[string]bool{}
}
