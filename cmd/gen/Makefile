.PHONY:gen
bench:gen
	go test . -bench=. -benchmem -v 
gen:
	go install github.com/donnol/tools/cmd/gen
	rm -rf testdata/testdata_querygen.go
	gen -r github.com/donnol/tools/cmd/gen/testdata
test:gen
	go test . -v
