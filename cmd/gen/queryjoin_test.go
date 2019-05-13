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
		Admin{AdminId: 23, Level: 5},
		Admin{AdminId: 3, Level: 30},
		Admin{AdminId: 4, Level: 7},
	}
	user := []User{
		User{UserId: 3, Name: "a"},
		User{UserId: 3, Name: "c"},
		User{UserId: 23, Name: "d"},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserId: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserId: 23, Name: "a"},
		User{UserId: 1},
		User{UserId: 1},
	}
	assert.Equal(t, query.LeftJoin(admin, user, "AdminId = UserId", func(admin Admin, user User) AdminUser {
		return AdminUser{
			AdminId:    admin.AdminId,
			Level:      admin.Level,
			Name:       user.Name,
			CreateTime: user.CreateTime,
		}
	}), []AdminUser{
		AdminUser{AdminId: 23, Level: 5, Name: "d"},
		AdminUser{AdminId: 23, Level: 5, Name: "c", CreateTime: time.Unix(29, 0)},
		AdminUser{AdminId: 23, Level: 5, Name: "g", CreateTime: time.Unix(1, 0)},
		AdminUser{AdminId: 23, Level: 5, Name: "h", CreateTime: time.Unix(33, 0)},
		AdminUser{AdminId: 23, Level: 5, Name: "a"},
		AdminUser{AdminId: 3, Level: 30, Name: "a"},
		AdminUser{AdminId: 3, Level: 30, Name: "c"},
		AdminUser{AdminId: 4, Level: 7},
	})
	assert.Equal(t, query.RightJoin(user, []int{23, 3, 4, 6, 7}, "UserId = .", func(left User, right int) User {
		return User{
			UserId:     right,
			Name:       left.Name,
			CreateTime: left.CreateTime,
		}
	}), []User{
		User{UserId: 3, Name: "a"},
		User{UserId: 3, Name: "c"},
		User{UserId: 23, Name: "d"},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 23, Name: "g", CreateTime: time.Unix(1, 0)},
		User{UserId: 23, Name: "h", CreateTime: time.Unix(33, 0)},
		User{UserId: 23, Name: "a"},
		User{UserId: 4},
		User{UserId: 6},
		User{UserId: 7},
	})
}

func initQueryJoinData() ([]User, []Admin) {
	user := make([]User, 1000, 1000)
	admin := make([]Admin, 1000, 1000)
	for i := range user {
		user[i].UserId = rand.Int()
		user[i].Age = rand.Int()
		admin[i].AdminId = rand.Int()
	}
	return user, admin
}

func BenchmarkQueryJoinHand(b *testing.B) {
	users, admins := initQueryJoinData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		adminMap := map[int][]Admin{}
		for _, admin := range admins {
			findAdmins, isExist := adminMap[admin.AdminId]
			if isExist == false {
				findAdmins = []Admin{}
			}
			findAdmins = append(findAdmins, admin)
			adminMap[admin.AdminId] = findAdmins
		}
		result := make([]AdminUser, 0, len(users))
		for _, user := range users {
			admins := adminMap[user.UserId]
			for _, admin := range admins {
				result = append(result, AdminUser{
					AdminId:    admin.AdminId,
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
		query.LeftJoin(admin, user, "AdminId = UserId", func(admin Admin, user User) AdminUser {
			return AdminUser{
				AdminId:    admin.AdminId,
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
		query.LeftJoin(admin, user, "AdminId = Age", func(admin Admin, user User) AdminUser {
			return AdminUser{
				AdminId:    admin.AdminId,
				Level:      admin.Level,
				Name:       user.Name,
				CreateTime: user.CreateTime,
			}
		})
	}
}
