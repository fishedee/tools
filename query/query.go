package query

import (
	"fmt"
	"log"
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
type WhereMacroHandler[T any] func(data []T, whereFunctor func(T) bool) []T

// WhereMacroRegister 注册器
func WhereMacroRegister[T any](data []T, whereFunctor func(T) bool, handler WhereMacroHandler[T]) {
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
func Where[T any](data []T, whereFuctor func(T) bool) []T {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), reflect.TypeOf(whereFuctor).String()})
	handler, isExist := whereMacroMapper[id]
	if isExist {
		return handler.(WhereMacroHandler[T])(data, whereFuctor)
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
type SortMacroHandler[T any] func(data []T, sortType string) []T

// SortMacroRegister 注册
func SortMacroRegister[T any](data []T, sortType string, handler SortMacroHandler[T]) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), sortType})
	sortMacroMapper[id] = handler
}

// Sort sort data from table,support multiple column,for Example: UserID desc,Age asc
//     * First Argument:table
//     * Second Argument:sort condition
//     result = query.Sort(users, "UserID asc")
//     sort := result.([]User)
func Sort[T any](data []T, sortType string) []T {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), sortType})
	handler, isExist := sortMacroMapper[id]
	if isExist {
		return handler.(SortMacroHandler[T])(data, sortType)
	}

	reflectWarn("QuerySort")
	return query.SortReflect(data, sortType)
}

// JoinMacroHandler 基础类函数QueryJoin
type JoinMacroHandler[L, R, LR any] func(leftData []L, rightData []R, joinPlace, joinType string, joinFuctor func(L, R) LR) []LR

// JoinMacroRegister 注册
func JoinMacroRegister[L, R, LR any](leftData []L, rightData []R, joinPlace, joinType string, joinFuctor func(L, R) LR, handler JoinMacroHandler[L, R, LR]) {
	id := registerQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	joinMacroMapper[id] = handler
}

// Join see LeftJoin
func Join[L, R, LR any](leftData []L, rightData []R, joinPlace, joinType string, joinFuctor func(L, R) LR) []LR {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), joinPlace, joinType, reflect.TypeOf(joinFuctor).String()})
	handler, isExist := joinMacroMapper[id]
	if isExist {
		return handler.(JoinMacroHandler[L, R, LR])(leftData, rightData, joinPlace, joinType, joinFuctor)
	}

	reflectWarn("QueryJoin")
	return query.JoinReflect(leftData, rightData, joinPlace, joinType, joinFuctor)
}

// GroupMacroHandler 基础类函数 QueryGroup
type GroupMacroHandler[T, E any, R *E | []E] func(data []T, groupType string, groupFunctor func([]T) E) R

// GroupMacroRegister 注册
func GroupMacroRegister[T, E any, R *E | []E](data []T, groupType string, groupFunctor func([]T) E, handler GroupMacroHandler[T, E, R]) {
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
func Group[T, E any, R *E | []E](data []T, groupType string, groupFunctor func([]T) E) R {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), groupType, reflect.TypeOf(groupFunctor).String()})
	handler, isExist := groupMacroMapper[id]
	if isExist {
		return handler.(GroupMacroHandler[T, E, R])(data, groupType, groupFunctor)
	}

	reflectWarn("QueryGroup")
	return query.GroupReflect[T, E, R](data, groupType, groupFunctor)
}

// ColumnMacroHandler 扩展类函数 QueryColumn
type ColumnMacroHandler[T, R any] func(data []T, column string) []R

// ColumnMacroRegister 注册
func ColumnMacroRegister[T, R any](data []T, column string, handler ColumnMacroHandler[T, R]) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMacroMapper[id] = handler
}

// Column extract column from table
//     * First Argument:table
//     * Second Argument:column name
//     result := query.Column(users, "UserID")
//     userIDs := result.([]int)
func Column[T, R any](data []T, column string) []R {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMacroMapper[id]
	if isExist {
		return handler.(ColumnMacroHandler[T, R])(data, column)
	}

	reflectWarn("QueryColumn")
	return query.ColumnReflect[T, R](data, column)
}

// ColumnMapMacroHandler 扩展类函数 QueryColumnMap
type ColumnMapMacroHandler[T any, K comparable, R map[K]T | map[K][]T] func(data []T, column string) R

// ColumnMapMacroRegister 注册
func ColumnMapMacroRegister[T any, K comparable, R map[K]T | map[K][]T](data []T, column string, handler ColumnMapMacroHandler[T, K, R]) {
	id := registerQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	columnMapMacroMapper[id] = handler
}

// ColumnMap generate a map from table,key is column value and value is it's row
//     * First Argument:table
//     * Second Argument:column name
//     result = query.ColumnMap(users, "UserID")
//     userMap := result.(map[int]User)
func ColumnMap[T any, K comparable, R map[K]T | map[K][]T](data []T, column string) R {
	id := getQueryTypeID([]string{reflect.TypeOf(data).String(), column})
	handler, isExist := columnMapMacroMapper[id]
	if isExist {
		return handler.(ColumnMapMacroHandler[T, K, R])(data, column)
	}

	reflectWarn("QueryColumnMap")
	return query.ColumnMapReflect[T, K, R](data, column)
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
func LeftJoin[L, R, LR any](leftData []L, rightData []R, joinType string, joinFuctor func(L, R) LR) []LR {
	return Join(leftData, rightData, "left", joinType, joinFuctor)
}

// RightJoin see LeftJoin
func RightJoin[L, R, LR any](leftData []L, rightData []R, joinType string, joinFuctor func(L, R) LR) []LR {
	return Join(leftData, rightData, "right", joinType, joinFuctor)
}

// InnerJoin see LeftJoin
func InnerJoin[L, R, LR any](leftData []L, rightData []R, joinType string, joinFuctor func(L, R) LR) []LR {
	return Join(leftData, rightData, "inner", joinType, joinFuctor)
}

// OuterJoin see LeftJoin
func OuterJoin[L, R, LR any](leftData []L, rightData []R, joinType string, joinFuctor func(L, R) LR) []LR {
	return Join(leftData, rightData, "outer", joinType, joinFuctor)
}

// CombineMacroHandler 扩展累函数 QueryCombine
type CombineMacroHandler[L, R, LR any] func(leftData []L, rightData []R, combineFuctor func(L, R) LR) []LR

// CombineMacroRegister 注册
func CombineMacroRegister[L, R, LR any](leftData []L, rightData []R, combineFuctor func(L, R) LR, handler CombineMacroHandler[L, R, LR]) {
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
func Combine[L, R, LR any](leftData []L, rightData []R, combineFuctor func(L, R) LR) []LR {
	id := getQueryTypeID([]string{reflect.TypeOf(leftData).String(), reflect.TypeOf(rightData).String(), reflect.TypeOf(combineFuctor).String()})
	handler, isExist := combineMacroMapper[id]
	if isExist {
		return handler.(CombineMacroHandler[L, R, LR])(leftData, rightData, combineFuctor)
	}

	reflectWarn("QueryCombine")
	return query.CombineReflect(leftData, rightData, combineFuctor)
}

// Reduce reduce from list to single
//     query.Reduce([]User{}, func(sum int, singleData User) int {
//         return 1
//     }, 0)
func Reduce[T, R any](data []T, reduceFuctor func(R, T) R, resultReduce R) R {
	datalen := len(data)
	for i := 0; i != datalen; i++ {
		resultReduce = reduceFuctor(resultReduce, data[i])
	}
	return resultReduce
}

type Number interface {
	~int | ~float32 | ~float64
}

// Sum get the sum of data.
// only support int, float32, float64 type
func Sum[T Number](data []T) T {
	return Reduce(data, func(sum T, single T) T {
		return sum + single
	}, 0)
}

// Max get max value from table.
// only support int, float32, float64 type
func Max[T Number](data []T) T {
	if len(data) == 0 {
		panic(fmt.Errorf("data length is 0"))
	}
	return Reduce(data, func(r T, t T) T {
		if t > r {
			return t
		}
		return r
	}, data[0])
}

// Min get min value from table.
// only support int, float32, float64 type
func Min[T Number](data []T) T {
	if len(data) == 0 {
		panic(fmt.Errorf("data length is 0"))
	}
	return Reduce(data, func(r T, t T) T {
		if t < r {
			return t
		}
		return r
	}, data[0])
}

// Reverse reverse data in table
//     query.Reverse(
//         []User{},
//     )
func Reverse[T any](data []T) []T {
	dataLen := len(data)
	result := make([]T, dataLen, dataLen)
	for i := 0; i != dataLen; i++ {
		result[dataLen-i-1] = data[i]
	}
	return result
}

// Distinct unique by column
//     result := query.Distinct([]User{}, "Name")
//     dis := result.([]User{})
func Distinct[T any](data []T, columnNames string) []T {
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
	return result.Interface().([]T)
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
	whereMacroMapper     = map[int64]any{}
	sortMacroMapper      = map[int64]any{}
	joinMacroMapper      = map[int64]any{}
	groupMacroMapper     = map[int64]any{}
	columnMacroMapper    = map[int64]any{}
	columnMapMacroMapper = map[int64]any{}
	combineMacroMapper   = map[int64]any{}
	typeIDMapper         = map[string]int64{}
	reflectWarning       = false
)
