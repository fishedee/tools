package main

import (
	"html/template"
	"strings"
)

// QueryColumnMapGen QueryColumnMapGen
func QueryColumnMapGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第二个参数
	secondArgValue := getContantStringValue(line, args[1].Value)
	column := strings.Trim(secondArgValue, " ")

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgElem := firstArgSlice.Elem()
	columnExtract, columnArgType := getExtendFieldType(line, firstArgElem, column)

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, true})
	if hasQueryColumnMapGenerate[signature] == true {
		return nil
	}
	hasQueryColumnMapGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgElem, importPackage)
	setImportPackage(line, columnArgType, importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, true})
	funcBody := excuteTemplate(queryColumnMapFuncTmpl, map[string]string{
		"signature":              signature,
		"firstArgElemType":       getTypeDeclareCode(line, firstArgElem),
		"firstArgElemColumnType": getTypeDeclareCode(line, columnArgType),
		"columnExtract":          columnExtract,
	})
	initBody := excuteTemplate(queryColumnMapInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      "queryColumnMap_" + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

var (
	queryColumnMapFuncTmpl    *template.Template
	queryColumnMapInitTmpl    *template.Template
	hasQueryColumnMapGenerate map[string]bool
)

func init() {
	var err error
	queryColumnMapFuncTmpl, err = template.New("name").Parse(`
	func queryColumnMap_{{ .signature }}(data interface{},column string)interface{}{
		dataIn := data.([]{{ .firstArgElemType }})
		result := make(map[{{ .firstArgElemColumnType }}]{{ .firstArgElemType }},len(dataIn))

		for _,single := range dataIn{
			result[single{{ .columnExtract }}] = single
		}
		return result
	}
	`)
	if err != nil {
		panic(err)
	}
	queryColumnMapInitTmpl, err = template.New("name").Parse(`
		query.ColumnMapMacroRegister({{.argumentDefine}},queryColumnMap_{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.ColumnMap", QueryColumnMapGen)
	hasQueryColumnMapGenerate = map[string]bool{}
}
