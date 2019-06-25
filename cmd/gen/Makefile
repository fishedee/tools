.PHONY:gen
bench:gen
	go test . -bench=. -benchmem -v 
gen:
	go install github.com/fishedee/tools/cmd/gen
	rm -rf testdata/testdata_querygen.go
	gen -r github.com/fishedee/tools/cmd/gen/testdata
	#cd testdata && gen -r .
test:gen
	go test . -v
