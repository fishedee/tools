# tools

[![Go Report Card](https://goreportcard.com/badge/github.com/fishedee/tools)](https://goreportcard.com/report/github.com/fishedee/tools)
[![GoDoc](https://pkg.go.dev/github.com/fishedee/tools?status.svg)](https://pkg.go.dev/github.com/fishedee/tools)

高性能便利的golang数据操作库，原始代码来源于这里[here](https://github.com/fishedee/fishgo/tree/master/src/github.com/fishedee/language).

**注意, 现已支持Go1.18泛型。**

## 原理

```golang
sort.Slice(people, func(i, j int) bool {
	return people[i].Age > people[j].Age
})
```

因为写后台服务时，经常是从数据库中取出数据，然后我们时常需要对多个表的数据在golang业务层进行group,join,sort等操作。这些操作大同小异，写得多了就感觉重复繁琐，容易出错。

```golang
pepole2 := query.Sort(people,"Age asc")
```

所以我们的想法是模拟数据库的sql操作，在golang上用统一的API来执行group,join,sort等操作。当然，由于golang的强类型特性，我们只能用interface{}来处理不同类型的输入，然后用reflect来实现算法。有了这个工具以后，我们在写业务代码时爽了很多，代码更简洁，更清晰。但是，我们随后发现了这样做对于超大数据量（10w行以上数据）时，会造成比较严重的性能问题。原因很简单，因为操作数据的算法都是用reflect来实现，有大量的时间浪费在reflect的类型解析和校验上了，而且，golang编译器无法像手工写代码一样能在编译时根据不同的代码执行最好的优化。

```golang
func querySortV10a439a196b4cc9dca0592a40a23aba8392203e4(data []People, sortType string) []People {
	dataIn := data
	newData := make([]People, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Age < newData[j].Age {
			return -1
		} else if newData[i].Age > newData[j].Age {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}
```

便利性解决了，我们就进一步解决性能问题。我们的想法是，在编译时通过分析ast树，查找出所有调用query.Sort的地方，提取出它的调用类型，和字符串常数，然后根据不同的类型和常数自动生成出具体类型的sort代码，然后将这部分自动生成的代码也一起编译进可执行文件里面。当运行时，执行query.Sort时，我们根据对应类型和字符串常数，来跳转到具体的不同的自动生成的query.Sort函数体内部去执行。于是，我们既实现了便利性，也实现了高性能。


```sh
gen -r [packageName] # 指定包名，不指定则使用当前目录所在包
```

而这一切，我们需要做的，仅仅就是在编译前，执行一行命令而已。


## 例子

```golang
// User结构体
type User struct {
    UserID     int
    Age        int
    Name       string
    CreateTime time.Time
}

// Admin结构体
type Admin struct {
    AdminID int
    Level   int
}

// AdminUser结构体
type AdminUser struct {
    AdminID    int
    Level      int
    Age        int
    Name       string
    CreateTime time.Time
}

// Department结构体
type Department struct {
    DepartmentID int
    Name         string
    Employees    []User
}

// Sex结构体
type Sex struct {
    IsMale bool
}

var users = make([]User, 1000, 1000)
var admins = make([]Admin, 1000, 1000)

// 获取指定列
// * 第一个参数为表格
// * 第二个参数为列名
userIDs := query.Column[User, int](users, "UserID") // []int

// 以某列的值生成映射
// * 第一个参数为表格
// * 第二个参数为列名
userMap = query.ColumnMap[User, int, map[int]User](users, "UserID") // map[int]User
// or
usersMap = query.ColumnMap[User, int, map[int][]User](users, "[]UserID") // map[int][]User

// 表格转换或提取操作
// * 第一个参数为表格
// * 第二个参数为转换规则
sel = query.Select(users, func(a User) Sex {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return Sex{IsMale: true}
    }
    return Sex{IsMale: false}
}) // []Sex

// 表格筛选操作
// * 第一个参数为表格
// * 第二个参数为筛选规则
where = query.Where(users, func(a User) bool {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return true
    }
    return false
}) // []User

// 两个表一对一合并
// * 第一个参数为左表
// * 第二个参数为右表
// * 第三个参数为合并规则
combine = query.Combine(admins, users, func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
}) // []AdminUser

// 表根据某列进行分组操作
// * 第一个参数为表格
// * 第二个参数为列名
// * 第三个参数为分组规则
group = query.Group[User, Department, []Department](users, "UserID", func(users []User) Department {
    return Department{
        Employees: users,
    }
}) // []Department
// or
group = query.Group[User, []Department, *[]Department](users, "UserID", func(users []User) []Department {
    return []Department{
        Employees: users,
    }
}) // *[]Department

// 两个表进行连接操作，支持LeftJoin,RightJoin,InnerJoin和OuterJoin
// * 第一个参数为左表
// * 第二个参数为右表
// * 第三个参数为连接操作的条件
// * 第四个参数为连接规则
join = query.LeftJoin(admins, users, "AdminID = UserID", func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
}) // []AdminUser

// 表格排序操作，支持多列排序，如UserId desc,Age asc
// * 第一个参数为表格
// * 第二个参数为排序规则
sort = query.Sort(users, "UserID asc") // []User
```

就以上这几个API，没有其他的了，参数一目了然。

## Gen工具

```shell
安装：
go install github.com/fishedee/tools/cmd/gen

使用:
gen -r [包名]
```

gen工具是性能提速的核心，使用方法很简单，gen加包名就可以了，如果加-r参数，其会递归查找包下的所有子包的源码。gen读取源码以后，其会自动分析调用query.XXXX的地方，然后就会生成一个xxxx_querygen.go代码来代替reflect的实现。最后，你就能如平常一样编译代码就可以了，就能享受到超便利和高性能的好处了。但要注意的是，当包下面的源代码发生改变后，你要重新执行gen来生成代码。

```
goos: linux
goarch: amd64
pkg: github.com/fishedee/tools/cmd/gen
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz

BenchmarkQueryColumnHand
BenchmarkQueryColumnHand-4                        215596              4776 ns/op     8192 B/op          1 allocs/op
BenchmarkQueryColumnMacro
BenchmarkQueryColumnMacro-4                       214627              4770 ns/op     8216 B/op          2 allocs/op
BenchmarkQueryColumnReflect
BenchmarkQueryColumnReflect-4                      50305             23254 ns/op     8296 B/op          5 allocs/op

BenchmarkQueryColumnMapHand
BenchmarkQueryColumnMapHand-4                      17434             72615 ns/op   147480 B/op          2 allocs/op
BenchmarkQueryColumnMapMacro
BenchmarkQueryColumnMapMacro-4                     17889             71254 ns/op   147552 B/op          4 allocs/op
BenchmarkQueryColumnMapReflect
BenchmarkQueryColumnMapReflect-4                    9592            114374 ns/op   147608 B/op          6 allocs/op

BenchmarkQueryColumnMapSliceHand
BenchmarkQueryColumnMapSliceHand-4                  7903            154506 ns/op   145944 B/op       1002 allocs/op
BenchmarkQueryColumnMapSliceMacro
BenchmarkQueryColumnMapSliceMacro-4                 6750            151835 ns/op   188649 B/op         12 allocs/op
BenchmarkQueryColumnMapSliceReflect
BenchmarkQueryColumnMapSliceReflect-4               3522            314086 ns/op   212818 B/op       1018 allocs/op

BenchmarkQueryCombineHand
BenchmarkQueryCombineHand-4                        37981             31381 ns/op    65536 B/op          1 allocs/op
BenchmarkQueryCombineMacro
BenchmarkQueryCombineMacro-4                       29568             37010 ns/op    65584 B/op          3 allocs/op
BenchmarkQueryCombineReflect
BenchmarkQueryCombineReflect-4                      1999            578415 ns/op        177657 B/op       3006 allocs/op

BenchmarkQueryGroupHand
BenchmarkQueryGroupHand-4                           4983            201296 ns/op  195098 B/op       1003 allocs/op
BenchmarkQueryGroupMacro
BenchmarkQueryGroupMacro-4                         11942            104433 ns/op  155810 B/op         10 allocs/op
BenchmarkQueryGroupReflect
BenchmarkQueryGroupReflect-4                        1810            643851 ns/op  275994 B/op       4017 allocs/op

BenchmarkQueryJoinHand
BenchmarkQueryJoinHand-4                            6081            180254 ns/op  256049 B/op       1031 allocs/op
BenchmarkQueryJoinMacro
BenchmarkQueryJoinMacro-4                           7669            132689 ns/op  156995 B/op         17 allocs/op
BenchmarkQueryJoinReflect
BenchmarkQueryJoinReflect-4                         1252            805528 ns/op  472557 B/op       4039 allocs/op

BenchmarkQuerySelectHand
BenchmarkQuerySelectHand-4                        360092              3073 ns/op    1024 B/op          1 allocs/op
BenchmarkQuerySelectMacro
BenchmarkQuerySelectMacro-4                       204452              5238 ns/op    1048 B/op          2 allocs/op
BenchmarkQuerySelectReflect
BenchmarkQuerySelectReflect-4                       3679            301481 ns/op   26096 B/op       2004 allocs/op

BenchmarkQuerySortHand
BenchmarkQuerySortHand-4                            2088            490302 ns/op   57480 B/op          4 allocs/op
BenchmarkQuerySortMacro
BenchmarkQuerySortMacro-4                           3784            321839 ns/op   57488 B/op          7 allocs/op
BenchmarkQuerySortReflect
BenchmarkQuerySortReflect-4                         1339            833588 ns/op   57848 B/op         20 allocs/op

BenchmarkQueryWhereHand
BenchmarkQueryWhereHand-4                          62037             17216 ns/op   57344 B/op          1 allocs/op
BenchmarkQueryWhereMacro
BenchmarkQueryWhereMacro-4                         49932             20469 ns/op   57368 B/op          2 allocs/op
BenchmarkQueryWhereReflect
BenchmarkQueryWhereReflect-4                        5196            233892 ns/op   25072 B/op       2003 allocs/op

PASS
ok      github.com/fishedee/tools/cmd/gen       38.660s
```

我们从性能测试的结果可以看出，自动生成的代码至少和手工写的代码一样优化，在一些特定的场景下（group,join,sort），自动生成的代码由于高度优化的算法实现，和仔细的内存分配，它甚至比手工写的代码要快1倍。想知道为什么，可以看一下自动生成的[代码](https://github.com/fishedee/tools/blob/master/query/testdata/testdata_querygen.go)

## 协议

看[这里](LICENSE).
