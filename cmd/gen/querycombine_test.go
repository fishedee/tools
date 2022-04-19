package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	gentestdata "github.com/fishedee/tools/cmd/gen/testdata"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQueryCombine(t *testing.T) {
	admin := []gentestdata.Admin{
		{AdminID: 23, Level: 5},
		{AdminID: 23, Level: 5},
		{AdminID: 23, Level: 5},
		{AdminID: 3, Level: 30},
		{AdminID: 4, Level: 7},
	}
	user := []gentestdata.User{
		{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		{UserID: 3, Name: "a"},
		{UserID: 4, Name: "j"},
	}
	assert.Equal(t, query.Combine(admin, user, func(admin gentestdata.Admin, user gentestdata.User) gentestdata.AdminUser {
		return gentestdata.AdminUser{
			AdminID:    admin.AdminID,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []gentestdata.AdminUser{
		{AdminID: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		{AdminID: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		{AdminID: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		{AdminID: 3, Level: 30, Name: "a"},
		{AdminID: 4, Level: 7, Name: "j"},
	})
	assert.Equal(t, query.Combine([]int{1, 2, 3, 4, 5}, user, func(admin int, user gentestdata.User) gentestdata.User {
		return gentestdata.User{
			UserID:     admin,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []gentestdata.User{
		{UserID: 1, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 2, Name: "g", CreateTime: time.Unix(1, 0)},
		{UserID: 3, Name: "h", CreateTime: time.Unix(33, 0)},
		{UserID: 4, Name: "a"},
		{UserID: 5, Name: "j"},
	})

	// 测试
	testCase := testdata.GetQueryCombineTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {
		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)
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
