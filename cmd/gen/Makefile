.PHONY:gen
bench_only:
	go test . -bench=. -benchmem -v
bench:gen bench_only	
gen:
	go install github.com/fishedee/tools/cmd/gen
	rm -rf testdata/testdata_querygen.go
	gen -r github.com/fishedee/tools/cmd/gen/testdata
test:gen
	go test . -v
