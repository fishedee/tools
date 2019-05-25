package testdata

import (
	"time"

	"github.com/fishedee/tools/query"
)

// TestCase 测试用例
type TestCase struct {
	Handler func() interface{}
	Target  interface{}
}

// ContentType 测试类型 支持bool,int,float,string和time.Time
type ContentType struct {
	Name      string
	Age       int
	Ok        bool
	Money     float32
	CardMoney float64
	Register  time.Time
}

// QueryInnerStruct QueryInnerStruct
type QueryInnerStruct struct {
	MM int
}

// QueryInnerStruct2 QueryInnerStruct2
type QueryInnerStruct2 struct {
	QueryInnerStruct
	MM int
	DD float32
}

// GetQueryColumnTestCase GetQueryColumnTestCase
func GetQueryColumnTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Column([]int{}, " . ")
			},
			[]int{},
		},
		{
			func() interface{} {
				return query.Column([]string{"1", "7", "8"}, " . ")
			},
			[]string{"1", "7", "8"},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{}, " Name ")
			},
			[]string{},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{
					ContentType{"a", 3, true, 0, 0, nowTime},
					ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
					ContentType{"1", 10, true, -2.2, -1.2, oldTime},
					ContentType{"-1", -2, false, 0, 0, zeroTime},
					ContentType{"z", 3, true, 0, 0, nowTime},
				}, "     Name         ")
			},
			[]string{"a", "0", "1", "-1", "z"},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{
					ContentType{"a", 3, true, 0, 0, nowTime},
					ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
					ContentType{"1", 10, true, -2.2, -1.2, oldTime},
					ContentType{"-1", -2, false, 0, 0, zeroTime},
					ContentType{"z", 3, true, 0, 0, nowTime},
				},
					"Age        ")
			},
			[]int{3, -1, 10, -2, 3},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{
					ContentType{"a", 3, true, 0, 0, nowTime},
					ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
					ContentType{"1", 10, true, -2.2, -1.2, oldTime},
					ContentType{"-1", -2, false, 0, 0, zeroTime},
					ContentType{"z", 3, true, 0, 0, nowTime},
				},
					"Ok        ")
			},
			[]bool{true, false, true, false, true},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{
					ContentType{"a", 3, true, 0, 0, nowTime},
					ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
					ContentType{"1", 10, true, -2.2, -1.2, oldTime},
					ContentType{"-1", -2, false, 0, 0, zeroTime},
					ContentType{"z", 3, true, 0, 0, nowTime},
				},
					"    Money  ")
			},
			[]float32{0, 1.1, -2.2, 0, 0},
		},
		{
			func() interface{} {
				return query.Column([]ContentType{
					ContentType{"a", 3, true, 0, 0, nowTime},
					ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
					ContentType{"1", 10, true, -2.2, -1.2, oldTime},
					ContentType{"-1", -2, false, 0, 0, zeroTime},
					ContentType{"z", 3, true, 0, 0, nowTime},
				},
					"    CardMoney")
			},
			[]float64{0, 1.1, -1.2, 0, 0},
		},
		// FIXME:
		// {
		// 	func() interface{} {
		// 		return query.Column([]QueryInnerStruct2{
		// 			QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
		// 			QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
		// 			QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
		// 		},
		// 			"QueryInnerStruct.MM")
		// 	},
		// 	[]int{1, 2, 3},
		// },
	}
	return testCase
}

/*
// SelectCase 测试用例
type SelectCase struct {
	Origin   interface{}
	Function interface{}
	Target   interface{}
}

// GetQuerySelectTestCase GetQuerySelectTestCase
func GetQuerySelectTestCase() []SelectCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []SelectCase{
		{
			[]ContentType{},
			func(singleData ContentType) ContentType {
				return singleData
			},
			[]ContentType{},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) ContentType {

				singleData.Name += "Edward"
				return singleData
			},
			[]ContentType{
				ContentType{"5Edward", 1, true, -1.1, -1.1, oldTime},
				ContentType{"Edward", 0, false, 0, 0, zeroTime},
				ContentType{"aEdward", -1, false, 1.1, 1.1, nowTime},
			},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) ContentType {

				singleData.Name += "Edward"
				return singleData
			},
			[]ContentType{
				ContentType{"5Edward", 1, true, -1.1, -1.1, oldTime},
				ContentType{"Edward", 0, false, 0, 0, zeroTime},
				ContentType{"aEdward", -1, false, 1.1, 1.1, nowTime},
			},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) string {

				return singleData.Name
			},
			[]string{"5", "", "a"},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) int {

				return singleData.Age
			},
			[]int{1, 0, -1},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) bool {

				return singleData.Ok
			},
			[]bool{true, false, false},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) float32 {

				return singleData.Money
			},
			[]float32{-1.1, 0, 1.1},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) float64 {

				return singleData.CardMoney
			},
			[]float64{-1.1, 0, 1.1},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) time.Time {

				return singleData.Register
			},
			[]time.Time{oldTime, zeroTime, nowTime},
		},
		{
			[]ContentType{
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, nowTime},
			},
			func(singleData ContentType) map[string]int {

				return map[string]int{singleData.Name: singleData.Age}
			},
			[]map[string]int{{"5": 1}, {"": 0}, {"a": -1}},
		},
	}

	return testCase
}

// GetQueryWhereTestCase GetQueryWhereTestCase
func GetQueryWhereTestCase() []SelectCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []SelectCase{
		{
			[]ContentType{},
			func(singleData ContentType) bool {
				return true
			},
			[]ContentType{},
		},
		{
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			func(singleData ContentType) bool {
				return singleData.Age >= 1
			},
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
		},
	}

	return testCase
}

// ReduceCase ReduceCase
type ReduceCase struct {
	Origin   interface{}
	Function interface{}
	InitNum  int
	Target   interface{}
}

// GetQueryReduceTestCase GetQueryReduceTestCase
func GetQueryReduceTestCase() []ReduceCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []ReduceCase{
		{
			[]ContentType{},
			func(sum int, singleData ContentType) int {
				return 1
			},
			0,
			0,
		},
		{
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			func(sum int, singleData ContentType) int {
				return singleData.Age + sum
			},
			0,
			15,
		},
		{
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 2.2, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			func(sum float32, singleData ContentType) float32 {
				return singleData.Money + sum
			},
			0,
			(float32)(1.1),
		},
		{
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 2.2, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -2.2, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			func(sum float64, singleData ContentType) float64 {
				return singleData.CardMoney + sum
			},
			0,
			-1.1,
		},
	}

	return testCase
}

// SortCase SortCase
type SortCase struct {
	SortName string
	Origin   interface{}
	Target   interface{}
}

// GetQuerySortTestCase GetQuerySortTestCase
func GetQuerySortTestCase() []SortCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []SortCase{
		//空集
		{
			"Name desc",
			[]ContentType{},
			[]ContentType{},
		},
		{
			". asc",
			[]int{},
			[]int{},
		},
		{
			". asc",
			[]int{3, 8, 2, 9, -1},
			[]int{-1, 2, 3, 8, 9},
		},
		{
			". desc",
			[]int{3, 8, 2, 9, -1},
			[]int{9, 8, 3, 2, -1},
		},
		{
			"Name desc",
			[]ContentType{
				ContentType{"5", 0, true, -1.1, -1.1, oldTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
			[]ContentType{
				ContentType{"z", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
			},
		},
		{
			"Age desc,Ok desc",
			[]ContentType{
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			"Money,Register desc",
			[]ContentType{
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			"CardMoney,Register desc",
			[]ContentType{
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			"Ok desc,Name",
			[]ContentType{
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			" Money desc,Age asc",
			[]ContentType{
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"z", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			" Money desc,Age asc,Name desc",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			" Money desc,Age asc,Name desc",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 5, false, 0, 0, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			"MM desc",
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 3.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 5, 2.1},
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 5, 2.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 3.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 2, 1.1},
			},
		},
		{
			"QueryInnerStruct.MM asc",
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 3.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 5, 2.1},
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 5, 2.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 3.1},
			},
		},
		{
			"Name asc",
			[]ContentType{
				ContentType{"4", 7, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", -9, true, 0, 0, nowTime},
				ContentType{"5", 3, true, 0, 0, nowTime},
				ContentType{"4", 13, true, 0, 0, nowTime},
				ContentType{"7", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 9, true, 0, 0, nowTime},
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"1", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"4", 6, true, 0, 0, nowTime},
				ContentType{"5", 2, false, 0, 0, zeroTime},
				ContentType{"5", 7, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"4", 7, true, 0, 0, nowTime},
				ContentType{"4", 13, true, 0, 0, nowTime},
				ContentType{"4", 6, true, 0, 0, nowTime},
				ContentType{"5", -9, true, 0, 0, nowTime},
				ContentType{"5", 3, true, 0, 0, nowTime},
				ContentType{"5", 9, true, 0, 0, nowTime},
				ContentType{"5", 1, true, -1.1, -1.1, oldTime},
				ContentType{"5", 2, false, 0, 0, zeroTime},
				ContentType{"5", 7, true, 0, 0, nowTime},
				ContentType{"7", -1, false, 1.1, 1.1, zeroTime},
			},
		},
	}

	return testCase
}

// UserType 测试类型 支持bool,int,float,string和time.Time
type UserType struct {
	Name      string
	Age       int
	Ok        bool
	Money     float64
	CardMoney float64
	Register  time.Time
}

// JoinCase JoinCase
type JoinCase struct {
	LeftData   interface{}
	RightData  interface{}
	JoinPlace  string
	JoinType   string
	JoinFuctor interface{}
	Target     interface{}
}

// GetQueryJoinTestCase GetQueryJoinTestCase
func GetQueryJoinTestCase() []JoinCase {

	type ContentType struct {
		UserName string
		Title    string
		Content  string
	}

	type BaseType struct {
		ContentID int
	}

	type ExtendType struct {
		BaseType
		Title string
	}

	type resultType struct {
		UserName  string
		Age       int
		Ok        bool
		Money     float64
		CardMoney float64
		Register  time.Time
		Title     string
		Content   string
	}

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []JoinCase{
		{
			[]string{},
			[]UserType{},
			"left",
			" . = Name",
			func(left string, right UserType) UserType {
				return UserType{}
			},
			[]UserType{},
		},
		{

			[]int{},
			[]ExtendType{},
			" left ",
			"  .  =  ContentID ",
			func(left int, right ExtendType) ExtendType {
				return ExtendType{}
			},
			[]ExtendType{},
		},
		{

			[]UserType{},
			[]UserType{},
			" left ",
			"  Name  =  Name ",
			func(left UserType, right UserType) UserType {
				return UserType{}
			},
			[]UserType{},
		},
		{

			[]ExtendType{},
			[]ExtendType{},
			" left ",
			"  ContentID  =  ContentID ",
			func(left ExtendType, right ExtendType) ExtendType {
				return ExtendType{}
			},
			[]ExtendType{},
		},
		{

			[]string{"edward", "fish", "jd"},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				//  ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				//  ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"left",
			"  .  =  UserName ",
			func(left string, right ContentType) ContentType {
				return ContentType{
					UserName: left,
					Title:    right.Title,
					Content:  right.Content,
				}
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				ContentType{"fish", "", ""},
				ContentType{"jd", "", ""},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				UserType{"fish", -1, true, -1.1, -1.1, zeroTime},
				UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				//  ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				//  ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"left",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  left.Name,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				resultType{"fish", -1, true, -1.1, -1.1, zeroTime, "", ""},
				resultType{"jd", 1, false, 0, 1.1, oldTime, "", ""},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				// UserType{"fish", -1, true, -1.1, -1.1, zeroTime},
				// UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"left",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  left.Name,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				// UserType{"fish", -1, true, -1.1, -1.1, nowTime},
				// UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"right",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  left.Name,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				resultType{"", 0, false, 0, 0, zeroTime, "曲奇制作", "制作方法非常简单"},
				resultType{"", 0, false, 0, 0, zeroTime, "马卡龙", "好吃好玩"},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				UserType{"fish", -1, true, -1.1, -1.1, nowTime},
				UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				// ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				// ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"right",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  left.Name,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				UserType{"fish", -1, true, -1.1, -1.1, nowTime},
				UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				// ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				// ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"inner",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  left.Name,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{

			[]UserType{
				UserType{"edward", 0, false, 1.1, 0, nowTime},
				UserType{"fish", -1, true, -1.1, -1.1, nowTime},
				UserType{"jd", 1, false, 0, 1.1, oldTime},
			},
			[]ContentType{
				ContentType{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				ContentType{"weinmey", "曲奇制作", "制作方法非常简单"},
				ContentType{"jd", "马卡龙", "好吃好玩"},
			},
			"outer",
			"  Name  =  UserName ",
			func(left UserType, right ContentType) resultType {
				return resultType{
					UserName:  right.UserName,
					Age:       left.Age,
					Ok:        left.Ok,
					Money:     left.Money,
					CardMoney: left.CardMoney,
					Register:  left.Register,
					Title:     right.Title,
					Content:   right.Content,
				}
			},
			[]resultType{
				resultType{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				resultType{"", -1, true, -1.1, -1.1, nowTime, "", ""},
				resultType{"jd", 1, false, 0, 1.1, oldTime, "马卡龙", "好吃好玩"},
				resultType{"weinmey", 0, false, 0, 0, zeroTime, "曲奇制作", "制作方法非常简单"},
			},
		},

		{
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, oldTime},
				UserType{"a", -1, true, -1.1, -1.1, nowTime},
				UserType{"", 1, false, 1.1, 1.1, oldTime},
				UserType{"", 0, true, 0, 0, zeroTime},
				UserType{"z", -1, false, 0, 0, oldTime},
			},
			[]UserType{
				UserType{"s", -1, true, 0, 0, nowTime},
				UserType{"a", 0, false, 1.1, 1.1, zeroTime},
				UserType{"", -1, true, -1.1, -1.1, oldTime},
				UserType{"", 1, false, 1, 1, zeroTime},
				UserType{"z", 1, true, -1, -1, nowTime},
			},
			"right",
			"Age=Age",
			func(left UserType, right UserType) UserType {
				return UserType{
					Name:      left.Name,
					Age:       right.Age,
					Ok:        left.Ok,
					Money:     right.Money,
					CardMoney: left.CardMoney,
					Register:  right.Register,
				}
			},
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, zeroTime},
				UserType{"a", -1, true, 0, -1.1, nowTime},
				UserType{"a", -1, true, -1.1, -1.1, oldTime},
				UserType{"", 1, false, 1, 1.1, zeroTime},
				UserType{"", 1, false, -1, 1.1, nowTime},
				UserType{"", 0, true, 1.1, 0, zeroTime},
				UserType{"z", -1, false, 0, 0, nowTime},
				UserType{"z", -1, false, -1.1, 0, oldTime},
			},
		},
		{
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, oldTime},
				UserType{"a", -1, true, -1.1, -1.1, nowTime},
				UserType{"", 1, false, 1.1, 1.1, oldTime},
			},
			[]UserType{
				UserType{"s", -1, true, 0, 0, nowTime},
				UserType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			"left",
			"Ok  =  Ok",
			func(left UserType, right UserType) UserType {
				return UserType{
					Name:      left.Name,
					Age:       right.Age,
					Ok:        left.Ok,
					Money:     right.Money,
					CardMoney: left.CardMoney,
					Register:  right.Register,
				}
			},
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, zeroTime},
				UserType{"a", -1, true, 0, -1.1, nowTime},
				UserType{"", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, oldTime},
				UserType{"a", -1, true, -1.1, -1.1, nowTime},
				UserType{"", 1, false, 0, 1.1, oldTime},
			},
			[]UserType{
				UserType{"s", -1, true, 0, 0, nowTime},
				UserType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			"left",
			" Money=Money ",
			func(left UserType, right UserType) UserType {
				return UserType{
					Name:      left.Name,
					Age:       right.Age,
					Ok:        left.Ok,
					Money:     right.Money,
					CardMoney: left.CardMoney,
					Register:  right.Register,
				}
			},
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, zeroTime},
				UserType{"a", 0, true, 0, -1.1, zeroTime},
				UserType{"", -1, false, 0, 1.1, nowTime},
			},
		},

		{
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, oldTime},
				UserType{"a", -1, true, -1.1, -1.1, nowTime},
				UserType{"", 1, false, 0, 1.1, oldTime},
			},
			[]UserType{
				UserType{"s", -1, true, 0, 0, nowTime},
				UserType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			"left",
			" CardMoney = Money ",
			func(left UserType, right UserType) UserType {
				return UserType{
					Name:      left.Name,
					Age:       right.Age,
					Ok:        left.Ok,
					Money:     right.Money,
					CardMoney: left.CardMoney,
					Register:  right.Register,
				}
			},
			[]UserType{
				UserType{"s", -1, false, 0, 0, nowTime},
				UserType{"a", 0, true, 0, -1.1, zeroTime},
				UserType{"", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			[]UserType{
				UserType{"s", 0, false, 1.1, 0, oldTime},
				UserType{"a", -1, true, -1.1, -1.1, nowTime},
				UserType{"", 1, false, 0, 1.1, oldTime},
			},
			[]UserType{
				UserType{"s", -1, true, 0, 0, nowTime},
				UserType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			"left",
			" Register = Register ",
			func(left UserType, right UserType) UserType {
				return UserType{
					Name:      left.Name,
					Age:       right.Age,
					Ok:        left.Ok,
					Money:     right.Money,
					CardMoney: left.CardMoney,
					Register:  right.Register,
				}
			},
			[]UserType{
				UserType{"s", 0, false, 0, 0, zeroTime},
				UserType{"a", -1, true, 0, -1.1, nowTime},
				UserType{"", 0, false, 0, 1.1, zeroTime},
			},
		},
		{
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 1, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 2, 2.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 3, 3.1},
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 4.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 5, 5.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 6, 6.1},
			},
			"left",
			"QueryInnerStruct.MM = QueryInnerStruct.MM",
			func(left QueryInnerStruct2, right QueryInnerStruct2) QueryInnerStruct2 {
				return QueryInnerStruct2{
					left.QueryInnerStruct,
					right.MM,
					left.DD,
				}
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 5, 2.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 6, 3.1},
			},
		},
	}

	return testCase
}

// GroupCase GroupCase
type GroupCase struct {
	Data        interface{}
	GroupType   string
	GroupFuctor interface{}
	Target      interface{}
}

// GetQueryGroupTestCase GetQueryGroupTestCase
func GetQueryGroupTestCase() []GroupCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []GroupCase{
		{
			[]int{},
			".",
			func(Data []int) int {
				return len(Data)
			},
			[]int{},
		},
		{
			[]ContentType{},
			" Ok ",
			func(list []ContentType) []ContentType {
				return []ContentType{}
			},
			[]ContentType{},
		},
		{
			[]string{"a", "a", "", "", "z"},
			".",
			func(list []string) ContentType {
				return ContentType{
					Name: list[0],
					Age:  len(list),
				}
			},
			[]ContentType{
				ContentType{"a", 2, false, 0, 0, zeroTime},
				ContentType{"", 2, false, 0, 0, zeroTime},
				ContentType{"z", 1, false, 0, 0, zeroTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -2.2, -1.2, oldTime},
				ContentType{"", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Name",
			func(list []ContentType) []ContentType {
				sum := query.Sum(query.Column(list, "  Money  "))
				list[0].Money = sum.(float32)
				return []ContentType{list[0]}
			},
			[]ContentType{
				ContentType{"a", 3, true, 1.1, 0, nowTime},
				ContentType{"", 10, true, -2.2, -1.2, oldTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -2.2, -1.2, oldTime},
				ContentType{"", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Name",
			func(list []ContentType) float32 {
				sum := query.Sum(query.Column(list, "  Money  ")).(float32)
				return sum
			},
			[]float32{1.1, -2.2, 0},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -2.2, -1.2, oldTime},
				ContentType{"", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Ok",
			func(list []ContentType) []ContentType {
				sum := query.Sum(query.Column(list, "CardMoney  "))
				list[0].CardMoney = sum.(float64)
				return []ContentType{list[0]}
			},
			[]ContentType{
				ContentType{"a", 3, true, 0, -1.2, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", -1, false, 0, 0, zeroTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
			},
			" Age ",
			func(list []ContentType) []ContentType {
				sum := query.Sum(query.Column(list, "  CardMoney  "))
				list[0].CardMoney = sum.(float64)
				return []ContentType{list[0]}
			},
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
			},
		},
		{
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", -1, false, 0, 0, zeroTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
			},
			" Age ",
			func(list []ContentType) float64 {
				sum := query.Sum(query.Column(list, "  CardMoney  ")).(float64)
				return sum

			},
			[]float64{0, 1.1, -1.1},
		},
		{
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", -1, false, 0, 0, zeroTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
			},
			" Age ",
			func(list []ContentType) []float64 {
				sum := query.Sum(query.Column(list, "  CardMoney  "))
				return []float64{sum.(float64)}

			},
			[]float64{0, 1.1, -1.1},
		},
		{
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", -1, false, 0, 0, zeroTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
			},
			"Register ",
			func(list []ContentType) int {
				sum := query.Sum(query.Column(list, "  Age  "))
				return sum.(int)

			},
			[]int{2, -2, 0},
		},
		{
			[]ContentType{
				ContentType{"s", 1, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
				ContentType{"", -1, false, 0, 0, zeroTime},
				ContentType{"z", 1, true, 0, 0, nowTime},
			},
			"Register ",
			func(list []ContentType) []ContentType {
				sum := query.Sum(query.Column(list, "  Age  "))
				list[0].Age = sum.(int)
				return []ContentType{list[0]}
			},
			[]ContentType{
				ContentType{"s", 2, true, 0, 0, nowTime},
				ContentType{"a", -2, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, true, -1.1, -1.1, oldTime},
			},
		},
		{
			[] ContentType{
				 ContentType{"s", 1, true, 0, 0, nowTime},
				 ContentType{"s", 1, true, 6.6, 6.6, nowTime},
				 ContentType{"", 0, true, -5.1, -5.1, oldTime},
				 ContentType{"", 0, true, 2.1, 2.1, oldTime},
				 ContentType{"", -1, false, -3.3, -3.3, zeroTime},
				 ContentType{"", -1, false, 4.3, 4.3, zeroTime},
			},
			" Name , Ok ",
			func(list [] ContentType) [] ContentType {
				sum := query.Sum(query.Column(list, "  Age  "))
				list[0].Age = sum.(int)
				return [] ContentType{list[0]}
			},
			[] ContentType{
				 ContentType{"", -2, false, -3.3, -3.3, zeroTime},
				 ContentType{"", 0, true, -5.1, -5.1, oldTime},
				 ContentType{"s", 2, true, 0, 0, nowTime},
			},
		},
		{
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 4.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 5, 5.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 6, 6.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 6, 6.1},
			},
			"QueryInnerStruct.MM",
			func(list []QueryInnerStruct2) []QueryInnerStruct2 {
				sum := query.Sum(query.Column(list, "  MM  "))
				list[0].MM = sum.(int)
				return []QueryInnerStruct2{list[0]}
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{3}, 4, 4.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 11, 5.1},
				QueryInnerStruct2{QueryInnerStruct{1}, 6, 6.1},
			},
		},
	}

	return testCase
}

// ColumnCase ColumnCase
type ColumnCase struct {
	Data   interface{}
	Column string
	Target interface{}
}

// GetQueryColumnTestCase GetQueryColumnTestCase
func GetQueryColumnTestCase() []ColumnCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []ColumnCase{
		{
			[]int{},
			" . ",
			[]int{},
		},
		{
			[]string{"1", "7", "8"},
			" . ",
			[]string{"1", "7", "8"},
		},
		{
			[]ContentType{},
			" Name ",
			[]string{},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"     Name         ",
			[]string{"a", "0", "1", "-1", "z"},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Age        ",
			[]int{3, -1, 10, -2, 3},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Ok        ",
			[]bool{true, false, true, false, true},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"    Money  ",
			[]float32{0, 1.1, -2.2, 0, 0},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"    CardMoney",
			[]float64{0, 1.1, -1.2, 0, 0},
		},
		{
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
			},
			"QueryInnerStruct.MM",
			[]int{1, 2, 3},
		},
	}

	return testCase
}

// ColumnMapCase ColumnMapCase
type ColumnMapCase struct {
	Data   interface{}
	Column string
	Target interface{}
}

// GetQueryColumnMapTestCase GetQueryColumnMapTestCase
func GetQueryColumnMapTestCase() []ColumnMapCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []ColumnMapCase{
		{
			[]int{},
			" . ",
			map[int]int{},
		},
		{
			[]string{"1", "7", "8"},
			" . ",
			map[string]string{"1": "1", "7": "7", "8": "8"},
		},
		{
			[]ContentType{},
			" Name ",
			map[string]ContentType{},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"     Name         ",
			map[string]ContentType{
				"a":  ContentType{"a", 3, true, 0, 0, nowTime},
				"0":  ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				"1":  ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				"-1": ContentType{"-1", -2, false, 0, 0, zeroTime},
				"z":  ContentType{"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Age        ",
			map[int]ContentType{
				3:  ContentType{"a", 3, true, 0, 0, nowTime},
				-1: ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				10: ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				-2: ContentType{"-1", -2, false, 0, 0, zeroTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"Ok        ",
			map[bool]ContentType{
				true:  ContentType{"a", 3, true, 0, 0, nowTime},
				false: ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"    Money  ",
			map[float32]ContentType{
				0:    ContentType{"a", 3, true, 0, 0, nowTime},
				1.1:  ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				-2.2: ContentType{"1", 10, true, -2.2, -1.2, oldTime},
			},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			"    CardMoney",
			map[float64]ContentType{
				0:    ContentType{"a", 3, true, 0, 0, nowTime},
				1.1:  ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				-1.2: ContentType{"1", 10, true, -2.2, -1.2, oldTime},
			},
		},
		{
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
			},
			"QueryInnerStruct.MM",
			map[int]QueryInnerStruct2{
				1: QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
				2: QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
				3: QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
			},
		},
	}

	return testCase
}

// ReverseCase ReverseCase
type ReverseCase struct {
	Data   interface{}
	Target interface{}
}

// GetQueryReverseTestCase GetQueryReverseTestCase
func GetQueryReverseTestCase() []ReverseCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []ReverseCase{
		{
			[]ContentType{},
			[]ContentType{},
		},
		{
			[]ContentType{
				ContentType{"a", 3, true, 0, 0, nowTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"z", 3, true, 0, 0, nowTime},
				ContentType{"-1", -2, false, 0, 0, zeroTime},
				ContentType{"1", 10, true, -2.2, -1.2, oldTime},
				ContentType{"0", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"a", 3, true, 0, 0, nowTime},
			},
		},
	}

	return testCase
}

// CombineCase CombineCase
type CombineCase struct {
	Origin  interface{}
	Origin2 interface{}
	Functor interface{}
	Target  interface{}
}

// GetQueryCombineTestCase GetQueryCombineTestCase
func GetQueryCombineTestCase() []CombineCase {
	testCase := []CombineCase{
		{
			[]ContentType{},
			[]ContentType{},
			func(left ContentType, right ContentType) ContentType {
				return ContentType{}
			},
			[]ContentType{},
		},
		{
			[]ContentType{
				ContentType{Name: "1"},
				ContentType{Name: "2"},
				ContentType{Name: "3"},
			},
			[]int{1, 2, 3},
			func(left ContentType, right int) ContentType {
				return ContentType{
					Name: left.Name,
					Age:  right,
				}
			},
			[]ContentType{
				ContentType{Name: "1", Age: 1},
				ContentType{Name: "2", Age: 2},
				ContentType{Name: "3", Age: 3},
			},
		},
	}

	return testCase
}

// DistinctCase DistinctCase
type DistinctCase struct {
	UniqueName string
	Origin     interface{}
	Target     interface{}
}

// GetQueryDistinctTestCase GetQueryDistinctTestCase
func GetQueryDistinctTestCase() []DistinctCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []DistinctCase{
		//空集
		{
			"",
			[]ContentType{},
			[]ContentType{},
		},
		{
			"   Name    ",
			[]ContentType{},
			[]ContentType{},
		},
		{
			"",
			[]int{},
			[]int{},
		},
		//默认值
		{
			"",
			[]ContentType{
				ContentType{"", 0, false, 0, 0, zeroTime},
			},
			[]ContentType{
				ContentType{"", 0, false, 0, 0, zeroTime},
			},
		},
		//单排除
		{
			" . ",
			[]string{"s", "a", "", "", "z"},
			[]string{"s", "a", "", "z"},
		},
		{
			"Name",
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"s", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"", 10, true, -1.1, -1.1, oldTime},
				ContentType{"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			"Ok",
			[]ContentType{
				ContentType{"b", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", 3, true, 0, 0, nowTime},
			},
			[]ContentType{
				ContentType{"b", 3, true, 0, 0, nowTime},
				ContentType{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			"   Age   ",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			"   Money",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			"   CardMoney",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			"Register   ",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		//多值传递
		{
			"  Age  ,  Money",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"a", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			"  Name  ,  Money,Register  ",
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"a", 15, true, 1.1, 1.1, zeroTime},
				ContentType{"5", 0, false, -1.1, -1.1, oldTime},
			},
			[]ContentType{
				ContentType{"b", -1, true, 0, 0, nowTime},
				ContentType{"", 0, false, 0, 0, zeroTime},
				ContentType{"h", -1, true, 0, 0, nowTime},
				ContentType{"5", 10, true, -1.1, -1.1, oldTime},
				ContentType{"a", 15, true, 1.1, 1.1, zeroTime},
			},
		},
		{
			"QueryInnerStruct.MM",
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 5, 4.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
			},
			[]QueryInnerStruct2{
				QueryInnerStruct2{QueryInnerStruct{1}, 2, 1.1},
				QueryInnerStruct2{QueryInnerStruct{2}, 4, 2.1},
				QueryInnerStruct2{QueryInnerStruct{3}, 5, 3.1},
			},
		},
	}

	return testCase
}
*/
