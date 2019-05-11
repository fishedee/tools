package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/donnol/tools/assert"
	"github.com/donnol/tools/query"
)

func TestQueryCombine(t *testing.T) {
	admin := []Admin{
		Admin{AdminId: 23, Level: 5},
		Admin{AdminId: 23, Level: 5},
		Admin{AdminId: 23, Level: 5},
		Admin{AdminId: 3, Level: 30},
		Admin{AdminId: 4, Level: 7},
	}
	user := []User{
		User{UserId: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserId: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserId: 3, Name: "a"},
		User{UserId: 4, Name: "j"},
	}
	assert.Equal(t, query.Combine(admin, user, func(admin Admin, user User) AdminUser {
		return AdminUser{
			AdminId:    admin.AdminId,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []AdminUser{
		AdminUser{AdminId: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		AdminUser{AdminId: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		AdminUser{AdminId: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		AdminUser{AdminId: 3, Level: 30, Name: "a"},
		AdminUser{AdminId: 4, Level: 7, Name: "j"},
	})
	assert.Equal(t, query.Combine([]int{1, 2, 3, 4, 5}, user, func(admin int, user User) User {
		return User{
			UserId:     admin,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []User{
		User{UserId: 1, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 2, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserId: 3, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserId: 4, Name: "a"},
		User{UserId: 5, Name: "j"},
	})
}

func initQueryCombineData() ([]User, []Admin) {
	user := make([]User, 1000, 1000)
	admin := make([]Admin, 1000, 1000)
	for i := range user {
		user[i].UserId = rand.Int()
		user[i].Age = rand.Int()
		admin[i].AdminId = rand.Int()
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
				AdminId:    admins[i].AdminId,
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
				AdminId:    admin.AdminId,
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
				AdminId:    admin.AdminId,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}
