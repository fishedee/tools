package main

import (
	"math/rand"
	"testing"

	"github.com/fishedee/tools/assert"
	gentestdata "github.com/fishedee/tools/cmd/gen/testdata"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQueryColumnMap(t *testing.T) {
	data := []gentestdata.User{
		gentestdata.User{UserID: 1},
		gentestdata.User{UserID: -2},
		gentestdata.User{UserID: 3},
	}
	assert.Equal(t, query.ColumnMap(data, "UserID"), map[int]gentestdata.User{
		1:  gentestdata.User{UserID: 1},
		-2: gentestdata.User{UserID: -2},
		3:  gentestdata.User{UserID: 3},
	})
	assert.Equal(t, query.ColumnMap([]int{5, 6, 8, 8, 0, 6}, "."), map[int]int{
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
		query.ColumnMap(data, "UserID")
	}
}

func BenchmarkQueryColumnMapReflect(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap(data, "Age")
	}
}
