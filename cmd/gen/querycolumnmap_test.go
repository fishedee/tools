package main

import (
	"math/rand"
	"testing"

	"github.com/fishedee/tools/assert"
	gentestdata "github.com/fishedee/tools/cmd/gen/testdata"
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

func TestQueryColumnMap(t *testing.T) {
	data := []gentestdata.User{
		{UserID: 1},
		{UserID: -2},
		{UserID: 3},
	}
	assert.Equal(t, query.ColumnMap[gentestdata.User, int](data, "UserID"), map[int]gentestdata.User{
		1:  {UserID: 1},
		-2: {UserID: -2},
		3:  {UserID: 3},
	})
	assert.Equal(t, query.ColumnMap[int, int]([]int{5, 6, 8, 8, 0, 6}, "."), map[int]int{
		5: 5,
		6: 6,
		8: 8,
		0: 0,
	})

	// 测试
	testCase := testdata.GetQueryColumnMapTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func initQueryColumnMapData() []gentestdata.User {
	data := make([]gentestdata.User, 1000, 1000)
	for i := range data {
		data[i].UserID = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQueryColumnMapHand(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make(map[int]gentestdata.User, len(data))
		for _, single := range data {
			newData[single.UserID] = single
		}
	}
}

func BenchmarkQueryColumnMapMacro(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap[gentestdata.User, int](data, "UserID")
	}
}

func BenchmarkQueryColumnMapReflect(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap[gentestdata.User, int](data, "Age")
	}
}

func BenchmarkQueryColumnMapSliceHand(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make(map[int][]gentestdata.User, len(data))
		for _, single := range data {
			temp := newData[single.UserID]
			temp = append(temp, single)
			newData[single.UserID] = temp
		}
	}
}

func BenchmarkQueryColumnMapSliceMacro(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap[gentestdata.User, int](data, "[]UserID")
	}
}

func BenchmarkQueryColumnMapSliceReflect(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap[gentestdata.User, int](data, "[]Age")
	}
}
