package testdata

import (
	"time"

	"github.com/fishedee/tools/query"
)

// Department Department
type Department struct {
	DepartmentId int
	Name         string
	Employees    []User
}

// User User
type User struct {
	UserId     int
	Age        int
	Name       string
	CreateTime time.Time
}

// Admin Admin
type Admin struct {
	AdminId int
	Level   int
}

// AdminUser AdminUser
type AdminUser struct {
	AdminId    int
	Level      int
	Age        int
	Name       string
	CreateTime time.Time
}

// Sex Sex
type Sex struct {
	IsMale bool
}

func logic() {
	query.Column([]User{}, "UserId")
	query.Column([]User{}, ".")
	query.Column([]int{}, ".")
	query.Select([]User{}, func(d User) Sex {
		return Sex{}
	})
	query.Where([]int{}, func(c int) bool {
		return c%2 == 0
	})
	query.Where([]User{}, func(c User) bool {
		return true
	})
	query.Sort([]User{}, "UserId desc,Name asc,CreateTime asc")
	query.Sort([]User{}, "UserId asc")
	query.Sort([]int{}, ". desc")
	query.ColumnMap([]User{}, "UserId")
	query.ColumnMap([]int{}, ".")
	query.Group([]User{}, "UserId", func(user []User) Department {
		return Department{}
	})
	query.Group([]User{}, "CreateTime", func(user []User) Department {
		return Department{}
	})
	query.Group([]int{}, ".", func(ids []int) Department {
		users := query.Select(ids, func(id int) User {
			return User{UserId: id}
		}).([]User)
		return Department{Employees: users}
	})
	query.LeftJoin([]Admin{}, []User{}, "AdminId = UserId", func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.RightJoin([]User{}, []int{}, "UserId = .", func(left User, right int) User {
		return User{}
	})
	query.Join([]Admin{}, []User{}, "inner", "AdminId = UserId", func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.Combine([]Admin{}, []User{}, func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.Combine([]int{}, []User{}, func(left int, right User) User {
		return User{}
	})
}
