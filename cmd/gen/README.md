# gen

这个工具是用来生成 query 里方法的相应版本的。

## 性能测试结果

```go
make

go install github.com/fishedee/tools/cmd/gen
rm -rf testdata/testdata_querygen.go
gen -r github.com/fishedee/tools/cmd/gen/testdata
go test . -bench=. -benchmem -v
QueryReflectWarning: false
=== RUN   TestQueryColumn
--- PASS: TestQueryColumn (0.00s)
=== RUN   TestQueryColumnMap
--- PASS: TestQueryColumnMap (0.00s)
=== RUN   TestQueryCombine
--- PASS: TestQueryCombine (0.00s)
=== RUN   TestQueryGroup
--- PASS: TestQueryGroup (0.00s)
=== RUN   TestQueryJoin
--- PASS: TestQueryJoin (0.00s)
=== RUN   TestQuerySelect
--- PASS: TestQuerySelect (0.00s)
=== RUN   TestQuerySort
--- PASS: TestQuerySort (0.00s)
=== RUN   TestQueryWhere
--- PASS: TestQueryWhere (0.00s)
goos: linux
goarch: amd64
pkg: github.com/fishedee/tools/cmd/gen
BenchmarkQueryColumnHand-8                200000
     5297 ns/op     8192 B/op          1 allocs/op
BenchmarkQueryColumnMacro-8               300000
     5223 ns/op     8256 B/op          3 allocs/op
BenchmarkQueryColumnReflect-8              50000
    38013 ns/op     8320 B/op          5 allocs/op
BenchmarkQueryColumnMapHand-8              20000
    66341 ns/op   147488 B/op          2 allocs/op
BenchmarkQueryColumnMapMacro-8             20000
    66511 ns/op   147568 B/op          4 allocs/op
BenchmarkQueryColumnMapReflect-8            5000
   239529 ns/op   314034 B/op         39 allocs/op
BenchmarkQueryCombineHand-8                50000
    26340 ns/op    65536 B/op          1 allocs/op
BenchmarkQueryCombineMacro-8               50000
    37597 ns/op    65632 B/op          4 allocs/op
BenchmarkQueryCombineReflect-8              5000
   371373 ns/op   241632 B/op       2004 allocs/op
BenchmarkQueryGroupHand-8                  10000
   174619 ns/op   195104 B/op       1003 allocs/op
BenchmarkQueryGroupMacro-8                 20000
   105002 ns/op   155866 B/op         11 allocs/op
BenchmarkQueryGroupReflect-8                2000
   589985 ns/op   332019 B/op       4016 allocs/op
BenchmarkQueryJoinHand-8                   10000
   221373 ns/op   264036 B/op       1031 allocs/op
BenchmarkQueryJoinMacro-8                  10000
   123249 ns/op   157075 B/op         18 allocs/op
BenchmarkQueryJoinReflect-8                 2000
   748088 ns/op   430941 B/op       3039 allocs/op
BenchmarkQuerySelectHand-8                500000
     3354 ns/op     1024 B/op          1 allocs/op
BenchmarkQuerySelectMacro-8               300000
     5402 ns/op     1088 B/op          3 allocs/op
BenchmarkQuerySelectReflect-8               5000
   286608 ns/op    97088 B/op       2003 allocs/op
BenchmarkQuerySortHand-8                    2000
   877680 ns/op    57488 B/op          4 allocs/op
BenchmarkQuerySortMacro-8                   3000
   403047 ns/op    57536 B/op          8 allocs/op
BenchmarkQuerySortReflect-8                 1000
  1332166 ns/op    57952 B/op         22 allocs/op
BenchmarkQueryWhereHand-8                 100000
    18973 ns/op    57344 B/op          1 allocs/op
BenchmarkQueryWhereMacro-8                100000
    19157 ns/op    57408 B/op          3 allocs/op
BenchmarkQueryWhereReflect-8               10000
   207993 ns/op    48064 B/op       2002 allocs/op
PASS
ok      github.com/fishedee/tools/cmd/gen       43.099s
```
