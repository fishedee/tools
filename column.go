package query

import (
	"fmt"
	"reflect"
	"sort"
)

// ArrayColumnSort ArrayColumnSort
func ArrayColumnSort(data interface{}, columnNames string) interface{} {
	return Sort(data, columnNames)
}

// ArrayColumnUnique ArrayColumnUnique
func ArrayColumnUnique(data interface{}, columnNames string) interface{} {
	return Distinct(data, columnNames)
}

type arrayColumnMapInfo struct {
	Index   []int
	Type    reflect.Type
	MapType reflect.Type
}

// ArrayColumnKey ArrayColumnKey
func ArrayColumnKey(data interface{}, columnName string) interface{} {
	return Column(data, columnName)
}

// ArrayColumnMap ArrayColumnMap
func ArrayColumnMap(data interface{}, columnNames string) interface{} {
	//提取信息
	name := Explode(columnNames, ",")
	nameInfo := []arrayColumnMapInfo{}
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	for _, singleName := range name {
		singleField, ok := getFieldByName(dataType, singleName)
		if !ok {
			panic(dataType.Name() + " struct has not field " + singleName)
		}
		nameInfo = append(nameInfo, arrayColumnMapInfo{
			Index: singleField.Index,
			Type:  singleField.Type,
		})
	}
	prevType := dataType
	for i := len(nameInfo) - 1; i >= 0; i-- {
		nameInfo[i].MapType = reflect.MapOf(
			nameInfo[i].Type,
			prevType,
		)
		prevType = nameInfo[i].MapType
	}

	//整合map
	result := reflect.MakeMap(nameInfo[0].MapType)
	dataLen := dataValue.Len()
	for i := 0; i != dataLen; i++ {
		singleValue := dataValue.Index(i)
		prevValue := result
		for singleNameIndex, singleNameInfo := range nameInfo {
			var nextValue reflect.Value
			singleField := singleValue.FieldByIndex(singleNameInfo.Index)
			nextValue = prevValue.MapIndex(singleField)
			if !nextValue.IsValid() {
				if singleNameIndex+1 < len(nameInfo) {
					nextValue = reflect.MakeMap(nameInfo[singleNameIndex+1].MapType)
				} else {
					nextValue = singleValue
				}
				prevValue.SetMapIndex(singleField, nextValue)
			}
			prevValue = nextValue
		}
	}
	return result.Interface()
}

// ArrayColumnTable ArrayColumnTable
func ArrayColumnTable(column interface{}, data interface{}) [][]string {
	result := [][]string{}

	columnMap := column.(map[string]string)
	columnKeys, _ := ArrayKeyAndValue(column)
	columnKeysReal := columnKeys.([]string)
	columnValuesReal := []string{}
	sort.Sort(sort.StringSlice(columnKeysReal))
	for _, singleKey := range columnKeysReal {
		columnValuesReal = append(columnValuesReal, columnMap[singleKey])
	}
	result = append(result, columnValuesReal)

	dataValue := reflect.ValueOf(data)
	dataLen := dataValue.Len()
	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		singleDataStringValue := ArrayToMap(singleDataValue.Interface(), "json")
		singleDataStringValueData := reflect.ValueOf(singleDataStringValue)
		singleResult := []string{}
		for _, singleColumn := range columnKeysReal {
			singleResultString := ""
			singleValue := singleDataStringValueData.MapIndex(reflect.ValueOf(singleColumn))
			if singleValue.IsValid() == false {
				singleResultString = ""
			} else {
				singleResultString = fmt.Sprintf("%v", singleValue)
			}
			singleResult = append(
				singleResult,
				singleResultString,
			)
		}
		result = append(result, singleResult)
	}
	return result
}
