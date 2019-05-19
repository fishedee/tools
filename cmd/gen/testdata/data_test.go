package testdata

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/fishedee/tools/query"
)

func BenchmarkQueryColumnHand(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]int, len(data), len(data))
		for i, single := range data {
			newData[i] = single.UserID
		}
	}
}

func BenchmarkQueryColumnMacro(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column(data, "UserID")
	}
}

func BenchmarkQueryColumnHandMany(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]int, len(data), len(data))
		newDataAge := make([]int, len(data), len(data))
		for i, single := range data {
			newData[i] = single.UserID
			newDataAge[i] = single.Age
		}
	}
}

func BenchmarkQueryColumnMacroMany(b *testing.B) {
	data := make([]User, 1000, 1000)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Column(data, "UserID")
		query.Column(data, "Age")
	}
}

func initQueryColumnMapData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserID = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQueryColumnMapHand(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make(map[int]User, len(data))
		for _, single := range data {
			newData[single.UserID] = single
		}
	}
}

func BenchmarkQueryColumnMapMacro(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap(data, "UserID")
	}
}

func BenchmarkQueryColumnMapHandMany(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make(map[int]User, len(data))
		newDataAge := make(map[int]User, len(data))
		for _, single := range data {
			newData[single.UserID] = single
			newDataAge[single.Age] = single
		}
	}
}

func BenchmarkQueryColumnMapMacroMany(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap(data, "UserID")
		query.ColumnMap(data, "Age")
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

func initQuerySortData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserID = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQuerySortHand(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]User, len(data), len(data))
		copy(newData, data)
		sort.SliceStable(newData, func(i int, j int) bool {
			return newData[i].UserID < newData[j].UserID
		})
	}
}

func BenchmarkQuerySortMacro(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Sort(data, "UserID asc")
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
