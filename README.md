# tools

[![Go Report Card](https://goreportcard.com/badge/github.com/fishedee/tools)](https://goreportcard.com/report/github.com/fishedee/tools)
[![GoDoc](https://pkg.go.dev/github.com/fishedee/tools?status.svg)](https://pkg.go.dev/github.com/fishedee/tools)

Efficient and convenient func for golang data processing.All original code comes from [here](https://github.com/fishedee/fishgo/tree/master/src/github.com/fishedee/language).

Chinese read [这里](README_zhCN.md)

**This library already supports Go1.18 generic.**

## Why

```golang
sort.Slice(people, func(i, j int) bool {
	return people[i].Age > people[j].Age
})
```

When we write backend code,we often extract data from database.and next step ,we need to group,join and sort multiple table data to generate result in golang service.The code like this is always repeat again and again, The code we write more,The boring we feel more.

```golang
pepole2 := query.Sort(people,"Age asc")
```

so we have an idea!we use a unify API to finish group,join and sort.but because of strong typed in golang,we only use interface{} to handle different type input data,and use reflect to achieve the algorithm. as we finish this tool , it is a new world for me,the code more short,more easy and more beautiful.And after a period of time, we find a new problem,when we handle one million rows in this API,it will slow.This is a simple reason,because all algorithm is achieved by reflect, not by hand.Compiler can not optimize the code by different type.The other reason is there are much time wasted in reflect package by type get and type check.

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

so we have further idea, to slove why we slow.We build a new tool,when in compile time,we get the AST tree and find all the code where call query.Sort function.And we auto generate the "concrete sort code" by differnt type and different const string in sortType.The last step is we compile the "concrete sort code" to our executable file.

So when we run query.Sort(xxxx),query.Sort will find the "concrete sort code" by differnt type and different const string in sortType.It just work!It is even faster than hand code


```sh
gen -r [packageName] # specify package，if not, use current package
```

all of this ,we only need to run this command before we compile.That is all.


## Sample

```golang
// User struct
type User struct {
    UserID     int
    Age        int
    Name       string
    CreateTime time.Time
}

// Admin struct
type Admin struct {
    AdminID int
    Level   int
}

// AdminUser struct
type AdminUser struct {
    AdminID    int
    Level      int
    Age        int
    Name       string
    CreateTime time.Time
}

// Department struct
type Department struct {
    DepartmentID int
    Name         string
    Employees    []User
}

// Sex struct
type Sex struct {
    IsMale bool
}

var users = make([]User, 1000, 1000)
var admins = make([]Admin, 1000, 1000)

// extract column from table
// * First Argument:table
// * Second Argument:column name
userIDs := query.Column[User, int](users, "UserID") // []int

// generate a map from table,key is column value and value is it's row
// * First Argument:table
// * Second Argument:column name
userMap = query.ColumnMap[User, int, map[int]User](users, "UserID") // map[int]User
// or
usersMap = query.ColumnMap[User, int, map[int][]User](users, "[]UserID") // map[int][]User

// select data from table
// * First Argument:table
// * Second Argument:select rule
sel = query.Select(users, func(a User) Sex {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return Sex{IsMale: true}
    }
    return Sex{IsMale: false}
}) // []Sex

// filter data from table
// * First Argument:table
// * Second Argument:filter rule
where = query.Where(users, func(a User) bool {
    if len(a.Name) >= 3 && a.Name[0:3] == "Man" {
        return true
    }
    return false
}) // []User

// combine data from two table , one by one
// * First Argument:left table
// * Second Argument:right table
// * Third Argument:combine rule
combine = query.Combine(admins, users, func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
}) // []AdminUser

// group data from table
// * First Argument: left table
// * Second Argument: group column name
// * Third Argument: group rule
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

// join data from two table，support LeftJoin,RightJoin,InnerJoin和OuterJoin
// * First Argument: left table
// * Second Argument: right table
// * Third Argument: join condition
// * Forth Argument: join rule
join = query.LeftJoin(admins, users, "AdminID = UserID", func(admin Admin, user User) AdminUser {
    return AdminUser{
        AdminID:    admin.AdminID,
        Level:      admin.Level,
        Name:       user.Name,
        CreateTime: user.CreateTime,
    }
}) // []AdminUser

// sort data from table,support multiple column,for Example: UserId desc,Age asc
// * First Argument:table
// * Second Argument:sort condition
sort = query.Sort(users, "UserID asc") // []User
```

This is All API , easy and less

## Gen tool

```shell
Install：
go install github.com/fishedee/tools/cmd/gen

Use:
gen -r [PackageName] # can ignore PackageName, when do that, use the current package
```

gen is the core of hight performance, it is easy to use , gen and packageName. If you add -r argument , it will read all code in children package(include children of children package, and so on).when gen tool read all code ,it will analyse where call query.XXXXX, and it will auto generate a file named xxxx_querygen.go in the same package.At last, compile all the code just as common and it work!

NOTICE,if you change the code in package , you should run gen tool again.

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

In Benchmark case , auto generate code is fast as same as hand code. In some specify scence(group,join,sort),auto generate code is two times faster than hand code, because of very optimized algorithm and very careful memory allocate.If you want to know why,you should see the auto generate [code](https://github.com/fishedee/tools/blob/master/query/testdata/testdata_querygen.go)

## License

See[this](LICENSE).
