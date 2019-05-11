package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/donnol/tools/macro"
	"github.com/donnol/tools/plode"
)

var (
	recursive      = flag.Bool("r", false, "generate package including sub package")
	queryGenMapper = map[string]queryGenHandler{}
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of querygen:\n")
	fmt.Fprintf(os.Stderr, "\tlarge improve performance language/query.go function \n")
	fmt.Fprintf(os.Stderr, "\tquerygen [flags] [packageName]\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttps://github.com/fishedee/fishgo/tree/master/src/github.com/fishedee/language/querygen\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func nodeString(fset *token.FileSet, n ast.Node) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, n)
	return buf.String()
}

type queryGenRequest struct {
	pkg    macro.MacroPackage
	expr   *ast.CallExpr
	caller *types.Func
	args   []types.TypeAndValue
}

type queryGenResponse struct {
	importPackage map[string]bool
	funcName      string
	funcBody      string
	initBody      string
}

type queryGenHandler func(request queryGenRequest) *queryGenResponse

func handleQueryGen(name string, request queryGenRequest) *queryGenResponse {
	handler, isExist := queryGenMapper[name]
	if isExist == false {
		return nil
	}
	return handler(request)
}

func registerQueryGen(name string, handler queryGenHandler) {
	queryGenMapper[name] = handler
}

func formatSource(data string) []byte {
	result, err := format.Source([]byte(data))
	if err != nil {
		Throw(1, "format source fail!%v,%v", err, data)
	}
	return result
}

func generate(packageName string, packagePath string, packages []queryGenResponse) {
	var fileDir string
	gopath, _ := os.LookupEnv("GOPATH")
	fileDir = gopath + "/src/" + packagePath
	fileSegment := plode.Explode(fileDir, "/")
	filePath := fileDir + "/" + fileSegment[len(fileSegment)-1] + "_querygen.go"

	//处理导入包
	importPackageMap := map[string]bool{}
	for _, singlePackage := range packages {
		for singleImport := range singlePackage.importPackage {
			importPackageMap[singleImport] = true
		}
	}
	importPackageMap["github.com/donnol/tools/query"] = true
	delete(importPackageMap, packagePath)
	importPackageList := []string{}
	for singlePackage := range importPackageMap {
		importPackageList = append(importPackageList, "\""+singlePackage+"\"")
	}
	sort.Slice(importPackageList, func(i int, j int) bool {
		return importPackageList[i] < importPackageList[j]
	})
	importBody := plode.Implode(importPackageList, "\n")

	//处理funcBody和initBody
	sort.Slice(packages, func(i int, j int) bool {
		return packages[i].funcName < packages[j].funcName
	})
	var funcBody bytes.Buffer
	var initBody bytes.Buffer
	for _, singlePackage := range packages {
		funcBody.WriteString(singlePackage.funcBody)
		initBody.WriteString(singlePackage.initBody)
	}

	//写入数据
	result := `package ` + packageName + "\n" +
		"import (\n" + importBody + ")\n" +
		funcBody.String() + "\n" +
		"func init(){\n" +
		initBody.String() + "\n" +
		"}\n"
	oldData, _ := ioutil.ReadFile(filePath)
	if string(oldData) == result {
		return
	}
	err := ioutil.WriteFile(filePath, formatSource(result), 0644)
	if err != nil {
		panic(err)
	}
}

func run() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
		panic("need package name")
	}

	macroObj := macro.NewMacro()
	if *recursive {
		err := macroObj.ImportRecursive(args[0])
		if err != nil {
			panic(err)
		}
	} else {
		err := macroObj.Import(args[0])
		if err != nil {
			panic(err)
		}
	}

	genPackage := []queryGenResponse{}
	initPackageName := ""
	globalGeneratePackagePath = args[0]
	err := macroObj.Walk(func(pkg macro.MacroPackage) {
		if pkg.Package().Path() == args[0] {
			initPackageName = pkg.Package().Name()
		}
		pkg.OnFuncCall(func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue) {
			callerFullName := caller.FullName()
			request := queryGenRequest{
				pkg:    pkg,
				expr:   expr,
				caller: caller,
				args:   args,
			}
			response := handleQueryGen(callerFullName, request)
			if response != nil {
				genPackage = append(genPackage, *response)
			}
		})
		pkg.Inspect()
	})
	if err != nil {
		panic(err)
	}

	generate(initPackageName, args[0], genPackage)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("querygen fail: ")
	run()
}
