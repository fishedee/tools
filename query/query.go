package query

import (
	"log"
	"math"
	"reflect"
	"runtime"
	"sort"

	"github.com/fishedee/tools/plode"
	"github.com/fishedee/tools/query/internal/query"
)

// SelectMacroHandler 基础类函数QuerySelect
type SelectMacroHandler[T, R any] func(data []T, selectFunctor func(a T) R) []R

// SelectMacroRegister 注册器
func SelectMacroRegister[T, R any](data []T, selectFunctor func(a T) R, handler SelectMacroHandler[T, R]) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(selectFunctor).String()})
	selectMacroMapper[id] = handler
}

// Select select data from table
//     * First Argument:table
//     * Second Argument:select rule
//     result = query.Select(users, func(a User) Sex {
//         if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
//             return Sex{IsMale: true}
//         }
//         return Sex{IsMale: false}
//     })
//     sel := result.([]Sex)
func Select[T, R any](data []T, selectFunctor func(a T) R) []R {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(selectFunctor).String()})
	handler, isExist := selectMacroMapper[id]
	if isExist {
		return handler.(SelectMacroHandler[T, R])(data, selectFunctor)
	}

	reflectWarn("QuerySelect")
	return query.SelectReflect(data, selectFunctor)
}

// WhereMacroHandler 基础类函数QueryWhere
type WhereMacroHandler func(data interface{}, whereFunctor interface{}) interface{}

// WhereMacroRegister 注册器
func WhereMacroRegister(data interface{}, whereFunctor interface{}, handler WhereMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(whereFunctor).String()})
	whereMacroMapper[id] = handler
}

// Where filter data from table
//     * First Argument:table
//     * Second Argument:filter rule
//     result = query.Where(users, func(a User) bool {
//         if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
//             return true
//         }
//         return false
//     })
//     where := result.([]User)
func Where(data interface{}, whereFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(whereFuctor).String()})
	handler, isExist := whereMacroMapper[id]
	if isExist {
		return handler(data, whereFuctor)
	}

	reflectWarn("QueryWhere")
	return query.WhereReflect(data, whereFuctor)
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

// Sort sort data from table,support multiple column,for Example: UserID desc,Age asc
//     * First Argument:table
//     * Second Argument:sort condition
//     result = query.Sort(users, "UserID asc")
//     sort := result.([]User)
func Sort(data interface{}, sortType string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), sortType})
	handler, isExist := sortMacroMapper[id]
	if isExist {
		return handler(data, sortType)
	}

	reflectWarn("QuerySort")
	return query.SortReflect(data, sortType)
}

// JoinMacroHandler 基础类函数QueryJoin
type JoinMacroHandler func(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}) interface{}

// JoinMacroRegister 注册
func JoinMacroRegister(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}, handler JoinMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	joinMacroMapper[id] = handler
}

// Join see LeftJoin
func Join(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	handler, isExist := joinMacroMapper[id]
	if isExist {
		return handler(leftData, rightData, joinPlace, joinType, joinFuctor)
	}

	reflectWarn("QueryJoin")
	return query.JoinReflect(leftData, rightData, joinPlace, joinType, joinFuctor)
}

// GroupMacroHandler 基础类函数 QueryGroup
type GroupMacroHandler func(data interface{}, groupType string, groupFunctor interface{}) interface{}

// GroupMacroRegister 注册
func GroupMacroRegister(data interface{}, groupType string, groupFunctor interface{}, handler GroupMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), groupType, reflect.TypeOf(groupFunctor).String()})
	groupMacroMapper[id] = handler
}

// Group group data from table
//     * First Argument: left table
//     * Second Argument: group column name
//     * Third Argument: group rule
//     result = query.Group(users, "UserID", func(users []User) Department {
//         return Department{
//             Employees: users,
//         }
//     })
//     group := result.([]Department)
func Group(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), groupType, reflect.TypeOf(groupFunctor).String()})
	handler, isExist := groupMacroMapper[id]
	if isExist {
		return handler(data, groupType, groupFunctor)
	}

	reflectWarn("QueryGroup")
	return query.GroupReflect(data, groupType, groupFunctor)
}

// ColumnMacroHandler 扩展类函数 QueryColumn
type ColumnMacroHandler func(data interface{}, column string) interface{}

// ColumnMacroRegister 注册
func ColumnMacroRegister(data interface{}, column string, handler ColumnMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMacroMapper[id] = handler
}

// Column extract column from table
//     * First Argument:table
//     * Second Argument:column name
//     result := query.Column(users, "UserID")
//     userIDs := result.([]int)
func Column(data interface{}, column string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMacroMapper[id]
	if isExist {
		return handler(data, column)
	}

	reflectWarn("QueryColumn")
	return query.ColumnReflect(data, column)
}

// ColumnMapMacroHandler 扩展类函数 QueryColumnMap
type ColumnMapMacroHandler func(data interface{}, column string) interface{}

// ColumnMapMacroRegister 注册
func ColumnMapMacroRegister(data interface{}, column string, handler ColumnMapMacroHandler) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMapMacroMapper[id] = handler
}

// ColumnMap generate a map from table,key is column value and value is it's row
//     * First Argument:table
//     * Second Argument:column name
//     result = query.ColumnMap(users, "UserID")
//     userMap := result.(map[int]User)
func ColumnMap(data interface{}, column string) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMapMacroMapper[id]
	if isExist {
		return handler(data, column)
	}

	reflectWarn("QueryColumnMap")
	return query.ColumnMapReflect(data, column)
}

// LeftJoin join data from two table，support LeftJoin,RightJoin,InnerJoin和OuterJoin
//     * First Argument: left table
//     * Second Argument: right table
//     * Third Argument: join condition
//     * Forth Argument: join rule
//     result = query.LeftJoin(admins, users, "AdminID = UserID", func(admin Admin, user User) AdminUser {
//         return AdminUser{
//             AdminID:    admin.AdminID,
//             Level:      admin.Level,
//             Name:       user.Name,
//             CreateTime: user.CreateTime,
//         }
//     })
//     join := result.([]AdminUser)
func LeftJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "left", joinType, joinFuctor)
}

// RightJoin see LeftJoin
func RightJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "right", joinType, joinFuctor)
}

// InnerJoin see LeftJoin
func InnerJoin(leftData interface{}, rightData interface{}, joinType string, joinFuctor interface{}) interface{} {
	return Join(leftData, rightData, "inner", joinType, joinFuctor)
}

// OuterJoin see LeftJoin
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

// Combine combine data from two table , one by one
//     * First Argument:left table
//     * Second Argument:right table
//     * Third Argument:combine rule
//     result = query.Combine(admins, users, func(admin Admin, user User) AdminUser {
//         return AdminUser{
//             AdminID:    admin.AdminID,
//             Level:      admin.Level,
//             Name:       user.Name,
//             CreateTime: user.CreateTime,
//         }
//     })
//     combine := result.([]AdminUser)
func Combine(leftData interface{}, rightData interface{}, combineFuctor interface{}) interface{} {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), reflect.TypeOf(combineFuctor).String()})
	handler, isExist := combineMacroMapper[id]
	if isExist {
		return handler(leftData, rightData, combineFuctor)
	}

	reflectWarn("QueryCombine")
	return query.CombineReflect(leftData, rightData, combineFuctor)
}

// Reduce reduce from list to single
//     query.Reduce([]User{}, func(sum int, singleData User) int {
//         return 1
//     }, 0)
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

// Sum get the sum of data.
// only support int, float32, float64 type
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

// Max get max value from table.
// only support int, float32, float64 type
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

// Min get min value from table.
// only support int, float32, float64 type
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

// Reverse reverse data in table
//     query.Reverse(
//         []User{},
//     )
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

// Distinct unique by column
//     result := query.Distinct([]User{}, "Name")
//     dis := result.([]User{})
func Distinct(data interface{}, columnNames string) interface{} {
	//提取信息
	name := plode.Explode(columnNames, ",")
	extractInfo := []query.QueryExtract{}
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type().Elem()
	for _, singleName := range name {
		_, extract := query.GetQueryExtract(dataType, singleName)
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
		if !isExist {
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
		if !isExist {
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
	selectMacroMapper    = map[int64]any{}
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
