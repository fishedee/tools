package testdata

import (
	"time"

	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
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
	IsMale  bool
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

// QueryInnerStruct QueryInnerStruct
type QueryInnerStruct = testdata.QueryInnerStruct

// QueryInnerStruct2 QueryInnerStruct2
type QueryInnerStruct2 = testdata.QueryInnerStruct2

func logic() {
	query.Column[User, int]([]User{}, "UserID")
	query.Column[User, string]([]User{}, "Name")
	query.Column[ContentType, string]([]ContentType{}, " Name ")
	query.Column[ContentType, string]([]ContentType{}, "     Name         ")
	query.Column[ContentType, int]([]ContentType{}, "Age        ")
	query.Column[ContentType, bool]([]ContentType{}, "Ok        ")
	query.Column[ContentType, float32]([]ContentType{}, "    Money  ")
	query.Column[ContentType, float64]([]ContentType{}, "    CardMoney")
	query.Column[QueryInnerStruct2, int]([]QueryInnerStruct2{}, "QueryInnerStruct.MM")
	query.Column[User, User]([]User{}, ".")
	query.Column[int, int]([]int{}, ".")
	query.Column[int, int]([]int{}, " . ")
	query.Column[string, string]([]string{}, " . ")
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
	query.Sort([]Admin{}, "IsMale asc")
	query.ColumnMap[User, int]([]User{}, "UserID")
	query.ColumnMap[User, string]([]User{}, "Name")
	query.ColumnMap[User, int]([]User{}, "[]UserID")
	query.ColumnMap[int, int]([]int{}, ".")
	query.ColumnMap[int, int]([]int{}, " . ")
	query.ColumnMap[string, string]([]string{}, " . ")
	query.ColumnMap[ContentType, string]([]ContentType{}, " Name ")
	query.ColumnMap[ContentType, string]([]ContentType{}, "     Name         ")
	query.ColumnMap[ContentType, int]([]ContentType{}, "Age        ")
	query.ColumnMap[ContentType, bool]([]ContentType{}, "Ok        ")
	query.ColumnMap[ContentType, float32]([]ContentType{}, "    Money  ")
	query.ColumnMap[ContentType, float64]([]ContentType{}, "    CardMoney")
	query.ColumnMap[QueryInnerStruct2, int]([]QueryInnerStruct2{}, "QueryInnerStruct.MM")
	query.Group[User, Department, []Department]([]User{}, "UserID", func(user []User) Department {
		return Department{}
	})
	query.Group[User, Department, []Department]([]User{}, "CreateTime", func(user []User) Department {
		return Department{}
	})
	query.Group[User, []Department, *[]Department]([]User{}, "CreateTime", func(user []User) []Department {
		return []Department{}
	})
	query.Group[int, Department, []Department]([]int{}, ".", func(ids []int) Department {
		users := query.Select[int, User](ids, func(id int) User {
			return User{UserID: id}
		})
		return Department{Employees: users}
	})
	query.Group[int, int, []int]([]int{}, ".", func(Data []int) int {
		return len(Data)
	})
	query.Group[ContentType, []ContentType, *[]ContentType]([]ContentType{}, " Ok ", func(list []ContentType) []ContentType {
		return []ContentType{}
	})
	query.Group[string, ContentType, []ContentType]([]string{"a", "a", "", "", "z"},
		".",
		func(list []string) ContentType {
			return ContentType{
				Name: list[0],
				Age:  len(list),
			}
		})
	query.Group[ContentType, []ContentType, *[]ContentType]([]ContentType{},
		"Name",
		func(list []ContentType) []ContentType {
			sum := query.Sum(query.Column[ContentType, float32](list, "  Money  "))
			list[0].Money = sum.(float32)
			return []ContentType{list[0]}
		})
	query.Group[ContentType, float32, []float32]([]ContentType{},
		"Name",
		func(list []ContentType) float32 {
			sum := query.Sum(query.Column[ContentType, float32](list, "  Money  ")).(float32)
			return sum
		})
	query.Group[ContentType, []ContentType, *[]ContentType]([]ContentType{},
		"Ok",
		func(list []ContentType) []ContentType {
			sum := query.Sum(query.Column[ContentType, float64](list, "CardMoney  "))
			list[0].CardMoney = sum.(float64)
			return []ContentType{list[0]}
		})
	query.Group[ContentType, []ContentType, *[]ContentType]([]ContentType{},
		" Age ",
		func(list []ContentType) []ContentType {
			sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  "))
			list[0].CardMoney = sum.(float64)
			return []ContentType{list[0]}
		})
	query.Group[ContentType, float64, []float64]([]ContentType{},
		" Age ",
		func(list []ContentType) float64 {
			sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  ")).(float64)
			return sum

		})
	query.Group[ContentType, []float64, *[]float64]([]ContentType{},
		" Age ",
		func(list []ContentType) []float64 {
			sum := query.Sum(query.Column[ContentType, float64](list, "  CardMoney  "))
			return []float64{sum.(float64)}
		})
	query.Group[ContentType, int, []int]([]ContentType{},
		"Register ",
		func(list []ContentType) int {
			sum := query.Sum(query.Column[ContentType, int](list, "  Age  "))
			return sum.(int)

		})
	query.Group[ContentType, []ContentType, *[]ContentType]([]ContentType{},
		"Register ",
		func(list []ContentType) []ContentType {
			sum := query.Sum(query.Column[ContentType, int](list, "  Age  "))
			list[0].Age = sum.(int)
			return []ContentType{list[0]}
		})
	query.Group[QueryInnerStruct2, []QueryInnerStruct2, *[]QueryInnerStruct2]([]QueryInnerStruct2{},
		"QueryInnerStruct.MM",
		func(list []QueryInnerStruct2) []QueryInnerStruct2 {
			sum := query.Sum(query.Column[QueryInnerStruct2, int](list, "  MM  "))
			list[0].MM = sum.(int)
			return []QueryInnerStruct2{list[0]}
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
	query.Combine([]ContentType{}, []ContentType{}, func(left ContentType, right ContentType) ContentType {
		return ContentType{}
	})
	query.Combine([]ContentType{}, []int{}, func(left ContentType, right int) ContentType {
		return ContentType{}
	})
	query.Combine([]int{}, []User{}, func(left int, right User) User {
		return User{}
	})
}
