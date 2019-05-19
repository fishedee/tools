# 测试结果

可以看到，生成的方法和手写的方法性能相差无几。

```go
go test -v -bench=. -benchmem

goos: linux
goarch: amd64
pkg: github.com/fishedee/tools/cmd/gen/testdata

BenchmarkQueryColumnHand-8 300000 5001 ns/op 8192 B/op 1 allocs/op
BenchmarkQueryColumnMacro-8 300000 5033 ns/op 8256 B/op 3 allocs/op

BenchmarkQueryColumnHandMany-8 2000 902803 ns/op  1605632 B/op 2 allocs/op
BenchmarkQueryColumnMacroMany-8 1000 1213797 ns/op  1605760 B/op 6 allocs/op

BenchmarkQueryColumnMapHand-8 20000 62174 ns/op 147489 B/op 2 allocs/op
BenchmarkQueryColumnMapMacro-8 20000 62215 ns/op 147568 B/op 4 allocs/op

BenchmarkQueryColumnMapHandMany-8 10000 126767 ns/op 294979 B/op 4 allocs/op
BenchmarkQueryColumnMapMacroMany-8 10000 119876 ns/op 295138 B/op 8 allocs/op

BenchmarkQueryCombineHand-8 50000 23138 ns/op 65536 B/op 1 allocs/op
BenchmarkQueryCombineMacro-8 50000 34255 ns/op 65632 B/op 4 allocs/op

BenchmarkQueryGroupHand-8 10000 187355 ns/op 195104 B/op 1003 allocs/op
BenchmarkQueryGroupMacro-8 20000 92391 ns/op 155866 B/op 11 allocs/op

BenchmarkQueryJoinHand-8 10000 205800 ns/op 248041 B/op 1031 allocs/op
BenchmarkQueryJoinMacro-8 10000 127462 ns/op 157074 B/op 18 allocs/op

BenchmarkQuerySelectHand-8 500000 3491 ns/op 1024 B/op 1 allocs/op
BenchmarkQuerySelectMacro-8 300000 5654 ns/op 1088 B/op 3 allocs/op

BenchmarkQuerySortHand-8 2000 788994 ns/op 57488 B/op 4 allocs/op
BenchmarkQuerySortMacro-8 3000 426857 ns/op 57536 B/op 8 allocs/op

BenchmarkQueryWhereHand-8 100000 15417 ns/op 57344 B/op 1 allocs/op
BenchmarkQueryWhereMacro-8 100000 19301 ns/op 57408 B/op 3 allocs/op

PASS
ok github.com/fishedee/tools/cmd/gen/testdata 28.909s
```
