package macro

import (
	"bytes"
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
	err := macro.ImportRecursive("github.com/fishedee/tools/macro")
	assert.Equal(t, err, nil)

	err = macro.Walk(func(pkg Package) {
		pkg.OnFuncCall(func(expr *ast.CallExpr, caller *types.Func, args []types.TypeAndValue) {
			packag := caller.Pkg()
			if packag == nil {
				return
			}
			if packag.Path() != "github.com/fishedee/tools/macro" {
				return
			}
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
	src, err := NewMacro().FormatSource(s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, src, `package main

import (
	"fmt"
)

func main() { fmt.Println("123") }
`)
}

func TestPkgDir(t *testing.T) {
	m := &Macro{}
	for _, cas := range []struct {
		pkg string
	}{
		{"github.com/fishedee/tools/macro"},
		{"database/sql"},
	} {

		r, err := m.getPkgBaseDir(cas.pkg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s\n", r)
	}
}
