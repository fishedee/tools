package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQuerySort(t *testing.T) {
	data := []User{
		{UserID: 3, Name: "a"},
		{UserID: 3, Name: "c"},
		{UserID: 23, Name: "d"},
		{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		{UserID: 23, Name: "a"},
		{UserID: 1},
		{UserID: 1},
	}
	assert.Equal(t, query.Sort(data, "UserID desc,Name asc,CreateTime asc"), []User{
		{UserID: 23, Name: "a"},
		{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		{UserID: 23, Name: "d"},
		{UserID: 3, Name: "a"},
		{UserID: 3, Name: "c"},
		{UserID: 1},
		{UserID: 1},
	})
	assert.Equal(t, query.Sort([]int{3, 2, 1, 7, -8}, ". desc"), []int{7, 3, 2, 1, -8})

	// 测试
	testCase := testdata.GetQuerySortTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

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
