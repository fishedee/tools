package exception_test

import (
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/exception"
	"github.com/fishedee/tools/plode"
)

func getCatchMessage(fun func()) (_last string) {
	defer exception.Catch(func(e exception.Exception) {
		_last = e.GetMessage()
	})
	fun()
	return ""
}

func getCatchCrashMessage(fun func()) (_last string) {
	defer exception.CatchCrash(func(e exception.Exception) {
		_last = e.GetMessage()
	})
	fun()
	return ""
}

type errorStruct struct {
}

func (es *errorStruct) Error() string {
	return "m2"
}

func TestCatch(t *testing.T) {
	testCase := []struct {
		origin func()
		target string
	}{
		{func() {
			panic("m1")
		}, "m1"},
		{func() {
			panic(&errorStruct{})
		}, "m2"},
		{func() {
			exception.Throw(1, "m3")
		}, "m3"},
	}

	for singleIndex, singleTestCase := range testCase {
		msg := getCatchCrashMessage(singleTestCase.origin)
		assert.Equal(t, msg, singleTestCase.target, singleIndex)
	}
}

func getLastStackTraceLine(e exception.Exception) string {
	lines := plode.Explode(e.GetStackTraceLine(0), "/")
	return lines[len(lines)-1]
}

func TestCatchStack1(t *testing.T) {
	defer exception.Catch(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:65")
	})
	exception.Throw(1, "test1")
}

func TestCatchStack2(t *testing.T) {
	defer exception.CatchCrash(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:72")
	})
	exception.Throw(1, "test2")
}

func TestCatchStack3(t *testing.T) {
	defer exception.CatchCrash(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:79")
	})
	panic("test3")
}

func TestCatchStack4(t *testing.T) {
	defer exception.CatchCrash(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:89")
	})
	defer exception.Catch(func(e exception.Exception) {
		assert.Equal(t, "should not be here!", false)
	})
	panic("test4")
}

func TestCatchStack5(t *testing.T) {
	defer exception.CatchCrash(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:99")
	})
	defer exception.Catch(func(e exception.Exception) {
		panic(&e)
	})
	exception.Throw(1, "test5")
}

func TestCatchStack6(t *testing.T) {
	defer exception.Catch(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:109")
	})
	defer exception.Catch(func(e exception.Exception) {
		panic(&e)
	})
	exception.Throw(1, "test6")
}

func TestCatchStack7(t *testing.T) {
	defer exception.CatchCrash(func(e exception.Exception) {
		assert.Equal(t, getLastStackTraceLine(e), "exception_test.go:119")
	})
	defer exception.CatchCrash(func(e exception.Exception) {
		panic(&e)
	})
	panic("test7")
}
