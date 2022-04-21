# tools

[![Go Report Card](https://goreportcard.com/badge/github.com/fishedee/tools)](https://goreportcard.com/report/github.com/fishedee/tools)
[![GoDoc](https://godoc.org/github.com/fishedee/tools?status.svg)](https://godoc.org/github.com/fishedee/tools)

Efficient and convenient func for golang data processing.All original code comes from [here](https://github.com/fishedee/fishgo/tree/master/src/github.com/fishedee/language).

Chinese read [这里](README_zhCN.md)

**Now, this library support Go1.18 generic.**

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
gen -r packageName
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
BenchmarkQueryColumnHand-8                200000 	5297 ns/op     8192 B/op          1 allocs/op
BenchmarkQueryColumnMacro-8               300000 	5223 ns/op     8256 B/op          3 allocs/op
BenchmarkQueryColumnReflect-8              50000	38013 ns/op     8320 B/op          5 allocs/op

BenchmarkQueryColumnMapHand-8              20000	66341 ns/op   147488 B/op          2 allocs/op
BenchmarkQueryColumnMapMacro-8             20000	66511 ns/op   147568 B/op          4 allocs/op
BenchmarkQueryColumnMapReflect-8            5000	239529 ns/op   314034 B/op         39 allocs/op

BenchmarkQueryCombineHand-8                50000	26340 ns/op    65536 B/op          1 allocs/op
BenchmarkQueryCombineMacro-8               50000	37597 ns/op    65632 B/op          4 allocs/op
BenchmarkQueryCombineReflect-8              5000	371373 ns/op   241632 B/op       2004 allocs/op

BenchmarkQueryGroupHand-8                  10000	174619 ns/op   195104 B/op       1003 allocs/op
BenchmarkQueryGroupMacro-8                 20000	105002 ns/op   155866 B/op         11 allocs/op
BenchmarkQueryGroupReflect-8                2000	589985 ns/op   332019 B/op       4016 allocs/op

BenchmarkQueryJoinHand-8                   10000	221373 ns/op   264036 B/op       1031 allocs/op
BenchmarkQueryJoinMacro-8                  10000	123249 ns/op   157075 B/op         18 allocs/op
BenchmarkQueryJoinReflect-8                 2000	748088 ns/op   430941 B/op       3039 allocs/op

BenchmarkQuerySelectHand-8                500000	3354 ns/op     1024 B/op          1 allocs/op
BenchmarkQuerySelectMacro-8               300000	5402 ns/op     1088 B/op          3 allocs/op
BenchmarkQuerySelectReflect-8               5000	286608 ns/op    97088 B/op       2003 allocs/op

BenchmarkQuerySortHand-8                    2000	877680 ns/op    57488 B/op          4 allocs/op
BenchmarkQuerySortMacro-8                   3000	403047 ns/op    57536 B/op          8 allocs/op
BenchmarkQuerySortReflect-8                 1000	1332166 ns/op    57952 B/op         22 allocs/op

BenchmarkQueryWhereHand-8                 100000	18973 ns/op    57344 B/op          1 allocs/op
BenchmarkQueryWhereMacro-8                100000	19157 ns/op    57408 B/op          3 allocs/op
BenchmarkQueryWhereReflect-8               10000	207993 ns/op    48064 B/op       2002 allocs/op
PASS
ok      github.com/fishedee/tools/cmd/gen       43.099s
```

In Benchmark case , auto generate code is fast as same as hand code. In some specify scence(group,join,sort),auto generate code is two times faster than hand code, because of very optimized algorithm and very careful memory allocate.If you want to know why,you should see the auto generate [code](https://github.com/fishedee/tools/blob/master/query/testdata/testdata_querygen.go)

## License

See[this](LICENSE).
