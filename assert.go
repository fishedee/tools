package query

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertEqual AssertEqual
func AssertEqual(t *testing.T, left interface{}, right interface{}, testCase ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(left, right) != true {
		t.Errorf("assert equal fail testcase:%v, %v != %v", testCase, left, right)
	}
}

// AssertError AssertError
func AssertError(t *testing.T, errorText string, function func(), testCase ...interface{}) {
	defer func() {
		r := fmt.Sprintf("%+v", recover())
		if r != errorText {
			t.Errorf("testCase:%v , assert fail: %v != %v", testCase, errorText, r)
		}
	}()
	function()
}
