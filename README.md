# tools

[![GoDoc](https://godoc.org/github.com/fishedee/tools?status.svg)](https://godoc.org/github.com/fishedee/tools)

Efficient and convenient func for data processing.

All code comes from [here](https://github.com/fishedee/fishgo/tree/master/src/github.com/fishedee/language). This repository make it easier to use.

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

## Query generate

我们都知道，使用反射来实现如 query.Column 的效率相对手写是非常低的，但是如果一直使用手写又太过啰嗦。得益于 Go 提供的标准库，以及官方 go generate 工具的启发。所以我们提供了 cmd/gen 工具来生成可以媲美手写版本性能的代码，从而我们既能享受到 query.Column 等方法的便利，又能保证性能。

```shell
安装：
go install github.com/fishedee/tools/cmd/gen

使用:
rm -rf testdata/testdata_querygen.go
gen -r github.com/fishedee/tools/cmd/gen/testdata

生成的源码文件: cmd/gen/testdata/testdata_querygen.go, 里面的代码就是我们在实际开发中需要用到的各个手写版本的替代。
```

## Author

[A very friendly, helpful, lazy, but to be honest, very good man.](https://github.com/fishedee)

## License

see [this](LICENSE).
