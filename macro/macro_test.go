package macro

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"testing"

	"github.com/fishedee/tools/assert"
)

func nodeString(fset *token.FileSet, n ast.Node) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, n)
	return buf.String()
}

func TestFuncCallInspect(t *testing.T) {
	macro := NewMacro()
	err := macro.ImportRecursive("github.com/fishedee/tools/query")
	assert.Equal(t, err, nil)

	err = macro.Walk(func(pkg Package) {
		pkg.OnFuncCall(func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue) {
			packag := caller.Pkg()
			if packag == nil {
				return
			}
			if packag.Path() != "github.com/fishedee/tools/query" {
				return
			}
			fmt.Printf("%v:%v\n", pkg.FileSet().Position(expr.Pos()), nodeString(pkg.FileSet(), expr))
		})
		pkg.Inspect()
	})
	assert.Equal(t, err, nil)
}

func TestFormatSource(t *testing.T) {
	s := `
	package main
	import ("fmt")
	func main(){fmt.Println("123")}
	`
	fmt.Println(NewMacro().FormatSource(s))
}
