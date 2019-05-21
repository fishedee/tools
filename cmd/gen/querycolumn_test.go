package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQueryColumn(t *testing.T) {
	data := []User{
		User{UserID: 1},
		User{UserID: -2},
		User{UserID: 3},
	}
	assert.Equal(t, query.Column(data, "UserID"), []int{1, -2, 3})
	assert.Equal(t, query.Column(data, "."), data)
	assert.Equal(t, query.Column([]int{1, -2, 3}, "."), []int{1, -2, 3})

	// 测试
	testCase := testdata.GetQueryColumnTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Column(singleTestCase.Data, singleTestCase.Column)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func BenchmarkQueryColumnHand(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]int, len(data), len(data))
		for i, single := range data {
			newData[i] = single.UserID
		}
	}
}

func BenchmarkQueryColumnMacro(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column(data, "UserID")
	}
}

func BenchmarkQueryColumnReflect(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column(data, "Age")
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
