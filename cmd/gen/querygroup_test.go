package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	gentestdata "github.com/fishedee/tools/cmd/gen/testdata"
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

func TestQueryGroup(t *testing.T) {
	data := []gentestdata.User{
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
	assert.Equal(t, query.Group[gentestdata.User, gentestdata.Department, []gentestdata.Department](data, "UserID", func(users []gentestdata.User) gentestdata.Department {
		return gentestdata.Department{
			Employees: users,
		}
	}), []gentestdata.Department{
		{Employees: []gentestdata.User{
			{UserID: 3, Name: "a"},
			{UserID: 3, Name: "c"},
		}},
		{Employees: []gentestdata.User{
			{UserID: 23, Name: "d"},
			{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
			{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
			{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
			{UserID: 23, Name: "a"},
		}},
		{Employees: []gentestdata.User{
			{UserID: 1},
			{UserID: 1},
		}},
	})
	assert.Equal(t, query.Group[int, gentestdata.Department, []gentestdata.Department]([]int{1, 3, 4, 4, 3, 3}, ".", func(ids []int) gentestdata.Department {
		users := query.Select(ids, func(id int) gentestdata.User {
			return gentestdata.User{UserID: id}
		})
		return gentestdata.Department{Employees: users}
	}), []gentestdata.Department{
		{Employees: []gentestdata.User{
			{UserID: 1},
		}},
		{Employees: []gentestdata.User{
			{UserID: 3},
			{UserID: 3},
			{UserID: 3},
		}},
		{Employees: []gentestdata.User{
			{UserID: 4},
			{UserID: 4},
		}},
	})

	// 测试
	testCase := testdata.GetQueryGroupTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func initQueryGroupData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserID = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQueryGroupHand(b *testing.B) {
	data := initQueryGroupData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		findMap := make(map[int][]User, len(data))
		result := make([]Department, 0, len(data))
		for _, single := range data {
			users, isExist := findMap[single.UserID]
			if isExist == false {
				users = []User{}
			}
			users = append(users, single)
			findMap[single.UserID] = users
		}
		for _, single := range data {
			users, isExist := findMap[single.UserID]
			if isExist {
				continue
			}
			delete(findMap, single.UserID)
			result = append(result, Department{
				Employees: users,
			})
		}
	}
}

func BenchmarkQueryGroupMacro(b *testing.B) {
	data := initQueryGroupData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Group[gentestdata.User, Department, []Department](data, "UserID", func(users []User) Department {
			return Department{
				Employees: users,
			}
		})
	}
}

func BenchmarkQueryGroupReflect(b *testing.B) {
	data := initQueryGroupData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Group[gentestdata.User, Department, []Department](data, "Age", func(users []User) Department {
			return Department{
				Employees: users,
			}
		})
	}
}
