package plode

import (
	"testing"

	"github.com/donnol/tools/assert"
)

func TestPlode(t *testing.T) {
	testCase := []struct {
		origin    string
		seperator string
		data      []string
	}{
		{"", ",", []string{}},
		{"mmx", ",", []string{"mmx"}},
		{"mmx,mmd,xxu", ",", []string{"mmx", "mmd", "xxu"}},
		{"mmx_mmd_xxu", "_", []string{"mmx", "mmd", "xxu"}},
	}

	//test explode
	for _, singleTestCase := range testCase {
		result := Explode(singleTestCase.origin, singleTestCase.seperator)
		assert.Equal(t, result, singleTestCase.data, singleTestCase)
	}

	//test implode
	for _, singleTestCase := range testCase {
		singleOrigin := Implode(singleTestCase.data, singleTestCase.seperator)
		assert.Equal(t, singleOrigin, singleTestCase.origin, singleTestCase)
	}
}

func TestPlodeInt(t *testing.T) {
	testCase := []struct {
		origin    string
		seperator string
		data      []int
	}{
		{"", ",", []int{}},
		{"1", ",", []int{1}},
		{"1,2,3", ",", []int{1, 2, 3}},
		{"11_22_33", "_", []int{11, 22, 33}},
	}

	//test explode
	for _, singleTestCase := range testCase {
		result := ExplodeInt(singleTestCase.origin, singleTestCase.seperator)
		assert.Equal(t, result, singleTestCase.data, singleTestCase)
	}

	//test implode
	for _, singleTestCase := range testCase {
		singleOrigin := ImplodeInt(singleTestCase.data, singleTestCase.seperator)
		assert.Equal(t, singleOrigin, singleTestCase.origin, singleTestCase)
	}
}
