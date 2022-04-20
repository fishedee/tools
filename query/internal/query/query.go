package query

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/fishedee/tools/decimal"
	"github.com/fishedee/tools/kind"
)

// SelectReflect 反射实现
func SelectReflect[T, R any](data []T, selectFuctor func(a T) R) []R {
	dataValue := reflect.ValueOf(data)
	dataLen := dataValue.Len()

	selectFuctorValue := reflect.ValueOf(selectFuctor)
	selectFuctorType := selectFuctorValue.Type()
	selectFuctorOuterType := selectFuctorType.Out(0)
	resultType := reflect.SliceOf(selectFuctorOuterType)
	resultValue := reflect.MakeSlice(resultType, dataLen, dataLen)
	callArgument := []reflect.Value{{}}

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		callArgument[0] = singleDataValue
		singleResultValue := selectFuctorValue.Call(callArgument)[0]
		resultValue.Index(i).Set(singleResultValue)
	}
	return resultValue.Interface().([]R)
}

// WhereReflect 反射
func WhereReflect[T any](data []T, whereFuctor func(T) bool) []T {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataLen := dataValue.Len()

	whereFuctorValue := reflect.ValueOf(whereFuctor)
	resultType := reflect.SliceOf(dataType.Elem())
	resultValue := reflect.MakeSlice(resultType, 0, 0)

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		singleResultValue := whereFuctorValue.Call([]reflect.Value{singleDataValue})[0]
		if singleResultValue.Bool() {
			resultValue = reflect.Append(resultValue, singleDataValue)
		}
	}
	return resultValue.Interface().([]T)
}

// SortReflect 反射
func SortReflect[T any](data []T, sortType string) []T {
	//拷贝一份
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataElemType := dataType.Elem()
	dataValueLen := dataValue.Len()

	dataResult := reflect.MakeSlice(dataType, dataValueLen, dataValueLen)
	reflect.Copy(dataResult, dataValue)

	//排序
	targetCompares := getQueryExtractAndCompares(dataElemType, sortType)
	targetCompare := combineQueryCompare(targetCompares)
	result := dataResult.Interface()
	swapper := reflect.Swapper(result)

	SortInternal(dataValueLen, func(i int, j int) int {
		left := dataResult.Index(i)
		right := dataResult.Index(j)
		return targetCompare(left, right)
	}, swapper)

	return result.([]T)
}

func getQueryExtractAndCompares(dataType reflect.Type, sortTypeStr string) []queryCompare {
	sortName, sortType := analyseSort(sortTypeStr)
	targetCompare := []queryCompare{}
	for index, singleSortName := range sortName {
		singleSortType := sortType[index]
		singleCompare := getQueryExtractAndCompare(dataType, singleSortName)
		if !singleSortType {
			//逆序
			singleTempCompare := singleCompare
			singleCompare = func(left reflect.Value, right reflect.Value) int {
				return singleTempCompare(right, left)
			}
		}
		targetCompare = append(targetCompare, singleCompare)
	}
	return targetCompare
}

func getQueryExtractAndCompare(dataType reflect.Type, name string) queryCompare {
	fieldType, extract := GetQueryExtract(dataType, name)
	compare := getQueryCompare(fieldType)
	return func(left reflect.Value, right reflect.Value) int {
		return compare(extract(left), extract(right))
	}
}

func getQueryCompare(fieldType reflect.Type) queryCompare {
	typeKind := kind.GetTypeKind(fieldType)
	if typeKind == kind.TypeKind.BOOL {
		return func(left reflect.Value, right reflect.Value) int {
			leftBool := left.Bool()
			rightBool := right.Bool()
			if leftBool == rightBool {
				return 0
			} else if !leftBool {
				return -1
			} else {
				return 1
			}
		}
	} else if typeKind == kind.TypeKind.INT {
		return func(left reflect.Value, right reflect.Value) int {
			leftInt := left.Int()
			rightInt := right.Int()
			if leftInt < rightInt {
				return -1
			} else if leftInt > rightInt {
				return 1
			} else {
				return 0
			}
		}
	} else if typeKind == kind.TypeKind.UINT {
		return func(left reflect.Value, right reflect.Value) int {
			leftUint := left.Uint()
			rightUint := right.Uint()
			if leftUint < rightUint {
				return -1
			} else if leftUint > rightUint {
				return 1
			} else {
				return 0
			}
		}
	} else if typeKind == kind.TypeKind.FLOAT {
		return func(left reflect.Value, right reflect.Value) int {
			leftFloat := left.Float()
			rightFloat := right.Float()
			if leftFloat < rightFloat {
				return -1
			} else if leftFloat > rightFloat {
				return 1
			} else {
				return 0
			}
		}
	} else if typeKind == kind.TypeKind.STRING {
		if fieldType == reflect.TypeOf(decimal.Decimal("")) {
			return func(left reflect.Value, right reflect.Value) int {
				leftDecimal := left.Interface().(decimal.Decimal)
				rightDecimal := right.Interface().(decimal.Decimal)
				return leftDecimal.Cmp(rightDecimal)
			}
		}

		return func(left reflect.Value, right reflect.Value) int {
			leftString := left.String()
			rightString := right.String()
			if leftString < rightString {
				return -1
			} else if leftString > rightString {
				return 1
			} else {
				return 0
			}
		}

	} else if typeKind == kind.TypeKind.STRUCT && fieldType == reflect.TypeOf(time.Time{}) {
		return func(left reflect.Value, right reflect.Value) int {
			leftTime := left.Interface().(time.Time)
			rightTime := right.Interface().(time.Time)
			if leftTime.Before(rightTime) {
				return -1
			} else if leftTime.After(rightTime) {
				return 1
			} else {
				return 0
			}
		}
	} else {
		panic(fieldType.Name() + " can not compare")
	}
}

type QueryExtract func(reflect.Value) reflect.Value

func GetQueryExtract(dataType reflect.Type, name string) (reflect.Type, QueryExtract) {
	if name == "." {
		return dataType, func(v reflect.Value) reflect.Value {
			return v
		}
	}

	field, ok := kind.GetFieldByName(dataType, name)
	if !ok {
		panic(dataType.Name() + " has not name " + name)
	}
	fieldIndex := field.Index
	fieldType := field.Type
	return fieldType, func(v reflect.Value) reflect.Value {
		return v.FieldByIndex(fieldIndex)
	}
}

func analyseSort(sortType string) (result1 []string, result2 []bool) {
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

type queryCompare func(reflect.Value, reflect.Value) int

func combineQueryCompare(targetCompare []queryCompare) queryCompare {
	return func(left reflect.Value, right reflect.Value) int {
		for _, singleCompare := range targetCompare {
			compareResult := singleCompare(left, right)
			if compareResult < 0 {
				return -1
			} else if compareResult > 0 {
				return 1
			}
		}
		return 0
	}
}

// SortInternal 排序
func SortInternal(length int, lessHandler func(i, j int) int, swapHandler func(i, j int)) {
	sort.Stable(&sortInterface{
		lenHandler: func() int {
			return length
		},
		lessHandler: func(i int, j int) bool {
			return lessHandler(i, j) < 0
		},
		swapHandler: swapHandler,
	})

}

// sortInterface 基础类函数QuerySort
type sortInterface struct {
	lenHandler  func() int
	lessHandler func(i int, j int) bool
	swapHandler func(i int, j int)
}

// Len 长度
func (si *sortInterface) Len() int {
	return si.lenHandler()
}

// Less 比较
func (si *sortInterface) Less(i int, j int) bool {
	return si.lessHandler(i, j)
}

// Swap 交换
func (si *sortInterface) Swap(i int, j int) {
	si.swapHandler(i, j)
}

// JoinReflect 联表
func JoinReflect[L, R, LR any](leftData []L, rightData []R, joinPlace, joinType string, joinFuctor func(L, R) LR) []LR {
	//解析配置
	leftJoinType, rightJoinType := analyseJoin(joinType)

	leftDataValue := reflect.ValueOf(leftData)
	leftDataType := leftDataValue.Type()
	leftDataElemType := leftDataType.Elem()
	leftDataValueLen := leftDataValue.Len()
	leftDataJoinType, leftDataJoinExtract := GetQueryExtract(leftDataElemType, leftJoinType)

	rightDataValue := reflect.ValueOf(rightData)
	rightDataType := rightDataValue.Type()
	rightDataElemType := rightDataType.Elem()
	rightDataValueLen := rightDataValue.Len()
	_, rightDataJoinExtract := GetQueryExtract(rightDataElemType, rightJoinType)

	joinFuctorValue := reflect.ValueOf(joinFuctor)
	joinFuctorType := joinFuctorValue.Type()

	resultValue := reflect.MakeSlice(reflect.SliceOf(joinFuctorType.Out(0)), 0, 0)

	//执行join
	emptyLeftValue := reflect.New(leftDataElemType).Elem()
	emptyRightValue := reflect.New(rightDataElemType).Elem()
	joinPlace = strings.Trim(strings.ToLower(joinPlace), " ")

	nextData := make([]int, rightDataValueLen)
	mapDataNext := reflect.MakeMapWithSize(reflect.MapOf(leftDataJoinType, reflect.TypeOf(1)), rightDataValueLen)
	mapDataFirst := reflect.MakeMapWithSize(reflect.MapOf(leftDataJoinType, reflect.TypeOf(1)), rightDataValueLen)
	tempValueInt := reflect.New(reflect.TypeOf(1)).Elem()

	for i := 0; i != rightDataValueLen; i++ {
		tempValueInt.SetInt(int64(i))
		fieldValue := rightDataJoinExtract(rightDataValue.Index(i))
		lastNextIndex := mapDataNext.MapIndex(fieldValue)
		if lastNextIndex.IsValid() {
			nextData[int(lastNextIndex.Int())] = i
		} else {
			mapDataFirst.SetMapIndex(fieldValue, tempValueInt)
		}
		nextData[i] = -1
		mapDataNext.SetMapIndex(fieldValue, tempValueInt)
	}
	rightHaveJoin := make([]bool, rightDataValueLen)
	for i := 0; i != leftDataValueLen; i++ {
		leftValue := leftDataValue.Index(i)
		fieldValue := leftDataJoinExtract(leftDataValue.Index(i))
		rightIndex := mapDataFirst.MapIndex(fieldValue)
		if rightIndex.IsValid() {
			//找到右值
			j := int(rightIndex.Int())
			for nextData[j] != -1 {
				singleResult := joinFuctorValue.Call([]reflect.Value{leftValue, rightDataValue.Index(j)})[0]
				resultValue = reflect.Append(resultValue, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFuctorValue.Call([]reflect.Value{leftValue, rightDataValue.Index(j)})[0]
			resultValue = reflect.Append(resultValue, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFuctorValue.Call([]reflect.Value{leftValue, emptyRightValue})[0]
				resultValue = reflect.Append(resultValue, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != rightDataValueLen; j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFuctorValue.Call([]reflect.Value{emptyLeftValue, rightDataValue.Index(j)})[0]
			resultValue = reflect.Append(resultValue, singleResult)
		}
	}
	return resultValue.Interface().([]LR)
}

func analyseJoin(joinType string) (string, string) {
	joinTypeArray := strings.Split(joinType, "=")
	leftJoinType := strings.Trim(joinTypeArray[0], " ")
	rightJoinType := strings.Trim(joinTypeArray[1], " ")
	return leftJoinType, rightJoinType
}

// GroupReflect 反射
func GroupReflect[T, E any, R *E | []E](data []T, groupType string, groupFunctor func([]T) E) R {
	groupFuctorValue := reflect.ValueOf(groupFunctor)
	groupFuctorType := groupFuctorValue.Type()

	//解析输入数据
	dataValueLen := reflect.ValueOf(data).Len()

	//计算最终数据
	var resultValue reflect.Value
	resultType := groupFuctorType.Out(0)
	if resultType.Kind() == reflect.Slice {
		resultValue = reflect.MakeSlice(resultType, 0, dataValueLen)
	} else {
		resultValue = reflect.MakeSlice(reflect.SliceOf(resultType), 0, dataValueLen)
	}

	//执行分组操作
	GroupWalkReflect(data, groupType, func(data reflect.Value) {
		singleResult := groupFuctorValue.Call([]reflect.Value{data})[0]
		if singleResult.Kind() == reflect.Slice {
			resultValue = reflect.AppendSlice(resultValue, singleResult)
		} else {
			resultValue = reflect.Append(resultValue, singleResult)
		}
	})

	if resultType.Kind() == reflect.Slice {
		res := reflect.New(resultType).Elem()
		res.Set(resultValue)
		return res.Addr().Interface().(R)
	}

	return resultValue.Interface().(R)
}

type GroupWalkHandler func(data reflect.Value)

func GroupWalkReflect(data interface{}, groupType string, groupWalkHandler GroupWalkHandler) {
	//解析输入数据
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataElemType := dataType.Elem()
	dataValueLen := dataValue.Len()

	//分组操作
	groupType = strings.Trim(groupType, " ")
	dataFieldType, dataFieldExtract := GetQueryExtract(dataElemType, groupType)
	findMap := reflect.MakeMapWithSize(reflect.MapOf(dataFieldType, reflect.TypeOf(1)), dataValueLen)
	bufferData := reflect.MakeSlice(dataType, dataValueLen, dataValueLen)
	tempValueInt := reflect.New(reflect.TypeOf(1)).Elem()

	nextData := make([]int, dataValueLen)
	for i := 0; i != dataValueLen; i++ {
		fieldValue := dataFieldExtract(dataValue.Index(i))
		lastIndex := findMap.MapIndex(fieldValue)
		if lastIndex.IsValid() {
			nextData[int(lastIndex.Int())] = i
		}
		nextData[i] = -1
		tempValueInt.SetInt(int64(i))
		findMap.SetMapIndex(fieldValue, tempValueInt)
	}
	k := 0
	for i := 0; i != dataValueLen; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData.Index(k).Set(dataValue.Index(j))
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData.Index(k).Set(dataValue.Index(j))
		k++
		nextData[j] = 0
		groupWalkHandler(bufferData.Slice(kbegin, k))
	}
}

// ColumnReflect 反射
func ColumnReflect[T, R any](data []T, column string) []R {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	dataLen := dataValue.Len()
	column = strings.Trim(column, " ")
	dataFieldType, dataFieldExtract := GetQueryExtract(dataType, column)

	resultValue := reflect.MakeSlice(reflect.SliceOf(dataFieldType), dataLen, dataLen)

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		singleResultValue := dataFieldExtract(singleDataValue)
		resultValue.Index(i).Set(singleResultValue)
	}
	return resultValue.Interface().([]R)
}

// ColumnMapReflect 反射
func ColumnMapReflect[T any, K comparable](data []T, column string) map[K]T {
	column = strings.Trim(column, " ")
	// 返回值类型不一致，暂时去掉
	// if len(column) >= 2 && column[0:2] == "[]" {
	// 	column = column[2:]
	// 	return columnMapReflectSlice[T, K](data, column)
	// } else {
	return columnMapReflectSingle[T, K](data, column)
	// }
}

func columnMapReflectSlice[T any, K comparable](data []T, column string) map[K][]T {
	dataValue := reflect.ValueOf(data)
	dataValueType := dataValue.Type()
	dataType := dataValue.Type().Elem()
	dataLen := dataValue.Len()
	column = strings.Trim(column, " ")
	dataFieldType, dataFieldExtract := GetQueryExtract(dataType, column)

	resultValue := reflect.MakeMapWithSize(reflect.MapOf(dataFieldType, dataValueType), dataLen)

	GroupWalkReflect(data, column, func(group reflect.Value) {
		singleResultValue := dataFieldExtract(group.Index(0))
		resultValue.SetMapIndex(singleResultValue, group)
	})
	return resultValue.Interface().(map[K][]T)
}

func columnMapReflectSingle[T any, K comparable](data []T, column string) map[K]T {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	dataLen := dataValue.Len()
	column = strings.Trim(column, " ")
	dataFieldType, dataFieldExtract := GetQueryExtract(dataType, column)

	resultValue := reflect.MakeMapWithSize(reflect.MapOf(dataFieldType, dataType), dataLen)
	for i := dataLen - 1; i >= 0; i-- {
		singleDataValue := dataValue.Index(i)
		singleResultValue := dataFieldExtract(singleDataValue)
		resultValue.SetMapIndex(singleResultValue, singleDataValue)
	}
	return resultValue.Interface().(map[K]T)
}

// CombineReflect 反射
func CombineReflect[L, R, LR any](leftData []L, rightData []R, combineFuctor func(L, R) LR) []LR {
	leftValue := reflect.ValueOf(leftData)
	rightValue := reflect.ValueOf(rightData)
	if leftValue.Len() != rightValue.Len() {
		panic(fmt.Sprintf("len dos not equal %v != %v", leftValue.Len(), rightValue.Len()))
	}
	dataLen := leftValue.Len()
	combineFuctorValue := reflect.ValueOf(combineFuctor)
	resultType := combineFuctorValue.Type().Out(0)
	result := reflect.MakeSlice(reflect.SliceOf(resultType), dataLen, dataLen)
	for i := 0; i != dataLen; i++ {
		singleResultValue := combineFuctorValue.Call([]reflect.Value{leftValue.Index(i), rightValue.Index(i)})
		result.Index(i).Set(singleResultValue[0])
	}
	return result.Interface().([]LR)
}
