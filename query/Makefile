.PHONY:test_macro test_reflect test_all
test_all:test_reflect test_macro 
	echo "test finish"
test_macro:
	go install github.com/fishedee/tools/cmd/gen
	rm -rf testdata/testdata_querygen.go
	gen -r github.com/fishedee/tools/query/testdata
	#cd testdata && gen -r .
	go test . -v -args macro
test_reflect:
	rm -rf testdata/testdata_querygen.go
	go test . -v -args reflect
