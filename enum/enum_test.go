package enum

import (
	"testing"

	"github.com/fishedee/tools/assert"
)

func getEnumStructError(handler func()) (lastErr string) {
	defer func() {
		err := recover()
		if err != nil {
			lastErr = err.(string)
		}
	}()
	handler()
	return ""
}

func TestStruct(t *testing.T) {
	var testCase struct {
		Struct
		ENUM1 int `enum:"1,枚举1"`
		ENUM2 int `enum:"2,枚举2"`
		ENUM3 int `enum:"3,枚举3"`
	}

	InitStruct(&testCase)

	//断言基本枚举值
	assert.Equal(t, testCase.ENUM1, 1)
	assert.Equal(t, testCase.ENUM2, 2)
	assert.Equal(t, testCase.ENUM3, 3)

	//断言函数
	assert.Equal(t, testCase.Names(), map[string]string{
		"1": "枚举1",
		"2": "枚举2",
		"3": "枚举3",
	})
	assert.Equal(t, testCase.Entrys(), map[int]string{
		1: "枚举1",
		2: "枚举2",
		3: "枚举3",
	})
	assert.Equal(t, testCase.Datas(), []Data{
		{1, "枚举1"},
		{2, "枚举2"},
		{3, "枚举3"},
	})
	assert.Equal(t, testCase.Keys(), []int{1, 2, 3})
	assert.Equal(t, testCase.Values(), []string{"枚举1", "枚举2", "枚举3"})

	//错误的enum
	var err string
	var realErr = "invalid enum struct"
	var testCaseErr1 struct {
		Struct
		ENUM1 int `enum:1,枚举1`
		ENUM2 int `enum:"2,枚举2"`
		ENUM3 int `enum:"3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStruct(&testCaseErr1)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)

	var testCaseErr2 struct {
		Struct
		ENUM1 int `enum:"枚举1"`
		ENUM2 int `enum:"2,枚举2"`
		ENUM3 int `enum:"3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStruct(&testCaseErr2)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)

	var testCaseErr3 struct {
		Struct
		ENUM1 int `enum:"1"`
		ENUM2 int `enum:"2,枚举2"`
		ENUM3 int `enum:"3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStruct(&testCaseErr3)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)

	var testCaseErr4 struct {
		Struct
		ENUM1 int `enum:"1z,枚举1"`
		ENUM2 int `enum:"2,枚举2"`
		ENUM3 int `enum:"3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStruct(&testCaseErr4)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)
}

func TestEnumStructString(t *testing.T) {
	var testCase struct {
		StructString
		ENUM1 string `enum:"/content/del1,枚举1"`
		ENUM2 string `enum:"/content/del2,枚举2"`
		ENUM3 string `enum:"/content/del3,枚举3"`
	}

	InitStructString(&testCase)

	//断言基本枚举值
	assert.Equal(t, testCase.ENUM1, "/content/del1")
	assert.Equal(t, testCase.ENUM2, "/content/del2")
	assert.Equal(t, testCase.ENUM3, "/content/del3")

	//断言函数
	assert.Equal(t, testCase.Names(), map[string]string{
		"/content/del1": "枚举1",
		"/content/del2": "枚举2",
		"/content/del3": "枚举3",
	})
	assert.Equal(t, testCase.Entrys(), map[string]string{
		"/content/del1": "枚举1",
		"/content/del2": "枚举2",
		"/content/del3": "枚举3",
	})
	assert.Equal(t, testCase.Datas(), []DataString{
		{"/content/del1", "枚举1"},
		{"/content/del2", "枚举2"},
		{"/content/del3", "枚举3"},
	})
	assert.Equal(t, testCase.Keys(), []string{"/content/del1", "/content/del2", "/content/del3"})
	assert.Equal(t, testCase.Values(), []string{"枚举1", "枚举2", "枚举3"})

	//错误的enum
	var err string
	var realErr = "invalid enum struct"
	var testCaseErr1 struct {
		StructString
		ENUM1 string `enum:/content/del1,枚举1`
		ENUM2 string `enum:"/content/del2,枚举2"`
		ENUM3 string `enum:"/content/del3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStructString(&testCaseErr1)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)

	var testCaseErr2 struct {
		StructString
		ENUM1 string `enum:"/content/del1"`
		ENUM2 string `enum:"/content/del2,枚举2"`
		ENUM3 string `enum:"/content/del3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStructString(&testCaseErr2)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)

	var testCaseErr3 struct {
		StructString
		ENUM1 string `enum:"枚举1"`
		ENUM2 string `enum:"/content/del2,枚举2"`
		ENUM3 string `enum:"/content/del3,枚举3"`
	}
	err = getEnumStructError(func() {
		InitStructString(&testCaseErr3)
	})
	assert.Equal(t, err[0:len(realErr)], realErr)
}
