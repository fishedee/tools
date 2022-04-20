package main

import (
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

func TestQuerySelect(t *testing.T) {
	data := []User{
		{Name: "Man_a"},
		{Name: "Woman_b"},
		{Name: "Man_c"},
	}

	assert.Equal(t, query.Select(data, func(a User) Sex {
		if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
			return Sex{IsMale: true}
		}
		return Sex{IsMale: false}
	}), []Sex{
		{IsMale: true},
		{IsMale: false},
		{IsMale: true},
	})
	assert.Equal(t, query.Select([]int{3, 4, 5, -1}, func(a int) User {
		return User{UserID: a}
	}), []User{
		{UserID: 3},
		{UserID: 4},
		{UserID: 5},
		{UserID: -1},
	})

	// 测试
	testCase := testdata.GetQuerySelectTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func BenchmarkQuerySelectHand(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]Sex, len(data), len(data))
		for i, single := range data {
			newData[i] = func(a User) Sex {
				if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
					return Sex{IsMale: true}
				}
				return Sex{IsMale: false}
			}(single)
		}
	}
}

func BenchmarkQuerySelectMacro(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Select(data, func(a User) Sex {
			if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
				return Sex{IsMale: true}
			}
			return Sex{IsMale: false}
		})
	}
}

func BenchmarkQuerySelectReflect(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Select(data, func(a User) bool {
			if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
				return true
			}
			return false
		})
	}
}
