package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

func TestQueryJoin(t *testing.T) {
	admin := []Admin{
		{AdminID: 23, Level: 5},
		{AdminID: 3, Level: 30},
		{AdminID: 4, Level: 7},
	}
	user := []User{
		{UserID: 3, Name: "a"},
		{UserID: 3, Name: "c"},
		{UserID: 23, Name: "d"},
		{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		{UserID: 23, Name: "a"},
		{UserID: 1},
		{UserID: 1},
	}
	assert.Equal(t, query.LeftJoin(admin, user, "AdminID = UserID", func(admin Admin, user User) AdminUser {
		return AdminUser{
			AdminID:    admin.AdminID,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []AdminUser{
		{AdminID: 23, Level: 5, Name: "d"},
		{AdminID: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		{AdminID: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		{AdminID: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		{AdminID: 23, Level: 5, Name: "a"},
		{AdminID: 3, Level: 30, Name: "a"},
		{AdminID: 3, Level: 30, Name: "c"},
		{AdminID: 4, Level: 7},
	})
	assert.Equal(t, query.RightJoin(user, []int{23, 3, 4, 6, 7}, "UserID = .", func(left User, right int) User {
		return User{
			UserID:     right,
			Name:       left.Name,
			CreateTime: left.CreateTime,
		}
	}), []User{
		{UserID: 3, Name: "a"},
		{UserID: 3, Name: "c"},
		{UserID: 23, Name: "d"},
		{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		{UserID: 23, Name: "a"},
		{UserID: 4},
		{UserID: 6},
		{UserID: 7},
	})

	// 测试
	testCase := testdata.GetQueryJoinTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)

	}
}

func initQueryJoinData() ([]User, []Admin) {
	user := make([]User, 1000, 1000)
	admin := make([]Admin, 1000, 1000)
	for i := range user {
		user[i].UserID = rand.Int()
		user[i].Age = rand.Int()
		admin[i].AdminID = rand.Int()
	}
	return user, admin
}

func BenchmarkQueryJoinHand(b *testing.B) {
	users, admins := initQueryJoinData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		adminMap := map[int][]Admin{}
		for _, admin := range admins {
			findAdmins, isExist := adminMap[admin.AdminID]
			if isExist == false {
				findAdmins = []Admin{}
			}
			findAdmins = append(findAdmins, admin)
			adminMap[admin.AdminID] = findAdmins
		}
		result := make([]AdminUser, 0, len(users))
		for _, user := range users {
			admins := adminMap[user.UserID]
			for _, admin := range admins {
				result = append(result, AdminUser{
					AdminID:    admin.AdminID,
					Level:      admin.Level,
					Name:       user.Name,
					CreateTime: user.CreateTime,
				})
			}
		}
	}
}

func BenchmarkQueryJoinMacro(b *testing.B) {
	user, admin := initQueryJoinData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.LeftJoin(admin, user, "AdminID = UserID", func(admin Admin, user User) AdminUser {
			return AdminUser{
				AdminID:    admin.AdminID,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}

func BenchmarkQueryJoinReflect(b *testing.B) {
	user, admin := initQueryJoinData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.LeftJoin(admin, user, "AdminID = Age", func(admin Admin, user User) AdminUser {
			return AdminUser{
				AdminID:    admin.AdminID,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}
