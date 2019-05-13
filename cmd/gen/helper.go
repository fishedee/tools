package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"go/constant"
	"go/types"
	"html/template"
	"log"
	"strings"

	"github.com/fishedee/tools/plode"
)

// Throw 临时使用
func Throw(code int, format string, args ...interface{}) {
	log.Printf("code: %d, "+format, code, args)
}

func getFunctionSignature(line string, arguments []types.TypeAndValue, isConstant []bool) string {
	var buffer bytes.Buffer
	for i, argument := range arguments {
		single := ""
		if isConstant[i] == true {
			single = getContantStringValue(line, argument.Value)
		} else {
			single = argument.Type.String()
		}
		buffer.WriteString("_" + single)
	}
	hash := sha1.New()
	hash.Write(buffer.Bytes())
	etag := hash.Sum(nil)
	etagString := hex.EncodeToString(etag)
	return etagString
}

func getContantStringValue(line string, value constant.Value) string {
	if value == nil {
		Throw(1, "%v:should be constant!%v", line, value)
	}
	return constant.StringVal(value)
}

func getNamedType(line string, t types.Type) *types.Named {
	t1, isNamed := t.(*types.Named)
	if isNamed == false {
		Throw(1, "%v:should be named type!%v", line, t)
	}
	return t1
}

func getFunctionType(line string, t types.Type) *types.Signature {
	t1, isFunc := t.(*types.Signature)
	if isFunc == false {
		Throw(1, "%v:should be function type!%v", line, t)
	}
	if t1.Recv() != nil {
		Throw(1, "%v:should be pure function", line)
	}
	if t1.Variadic() == true {
		Throw(1, "%v:should not variadic function")
	}
	return t1
}

func getArgumentType(line string, t *types.Signature) []types.Type {
	arguments := t.Params()
	length := arguments.Len()
	r := make([]types.Type, length, length)
	for i := 0; i != length; i++ {
		r[i] = arguments.At(i).Type()
	}
	return r
}

func getReturnType(line string, t *types.Signature) []types.Type {
	arguments := t.Results()
	length := arguments.Len()
	r := make([]types.Type, length, length)
	for i := 0; i != length; i++ {
		r[i] = arguments.At(i).Type()
	}
	return r
}

func getSliceType(line string, t types.Type) *types.Slice {
	t1, isSlice := t.(*types.Slice)
	if isSlice == false {
		Throw(1, "%v:should be slice type!%v", line, t)
	}
	return t1
}

func getStructType(line string, t types.Type) *types.Struct {
	t1, isStruct := t.(*types.Struct)
	if isStruct == false {
		Throw(1, "%v:should be struct type!%v", line, t)
	}
	return t1
}

func getFieldType(line string, tStruct *types.Struct, column string) types.Type {
	for i := 0; i != tStruct.NumFields(); i++ {
		field := tStruct.Field(i)
		if field.Name() == column {
			return field.Type()
		}
	}
	Throw(1, "%v:%v has not found column %v", line, tStruct, column)
	return nil
}

func getExtendFieldType(line string, t types.Type, column string) (string, types.Type) {
	column = strings.Trim(column, " ")
	if column == "." {
		return "", t
	}

	tNamed, isNamed := t.(*types.Named)
	if isNamed == false {
		Throw(1, "%v:should be named type!because column is not comma :%v,%v", line, t, column)
	}
	tStruct, isStruct := tNamed.Underlying().(*types.Struct)
	if isStruct == false {
		Throw(1, "%v:should be struct type!because column is not comma :%v,%v", line, t, column)
	}
	for i := 0; i != tStruct.NumFields(); i++ {
		field := tStruct.Field(i)
		if field.Name() == column {
			return "." + column, field.Type()
		}
	}
	Throw(1, "%v:%v has not found column %v", line, tStruct, column)
	return "", nil
}

func getTypeDeclareCode(line string, t types.Type) string {
	if tBasic, ok := t.(*types.Basic); ok {
		switch tBasic.Kind() {
		case types.Bool:
			return "bool"
		case types.Int:
			return "int"
		case types.String:
			return "string"
		default:
			Throw(1, "%v:unknown basic type %v", line, t.String())
			return ""
		}
	} else if tSlice, ok := t.(*types.Slice); ok {
		elemType := tSlice.Elem()
		return "[]" + getTypeDeclareCode(line, elemType)
	} else if tMap, ok := t.(*types.Map); ok {
		keyType := getTypeDeclareCode(line, tMap.Key())
		elemType := getTypeDeclareCode(line, tMap.Elem())
		return "map[" + keyType + "]" + elemType
	} else if tSignature, ok := t.(*types.Signature); ok {
		argumentTypes := getArgumentType(line, tSignature)
		returnTypes := getReturnType(line, tSignature)
		argumentTypeCode := []string{}
		for _, argumentType := range argumentTypes {
			argumentTypeCode = append(argumentTypeCode, getTypeDeclareCode(line, argumentType))
		}
		argumentTypeCodeString := "(" + plode.Implode(argumentTypeCode, ",") + ")"
		returnTypeCode := []string{}
		for _, returnType := range returnTypes {
			returnTypeCode = append(returnTypeCode, getTypeDeclareCode(line, returnType))
		}
		returnTypeCodeString := ""
		if len(returnTypeCode) == 0 {
			returnTypeCodeString = ""
		} else if len(returnTypeCode) == 1 {
			returnTypeCodeString = returnTypeCode[0]
		} else {
			returnTypeCodeString = "(" + plode.Implode(returnTypeCode, ",") + ")"
		}
		return "func" + argumentTypeCodeString + returnTypeCodeString
	} else if tNamed, ok := t.(*types.Named); ok {
		obj := tNamed.Obj()
		if obj.Pkg().Path() == globalGeneratePackagePath {
			return obj.Name()
		}
		return obj.Pkg().Name() + "." + obj.Name()
	} else {
		Throw(1, "%v:unknown type to declare: %v", line, t.String())
		return ""
	}

}

func getTypeDefineCode(line string, t types.Type) string {
	if tBasic, ok := t.(*types.Basic); ok {
		switch tBasic.Kind() {
		case types.Bool:
			return "false"
		case types.Int:
			return "0"
		case types.String:
			return "\"\""
		default:
			Throw(1, "%v:unknown basic type %v", line, t.String())
			return ""
		}
	} else if tSlice, ok := t.(*types.Slice); ok {
		return getTypeDeclareCode(line, tSlice) + "{}"
	} else if tMap, ok := t.(*types.Map); ok {
		return getTypeDeclareCode(line, tMap) + "{}"
	} else if tSignature, ok := t.(*types.Signature); ok {
		return "(" + getTypeDeclareCode(line, tSignature) + ")(nil)"
	} else if tNamed, ok := t.(*types.Named); ok {
		obj := tNamed.Obj()
		underType := tNamed.Underlying()
		if _, isStruct := underType.(*types.Struct); isStruct {
			declareName := ""
			if obj.Pkg().Path() == globalGeneratePackagePath {
				declareName = obj.Name()
			} else {
				declareName = obj.Pkg().Name() + "." + obj.Name()
			}
			return declareName + "{}"
		}

		underTypeDefine := getTypeDefineCode(line, underType)
		if obj.Pkg().Path() == globalGeneratePackagePath {
			return obj.Name() + "(" + underTypeDefine + ")"
		}

		return obj.Pkg().Name() + "." + obj.Name() + "(" + underTypeDefine + ")"
	} else {
		Throw(1, "%v:unknown type to define %v", line, t.String())
		return ""
	}
}

func setImportPackage(line string, t types.Type, importPkg map[string]bool) {
	if _, ok := t.(*types.Basic); ok {
		return
	} else if tSlice, ok := t.(*types.Slice); ok {
		elemType := tSlice.Elem()
		setImportPackage(line, elemType, importPkg)
	} else if tMap, ok := t.(*types.Map); ok {
		keyType := tMap.Key()
		elemType := tMap.Elem()
		setImportPackage(line, keyType, importPkg)
		setImportPackage(line, elemType, importPkg)
	} else if tNamed, ok := t.(*types.Named); ok {
		obj := tNamed.Obj()
		pkg := obj.Pkg()
		importPkg[pkg.Path()] = true
	} else {
		Throw(1, "%v:unknown type to define %v", line, t.String())
	}
}

func getFunctionArgumentCode(line string, arguments []types.TypeAndValue, isConstant []bool) string {
	argvs := []string{}
	for i, argument := range arguments {
		if isConstant[i] == true {
			argvs = append(argvs, "\""+getContantStringValue(line, argument.Value)+"\"")
		} else {
			argvs = append(argvs, getTypeDefineCode(line, argument.Type))
		}
	}
	return plode.Implode(argvs, ",")
}

func excuteTemplate(tmpl *template.Template, data map[string]string) string {
	newData := make(map[string]template.HTML, len(data))
	for key, value := range data {
		newData[key] = template.HTML(value)
	}
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, newData)
	if err != nil {
		Throw(1, "execute fail %v", err)
	}
	return buffer.String()
}

func analyseSortType(sortType string) (result1 []string, result2 []bool) {
	sortTypeArray := strings.Split(sortType, ",")
	for _, singleSortTypeArray := range sortTypeArray {
		singleSortTypeArrayTemp := strings.Split(singleSortTypeArray, " ")
		singleSortTypeArray := []string{}
		for _, singleSort := range singleSortTypeArrayTemp {
			singleSort = strings.Trim(singleSort, " ")
			if singleSort == "" {
				continue
			}
			singleSortTypeArray = append(singleSortTypeArray, singleSort)
		}
		var singleSortName string
		var singleSortType bool
		if len(singleSortTypeArray) >= 2 {
			singleSortName = singleSortTypeArray[0]
			singleSortType = (strings.ToLower(strings.Trim(singleSortTypeArray[1], " ")) == "asc")
		} else {
			singleSortName = singleSortTypeArray[0]
			singleSortType = true
		}
		result1 = append(result1, singleSortName)
		result2 = append(result2, singleSortType)
	}
	return result1, result2
}

func getLessCompareCode(line string, name1 string, extractFieldName1 string, name2 string, extractFieldName2 string, sortFieldIsAsc bool, sortFieldType types.Type) string {
	lessTrueCode := ""
	lessFalseCode := ""
	if sortFieldIsAsc {
		lessTrueCode = "-1"
		lessFalseCode = "1"
	} else {
		lessTrueCode = "1"
		lessFalseCode = "-1"
	}
	_, isBasic := sortFieldType.(*types.Basic)
	if isBasic {
		return "if " + name1 + extractFieldName1 + "<" + name2 + extractFieldName2 + "{\n" +
			"return " + lessTrueCode + "\n" +
			"} else if " + name1 + extractFieldName1 + ">" + name2 + extractFieldName2 + "{\n" +
			"return " + lessFalseCode + "\n" +
			"}\n"
	}

	tNamed, isNamed := sortFieldType.(*types.Named)
	if isNamed && tNamed.String() == "time.Time" {
		return "if " + name1 + extractFieldName1 + ".Before(" + name2 + extractFieldName2 + "){\n" +
			"return " + lessTrueCode + "\n" +
			"} else if " + name1 + extractFieldName1 + ".After(" + name2 + extractFieldName2 + "){\n" +
			"return " + lessFalseCode + "\n" +
			"}\n"
	}

	Throw(1, "line:unknown how to sort type : %v", line, sortFieldType.String)
	return ""
}

func getCombineLessCompareCode(line string, name1 string, name2 string, sortFieldExtracts []string, sortFieldIsAscs []bool, sortFieldTypes []types.Type) string {
	code := []string{}
	for i := range sortFieldExtracts {
		singleCode := getLessCompareCode(line, name1, sortFieldExtracts[i], name2, sortFieldExtracts[i], sortFieldIsAscs[i], sortFieldTypes[i])
		code = append(code, singleCode)
	}
	return plode.Implode(code, "\n")
}

var (
	globalGeneratePackagePath = ""
)
