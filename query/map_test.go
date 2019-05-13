package query

import (
	"reflect"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/decimal"
)

func TestArrayToMapBasic(t *testing.T) {
	now := time.Now()
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		{nil, nil},
		{true, true},
		{1, 1},
		{1.2, 1.2},
		{"12", "12"},
		{decimal.Decimal("1.2"), "1.2"},
		{decimal.Decimal(""), "0"},
		{now, now.Format("2006-01-02 15:04:05")},
		{[]int{1, 2, 3}, []interface{}{1, 2, 3}},
		{map[string]string{
			"1": "2",
			"3": "5",
		},
			map[string]interface{}{
				"1": "2",
				"3": "5",
			}},
	}
	for singleTestKey, singleTestCase := range testCase {
		data := ArrayToMap(singleTestCase.origin, "json")
		assert.Equal(t, data, singleTestCase.target, singleTestKey)
	}
}

type AnaymonusStruct struct {
	First  string
	Second string
}

type AnaymonusMap map[int]string

func TestArrayToMapStruct(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		{struct {
			First   string
			Second  string
			Third   string    `json:"Third"`
			Forth   string    `json:"Forth,omitempty"`
			Fifth   string    `json:"Fifth,omitempty"`
			Sixth   string    `json:"-"`
			Seventh time.Time `json:"seventh,omitempty"`
			Eigth   string    `json:"->"`
			Ninth   string    `json:"<-"`
		}{"1", "2", "3", "4", "", "6", time.Time{}, "8", "9"}, map[string]interface{}{
			"first":  "1",
			"second": "2",
			"Third":  "3",
			"Forth":  "4",
			"eigth":  "8",
		}},
		{struct {
			First  string
			Second string
			Third  string `json:"Third"`
			Forth  string `json:"Forth,omitempty"`
			Fifth  string `json:"Fifth,omitempty"`
			Sixth  string `json:"-"`
		}{"1", "2", "3", "4", "", "6"}, map[string]interface{}{
			"first":  "1",
			"second": "2",
			"Third":  "3",
			"Forth":  "4",
		}},
		{struct {
			AnaymonusStruct
			Third string
		}{AnaymonusStruct{"1", "2"}, "3"}, map[string]interface{}{
			"first":  "1",
			"second": "2",
			"third":  "3",
		}},
		{struct {
			AnaymonusMap
			Third string
		}{AnaymonusMap{23: "1", 79: "2"}, "3"}, map[string]interface{}{
			"23":    "1",
			"79":    "2",
			"third": "3",
		}},
		{struct {
			First string
			AnaymonusStruct
			Third string
		}{"23", AnaymonusStruct{"1", "2"}, "3"}, map[string]interface{}{
			"first":  "23",
			"second": "2",
			"third":  "3",
		}},
		{struct {
			AnaymonusStruct
			First string
			Third string
		}{AnaymonusStruct{"1", "2"}, "23", "3"}, map[string]interface{}{
			"first":  "23",
			"second": "2",
			"third":  "3",
		}},
	}
	for singleTestKey, singleTestCase := range testCase {
		data := ArrayToMap(singleTestCase.origin, "json")
		assert.Equal(t, data, singleTestCase.target, singleTestKey)
	}
}

func TestArrayToMapTotal(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		{[]struct {
			First  string
			Second string
			Third  int    `json:"Third"`
			Forth  string `json:"Forth,omitempty"`
			Fifth  string `json:"Fifth,omitempty"`
			Sixth  string `json:"-"`
		}{
			{"1", "2", 3, "4", "", "6"},
			{"11", "22", 33, "44", "55", "66"},
		},
			[]interface{}{
				map[string]interface{}{
					"first":  "1",
					"second": "2",
					"Third":  3,
					"Forth":  "4",
				},
				map[string]interface{}{
					"first":  "11",
					"second": "22",
					"Third":  33,
					"Forth":  "44",
					"Fifth":  "55",
				},
			}},
		{
			struct {
				First  string
				Second interface{}
			}{
				"1",
				struct {
					Third string `json:"Third"`
				}{"dd"},
			},
			map[string]interface{}{
				"first": "1",
				"second": map[string]interface{}{
					"Third": "dd",
				},
			},
		},
	}
	for singleTestKey, singleTestCase := range testCase {
		data := ArrayToMap(singleTestCase.origin, "json")
		assert.Equal(t, data, singleTestCase.target, singleTestKey)
	}
}

func TestMapToArrayBasic(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		//basic type
		{true, true},
		{false, false},
		{"true", true},
		{"false", false},
		{-1, -1},
		{"-1", -1},
		{uint(1), int(1)},
		{float64(-1), int(-1)},
		{uint(1), uint(1)},
		{int(1), uint(1)},
		{float64(1), uint(1)},
		{"12345678123456789", uint64(12345678123456789)},
		{"12345678123456789", int64(12345678123456789)},
		{"-12345678123456789", int64(-12345678123456789)},
		{-1234567812345678.0, int64(-1234567812345678)},
		{1234567812345678.0, int64(1234567812345678)},
		{1234567812345678.0, uint64(1234567812345678)},
		{1.2, 1.2},
		{"1.2", 1.2},
		{"1", 1.0},
		{int(1), float64(1)},
		{uint(1), float64(1)},
		{true, "true"},
		{-1, "-1"},
		{uint(1), "1"},
		{1.2, "1.2"},
		{"abc", "abc"},
		//array type
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []string{"1", "2", "3"}},
		//map type
		{map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
		},
			map[int]int{
				1: 1,
				2: 2,
				3: 3,
				4: 4,
			}},
		{map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
		},
			map[int]string{
				1: "1",
				2: "2",
				3: "3",
				4: "4",
			}},
		{map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
		},
			map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
				"4": 4,
			}},
		{map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
		},
			map[string]string{
				"1": "1",
				"2": "2",
				"3": "3",
				"4": "4",
			}},
	}
	//普通测试
	for _, singleTestCase := range testCase {
		origin := singleTestCase.origin
		target := singleTestCase.target
		targetType := reflect.TypeOf(target)
		result := reflect.New(targetType)
		err := MapToArray(origin, result.Interface(), "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, result.Elem().Interface(), target)
	}
	//指针测试
	for _, singleTestCase := range testCase {
		origin := singleTestCase.origin
		target := singleTestCase.target
		targetType := reflect.PtrTo(reflect.TypeOf(target))
		result := reflect.New(targetType)
		err := MapToArray(origin, result.Interface(), "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, result.Elem().Elem().Interface(), target)
	}
	//interface测试
	for _, singleTestCase := range testCase {
		var result interface{}
		origin := singleTestCase.origin
		err := MapToArray(origin, &result, "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, result, origin)
	}
}

func TestMapToArrayStruct(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		{map[string]interface{}{
			"first":   "1",
			"second":  "2",
			"Third":   "3",
			"Forth":   "4",
			"fifth":   "5",
			"sixth":   "6",
			"seventh": "7",
		}, struct {
			First   string
			Second  int
			Third   string `json:"Third"`
			Forth   string `json:"Forth,omitempty"`
			Fifth   string `json:"-"`
			Sixth   string `json:"->"`
			Seventh string `json:"<-"`
		}{"1", 2, "3", "4", "", "", "7"}},
		{map[interface{}]interface{}{
			"first":  "1",
			"second": "2",
			"Third":  "3",
			"Forth":  "4",
			"fifth":  "5",
		}, struct {
			AnaymonusStruct
			Third string `json:"Third"`
			Forth string `json:"Forth,omitempty"`
			Fifth string `json:"-"`
			Sixth int
		}{AnaymonusStruct{"1", "2"}, "3", "4", "", 0}},
	}
	for _, singleTestCase := range testCase {
		origin := singleTestCase.origin
		target := singleTestCase.target
		targetType := reflect.TypeOf(target)
		result := reflect.New(targetType)
		err := MapToArray(origin, result.Interface(), "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, result.Elem().Interface(), target)
	}
}

type totalTempStruct struct {
	A string
	B int
}

func TestMapToArrayTotal(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
	}{
		{map[string]interface{}{
			"first":  "1",
			"second": "2",
			"Third":  "3",
			"Forth":  "4",
			"fifth":  "5",
			"sixth": []map[string]interface{}{
				{"a": "1", "b": "2"},
				{"a": "3", "b": "4"},
			},
		}, struct {
			First   string
			Second  int
			Third   string `json:"Third"`
			Forth   string `json:"Forth,omitempty"`
			Fifth   string `json:"-"`
			Sixth   []totalTempStruct
			Seventh []int
		}{
			"1",
			2,
			"3",
			"4",
			"",
			[]totalTempStruct{
				{"1", 2},
				{"3", 4},
			},
			nil,
		}},
	}
	for _, singleTestCase := range testCase {
		origin := singleTestCase.origin
		target := singleTestCase.target
		targetType := reflect.TypeOf(target)
		result := reflect.New(targetType)
		err := MapToArray(origin, result.Interface(), "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, result.Elem().Interface(), target)
	}
}

func TestMapToArrayHalfEmpty(t *testing.T) {
	testCase := []struct {
		Origin  interface{}
		Origin2 interface{}
		Target  interface{}
	}{
		//array
		{[]int{1, 2, 3, 4}, [0]int{}, [0]int{}},
		{[]int{1, 2, 3}, [4]int{3, 6, 7, 8}, [4]int{1, 2, 3, 0}},
		{[]int{1, 2, 3, 4}, [2]int{3, 6}, [2]int{1, 2}},
		//slice
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, []interface{}{3, 6, 7}, []interface{}{1, 2, 3, 4}},
		//map
		{map[int]string{1: "a", 2: "b"}, map[int]string{}, map[int]string{1: "a", 2: "b"}},
		{map[int]string{1: "a", 2: "b"}, map[int]string{3: "c"}, map[int]string{1: "a", 2: "b", 3: "c"}},
		{map[int]string{1: "a", 2: "b"}, map[int]interface{}{3: "c"}, map[int]interface{}{1: "a", 2: "b", 3: "c"}},
		{map[int]string{1: "a", 2: "b", 3: "yy"}, map[int]string{3: "c"}, map[int]string{1: "a", 2: "b", 3: "yy"}},
	}

	for _, singleTestCase := range testCase {
		origin := singleTestCase.Origin
		origin2 := reflect.ValueOf(&singleTestCase.Origin2)
		target := singleTestCase.Target
		err := MapToArray(origin, origin2.Interface(), "json")
		assert.Equal(t, err, nil)
		assert.Equal(t, origin2.Elem().Interface(), target)
	}
}

func TestMapToArrayError(t *testing.T) {
	testCase := []struct {
		origin interface{}
		target interface{}
		err    string
	}{
		{"zz", time.Now(), "不是时间，其值为[zz]"},
		{"1c", 1, "不是整数，其值为[1c]"},
		{"1.2c", 1.2, "不是浮点数，其值为[1.2c]"},
		{"1.2d", decimal.Decimal("1.2d"), "不是十进制数字，其值为[1.2d]"},
		{map[string]interface{}{
			"first": "1m",
		}, struct {
			First int
		}{1},
			"参数first不是整数，其值为[1m]"},
	}
	for _, singleTestCase := range testCase {
		origin := singleTestCase.origin
		target := singleTestCase.target
		targetType := reflect.TypeOf(target)
		result := reflect.New(targetType)
		err := MapToArray(origin, result.Interface(), "json")
		assert.Equal(t, err != nil, true)
		assert.Equal(t, err.Error(), singleTestCase.err)
	}
}
