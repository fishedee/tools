package testdata

import (
	"github.com/fishedee/tools/query"
)

func queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV1a5b7250371597524e364f0c816390c77a8b3331(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumnV3923b792e276005e09637544ecb3aec8be870f41(data interface{}, column string) interface{} {
	dataIn := data.([]string)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func init() {

	query.ColumnMacroRegister([]ContentType{}, "     Name         ", queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f)

	query.ColumnMacroRegister([]ContentType{}, " Name ", queryColumnV1a5b7250371597524e364f0c816390c77a8b3331)

	query.ColumnMacroRegister([]ContentType{}, "Ok        ", queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c)

	query.ColumnMacroRegister([]string{}, " . ", queryColumnV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMacroRegister([]ContentType{}, "    Money  ", queryColumnV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7)

	query.ColumnMacroRegister([]ContentType{}, "    CardMoney", queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd)

	query.ColumnMacroRegister([]int{}, " . ", queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMacroRegister([]ContentType{}, "Age        ", queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3)

}
