package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
)

func TestQueryGroup(t *testing.T) {
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
	assert.Equal(t, query.Group(data, "UserID", func(users []User) Department {
		return Department{
			Employees: users,
		}
	}), []Department{
		Department{Employees: []User{
			User{UserID: 3, Name: "a"},
			User{UserID: 3, Name: "c"},
		}},
		Department{Employees: []User{
			User{UserID: 23, Name: "d"},
			User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
			User{UserID: 23, Name: "c", CreateTime: time.Unix(1, 0)},
			User{UserID: 23, Name: "c", CreateTime: time.Unix(33, 0)},
			User{UserID: 23, Name: "a"},
		}},
		Department{Employees: []User{
			User{UserID: 1},
			User{UserID: 1},
		}},
	})
	assert.Equal(t, query.Group([]int{1, 3, 4, 4, 3, 3}, ".", func(ids []int) Department {
		users := query.Select(ids, func(id int) User {
			return User{UserID: id}
		}).([]User)
		return Department{Employees: users}
	}), []Department{
		Department{Employees: []User{
			User{UserID: 1},
		}},
		Department{Employees: []User{
			User{UserID: 3},
			User{UserID: 3},
			User{UserID: 3},
		}},
		Department{Employees: []User{
			User{UserID: 4},
			User{UserID: 4},
		}},
	})
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
		query.Group(data, "UserID", func(users []User) Department {
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
		query.Group(data, "Age", func(users []User) Department {
			return Department{
				Employees: users,
			}
		})
	}
}
