package main

import (
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQueryWhere(t *testing.T) {
	data := []User{
		{Name: "Man_a"},
		{Name: "Woman_b"},
		{Name: "Man_c"},
	}

	assert.Equal(t, query.Where(data, func(a User) bool {
		if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
			return true
		}
		return false
	}), []User{
		{Name: "Man_a"},
		{Name: "Man_c"},
	})
	assert.Equal(t, query.Where([]int{3, 2, 3, 5, 9, 4}, func(c int) bool {
		return c%2 == 0
	}), []int{2, 4})

	// 测试
	testCase := testdata.GetQueryWhereTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func BenchmarkQueryWhereHand(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]User, 0, len(data))
		for _, single := range data {
			isMan := func(a User) bool {
				if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
					return true
				}
				return false
			}(single)

			if isMan {
				newData = append(newData, single)
			}
		}
	}
}

func BenchmarkQueryWhereMacro(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Where(data, func(a User) bool {
			if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
				return true
			}
			return false
		})
	}
}

func BenchmarkQueryWhereReflect(b *testing.B) {
	data := make([]Sex, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Where(data, func(a Sex) bool {
			if a.IsMale == true {
				return true
			}
			return false
		})
	}
}
