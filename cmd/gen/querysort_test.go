package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
)

func TestQuerySort(t *testing.T) {
	data := []User{
		User{UserID: 3, Name: "a"},
		User{UserID: 3, Name: "c"},
		User{UserID: 23, Name: "d"},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		User{UserID: 23, Name: "a"},
		User{UserID: 1},
		User{UserID: 1},
	}
	assert.Equal(t, query.Sort(data, "UserID desc,Name asc,CreateTime asc"), []User{
		User{UserID: 23, Name: "a"},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		User{UserID: 23, Name: "d"},
		User{UserID: 3, Name: "a"},
		User{UserID: 3, Name: "c"},
		User{UserID: 1},
		User{UserID: 1},
	})
	assert.Equal(t, query.Sort([]int{3, 2, 1, 7, -8}, ". desc"), []int{7, 3, 2, 1, -8})

	// 测试
	testCase := GetQuerySortTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Sort(singleTestCase.Origin, singleTestCase.SortName)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func initQuerySortData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserID = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQuerySortHand(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]User, len(data), len(data))
		copy(newData, data)
		sort.SliceStable(newData, func(i int, j int) bool {
			return newData[i].UserID < newData[j].UserID
		})
	}
}

func BenchmarkQuerySortMacro(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Sort(data, "UserID asc")
	}
}

func BenchmarkQuerySortReflect(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Sort(data, "Age asc")
	}
}
