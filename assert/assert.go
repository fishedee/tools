package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// Equal Equal
func Equal(t *testing.T, left interface{}, right interface{}, testCase ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(left, right) != true {
		t.Errorf("assert equal fail testcase:%v, %v != %v", testCase, left, right)
	}
}

// Error Error
func Error(t *testing.T, errorText string, function func(), testCase ...interface{}) {
	defer func() {
		r := fmt.Sprintf("%+v", recover())
		if r != errorText {
			t.Errorf("testCase:%v , assert fail: %v != %v", testCase, errorText, r)
		}
	}()
	function()
}
