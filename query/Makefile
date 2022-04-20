.PHONY:test_macro test_reflect test_all
test_all:test_reflect test_macro 
	echo "test finish"
test_macro:
	go install github.com/fishedee/tools/cmd/gen
	rm -rf test_data/testdata_querygen.go
	cd test_data && gen -r 
	go test . -v -args macro
test_reflect:
	rm -rf test_data/testdata_querygen.go
	go test . -v -args reflect
