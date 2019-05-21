package testdata

import (
	"time"

	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

// Department Department
type Department struct {
	DepartmentID int
	Name         string
	Employees    []User
}

// User User
type User struct {
	UserID     int
	Age        int
	Name       string
	CreateTime time.Time
}

// Admin Admin
type Admin struct {
	AdminID int
	Level   int
}

// AdminUser AdminUser
type AdminUser struct {
	AdminID    int
	Level      int
	Age        int
	Name       string
	CreateTime time.Time
}

// Sex Sex
type Sex struct {
	IsMale bool
}

// ContentType 测试类型 支持bool,int,float,string和time.Time
type ContentType = testdata.ContentType

func logic() {
	query.Column([]User{}, "UserID")
	query.Column([]User{}, "Age")
	query.Column([]ContentType{}, " Name ")
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
	query.Sort([]User{}, "UserID desc,Name asc,CreateTime asc")
	query.Sort([]User{}, "UserID asc")
	query.Sort([]int{}, ". desc")
	query.ColumnMap([]User{}, "UserID")
	query.ColumnMap([]User{}, "Age")
	query.ColumnMap([]int{}, ".")
	query.Group([]User{}, "UserID", func(user []User) Department {
		return Department{}
	})
	query.Group([]User{}, "CreateTime", func(user []User) Department {
		return Department{}
	})
	query.Group([]int{}, ".", func(ids []int) Department {
		users := query.Select(ids, func(id int) User {
			return User{UserID: id}
		}).([]User)
		return Department{Employees: users}
	})
	query.LeftJoin([]Admin{}, []User{}, "AdminID = UserID", func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.RightJoin([]User{}, []int{}, "UserID = .", func(left User, right int) User {
		return User{}
	})
	query.Join([]Admin{}, []User{}, "inner", "AdminID = UserID", func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.Combine([]Admin{}, []User{}, func(left Admin, right User) AdminUser {
		return AdminUser{}
	})
	query.Combine([]int{}, []User{}, func(left int, right User) User {
		return User{}
	})
}
