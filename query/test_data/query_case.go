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
				return query.Column[int, int]([]int{}, " . ")
			},
			[]int{},
		},
		{
			func() interface{} {
				return query.Column[string, string]([]string{"1", "7", "8"}, " . ")
			},
			[]string{"1", "7", "8"},
		},
		{
			func() interface{} {
				return query.Column[ContentType, string]([]ContentType{}, " Name ")
			},
			[]string{},
		},
		{
			func() interface{} {
				return query.Column[ContentType, string]([]ContentType{
					{"a", 3, true, 0, 0, nowTime},
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				}, "     Name         ")
			},
			[]string{"a", "0", "1", "-1", "z"},
		},
		{
			func() interface{} {
				return query.Column[ContentType, int]([]ContentType{
					{"a", 3, true, 0, 0, nowTime},
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
					"Age        ")
			},
			[]int{3, -1, 10, -2, 3},
		},
		{
			func() interface{} {
				return query.Column[ContentType, bool]([]ContentType{
					{"a", 3, true, 0, 0, nowTime},
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
					"Ok        ")
			},
			[]bool{true, false, true, false, true},
		},
		{
			func() interface{} {
				return query.Column[ContentType, float32]([]ContentType{
					{"a", 3, true, 0, 0, nowTime},
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
					"    Money  ")
			},
			[]float32{0, 1.1, -2.2, 0, 0},
		},
		{
			func() interface{} {
				return query.Column[ContentType, float64]([]ContentType{
					{"a", 3, true, 0, 0, nowTime},
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
					"    CardMoney")
			},
			[]float64{0, 1.1, -1.2, 0, 0},
		},
		{
			func() interface{} {
				return query.Column[QueryInnerStruct2, int]([]QueryInnerStruct2{
					{QueryInnerStruct{1}, 2, 1.1},
					{QueryInnerStruct{2}, 4, 2.1},
					{QueryInnerStruct{3}, 5, 3.1},
				},
					"QueryInnerStruct.MM")
			},
			[]int{1, 2, 3},
		},
	}
	return testCase
}

// GetQuerySelectTestCase GetQuerySelectTestCase
func GetQuerySelectTestCase() []TestCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Select([]ContentType{},
					func(singleData ContentType) ContentType {
						return singleData
					})
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) ContentType {

						singleData.Name += "Edward"
						return singleData
					})
			},
			[]ContentType{
				{"5Edward", 1, true, -1.1, -1.1, oldTime},
				{"Edward", 0, false, 0, 0, zeroTime},
				{"aEdward", -1, false, 1.1, 1.1, nowTime},
			},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) ContentType {
						singleData.Name += "Edward"
						return singleData
					})
			},
			[]ContentType{
				{"5Edward", 1, true, -1.1, -1.1, oldTime},
				{"Edward", 0, false, 0, 0, zeroTime},
				{"aEdward", -1, false, 1.1, 1.1, nowTime},
			},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) string {
						return singleData.Name
					})
			},
			[]string{"5", "", "a"},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) int {
						return singleData.Age
					})
			},
			[]int{1, 0, -1},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) bool {
						return singleData.Ok
					})
			},
			[]bool{true, false, false},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) float32 {
						return singleData.Money
					})
			},
			[]float32{-1.1, 0, 1.1},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) float64 {
						return singleData.CardMoney
					})
			},
			[]float64{-1.1, 0, 1.1},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) time.Time {
						return singleData.Register
					})
			},
			[]time.Time{oldTime, zeroTime, nowTime},
		},
		{
			func() interface{} {
				return query.Select(
					[]ContentType{
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, nowTime},
					},
					func(singleData ContentType) map[string]int {
						return map[string]int{singleData.Name: singleData.Age}
					})
			},
			[]map[string]int{{"5": 1}, {"": 0}, {"a": -1}},
		},
	}

	return testCase
}

// GetQueryWhereTestCase GetQueryWhereTestCase
func GetQueryWhereTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Where(
					[]ContentType{},
					func(singleData ContentType) bool {
						return true
					})
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Where(
					[]ContentType{
						{"s", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					func(singleData ContentType) bool {
						return singleData.Age >= 1
					})
			},
			[]ContentType{
				{"s", 3, true, 0, 0, nowTime},
				{"", 10, true, -1.1, -1.1, oldTime},
				{"z", 3, true, 0, 0, nowTime},
			},
		},
	}

	return testCase
}

// GetQueryReduceTestCase GetQueryReduceTestCase
func GetQueryReduceTestCase() []TestCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Reduce(
					[]ContentType{},
					func(sum int, singleData ContentType) int {
						return 1
					},
					0)
			},
			0,
		},
		{
			func() interface{} {
				return query.Reduce(
					[]ContentType{
						{"s", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					func(sum int, singleData ContentType) int {
						return singleData.Age + sum
					},
					0)
			},
			15,
		},
		{
			func() interface{} {
				return query.Reduce(
					[]ContentType{
						{"s", 3, true, 0, 0, nowTime},
						{"a", -1, false, 2.2, 1.1, zeroTime},
						{"", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					func(sum float32, singleData ContentType) float32 {
						return singleData.Money + sum
					},
					0)
			},
			(float32)(1.1),
		},
		{
			func() interface{} {
				return query.Reduce(
					[]ContentType{
						{"s", 3, true, 0, 0, nowTime},
						{"a", -1, false, 2.2, 1.1, zeroTime},
						{"", 10, true, -1.1, -2.2, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					func(sum float64, singleData ContentType) float64 {
						return singleData.CardMoney + sum
					},
					0)
			},
			-1.1,
		},
	}

	return testCase
}

func GetQuerySumTestCase() []TestCase {

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Sum(
					[]int{1, 2, 3, 4, 5},
				)
			},
			15,
		},
		{
			func() interface{} {
				return query.Sum(
					[]float32{1.2, 2.3, 4.5, 2.6, 10.9},
				)
			},
			float32(21.5),
		},
		{
			func() interface{} {
				return query.Sum(
					[]float64{1.2, 2.3, 4.5, 2.6, 10.9, 100000000.200000000},
				)
			},
			100000021.7,
		},
	}

	return testCase
}

func GetQueryMaxTestCase() []TestCase {

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Max(
					[]int{1, 2, 3, 4, 5},
				)
			},
			5,
		},
		{
			func() interface{} {
				return query.Max(
					[]int{-1, -2, -3, -4, -5},
				)
			},
			-1,
		},
		{
			func() interface{} {
				return query.Max(
					[]int{-1, -2, 3, 4, 5},
				)
			},
			5,
		},
		{
			func() interface{} {
				return query.Max(
					[]float32{1.2, 2.3, 4.5, 2.6, 10.9},
				)
			},
			float32(10.9),
		},
		{
			func() interface{} {
				return query.Max(
					[]float64{1.2, 2.3, 4.5, 2.6, 10.9, 100000000.200000000},
				)
			},
			100000000.2,
		},
	}

	return testCase
}

func GetQueryMinTestCase() []TestCase {

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Min(
					[]int{1, 2, 3, 4, 5},
				)
			},
			1,
		},
		{
			func() interface{} {
				return query.Min(
					[]int{-1, -2, -3, -4, -5},
				)
			},
			-5,
		},
		{
			func() interface{} {
				return query.Min(
					[]int{-1, -2, 3, 4, 5},
				)
			},
			-2,
		},
		{
			func() interface{} {
				return query.Min(
					[]float32{1.2, 2.3, 4.5, 2.6, 10.9},
				)
			},
			float32(1.2),
		},
		{
			func() interface{} {
				return query.Min(
					[]float64{1.2, 2.3, 4.5, 2.6, 10.9, 100000000.200000000},
				)
			},
			1.2,
		},
	}

	return testCase
}

// GetQuerySortTestCase GetQuerySortTestCase
func GetQuerySortTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		//空集
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{},
					"Name desc",
				)
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Sort(
					[]int{},
					". asc")
			},
			[]int{},
		},
		{
			func() interface{} {
				return query.Sort(
					[]int{3, 8, 2, 9, -1},
					". asc")
			},
			[]int{-1, 2, 3, 8, 9},
		},
		{
			func() interface{} {
				return query.Sort(
					[]int{3, 8, 2, 9, -1},
					". desc")
			},
			[]int{9, 8, 3, 2, -1},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"5", 0, true, -1.1, -1.1, oldTime},
						{"z", 1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
					},
					"Name desc")
			},
			[]ContentType{
				{"z", 1, true, 0, 0, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"5", 0, true, -1.1, -1.1, oldTime},
				{"", 0, false, 0, 0, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"z", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
					},
					"Age desc,Ok desc")
			},
			[]ContentType{
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"", 5, false, 0, 0, zeroTime},
				{"z", -1, true, 0, 0, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"z", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
					},
					"Money,Register desc")
			},
			[]ContentType{
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"z", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"z", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
					},
					"CardMoney,Register desc")
			},
			[]ContentType{
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"z", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"z", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
					},
					"Ok desc,Name")
			},
			[]ContentType{
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"z", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"z", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
					},
					" Money desc,Age asc")
			},
			[]ContentType{
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"z", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
					},
					" Money desc,Age asc,Name desc")
			},
			[]ContentType{
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"h", -1, true, 0, 0, nowTime},
				{"b", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 5, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
					},
					" Money desc,Age asc,Name desc")
			},
			[]ContentType{
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"h", -1, true, 0, 0, nowTime},
				{"b", -1, true, 0, 0, nowTime},
				{"", 5, false, 0, 0, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]QueryInnerStruct2{
						{QueryInnerStruct{3}, 4, 3.1},
						{QueryInnerStruct{2}, 2, 1.1},
						{QueryInnerStruct{1}, 5, 2.1},
					},
					"MM desc")
			},
			[]QueryInnerStruct2{
				{QueryInnerStruct{1}, 5, 2.1},
				{QueryInnerStruct{3}, 4, 3.1},
				{QueryInnerStruct{2}, 2, 1.1},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]QueryInnerStruct2{
						{QueryInnerStruct{3}, 4, 3.1},
						{QueryInnerStruct{2}, 2, 1.1},
						{QueryInnerStruct{1}, 5, 2.1},
					},
					"QueryInnerStruct.MM asc")
			},
			[]QueryInnerStruct2{
				{QueryInnerStruct{1}, 5, 2.1},
				{QueryInnerStruct{2}, 2, 1.1},
				{QueryInnerStruct{3}, 4, 3.1},
			},
		},
		{
			func() interface{} {
				return query.Sort(
					[]ContentType{
						{"4", 7, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"5", -9, true, 0, 0, nowTime},
						{"5", 3, true, 0, 0, nowTime},
						{"4", 13, true, 0, 0, nowTime},
						{"7", -1, false, 1.1, 1.1, zeroTime},
						{"5", 9, true, 0, 0, nowTime},
						{"5", 1, true, -1.1, -1.1, oldTime},
						{"1", -1, false, 1.1, 1.1, zeroTime},
						{"4", 6, true, 0, 0, nowTime},
						{"5", 2, false, 0, 0, zeroTime},
						{"5", 7, true, 0, 0, nowTime},
					},
					"Name asc")
			},
			[]ContentType{
				{"0", -1, false, 1.1, 1.1, zeroTime},
				{"1", -1, false, 1.1, 1.1, zeroTime},
				{"4", 7, true, 0, 0, nowTime},
				{"4", 13, true, 0, 0, nowTime},
				{"4", 6, true, 0, 0, nowTime},
				{"5", -9, true, 0, 0, nowTime},
				{"5", 3, true, 0, 0, nowTime},
				{"5", 9, true, 0, 0, nowTime},
				{"5", 1, true, -1.1, -1.1, oldTime},
				{"5", 2, false, 0, 0, zeroTime},
				{"5", 7, true, 0, 0, nowTime},
				{"7", -1, false, 1.1, 1.1, zeroTime},
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

// ContentType2 ContentType2
type ContentType2 struct {
	UserName string
	Title    string
	Content  string
}

// BaseType BaseType
type BaseType struct {
	ContentID int
}

// ExtendType ExtendType
type ExtendType struct {
	BaseType
	Title     string
	ContentID int
}

// resultType resultType
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

// GetQueryJoinTestCase GetQueryJoinTestCase
func GetQueryJoinTestCase() []TestCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Join(
					[]string{},
					[]UserType{},
					"left",
					" . = Name",
					func(left string, right UserType) UserType {
						return UserType{}
					})
			},
			[]UserType{},
		},
		{
			func() interface{} {
				return query.Join(
					[]int{},
					[]ExtendType{},
					" left ",
					"  .  =  ContentID ",
					func(left int, right ExtendType) ExtendType {
						return ExtendType{}
					})
			},
			[]ExtendType{},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{},
					[]UserType{},
					" left ",
					"  Name  =  Name ",
					func(left UserType, right UserType) UserType {
						return UserType{}
					})
			},
			[]UserType{},
		},
		{
			func() interface{} {
				return query.Join(
					[]ExtendType{},
					[]ExtendType{},
					" left ",
					"  ContentID  =  ContentID ",
					func(left ExtendType, right ExtendType) ExtendType {
						return ExtendType{}
					})
			},
			[]ExtendType{},
		},
		{
			func() interface{} {
				return query.Join(
					[]string{"edward", "fish", "jd"},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						//  {"weinmey", "曲奇制作", "制作方法非常简单"},
						//  {"jd", "马卡龙", "好吃好玩"},
					},
					"left",
					"  .  =  UserName ",
					func(left string, right ContentType2) ContentType2 {
						return ContentType2{
							UserName: left,
							Title:    right.Title,
							Content:  right.Content,
						}
					})
			},
			[]ContentType2{
				{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
				{"fish", "", ""},
				{"jd", "", ""},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						{"fish", -1, true, -1.1, -1.1, zeroTime},
						{"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						//  {"weinmey", "曲奇制作", "制作方法非常简单"},
						//  {"jd", "马卡龙", "好吃好玩"},
					},
					"left",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				{"fish", -1, true, -1.1, -1.1, zeroTime, "", ""},
				{"jd", 1, false, 0, 1.1, oldTime, "", ""},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						// {"fish", -1, true, -1.1, -1.1, zeroTime},
						// {"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						{"weinmey", "曲奇制作", "制作方法非常简单"},
						{"jd", "马卡龙", "好吃好玩"},
					},
					"left",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						// UserType{"fish", -1, true, -1.1, -1.1, nowTime},
						// UserType{"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						{"weinmey", "曲奇制作", "制作方法非常简单"},
						{"jd", "马卡龙", "好吃好玩"},
					},
					"right",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				{"", 0, false, 0, 0, zeroTime, "曲奇制作", "制作方法非常简单"},
				{"", 0, false, 0, 0, zeroTime, "马卡龙", "好吃好玩"},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						{"fish", -1, true, -1.1, -1.1, nowTime},
						{"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						// {"weinmey", "曲奇制作", "制作方法非常简单"},
						// {"jd", "马卡龙", "好吃好玩"},
					},
					"right",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						{"fish", -1, true, -1.1, -1.1, nowTime},
						{"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						// {"weinmey", "曲奇制作", "制作方法非常简单"},
						// {"jd", "马卡龙", "好吃好玩"},
					},
					"inner",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"edward", 0, false, 1.1, 0, nowTime},
						{"fish", -1, true, -1.1, -1.1, nowTime},
						{"jd", 1, false, 0, 1.1, oldTime},
					},
					[]ContentType2{
						{"edward", "威风蛋糕", "威风蛋糕好好吃野！"},
						{"weinmey", "曲奇制作", "制作方法非常简单"},
						{"jd", "马卡龙", "好吃好玩"},
					},
					"outer",
					"  Name  =  UserName ",
					func(left UserType, right ContentType2) resultType {
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
					})
			},
			[]resultType{
				{"edward", 0, false, 1.1, 0, nowTime, "威风蛋糕", "威风蛋糕好好吃野！"},
				{"", -1, true, -1.1, -1.1, nowTime, "", ""},
				{"jd", 1, false, 0, 1.1, oldTime, "马卡龙", "好吃好玩"},
				{"weinmey", 0, false, 0, 0, zeroTime, "曲奇制作", "制作方法非常简单"},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"s", 0, false, 1.1, 0, oldTime},
						{"a", -1, true, -1.1, -1.1, nowTime},
						{"", 1, false, 1.1, 1.1, oldTime},
						{"", 0, true, 0, 0, zeroTime},
						{"z", -1, false, 0, 0, oldTime},
					},
					[]UserType{
						{"s", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
						{"", -1, true, -1.1, -1.1, oldTime},
						{"", 1, false, 1, 1, zeroTime},
						{"z", 1, true, -1, -1, nowTime},
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
					})
			},
			[]UserType{
				{"s", 0, false, 1.1, 0, zeroTime},
				{"a", -1, true, 0, -1.1, nowTime},
				{"a", -1, true, -1.1, -1.1, oldTime},
				{"", 1, false, 1, 1.1, zeroTime},
				{"", 1, false, -1, 1.1, nowTime},
				{"", 0, true, 1.1, 0, zeroTime},
				{"z", -1, false, 0, 0, nowTime},
				{"z", -1, false, -1.1, 0, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"s", 0, false, 1.1, 0, oldTime},
						{"a", -1, true, -1.1, -1.1, nowTime},
						{"", 1, false, 1.1, 1.1, oldTime},
					},
					[]UserType{
						{"s", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
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
					})
			},
			[]UserType{
				{"s", 0, false, 1.1, 0, zeroTime},
				{"a", -1, true, 0, -1.1, nowTime},
				{"", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"s", 0, false, 1.1, 0, oldTime},
						{"a", -1, true, -1.1, -1.1, nowTime},
						{"", 1, false, 0, 1.1, oldTime},
					},
					[]UserType{
						{"s", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
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
					})
			},
			[]UserType{
				{"s", 0, false, 1.1, 0, zeroTime},
				{"a", 0, true, 0, -1.1, zeroTime},
				{"", -1, false, 0, 1.1, nowTime},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"s", 0, false, 1.1, 0, oldTime},
						{"a", -1, true, -1.1, -1.1, nowTime},
						{"", 1, false, 0, 1.1, oldTime},
					},
					[]UserType{
						{"s", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
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
					})
			},
			[]UserType{
				{"s", -1, false, 0, 0, nowTime},
				{"a", 0, true, 0, -1.1, zeroTime},
				{"", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]UserType{
						{"s", 0, false, 1.1, 0, oldTime},
						{"a", -1, true, -1.1, -1.1, nowTime},
						{"", 1, false, 0, 1.1, oldTime},
					},
					[]UserType{
						{"s", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
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
					})
			},
			[]UserType{
				{"s", 0, false, 0, 0, zeroTime},
				{"a", -1, true, 0, -1.1, nowTime},
				{"", 0, false, 0, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Join(
					[]QueryInnerStruct2{
						{QueryInnerStruct{3}, 1, 1.1},
						{QueryInnerStruct{2}, 2, 2.1},
						{QueryInnerStruct{1}, 3, 3.1},
					},
					[]QueryInnerStruct2{
						{QueryInnerStruct{3}, 4, 4.1},
						{QueryInnerStruct{2}, 5, 5.1},
						{QueryInnerStruct{1}, 6, 6.1},
					},
					"left",
					"QueryInnerStruct.MM = QueryInnerStruct.MM",
					func(left QueryInnerStruct2, right QueryInnerStruct2) QueryInnerStruct2 {
						return QueryInnerStruct2{
							left.QueryInnerStruct,
							right.MM,
							left.DD,
						}
					})
			},
			[]QueryInnerStruct2{
				{QueryInnerStruct{3}, 4, 1.1},
				{QueryInnerStruct{2}, 5, 2.1},
				{QueryInnerStruct{1}, 6, 3.1},
			},
		},
	}

	return testCase
}

// GetQueryGroupTestCase GetQueryGroupTestCase
func GetQueryGroupTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Group[int, int, []int](
					[]int{},
					".",
					func(Data []int) int {
						return len(Data)
					})
			},
			[]int{},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []ContentType, *[]ContentType](
					[]ContentType{},
					" Ok ",
					func(list []ContentType) []ContentType {
						return []ContentType{}
					})
			},
			&[]ContentType{},
		},
		{
			func() interface{} {
				return query.Group[string, ContentType, []ContentType](
					[]string{"a", "a", "", "", "z"},
					".",
					func(list []string) ContentType {
						return ContentType{
							Name: list[0],
							Age:  len(list),
						}
					})
			},
			[]ContentType{
				{"a", 2, false, 0, 0, zeroTime},
				{"", 2, false, 0, 0, zeroTime},
				{"z", 1, false, 0, 0, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []ContentType, *[]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -2.2, -1.2, oldTime},
						{"", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Name",
					func(list []ContentType) []ContentType {
						sum := query.Sum(query.Column[ContentType, float32](list, "  Money  "))
						list[0].Money = sum
						return []ContentType{list[0]}
					})
			},
			&[]ContentType{
				{"a", 3, true, 1.1, 0, nowTime},
				{"", 10, true, -2.2, -1.2, oldTime},
				{"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			func() interface{} {
				return query.Group[ContentType, float32, []float32](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -2.2, -1.2, oldTime},
						{"", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Name",
					func(list []ContentType) float32 {
						sum := query.Sum(query.Column[ContentType, float32](list, "  Money  "))
						return sum
					})
			},
			[]float32{1.1, -2.2, 0},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []ContentType, *[]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -2.2, -1.2, oldTime},
						{"", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Ok",
					func(list []ContentType) []ContentType {
						sum := query.Sum(query.Column[ContentType, float64](list, "CardMoney  "))
						list[0].CardMoney = sum
						return []ContentType{list[0]}
					})
			},
			&[]ContentType{
				{"a", 3, true, 0, -1.2, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []ContentType, *[]ContentType](
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 0, true, -1.1, -1.1, oldTime},
						{"", -1, false, 0, 0, zeroTime},
						{"z", 1, true, 0, 0, nowTime},
					},
					" Age ",
					func(list []ContentType) []ContentType {
						sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  "))
						list[0].CardMoney = sum
						return []ContentType{list[0]}
					})
			},
			&[]ContentType{
				{"s", 1, true, 0, 0, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"", 0, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Group[ContentType, float64, []float64](
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 0, true, -1.1, -1.1, oldTime},
						{"", -1, false, 0, 0, zeroTime},
						{"z", 1, true, 0, 0, nowTime},
					},
					" Age ",
					func(list []ContentType) float64 {
						sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  "))
						return sum

					})
			},
			[]float64{0, 1.1, -1.1},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []float64, *[]float64](
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 0, true, -1.1, -1.1, oldTime},
						{"", -1, false, 0, 0, zeroTime},
						{"z", 1, true, 0, 0, nowTime},
					},
					" Age ",
					func(list []ContentType) []float64 {
						sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  "))
						return []float64{sum}

					})
			},
			&[]float64{0, 1.1, -1.1},
		},
		{
			func() interface{} {
				return query.Group[ContentType, int, []int](
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 0, true, -1.1, -1.1, oldTime},
						{"", -1, false, 0, 0, zeroTime},
						{"z", 1, true, 0, 0, nowTime},
					},
					"Register ",
					func(list []ContentType) int {
						sum := query.Sum(query.Column[ContentType, int](list, "  Age  "))
						return sum

					})
			},
			[]int{2, -2, 0},
		},
		{
			func() interface{} {
				return query.Group[ContentType, []ContentType, *[]ContentType](
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 0, true, -1.1, -1.1, oldTime},
						{"", -1, false, 0, 0, zeroTime},
						{"z", 1, true, 0, 0, nowTime},
					},
					"Register ",
					func(list []ContentType) []ContentType {
						sum := query.Sum(query.Column[ContentType, int](list, "  Age  "))
						list[0].Age = sum
						return []ContentType{list[0]}
					})
			},
			&[]ContentType{
				{"s", 2, true, 0, 0, nowTime},
				{"a", -2, false, 1.1, 1.1, zeroTime},
				{"", 0, true, -1.1, -1.1, oldTime},
			},
		},
		/*DO NOT SUPPORT MULTIPLE COLUMN GROUP
		{
			func() interface{} {
				return query.Group(
					[]ContentType{
						{"s", 1, true, 0, 0, nowTime},
						{"s", 1, true, 6.6, 6.6, nowTime},
						{"", 0, true, -5.1, -5.1, oldTime},
						{"", 0, true, 2.1, 2.1, oldTime},
						{"", -1, false, -3.3, -3.3, zeroTime},
						{"", -1, false, 4.3, 4.3, zeroTime},
					},
					" Name , Ok ",
					func(list []ContentType) []ContentType {
						sum := query.Sum(query.Column(list, "  Age  "))
						list[0].Age = sum.(int)
						return []ContentType{list[0]}
					})
			},
			[]ContentType{
				{"", -2, false, -3.3, -3.3, zeroTime},
				{"", 0, true, -5.1, -5.1, oldTime},
				{"s", 2, true, 0, 0, nowTime},
			},
		},
		*/
		{
			func() interface{} {
				return query.Group[QueryInnerStruct2, []QueryInnerStruct2, *[]QueryInnerStruct2](
					[]QueryInnerStruct2{
						{QueryInnerStruct{3}, 4, 4.1},
						{QueryInnerStruct{2}, 5, 5.1},
						{QueryInnerStruct{1}, 6, 6.1},
						{QueryInnerStruct{2}, 6, 6.1},
					},
					"QueryInnerStruct.MM",
					func(list []QueryInnerStruct2) []QueryInnerStruct2 {
						sum := query.Sum(query.Column[QueryInnerStruct2, int](list, "  MM  "))
						list[0].MM = sum
						return []QueryInnerStruct2{list[0]}
					})
			},
			&[]QueryInnerStruct2{
				{QueryInnerStruct{3}, 4, 4.1},
				{QueryInnerStruct{2}, 11, 5.1},
				{QueryInnerStruct{1}, 6, 6.1},
			},
		},
	}

	return testCase
}

// GetQueryColumnMapTestCase GetQueryColumnMapTestCase
func GetQueryColumnMapTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		//querycolumn single
		{
			func() interface{} {
				return query.ColumnMap[int, int, map[int]int](
					[]int{},
					" . ")
			},
			map[int]int{},
		},
		{
			func() interface{} {
				return query.ColumnMap[string, string, map[string]string](
					[]string{"1", "7", "8"},
					" . ")
			},
			map[string]string{"1": "1", "7": "7", "8": "8"},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, string, map[string]ContentType](
					[]ContentType{},
					" Name ")
			},
			map[string]ContentType{},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, string, map[string]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"     Name         ")
			},
			map[string]ContentType{
				"a":  {"a", 3, true, 0, 0, nowTime},
				"0":  {"0", -1, false, 1.1, 1.1, zeroTime},
				"1":  {"1", 10, true, -2.2, -1.2, oldTime},
				"-1": {"-1", -2, false, 0, 0, zeroTime},
				"z":  {"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, int, map[int]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Age        ")
			},
			map[int]ContentType{
				3:  {"a", 3, true, 0, 0, nowTime},
				-1: {"0", -1, false, 1.1, 1.1, zeroTime},
				10: {"1", 10, true, -2.2, -1.2, oldTime},
				-2: {"-1", -2, false, 0, 0, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, bool, map[bool]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Ok        ")
			},
			map[bool]ContentType{
				true:  {"a", 3, true, 0, 0, nowTime},
				false: {"0", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, float32, map[float32]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"    Money  ")
			},
			map[float32]ContentType{
				0:    {"a", 3, true, 0, 0, nowTime},
				1.1:  {"0", -1, false, 1.1, 1.1, zeroTime},
				-2.2: {"1", 10, true, -2.2, -1.2, oldTime},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, float64, map[float64]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"    CardMoney")
			},
			map[float64]ContentType{
				0:    {"a", 3, true, 0, 0, nowTime},
				1.1:  {"0", -1, false, 1.1, 1.1, zeroTime},
				-1.2: {"1", 10, true, -2.2, -1.2, oldTime},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[QueryInnerStruct2, int, map[int]QueryInnerStruct2](
					[]QueryInnerStruct2{
						{QueryInnerStruct{1}, 2, 1.1},
						{QueryInnerStruct{2}, 4, 2.1},
						{QueryInnerStruct{3}, 5, 3.1},
					},
					"QueryInnerStruct.MM")
			},
			map[int]QueryInnerStruct2{
				1: {QueryInnerStruct{1}, 2, 1.1},
				2: {QueryInnerStruct{2}, 4, 2.1},
				3: {QueryInnerStruct{3}, 5, 3.1},
			},
		},
		// queryColumn []slice
		{
			func() interface{} {
				return query.ColumnMap[ContentType, string, map[string][]ContentType](
					[]ContentType{},
					" []Name ")
			},
			map[string][]ContentType{},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, string, map[string][]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"     [] Name         ")
			},
			map[string][]ContentType{
				"a": {
					{"a", 3, true, 0, 0, nowTime},
				},
				"0": {
					{"0", -1, false, 1.1, 1.1, zeroTime},
				},
				"1": {
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"1", -2, false, 0, 0, zeroTime},
				},
				"z": {
					{"z", 3, true, 0, 0, nowTime},
				},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, int, map[int][]ContentType](
					[]ContentType{
						{"a", 10, true, 0, 0, nowTime},
						{"0", 10, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"[]Age        ")
			},
			map[int][]ContentType{
				10: {
					{"a", 10, true, 0, 0, nowTime},
					{"0", 10, false, 1.1, 1.1, zeroTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
				},
				-2: {
					{"-1", -2, false, 0, 0, zeroTime},
				},
				3: {
					{"z", 3, true, 0, 0, nowTime},
				},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, bool, map[bool][]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"[]Ok        ")
			},
			map[bool][]ContentType{
				true: {
					{"a", 3, true, 0, 0, nowTime},
					{"1", 10, true, -2.2, -1.2, oldTime},
					{"z", 3, true, 0, 0, nowTime},
				},
				false: {
					{"0", -1, false, 1.1, 1.1, zeroTime},
					{"-1", -2, false, 0, 0, zeroTime},
				},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, float32, map[float32][]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"    []Money  ")
			},
			map[float32][]ContentType{
				0: {
					{"a", 3, true, 0, 0, nowTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
				1.1: {
					{"0", -1, false, 1.1, 1.1, zeroTime},
				},
				-2.2: {
					{"1", 10, true, -2.2, -1.2, oldTime},
				},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[ContentType, float64, map[float64][]ContentType](
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"    []CardMoney")
			},
			map[float64][]ContentType{
				0: {
					{"a", 3, true, 0, 0, nowTime},
					{"-1", -2, false, 0, 0, zeroTime},
					{"z", 3, true, 0, 0, nowTime},
				},
				1.1: {
					{"0", -1, false, 1.1, 1.1, zeroTime},
				},
				-1.2: {
					{"1", 10, true, -2.2, -1.2, oldTime},
				},
			},
		},
		{
			func() interface{} {
				return query.ColumnMap[QueryInnerStruct2, int, map[int][]QueryInnerStruct2](
					[]QueryInnerStruct2{
						{QueryInnerStruct{1}, 2, 1.1},
						{QueryInnerStruct{2}, 4, 2.1},
						{QueryInnerStruct{3}, 5, 3.1},
					},
					"[]QueryInnerStruct.MM")
			},
			map[int][]QueryInnerStruct2{
				1: {{QueryInnerStruct{1}, 2, 1.1}},
				2: {{QueryInnerStruct{2}, 4, 2.1}},
				3: {{QueryInnerStruct{3}, 5, 3.1}},
			},
		},
	}

	return testCase
}

// GetQueryReverseTestCase GetQueryReverseTestCase
func GetQueryReverseTestCase() []TestCase {
	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		{
			func() interface{} {
				return query.Reverse(
					[]ContentType{},
				)
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Reverse(
					[]ContentType{
						{"a", 3, true, 0, 0, nowTime},
						{"0", -1, false, 1.1, 1.1, zeroTime},
						{"1", 10, true, -2.2, -1.2, oldTime},
						{"-1", -2, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					})
			},
			[]ContentType{
				{"z", 3, true, 0, 0, nowTime},
				{"-1", -2, false, 0, 0, zeroTime},
				{"1", 10, true, -2.2, -1.2, oldTime},
				{"0", -1, false, 1.1, 1.1, zeroTime},
				{"a", 3, true, 0, 0, nowTime},
			},
		},
	}

	return testCase
}

// GetQueryCombineTestCase GetQueryCombineTestCase
func GetQueryCombineTestCase() []TestCase {
	testCase := []TestCase{
		{
			func() interface{} {
				return query.Combine(
					[]ContentType{},
					[]ContentType{},
					func(left ContentType, right ContentType) ContentType {
						return ContentType{}
					})
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Combine(
					[]ContentType{
						{Name: "1"},
						{Name: "2"},
						{Name: "3"},
					},
					[]int{1, 2, 3},
					func(left ContentType, right int) ContentType {
						return ContentType{
							Name: left.Name,
							Age:  right,
						}
					})
			},
			[]ContentType{
				{Name: "1", Age: 1},
				{Name: "2", Age: 2},
				{Name: "3", Age: 3},
			},
		},
	}

	return testCase
}

// GetQueryDistinctTestCase GetQueryDistinctTestCase
func GetQueryDistinctTestCase() []TestCase {

	nowTime := time.Now()
	oldTime := nowTime.AddDate(-1, 0, 1)
	zeroTime := time.Time{}

	testCase := []TestCase{
		//空集
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{},
					"",
				)
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{},
					"   Name    ",
				)
			},
			[]ContentType{},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]int{},
					"",
				)
			},
			[]int{},
		},
		//默认值
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"", 0, false, 0, 0, zeroTime},
					},
					"",
				)
			},
			[]ContentType{
				{"", 0, false, 0, 0, zeroTime},
			},
		},
		//单排除
		{
			func() interface{} {
				return query.Distinct(
					[]string{"s", "a", "", "", "z"},
					" . ",
				)
			},
			[]string{"s", "a", "", "z"},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"s", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"z", 3, true, 0, 0, nowTime},
					},
					"Name",
				)
			},
			[]ContentType{
				{"s", 3, true, 0, 0, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
				{"", 10, true, -1.1, -1.1, oldTime},
				{"z", 3, true, 0, 0, nowTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", 3, true, 0, 0, nowTime},
						{"a", -1, false, 1.1, 1.1, zeroTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"h", 3, true, 0, 0, nowTime},
					},
					"Ok",
				)
			},
			[]ContentType{
				{"b", 3, true, 0, 0, nowTime},
				{"a", -1, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
						{"", 0, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
					},
					"   Age   ",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"a", 0, false, 1.1, 1.1, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
					},
					"   Money",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"a", 0, false, 1.1, 1.1, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
					},
					"   CardMoney",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"a", 0, false, 1.1, 1.1, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
					},
					"Register   ",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"", 0, false, 0, 0, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
			},
		},
		//多值传递
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"a", 0, false, 1.1, 1.1, zeroTime},
					},
					"  Age  ,  Money",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"", 0, false, 0, 0, zeroTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"a", 0, false, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]ContentType{
						{"b", -1, true, 0, 0, nowTime},
						{"", 0, false, 0, 0, zeroTime},
						{"h", -1, true, 0, 0, nowTime},
						{"5", 10, true, -1.1, -1.1, oldTime},
						{"", 0, false, 0, 0, zeroTime},
						{"a", 15, true, 1.1, 1.1, zeroTime},
						{"5", 0, false, -1.1, -1.1, oldTime},
					},
					"  Name  ,  Money,Register  ",
				)
			},
			[]ContentType{
				{"b", -1, true, 0, 0, nowTime},
				{"", 0, false, 0, 0, zeroTime},
				{"h", -1, true, 0, 0, nowTime},
				{"5", 10, true, -1.1, -1.1, oldTime},
				{"a", 15, true, 1.1, 1.1, zeroTime},
			},
		},
		{
			func() interface{} {
				return query.Distinct(
					[]QueryInnerStruct2{
						{QueryInnerStruct{1}, 2, 1.1},
						{QueryInnerStruct{2}, 4, 2.1},
						{QueryInnerStruct{2}, 5, 4.1},
						{QueryInnerStruct{3}, 5, 3.1},
					},
					"QueryInnerStruct.MM",
				)
			},
			[]QueryInnerStruct2{
				{QueryInnerStruct{1}, 2, 1.1},
				{QueryInnerStruct{2}, 4, 2.1},
				{QueryInnerStruct{3}, 5, 3.1},
			},
		},
	}

	return testCase
}
