package query

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

// ArrayKeyAndValue 分别返回键和值
func ArrayKeyAndValue(data interface{}) (interface{}, interface{}) {
	//解析data
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Map {
		panic("need a map for arrayKeyAndValue")
	}
	dataKeyType := dataType.Key()
	dataValueType := dataType.Elem()

	//合并数据
	dataKeySlice := reflect.MakeSlice(reflect.SliceOf(dataKeyType), 0, 0)
	dataValueSlice := reflect.MakeSlice(reflect.SliceOf(dataValueType), 0, 0)
	dataValue := reflect.ValueOf(data)
	for _, singleKey := range dataValue.MapKeys() {
		dataKeySlice = reflect.Append(dataKeySlice, singleKey)
		dataValueSlice = reflect.Append(dataValueSlice, dataValue.MapIndex(singleKey))
	}
	return dataKeySlice.Interface(), dataValueSlice.Interface()
}

// ArrayReverse 反转
func ArrayReverse(data interface{}) interface{} {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataLen := dataValue.Len()
	result := reflect.MakeSlice(dataType, dataLen, dataLen)

	for i := 0; i != dataLen; i++ {
		result.Index(dataLen - i - 1).Set(dataValue.Index(i))
	}
	return result.Interface()
}

// ArrayIn 存在
func ArrayIn(arrayData interface{}, findData interface{}) int {
	var findIndex int
	findIndex = -1
	arrayDataValue := reflect.ValueOf(arrayData)
	arrayDataValueLen := arrayDataValue.Len()
	for i := 0; i != arrayDataValueLen; i++ {
		singleArrayDataValue := arrayDataValue.Index(i).Interface()
		if singleArrayDataValue == findData {
			findIndex = i
			break
		}
	}
	return findIndex
}

// ArrayUnique 去重
func ArrayUnique(arrayData interface{}) interface{} {
	arrayValue := reflect.ValueOf(arrayData)
	arrayType := arrayValue.Type()
	arrayLen := arrayValue.Len()

	result := reflect.MakeSlice(arrayType, 0, 0)
	resultTemp := map[interface{}]bool{}

	for i := 0; i != arrayLen; i++ {
		singleArrayDataValue := arrayValue.Index(i)
		singleArrayDataValueInterface := singleArrayDataValue.Interface()
		_, isExist := resultTemp[singleArrayDataValueInterface]
		if isExist == true {
			continue
		}
		resultTemp[singleArrayDataValueInterface] = true
		result = reflect.Append(result, singleArrayDataValue)
	}
	return result.Interface()
}

func sliceToMap(arrayData []interface{}) map[interface{}]bool {
	result := map[interface{}]bool{}

	for _, singleArray := range arrayData {
		arrayValue := reflect.ValueOf(singleArray)
		arrayLen := arrayValue.Len()

		for i := 0; i != arrayLen; i++ {
			result[arrayValue.Index(i).Interface()] = true
		}
	}

	return result
}

// ArrayDiff 差集
func ArrayDiff(arrayData interface{}, arrayData2 interface{}, arrayOther ...interface{}) interface{} {
	arrayOther = append([]interface{}{arrayData2}, arrayOther...)
	arrayOtherMap := sliceToMap(arrayOther)

	arrayValue := reflect.ValueOf(arrayData)
	arrayType := arrayValue.Type()
	arrayLen := arrayValue.Len()
	result := reflect.MakeSlice(arrayType, 0, 0)

	for i := 0; i != arrayLen; i++ {
		singleArrayDataValue := arrayValue.Index(i)
		singleArrayDataValueInterface := singleArrayDataValue.Interface()

		_, isExist := arrayOtherMap[singleArrayDataValueInterface]
		if isExist == true {
			continue
		}
		result = reflect.Append(result, singleArrayDataValue)
		arrayOtherMap[singleArrayDataValueInterface] = true
	}

	return result.Interface()
}

// ArrayIntersect 交集
func ArrayIntersect(arrayData interface{}, arrayData2 interface{}, arrayOther ...interface{}) interface{} {
	arrayOther = append([]interface{}{arrayData2}, arrayOther...)
	arrayOtherMap := sliceToMap(arrayOther)

	arrayValue := reflect.ValueOf(arrayData)
	arrayType := arrayValue.Type()
	arrayLen := arrayValue.Len()
	result := reflect.MakeSlice(arrayType, 0, 0)

	for i := 0; i != arrayLen; i++ {
		singleArrayDataValue := arrayValue.Index(i)
		singleArrayDataValueInterface := singleArrayDataValue.Interface()

		isFirst, isExist := arrayOtherMap[singleArrayDataValueInterface]
		if isExist == false || isFirst == false {
			continue
		}
		result = reflect.Append(result, singleArrayDataValue)
		arrayOtherMap[singleArrayDataValueInterface] = false
	}

	return result.Interface()
}

// ArrayMerge 合并
func ArrayMerge(arrayData interface{}, arrayData2 interface{}, arrayOther ...interface{}) interface{} {
	arrayOther = append([]interface{}{arrayData2}, arrayOther...)
	arrayOtherMap := sliceToMap(arrayOther)

	arrayValue := reflect.ValueOf(arrayData)
	arrayType := arrayValue.Type()
	arrayLen := arrayValue.Len()
	result := reflect.MakeSlice(arrayType, 0, 0)

	for i := 0; i != arrayLen; i++ {
		singleArrayDataValue := arrayValue.Index(i)
		singleArrayDataValueInterface := singleArrayDataValue.Interface()

		isFirst, isExist := arrayOtherMap[singleArrayDataValueInterface]
		if isExist == true && isFirst == false {
			continue
		}
		result = reflect.Append(result, singleArrayDataValue)
		arrayOtherMap[singleArrayDataValueInterface] = false
	}

	for single, isFirst := range arrayOtherMap {
		if isFirst == false {
			continue
		}
		result = reflect.Append(result, reflect.ValueOf(single))
	}

	return result.Interface()
}

// ArraySort 排序
func ArraySort(data interface{}) interface{} {
	//建立一份拷贝数据
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataValueLen := dataValue.Len()

	dataResult := reflect.MakeSlice(dataType, dataValueLen, dataValueLen)
	reflect.Copy(dataResult, dataValue)

	//排序
	dataElemType := dataType.Elem()
	var result interface{}
	if dataElemType.Kind() == reflect.Int {
		intArray := dataResult.Interface().([]int)
		sort.Sort(sort.IntSlice(intArray))
		result = intArray
	} else if dataElemType.Kind() == reflect.String {
		stringArray := dataResult.Interface().([]string)
		sort.Sort(sort.StringSlice(stringArray))
		result = stringArray
	} else {
		panic("invalid sort type " + fmt.Sprintf("%v", dataElemType))
	}
	return result
}

// ArrayShuffle 洗牌
func ArrayShuffle(data interface{}) interface{} {
	//建立一份拷贝数据
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataValueLen := dataValue.Len()

	dataResult := reflect.MakeSlice(dataType, dataValueLen, dataValueLen)

	//打乱
	perm := rand.Perm(dataValueLen)
	for index, newIndex := range perm {
		dataResult.Index(newIndex).Set(dataValue.Index(index))
	}
	return dataResult.Interface()
}

// ArraySlice 数组切片
func ArraySlice(data interface{}, beginIndex int, endIndexArray ...int) interface{} {
	//建立一份拷贝数据
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataValueLen := dataValue.Len()

	//计算size
	endIndex := 0
	if len(endIndexArray) >= 1 {
		endIndex = endIndexArray[0]
	} else {
		endIndex = dataValueLen
	}
	size := 0
	if beginIndex >= endIndex {
		//逆向
		size = 0
	} else if endIndex <= 0 {
		//在左边
		size = 0
	} else if beginIndex >= dataValueLen {
		//在右边
		size = 0
	} else {
		//有交集
		if beginIndex <= 0 {
			beginIndex = 0
		}
		if endIndex >= dataValueLen {
			endIndex = dataValueLen
		}
		size = endIndex - beginIndex
	}

	//拷贝
	dataResult := reflect.MakeSlice(dataType, size, size)
	for i := 0; i != size; i++ {
		dataResult.Index(i).Set(dataValue.Index(i + beginIndex))
	}
	return dataResult.Interface()
}
