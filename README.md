# tools

[![GoDoc](https://godoc.org/github.com/fishedee/tools?status.svg)](https://godoc.org/github.com/fishedee/tools)

Efficient and convenient func for data processing.

## Query

```go
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

// Department Department
type Department struct {
    DepartmentID int
    Name         string
    Employees    []User
}

// Sex Sex
type Sex struct {
    IsMale bool
}

var users = make([]User, 1000, 1000)
var admins = make([]Admin, 1000, 1000)

// 获取指定列
result := query.Column(users, "UserID")
userIDs := result.([]int)

// 以某列做映射
result = query.ColumnMap(users, "UserID")
userMap := result.(map[int]User)

// 结合
result = query.Combine(admins, users, func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
})
combine := result.([]AdminUser)

// 分组
result = query.Group(users, "UserID", func(users []User) Department {
    return Department{
        Employees: users,
    }
})
group := result.([]Department)

// 连接
result = query.LeftJoin(admins, users, "AdminID = UserID", func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
})
join := result.([]AdminUser)

// 选择
result = query.Select(users, func(a User) Sex {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return Sex{IsMale: true}
    }
    return Sex{IsMale: false}
})
sel := result.([]Sex)

// 排序
result = query.Sort(users, "UserID asc")
sort := result.([]User)

// 条件
result = query.Where(users, func(a User) bool {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return true
    }
    return false
})
where := result.([]User)
```
