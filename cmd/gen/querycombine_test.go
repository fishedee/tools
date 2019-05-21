package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQueryCombine(t *testing.T) {
	admin := []Admin{
		Admin{AdminID: 23, Level: 5},
		Admin{AdminID: 23, Level: 5},
		Admin{AdminID: 23, Level: 5},
		Admin{AdminID: 3, Level: 30},
		Admin{AdminID: 4, Level: 7},
	}
	user := []User{
		User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserID: 3, Name: "a"},
		User{UserID: 4, Name: "j"},
	}
	assert.Equal(t, query.Combine(admin, user, func(admin Admin, user User) AdminUser {
		return AdminUser{
			AdminID:    admin.AdminID,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []AdminUser{
		AdminUser{AdminID: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		AdminUser{AdminID: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		AdminUser{AdminID: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		AdminUser{AdminID: 3, Level: 30, Name: "a"},
		AdminUser{AdminID: 4, Level: 7, Name: "j"},
	})
	assert.Equal(t, query.Combine([]int{1, 2, 3, 4, 5}, user, func(admin int, user User) User {
		return User{
			UserID:     admin,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []User{
		User{UserID: 1, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 2, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserID: 3, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserID: 4, Name: "a"},
		User{UserID: 5, Name: "j"},
	})

	// 测试
	testCase := testdata.GetQueryCombineTestCase()

	for _, singleTestCase := range testCase {
		result := query.Combine(singleTestCase.Origin, singleTestCase.Origin2, singleTestCase.Functor)
		assert.Equal(t, result, singleTestCase.Target)
	}
}

func initQueryCombineData() ([]User, []Admin) {
	user := make([]User, 1000, 1000)
	admin := make([]Admin, 1000, 1000)
	for i := range user {
		user[i].UserID = rand.Int()
		user[i].Age = rand.Int()
		admin[i].AdminID = rand.Int()
	}
	return user, admin
}

func BenchmarkQueryCombineHand(b *testing.B) {
	users, admins := initQueryCombineData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		result := make([]AdminUser, len(users), len(users))
		for i := range users {
			result[i] = AdminUser{
				AdminID:    admins[i].AdminID,
				Level:      admins[i].Level,
				Name:       users[i].Name,
				CreateTime: users[i].CreateTime,
			}
		}
	}
}

func BenchmarkQueryCombineMacro(b *testing.B) {
	user, admin := initQueryCombineData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Combine(admin, user, func(admin Admin, user User) AdminUser {
			return AdminUser{
				AdminID:    admin.AdminID,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}

func BenchmarkQueryCombineReflect(b *testing.B) {
	user, admin := initQueryCombineData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Combine(user, admin, func(user User, admin Admin) AdminUser {
			return AdminUser{
				AdminID:    admin.AdminID,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}
