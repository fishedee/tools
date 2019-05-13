package kind

import (
	"reflect"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
)

func TestIsEmptyValue(t *testing.T) {
	testCase := []struct {
		data    interface{}
		isEmpty bool
	}{
		{false, true},
		{true, false},
		{0, true},
		{1, false},
		{0.0, true},
		{0.1, false},
		{"", true},
		{"a", false},
		{time.Time{}, true},
		{time.Now(), false},
	}

	for _, singleTestCase := range testCase {
		dataValue := reflect.ValueOf(singleTestCase.data)
		result := IsEmptyValue(dataValue)
		assert.Equal(t, result, singleTestCase.isEmpty)
	}
}
