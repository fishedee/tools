package main

import (
	"go/types"
	"html/template"
)

// QuerySortGen QuerySortGen
func QuerySortGen(request QueryGenRequest) *QueryGenResponse {
	args := request.args
	line := request.pkg.FileSet().Position(request.expr.Pos()).String()

	//解析第一个参数
	firstArgSlice := getSliceType(line, args[0].Type)
	firstArgSliceElem := firstArgSlice.Elem()

	//解析第二个参数
	secondArgSortType := getContantStringValue(line, args[1].Value)
	sortFieldNames, sortFieldIsAscs := analyseSortType(secondArgSortType)
	sortFieldTypes := make([]types.Type, len(sortFieldNames), len(sortFieldNames))
	sortFieldExtracts := make([]string, len(sortFieldNames), len(sortFieldNames))
	for i, sortFieldName := range sortFieldNames {
		sortFieldExtracts[i], sortFieldTypes[i] = getExtendFieldType(line, firstArgSliceElem, sortFieldName)
	}

	//生成函数
	signature := getFunctionSignature(line, args, []bool{false, true})
	if hasQuerySortGenerate[signature] == true {
		return nil
	}
	hasQuerySortGenerate[signature] = true
	importPackage := map[string]bool{}
	setImportPackage(line, firstArgSliceElem, importPackage)
	argumentDefine := getFunctionArgumentCode(line, args, []bool{false, true})
	funcBody := excuteTemplate(querySortFuncTmpl, map[string]string{
		"signature":        signature,
		"firstArgElemType": getTypeDeclareCode(line, firstArgSliceElem),
		"sortCode":         getCombineLessCompareCode(line, "newData[i]", "newData[j]", sortFieldExtracts, sortFieldIsAscs, sortFieldTypes),
	})
	initBody := excuteTemplate(querySortInitTmpl, map[string]string{
		"signature":      signature,
		"argumentDefine": argumentDefine,
	})
	return &QueryGenResponse{
		importPackage: importPackage,
		funcName:      sortFuncPrefix + signature,
		funcBody:      funcBody,
		initBody:      initBody,
	}
}

const (
	sortFuncPrefix = "querySort"
)

var (
	querySortFuncTmpl    *template.Template
	querySortInitTmpl    *template.Template
	hasQuerySortGenerate map[string]bool
)

func init() {
	var err error
	querySortFuncTmpl, err = template.New("name").Parse(`
	func ` + sortFuncPrefix + `{{ .signature }}(data interface{},sortType string)interface{}{
		dataIn := data.([]{{ .firstArgElemType }})
		newData := make([]{{ .firstArgElemType }},len(dataIn),len(dataIn))
		copy(newData,dataIn)

		query.SortInternal(len(newData),func(i int, j int)int{
			{{ .sortCode }}
			return 0
		},func(i int,j int){
			newData[j] , newData[i] = newData[i],newData[j]
		})
		return newData
	}
	`)
	if err != nil {
		panic(err)
	}
	querySortInitTmpl, err = template.New("name").Parse(`
		query.SortMacroRegister({{.argumentDefine}},` + sortFuncPrefix + `{{.signature}})
	`)
	if err != nil {
		panic(err)
	}
	registerQueryGen("github.com/fishedee/tools/query.Sort", QuerySortGen)
	hasQuerySortGenerate = map[string]bool{}
}
