package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/fishedee/tools/assert"
	gentestdata "github.com/fishedee/tools/cmd/gen/testdata"
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

func TestQueryColumn(t *testing.T) {
	data := []gentestdata.User{
		{UserID: 1},
		{UserID: -2},
		{UserID: 3},
	}
	assert.Equal(t, query.Column[gentestdata.User, int](data, "UserID"), []int{1, -2, 3})
	assert.Equal(t, query.Column[gentestdata.User, gentestdata.User](data, "."), data)
	assert.Equal(t, query.Column[int, int]([]int{1, -2, 3}, "."), []int{1, -2, 3})

	// 测试
	testCase := testdata.GetQueryColumnTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func BenchmarkQueryColumnHand(b *testing.B) {
	data := make([]gentestdata.User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]int, len(data), len(data))
		for i, single := range data {
			newData[i] = single.UserID
		}
	}
}

func BenchmarkQueryColumnMacro(b *testing.B) {
	data := make([]gentestdata.User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column[gentestdata.User, int](data, "UserID")
	}
}

func BenchmarkQueryColumnReflect(b *testing.B) {
	data := make([]gentestdata.User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column[gentestdata.User, int](data, "Age")
	}
}

func init() {
	args := os.Args
	isWarning := true
	for _, arg := range args {
		if arg == "-test.benchmem=true" {
			isWarning = false
			break
		}
	}
	fmt.Println("QueryReflectWarning:", isWarning)
	query.ReflectWarning(isWarning)
}
