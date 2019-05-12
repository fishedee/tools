package macro

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"

	"golang.org/x/tools/go/loader"
)

// MacroFuncCallInspector MacroFuncCallInspector
type MacroFuncCallInspector func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue)

// MacroPackage MacroPackage
type MacroPackage struct {
	fset     *token.FileSet
	pkg      *loader.PackageInfo
	funcCall MacroFuncCallInspector
}

// Package Package
func (mp *MacroPackage) Package() *types.Package {
	return mp.pkg.Pkg
}

// FileSet FileSet
func (mp *MacroPackage) FileSet() *token.FileSet {
	return mp.fset
}

// TypeInfo TypeInfo
func (mp *MacroPackage) TypeInfo() types.Info {
	return mp.pkg.Info
}

// OnFuncCall OnFuncCall
func (mp *MacroPackage) OnFuncCall(funcCall MacroFuncCallInspector) {
	mp.funcCall = funcCall
}

func (mp *MacroPackage) fireFuncCall(n ast.Node) {
	if mp.funcCall == nil {
		return
	}

	expr, ok := n.(*ast.CallExpr)
	if ok == false {
		return
	}

	//获取caller信息
	exprIdent, ok := expr.Fun.(*ast.Ident)
	if ok == false {
		selectorExpr, ok := expr.Fun.(*ast.SelectorExpr)
		if ok == false {
			return
		}
		exprIdent = selectorExpr.Sel
	}

	info := mp.pkg.Info
	funcObj, ok := info.Uses[exprIdent].(*types.Func)
	if ok == false {
		return
	}

	//获取argument信息
	typeAndValues := []types.TypeAndValue{}
	for _, arg := range expr.Args {
		t1, isExist := info.Types[arg]
		if isExist == false {
			panic(fmt.Sprintf("unknown argument type:%v", expr.Args))
		}
		typeAndValues = append(typeAndValues, t1)
	}

	//触发
	mp.funcCall(expr, funcObj, typeAndValues)
}

// Inspect Inspect
func (mp *MacroPackage) Inspect() {
	for _, file := range mp.pkg.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			//检查函数
			mp.fireFuncCall(n)
			return true
		})
	}
}

// MacroWalker MacroWalker
type MacroWalker func(inspector MacroPackage)

// Macro Macro
type Macro struct {
	packages map[string]bool
}

// NewMacro NewMacro
func NewMacro() *Macro {
	return &Macro{
		packages: map[string]bool{},
	}
}

// Import Import
func (m *Macro) Import(pkg string) error {
	m.packages[pkg] = true
	return nil
}

func (m *Macro) getAllDir(baseDir string, pkgName string) ([]string, error) {
	files, err := ioutil.ReadDir(baseDir + "/" + pkgName)
	if err != nil {
		return nil, err
	}
	result := []string{}
	result = append(result, pkgName)
	for _, file := range files {
		if file.IsDir() {
			subPackageName := pkgName + "/" + file.Name()
			subPackages, err := m.getAllDir(baseDir, subPackageName)
			if err != nil {
				return nil, err
			}
			result = append(result, subPackages...)
		}
	}
	return result, nil
}

// ImportRecursive ImportRecursive
func (m *Macro) ImportRecursive(pkg string) error {
	gopath, _ := os.LookupEnv("GOPATH")
	allPackage, err := m.getAllDir(gopath+"/src", pkg)
	if err != nil {
		return err
	}
	for _, packageSingle := range allPackage {
		m.packages[packageSingle] = true
	}
	return nil
}

// Walk Walk
func (m *Macro) Walk(walker MacroWalker) error {
	var conf loader.Config
	if len(m.packages) == 0 {
		return errors.New("none package have to load")
	}
	for singlePackage := range m.packages {
		conf.Import(singlePackage)
	}
	lprog, err := conf.Load()
	if err != nil {
		return err
	}
	for _, singlePackage := range lprog.Imported {
		if len(singlePackage.Errors) != 0 {
			return singlePackage.Errors[0]
		}
		inspector := MacroPackage{
			fset: lprog.Fset,
			pkg:  singlePackage,
		}
		walker(inspector)
	}
	return nil
}

// FormatSource FormatSource
func (m *Macro) FormatSource(data string) (string, error) {
	result, err := format.Source([]byte(data))
	if err != nil {
		return "", errors.New(err.Error() + "," + data)
	}
	return string(result), nil
}
