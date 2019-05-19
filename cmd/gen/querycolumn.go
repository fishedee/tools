package main

import (
	"html/template"
	"strings"
)

// QueryColumnGen QueryColumnGen
func QueryColumnGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第二个参数
	secondArgValue := getContantStringValue(line, args[1].Value)
	column := strings.Trim(secondArgValue, " ")

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgElem := firstArgSlice.Elem()
	columnArgExtract, columnArgType := getExtendFieldType(line, firstArgElem, column)

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, true})
	if hasQueryColumnGenerate[signature] == true {
		return nil
	}
	hasQueryColumnGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgElem, importPackage)
	setImportPackage(line, columnArgType, importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, true})
	funcBody := excuteTemplate(queryColumnFuncTmpl, map[string]string{
		"signature":                 signature,
		"firstArgElemType":          getTypeDeclareCode(line, firstArgElem),
		"firstArgElemColumnType":    getTypeDeclareCode(line, columnArgType),
		"firstArgElemColumnExtract": columnArgExtract,
	})
	initBody := excuteTemplate(queryColumnInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      columnFuncPrefix + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

const (
	columnFuncPrefix = "queryColumnV"
)

var (
	queryColumnFuncTmpl    *template.Template
	queryColumnInitTmpl    *template.Template
	hasQueryColumnGenerate map[string]bool
)

func init() {
	var err error
	queryColumnFuncTmpl, err = template.New("name").Parse(`
	func ` + columnFuncPrefix + `{{ .signature }}(data interface{},column string)interface{}{
		dataIn := data.([]{{ .firstArgElemType }})
		result := make([]{{ .firstArgElemColumnType }},len(dataIn),len(dataIn))

		for i,single := range dataIn{
			result[i] = single{{ .firstArgElemColumnExtract }}
		}
		return result
	}
	`)
	if err != nil {
		panic(err)
	}
	queryColumnInitTmpl, err = template.New("name").Parse(`
		query.ColumnMacroRegister({{.argumentDefine}},` + columnFuncPrefix + `{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.Column", QueryColumnGen)
	hasQueryColumnGenerate = map[string]bool{}
}
