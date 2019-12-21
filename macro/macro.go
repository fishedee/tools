package macro

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"strings"

	"golang.org/x/tools/go/loader"
)

// FuncCallInspector FuncCallInspector
type FuncCallInspector func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue)

// Package Package
type Package struct {
	fset     *token.FileSet
	pkg      *loader.PackageInfo
	funcCall FuncCallInspector
}

// Package Package
func (mp *Package) Package() *types.Package {
	return mp.pkg.Pkg
}

// FileSet FileSet
func (mp *Package) FileSet() *token.FileSet {
	return mp.fset
}

// TypeInfo TypeInfo
func (mp *Package) TypeInfo() types.Info {
	return mp.pkg.Info
}

// OnFuncCall OnFuncCall
func (mp *Package) OnFuncCall(funcCall FuncCallInspector) {
	mp.funcCall = funcCall
}

func (mp *Package) fireFuncCall(n ast.Node) {
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
func (mp *Package) Inspect() {
	for _, file := range mp.pkg.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			//检查函数
			mp.fireFuncCall(n)
			return true
		})
	}
}

// Walker Walker
type Walker func(inspector Package)

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
	baseDir, err := m.getPkgBaseDir(pkg)
	if err != nil {
		return err
	}

	allPackage, err := m.getAllDir(baseDir, pkg)
	if err != nil {
		return err
	}
	for _, packageSingle := range allPackage {
		m.packages[packageSingle] = true
	}
	return nil
}

func (m *Macro) getPkgBaseDir(pkg string) (string, error) {
	buildPkg, err := build.Import(pkg, "", build.ImportComment)
	if err != nil {
		return "", err
	}
	dir := buildPkg.Dir

	index := strings.LastIndex(dir, pkg)
	if index == -1 {
		return "", fmt.Errorf("路径中找不到包: %s, %s", dir, pkg)
	}
	if index == 0 {
		return "", fmt.Errorf("路径与包相同: %s, %s", dir, pkg)
	}
	baseDir := dir[:index-1]

	return baseDir, nil
}

// Walk Walk
func (m *Macro) Walk(walker Walker) error {
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
		inspector := Package{
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
