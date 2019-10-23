package quicktag

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
)

var (
	jsonQuickTag *QuickTag
)

func jsonMarshal(data interface{}) ([]byte, error) {
	quickTagInstance := jsonQuickTag.GetTagInstance(data)

	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	err := encoder.Encode(quickTagInstance)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func jsonUnmarshal(in []byte, data interface{}) error {
	quickTagInstance := jsonQuickTag.GetTagInstance(data)

	err := json.Unmarshal(in, quickTagInstance)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	jsonQuickTag = NewQuickTag("json")
}

type User struct {
	UserID     int
	Name       string
	CreateTime time.Time
}

type Users struct {
	Count int
	Data  []User
}

type Class struct {
	ClassID  int
	Name     string
	Students []User
	Score    int `json:"sss"`
	Level    int `json:"-"`
}

func TestNil(t *testing.T) {
	var data interface{}

	data = nil

	assert.Equal(t, jsonQuickTag.GetTagInstance(data), nil)
}

func TestMarshalAndUnmarshal(t *testing.T) {
	testCases := []struct {
		data interface{}
		str  string
	}{
		{
			User{1, "fish", time.Unix(0, 0)},
			`{"userID":1,"name":"fish","createTime":"1970-01-01 08:00:00"}`,
		},
		{
			Users{
				Count: 2,
				Data: []User{
					User{3, "fish", time.Unix(1, 0)},
					User{4, "cat", time.Unix(2, 0)},
				},
			},
			`{"count":2,"data":[{"userID":3,"name":"fish","createTime":"1970-01-01 08:00:01"},{"userID":4,"name":"cat","createTime":"1970-01-01 08:00:02"}]}`,
		},
		{
			Class{
				ClassID: 5,
				Name:    "class1",
				Students: []User{
					User{6, "dog", time.Unix(3, 0)},
					User{7, "apple", time.Unix(4, 0)},
				},
				Score: 78,
			},
			`{"classID":5,"name":"class1","students":[{"userID":6,"name":"dog","createTime":"1970-01-01 08:00:03"},{"userID":7,"name":"apple","createTime":"1970-01-01 08:00:04"}],"sss":78}`,
		},
	}

	//序列化
	for _, singleTestCase := range testCases {
		str, err := jsonMarshal(singleTestCase.data)
		assert.Equal(t, err, nil)
		assert.Equal(t, string(str), singleTestCase.str+"\n")
	}

	//反序列化
	for _, singleTestCase := range testCases {
		typ := reflect.TypeOf(singleTestCase.data)
		temp := reflect.New(typ).Interface()
		err := jsonUnmarshal([]byte(singleTestCase.str), temp)
		assert.Equal(t, err, nil)
		assert.Equal(t, reflect.ValueOf(temp).Elem().Interface(), singleTestCase.data)
	}
}
