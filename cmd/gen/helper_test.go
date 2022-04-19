package main

import "testing"

func TestGetPkgPathFromDir(t *testing.T) {
	for _, cas := range []struct {
		name string
		want string
	}{
		{name: "test", want: "github.com/fishedee/tools/cmd/gen"},
	} {
		_ = cas
		modPath, pkgPath := getPkgPathFromDir()
		t.Logf("modPath: %s\n", modPath)
		if pkgPath != cas.want {
			t.Fatalf("No.%s Bad pkgPath: %s != %s\n", cas.name, pkgPath, cas.want)
		}
	}
}
