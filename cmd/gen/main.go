package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fishedee/tools/exception"
	"github.com/fishedee/tools/macro"
	"github.com/fishedee/tools/plode"
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

// QueryGenRequest QueryGenRequest
type QueryGenRequest struct {
	pkg    macro.Package
	expr   *ast.CallExpr
	caller *types.Func
	args   []types.TypeAndValue
}

// QueryGenResponse QueryGenResponse
type QueryGenResponse struct {
	importPackage map[string]bool
	funcName      string
	funcBody      string
	initBody      string
}

type queryGenHandler func(request QueryGenRequest) *QueryGenResponse

func handleQueryGen(name string, request QueryGenRequest) *QueryGenResponse {
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
		exception.Throw(1, "format source fail!%v,%v", err, data)
	}
	return result
}

func generate(packageName string, packagePath string, packages []QueryGenResponse) {
	var fileDir string
	buildPkg, err := build.Import(packagePath, "", build.ImportComment)
	if err != nil {
		panic(err)
	}
	fileDir = buildPkg.Dir

	fileSegment := plode.Explode(fileDir, "/")
	filePath := fileDir + "/" + fileSegment[len(fileSegment)-1] + "_querygen.go"

	//处理导入包
	importPackageMap := map[string]bool{}
	for _, singlePackage := range packages {
		for singleImport := range singlePackage.importPackage {
			importPackageMap[singleImport] = true
		}
	}
	importPackageMap["github.com/fishedee/tools/query"] = true
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
	oldData, err := ioutil.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if string(oldData) == result {
		return
	}
	err = ioutil.WriteFile(filePath, formatSource(result), 0644)
	if err != nil {
		panic(err)
	}
}

func isInVendorPath(path string) bool {
	pathInfo := plode.Explode(path, "/")
	for i := 0; i != len(pathInfo); i++ {
		if pathInfo[i] == "vendor" {
			return true
		}
	}
	return false
}

func getRealFullName(callerFullName string) string {
	fullNameInfo := plode.Explode(callerFullName, "/")
	i := len(fullNameInfo) - 1
	for ; i >= 0; i-- {
		if fullNameInfo[i] == "vendor" {
			break
		}
	}
	realFullName := plode.Implode(fullNameInfo[i+1:], "/")
	return realFullName
}

func run() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	var modPath, pkgPath string
	if len(args) == 0 {
		// 没有参数时，尝试获取当前目录所在模块的go.mod文件，从而获取包路径
		modPath, pkgPath = getPkgPathFromDir()

		if pkgPath == "" {
			usage()
			panic("need package name")
		}
	} else {
		pkgPath = args[0]
	}

	macroObj := macro.NewMacro()
	if *recursive {
		err := macroObj.ImportRecursive(modPath, pkgPath)
		if err != nil {
			panic(err)
		}
	} else {
		err := macroObj.Import(modPath, pkgPath)
		if err != nil {
			panic(err)
		}
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.FinalMSG = "gen finish.\n"
	s.Start() // Start the spinner
	defer func() {
		s.Stop()
	}()

	genPackage := []QueryGenResponse{}
	initPackageName := ""
	globalGeneratePackagePath = pkgPath
	err := macroObj.Walk(func(pkg macro.Package) {
		if pkg.Package().Path() == pkgPath {
			initPackageName = pkg.Package().Name()
		}
		if isInVendorPath(pkg.Package().Path()) {
			return
		}
		pkg.OnFuncCall(func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue) {
			callerFullName := caller.FullName()
			callerFullName = getRealFullName(callerFullName)
			request := QueryGenRequest{
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

	generate(initPackageName, pkgPath, genPackage)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("querygen fail: ")
	defer exception.CatchCrash(func(e exception.Exception) {
		log.Fatal(e.GetMessage())
	})
	run()
}
