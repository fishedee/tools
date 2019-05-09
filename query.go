package query

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

// SelectMacroHandler 基础类函数QuerySelect
type SelectMacroHandler func(data interface{}, selectFunctor interface{}) interface{}

// SelectMacroRegister 注册器
func SelectMacroRegister(data interface{}, selectFunctor interface{}, handler SelectMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(selectFunctor).String()})
	selectMacroMapper[id] = handler
}

// SelectReflect 反射实现
func SelectReflect(data interface{}, selectFuctor interface{}) interface{} {
	dataValue := reflect.ValueOf(data)
	dataLen := dataValue.Len()

	selectFuctorValue := reflect.ValueOf(selectFuctor)
	selectFuctorType := selectFuctorValue.Type()
	selectFuctorOuterType := selectFuctorType.Out(0)
	resultType := reflect.SliceOf(selectFuctorOuterType)
	resultValue := reflect.MakeSlice(resultType, dataLen, dataLen)
	callArgument := []reflect.Value{reflect.Value{}}

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		callArgument[0] = singleDataValue
		singleResultValue := selectFuctorValue.Call(callArgument)[0]
		resultValue.Index(i).Set(singleResultValue)
	}
	return resultValue.Interface()
}

// Select 选择
func Select(data interface{}, selectFunctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(selectFunctor).String()})
	handler, isExist := selectMacroMapper[id]
	if isExist {
		return handler(data, selectFunctor)
	}

	reflectWarn("QuerySelect")
	return SelectReflect(data, selectFunctor)
}

// WhereMacroHandler 基础类函数QueryWhere
type WhereMacroHandler func(data interface{}, whereFunctor interface{}) interface{}

// WhereMacroRegister 注册器
func WhereMacroRegister(data interface{}, whereFunctor interface{}, handler WhereMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(whereFunctor).String()})
	whereMacroMapper[id] = handler
}

// WhereReflect 反射
func WhereReflect(data interface{}, whereFuctor interface{}) interface{} {
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
	return resultValue.Interface()
}

// Where 条件
func Where(data interface{}, whereFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(whereFuctor).String()})
	handler, isExist := whereMacroMapper[id]
	if isExist {
		return handler(data, whereFuctor)
	}

	reflectWarn("QueryWhere")
	return WhereReflect(data, whereFuctor)
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

// SortMacroHandler 处理器
type SortMacroHandler func(data interface{}, sortType string) interface{}

// SortMacroRegister 注册
func SortMacroRegister(data interface{}, sortType string, handler SortMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), sortType})
	sortMacroMapper[id] = handler
}

// SortReflect 反射
func SortReflect(data interface{}, sortType string) interface{} {
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

	return result
}

// Sort 排序
func Sort(data interface{}, sortType string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), sortType})
	handler, isExist := sortMacroMapper[id]
	if isExist {
		return handler(data, sortType)
	}

	reflectWarn("QuerySort")
	return SortReflect(data, sortType)
}

// JoinMacroHandler 基础类函数QueryJoin
type JoinMacroHandler func(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}) interface{}

// JoinMacroRegister 注册
func JoinMacroRegister(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}, handler JoinMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	joinMacroMapper[id] = handler
}

// JoinReflect 联表
func JoinReflect(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}) interface{} {
	//解析配置
	leftJoinType, rightJoinType := analyseJoin(joinType)

	leftDataValue := reflect.ValueOf(leftData)
	leftDataType := leftDataValue.Type()
	leftDataElemType := leftDataType.Elem()
	leftDataValueLen := leftDataValue.Len()
	leftDataJoinType, leftDataJoinExtract := getQueryExtract(leftDataElemType, leftJoinType)

	rightDataValue := reflect.ValueOf(rightData)
	rightDataType := rightDataValue.Type()
	rightDataElemType := rightDataType.Elem()
	rightDataValueLen := rightDataValue.Len()
	_, rightDataJoinExtract := getQueryExtract(rightDataElemType, rightJoinType)

	joinFuctorValue := reflect.ValueOf(joinFuctor)
	joinFuctorType := joinFuctorValue.Type()

	resultValue := reflect.MakeSlice(reflect.SliceOf(joinFuctorType.Out(0)), 0, 0)

	//执行join
	emptyLeftValue := reflect.New(leftDataElemType).Elem()
	emptyRightValue := reflect.New(rightDataElemType).Elem()
	joinPlace = strings.Trim(strings.ToLower(joinPlace), " ")

	nextData := make([]int, rightDataValueLen, rightDataValueLen)
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
	rightHaveJoin := make([]bool, rightDataValueLen, rightDataValueLen)
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
	return resultValue.Interface()
}

// Join 联表
func Join(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	handler, isExist := joinMacroMapper[id]
	if isExist {
		return handler(leftData, rightData, joinPlace, joinType, joinFuctor)
	}

	reflectWarn("QueryJoin")
	return JoinReflect(leftData, rightData, joinPlace, joinType, joinFuctor)
}

// GroupMacroHandler 基础类函数 QueryGroup
type GroupMacroHandler func(data interface{}, groupType string, groupFunctor interface{}) interface{}

// GroupMacroRegister 注册
func GroupMacroRegister(data interface{}, groupType string, groupFunctor interface{}, handler GroupMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), groupType, reflect.TypeOf(groupFunctor).String()})
	groupMacroMapper[id] = handler
}

// GroupReflect 反射
func GroupReflect(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	//解析输入数据
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataElemType := dataType.Elem()
	dataValueLen := dataValue.Len()
	groupFuctorValue := reflect.ValueOf(groupFunctor)
	groupFuctorType := groupFuctorValue.Type()

	//计算最终数据
	var resultValue reflect.Value
	resultType := groupFuctorType.Out(0)
	if resultType.Kind() == reflect.Slice {
		resultValue = reflect.MakeSlice(resultType, 0, dataValueLen)
	} else {
		resultValue = reflect.MakeSlice(reflect.SliceOf(resultType), 0, dataValueLen)
	}

	//分组操作
	groupType = strings.Trim(groupType, " ")
	dataFieldType, dataFieldExtract := getQueryExtract(dataElemType, groupType)
	findMap := reflect.MakeMapWithSize(reflect.MapOf(dataFieldType, reflect.TypeOf(1)), dataValueLen)
	bufferData := reflect.MakeSlice(dataType, dataValueLen, dataValueLen)
	tempValueInt := reflect.New(reflect.TypeOf(1)).Elem()

	nextData := make([]int, dataValueLen, dataValueLen)
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
		singleResult := groupFuctorValue.Call([]reflect.Value{bufferData.Slice(kbegin, k)})[0]
		if singleResult.Kind() == reflect.Slice {
			resultValue = reflect.AppendSlice(resultValue, singleResult)
		} else {
			resultValue = reflect.Append(resultValue, singleResult)
		}
	}
	return resultValue.Interface()
}

// Group 分组
func Group(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), groupType, reflect.TypeOf(groupFunctor).String()})
	handler, isExist := groupMacroMapper[id]
	if isExist {
		return handler(data, groupType, groupFunctor)
	}

	reflectWarn("QueryGroup")
	return GroupReflect(data, groupType, groupFunctor)
}

func analyseJoin(joinType string) (string, string) {
	joinTypeArray := strings.Split(joinType, "=")
	leftJoinType := strings.Trim(joinTypeArray[0], " ")
	rightJoinType := strings.Trim(joinTypeArray[1], " ")
	return leftJoinType, rightJoinType
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

func getQueryCompare(fieldType reflect.Type) queryCompare {
	typeKind := GetTypeKind(fieldType)
	if typeKind == TypeKind.BOOL {
		return func(left reflect.Value, right reflect.Value) int {
			leftBool := left.Bool()
			rightBool := right.Bool()
			if leftBool == rightBool {
				return 0
			} else if leftBool == false {
				return -1
			} else {
				return 1
			}
		}
	} else if typeKind == TypeKind.INT {
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
	} else if typeKind == TypeKind.UINT {
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
	} else if typeKind == TypeKind.FLOAT {
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
	} else if typeKind == TypeKind.STRING {
		if fieldType == reflect.TypeOf(Decimal("")) {
			return func(left reflect.Value, right reflect.Value) int {
				leftDecimal := left.Interface().(Decimal)
				rightDecimal := right.Interface().(Decimal)
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

	} else if typeKind == TypeKind.STRUCT && fieldType == reflect.TypeOf(time.Time{}) {
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

type queryExtract func(reflect.Value) reflect.Value

func getQueryExtract(dataType reflect.Type, name string) (reflect.Type, queryExtract) {
	if name == "." {
		return dataType, func(v reflect.Value) reflect.Value {
			return v
		}
	}

	field, ok := getFieldByName(dataType, name)
	if !ok {
		panic(dataType.Name() + " has not name " + name)
	}
	fieldIndex := field.Index
	fieldType := field.Type
	return fieldType, func(v reflect.Value) reflect.Value {
		return v.FieldByIndex(fieldIndex)
	}
}

func getQueryExtractAndCompare(dataType reflect.Type, name string) queryCompare {
	fieldType, extract := getQueryExtract(dataType, name)
	compare := getQueryCompare(fieldType)
	return func(left reflect.Value, right reflect.Value) int {
		return compare(extract(left), extract(right))
	}
}

// ColumnMacroHandler 扩展类函数 QueryColumn
type ColumnMacroHandler func(data interface{}, column string) interface{}

// ColumnMacroRegister 注册
func ColumnMacroRegister(data interface{}, column string, handler ColumnMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMacroMapper[id] = handler
}

// ColumnReflect 反射
func ColumnReflect(data interface{}, column string) interface{} {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	dataLen := dataValue.Len()
	column = strings.Trim(column, " ")
	dataFieldType, dataFieldExtract := getQueryExtract(dataType, column)

	resultValue := reflect.MakeSlice(reflect.SliceOf(dataFieldType), dataLen, dataLen)

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		singleResultValue := dataFieldExtract(singleDataValue)
		resultValue.Index(i).Set(singleResultValue)
	}
	return resultValue.Interface()
}

// Column 列
func Column(data interface{}, column string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMacroMapper[id]
	if isExist {
		return handler(data, column)
	}

	reflectWarn("QueryColumn")
	return ColumnReflect(data, column)
}

// ColumnMapMacroHandler 扩展类函数 QueryColumnMap
type ColumnMapMacroHandler func(data interface{}, column string) interface{}

// ColumnMapMacroRegister 注册
func ColumnMapMacroRegister(data interface{}, column string, handler ColumnMapMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMapMacroMapper[id] = handler
}

// ColumnMapReflect 反射
func ColumnMapReflect(data interface{}, column string) interface{} {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	dataLen := dataValue.Len()
	column = strings.Trim(column, " ")
	dataFieldType, dataFieldExtract := getQueryExtract(dataType, column)

	resultValue := reflect.MakeMap(reflect.MapOf(dataFieldType, dataType))
	for i := dataLen - 1; i >= 0; i-- {
		singleDataValue := dataValue.Index(i)
		singleResultValue := dataFieldExtract(singleDataValue)
		resultValue.SetMapIndex(singleResultValue, singleDataValue)
	}
	return resultValue.Interface()
}

// ColumnMap 映射
func ColumnMap(data interface{}, column string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMapMacroMapper[id]
	if isExist {
		return handler(data, column)
	}

	reflectWarn("QueryColumnMap")
	return ColumnMapReflect(data, column)
}

// LeftJoin 扩展类函数
func LeftJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "left", joinType, joinFuctor)
}

// RightJoin 右
func RightJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "right", joinType, joinFuctor)
}

// InnerJoin 内
func InnerJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "inner", joinType, joinFuctor)
}

// OuterJoin 外
func OuterJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "outer", joinType, joinFuctor)
}

// CombineMacroHandler 扩展累函数 QueryCombine
type CombineMacroHandler func(leftData interface{}, rightData interface{}, combineFuctor interface{}) interface{}

// CombineMacroRegister 注册
func CombineMacroRegister(leftData interface{}, rightData interface{}, combineFuctor interface{}, handler CombineMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), reflect.TypeOf(combineFuctor).String()})
	combineMacroMapper[id] = handler
}

// CombineReflect 反射
func CombineReflect(leftData interface{}, rightData interface{}, combineFuctor interface{}) interface{} {
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
	return result.Interface()
}

// Combine 联合
func Combine(leftData interface{}, rightData interface{}, combineFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), reflect.TypeOf(combineFuctor).String()})
	handler, isExist := combineMacroMapper[id]
	if isExist {
		return handler(leftData, rightData, combineFuctor)
	}

	reflectWarn("QueryCombine")
	return CombineReflect(leftData, rightData, combineFuctor)
}

// Reduce 减少
func Reduce(data interface{}, reduceFuctor interface{}, resultReduce interface{}) interface{} {
	dataValue := reflect.ValueOf(data)
	dataLen := dataValue.Len()

	reduceFuctorValue := reflect.ValueOf(reduceFuctor)
	resultReduceType := reduceFuctorValue.Type().In(0)
	resultReduceValue := reflect.New(resultReduceType)
	err := MapToArray(resultReduce, resultReduceValue.Interface(), "json")
	if err != nil {
		panic(err)
	}
	resultReduceValue = resultReduceValue.Elem()

	for i := 0; i != dataLen; i++ {
		singleDataValue := dataValue.Index(i)
		resultReduceValue = reduceFuctorValue.Call([]reflect.Value{resultReduceValue, singleDataValue})[0]
	}
	return resultReduceValue.Interface()
}

// Sum 求和
func Sum(data interface{}) interface{} {
	dataType := reflect.TypeOf(data).Elem()
	if dataType.Kind() == reflect.Int {
		return Reduce(data, func(sum int, single int) int {
			return sum + single
		}, 0)
	} else if dataType.Kind() == reflect.Float32 {
		return Reduce(data, func(sum float32, single float32) float32 {
			return sum + single
		}, (float32)(0.0))
	} else if dataType.Kind() == reflect.Float64 {
		return Reduce(data, func(sum float64, single float64) float64 {
			return sum + single
		}, 0.0)
	} else {
		panic("invalid type " + dataType.String())
	}
}

// Max 最大值
func Max(data interface{}) interface{} {
	dataType := reflect.TypeOf(data).Elem()
	if dataType.Kind() == reflect.Int {
		return Reduce(data, func(max int, single int) int {
			if single > max {
				return single
			}
			return max
		}, math.MinInt32)
	} else if dataType.Kind() == reflect.Float32 {
		return Reduce(data, func(max float32, single float32) float32 {
			if single > max {
				return single
			}
			return max
		}, math.SmallestNonzeroFloat32)
	} else if dataType.Kind() == reflect.Float64 {
		return Reduce(data, func(max float64, single float64) float64 {
			if single > max {
				return single
			}
			return max
		}, math.SmallestNonzeroFloat64)
	} else {
		panic("invalid type " + dataType.String())
	}
}

// Min 最小值
func Min(data interface{}) interface{} {
	dataType := reflect.TypeOf(data).Elem()
	if dataType.Kind() == reflect.Int {
		return Reduce(data, func(min int, single int) int {
			if single < min {
				return single
			}
			return min
		}, math.MaxInt32)
	} else if dataType.Kind() == reflect.Float32 {
		return Reduce(data, func(min float32, single float32) float32 {
			if single < min {
				return single
			}
			return min
		}, math.MaxFloat32)
	} else if dataType.Kind() == reflect.Float64 {
		return Reduce(data, func(min float64, single float64) float64 {
			if single < min {
				return single
			}
			return min
		}, math.MaxFloat64)
	} else {
		panic("invalid type " + dataType.String())
	}
}

// Reverse 反转
func Reverse(data interface{}) interface{} {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	dataLen := dataValue.Len()
	result := reflect.MakeSlice(dataType, dataLen, dataLen)

	for i := 0; i != dataLen; i++ {
		result.Index(dataLen - i - 1).Set(dataValue.Index(i))
	}
	return result.Interface()
}

// Distinct 唯一
func Distinct(data interface{}, columnNames string) interface{} {
	//提取信息
	name := Explode(columnNames, ",")
	extractInfo := []queryExtract{}
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	for _, singleName := range name {
		_, extract := getQueryExtract(dataType, singleName)
		extractInfo = append(extractInfo, extract)
	}

	//整合map
	existsMap := map[interface{}]bool{}
	result := reflect.MakeSlice(dataValue.Type(), 0, 0)
	dataLen := dataValue.Len()
	for i := 0; i != dataLen; i++ {
		singleValue := dataValue.Index(i)
		newData := reflect.New(dataType).Elem()
		for _, singleExtract := range extractInfo {
			singleField := singleExtract(singleValue)
			singleExtract(newData).Set(singleField)
		}
		newDataValue := newData.Interface()
		_, isExist := existsMap[newDataValue]
		if isExist {
			continue
		}
		result = reflect.Append(result, singleValue)
		existsMap[newDataValue] = true
	}
	return result.Interface()
}

func registerQueryTypeID(data []string) int64 {
	var result int64
	for _, m := range data {
		id, isExist := typeIDMapper[m]
		if isExist == false {
			id = int64(len(typeIDMapper)) + 1
			typeIDMapper[m] = id
		}
		result = result<<10 + id
	}
	return result
}

func getQueryTypeID(data []string) int64 {
	var result int64
	for _, m := range data {
		id, isExist := typeIDMapper[m]
		if isExist == false {
			return -1
		}
		result = result<<10 + id
	}
	return result
}

func reflectWarn(funcName string) {
	if reflectWarning {
		_, file, line, _ := runtime.Caller(2)
		log.Printf("%s:%d use %v reflect version,you should use querygen to avoid this warning", file, line, funcName)
	}

}

// ReflectWarning 警告
func ReflectWarning(isWarning bool) {
	reflectWarning = isWarning
}

var (
	selectMacroMapper    = map[int64]SelectMacroHandler{}
	whereMacroMapper     = map[int64]WhereMacroHandler{}
	sortMacroMapper      = map[int64]SortMacroHandler{}
	joinMacroMapper      = map[int64]JoinMacroHandler{}
	groupMacroMapper     = map[int64]GroupMacroHandler{}
	columnMacroMapper    = map[int64]ColumnMacroHandler{}
	columnMapMacroMapper = map[int64]ColumnMapMacroHandler{}
	combineMacroMapper   = map[int64]CombineMacroHandler{}
	typeIDMapper         = map[string]int64{}
	reflectWarning       = false
)
