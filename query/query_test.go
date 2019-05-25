package query_test

import (
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
	"os"
)

func TestQueryColumn(t *testing.T) {
	testCase := testdata.GetQueryColumnTestCase()
	for singleTestCaseIndex, singleTestCase := range testCase {
		assert.Equal(t, singleTestCase.Target, singleTestCase.Handler(), singleTestCaseIndex)
	}
}

func init() {
	args := os.Args[len(os.Args)-1]
	if args == "reflect" {
		query.ReflectWarning(false)
	} else if args == "macro" {
		query.ReflectWarning(true)
	} else {
		panic("you should pass args to go test -args")
	}
}
