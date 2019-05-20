package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
)

func TestQueryJoin(t *testing.T) {
	admin := []Admin{
		Admin{AdminID: 23, Level: 5},
		Admin{AdminID: 3, Level: 30},
		Admin{AdminID: 4, Level: 7},
	}
	user := []User{
		User{UserID: 3, Name: "a"},
		User{UserID: 3, Name: "c"},
		User{UserID: 23, Name: "d"},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserID: 23, Name: "a"},
		User{UserID: 1},
		User{UserID: 1},
	}
	assert.Equal(t, query.LeftJoin(admin, user, "AdminID = UserID", func(admin Admin, user User) AdminUser {
		return AdminUser{
			AdminID:    admin.AdminID,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []AdminUser{
		AdminUser{AdminID: 23, Level: 5, Name: "d"},
		AdminUser{AdminID: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		AdminUser{AdminID: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		AdminUser{AdminID: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		AdminUser{AdminID: 23, Level: 5, Name: "a"},
		AdminUser{AdminID: 3, Level: 30, Name: "a"},
		AdminUser{AdminID: 3, Level: 30, Name: "c"},
		AdminUser{AdminID: 4, Level: 7},
	})
	assert.Equal(t, query.RightJoin(user, []int{23, 3, 4, 6, 7}, "UserID = .", func(left User, right int) User {
		return User{
			UserID:     right,
			Name:       left.Name,
			CreateTime: left.CreateTime,
		}
	}), []User{
		User{UserID: 3, Name: "a"},
		User{UserID: 3, Name: "c"},
		User{UserID: 23, Name: "d"},
		User{UserID: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserID: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserID: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserID: 23, Name: "a"},
		User{UserID: 4},
		User{UserID: 6},
		User{UserID: 7},
	})

	// 测试
	testCase := GetQueryJoinTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Join(singleTestCase.LeftData, singleTestCase.RightData, singleTestCase.JoinPlace, singleTestCase.JoinType, singleTestCase.JoinFuctor)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

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
