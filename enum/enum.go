package enum

import (
	"reflect"
	"strconv"
	"strings"
)

// Data 枚举数据
type Data struct {
	ID   int
	Name string
}

// Struct 枚举结构体
type Struct struct {
	names map[string]string
	datas []Data
}

// InitStruct 初始化
func InitStruct(this interface{}) {
	enumInfo := reflect.TypeOf(this).Elem()
	enumValue := reflect.ValueOf(this)
	result := enumValue.Elem().FieldByName("Struct").Addr().Interface().(*Struct)
	result.names = map[string]string{}
	result.datas = []Data{}

	for i := 0; i != enumInfo.NumField(); i++ {
		singleField := enumInfo.Field(i)
		if singleField.PkgPath != "" || singleField.Anonymous {
			continue
		}

		singleFieldName := singleField.Name
		singleFieldTag := singleField.Tag.Get("enum")
		singleFieldTagArray := strings.Split(singleFieldTag, ",")
		if len(singleFieldTagArray) != 2 {
			panic("invalid enum " + enumInfo.String() + ":" + singleFieldName)
		}

		singleFieldTagValue, err := strconv.Atoi(singleFieldTagArray[0])
		if err != nil {
			panic("invalid enum " + enumInfo.String() + ":" + singleFieldName)
		}
		singleFieldTagSeeName := singleFieldTagArray[1]
		if singleFieldTagSeeName == "" {
			panic("invalid enum " + enumInfo.String() + ":" + singleFieldName)
		}

		result.names[singleFieldTagArray[0]] = singleFieldTagSeeName
		result.datas = append(result.datas, Data{
			ID:   singleFieldTagValue,
			Name: singleFieldTagSeeName,
		})
		enumValue.Elem().Field(i).SetInt(int64(singleFieldTagValue))
	}
}

// Names 名字
func (es *Struct) Names() map[string]string {
	return es.names
}

// Entrys 条目
func (es *Struct) Entrys() map[int]string {
	result := map[int]string{}
	for key, value := range es.names {
		singleKey, _ := strconv.Atoi(key)
		result[singleKey] = value
	}
	return result
}

// Datas 数据
func (es *Struct) Datas() []Data {
	return es.datas
}

// Keys 键
func (es *Struct) Keys() []int {
	result := []int{}
	for _, singleEnum := range es.datas {
		result = append(
			result,
			singleEnum.ID,
		)
	}
	return result
}

// Values 值
func (es *Struct) Values() []string {
	result := []string{}
	for _, singleEnum := range es.datas {
		result = append(
			result,
			singleEnum.Name,
		)
	}
	return result
}

// DataString 字符串类型
type DataString struct {
	ID   string
	Name string
}

// StructString 字符串类型
type StructString struct {
	names map[string]string
	datas []DataString
}

// InitStructString 初始化
func InitStructString(this interface{}) {
	enumInfo := reflect.TypeOf(this).Elem()
	enumValue := reflect.ValueOf(this)
	result := enumValue.Elem().FieldByName("StructString").Addr().Interface().(*StructString)
	result.names = map[string]string{}
	result.datas = []DataString{}

	for i := 0; i != enumInfo.NumField(); i++ {
		singleField := enumInfo.Field(i)
		if singleField.PkgPath != "" || singleField.Anonymous {
			continue
		}
		singleFieldName := singleField.Name

		singleFieldTag := singleField.Tag.Get("enum")
		singleFieldTagArray := strings.Split(singleFieldTag, ",")
		if len(singleFieldTagArray) != 2 {
			panic("invalid enum " + enumInfo.String() + ":" + singleFieldName)
		}

		singleFieldTagValue := singleFieldTagArray[0]
		singleFieldTagSeeName := singleFieldTagArray[1]
		if singleFieldTagValue == "" || singleFieldTagSeeName == "" {
			panic("invalid enum " + enumInfo.String() + ":" + singleFieldName)
		}

		result.names[singleFieldTagValue] = singleFieldTagSeeName
		result.datas = append(result.datas, DataString{
			ID:   singleFieldTagValue,
			Name: singleFieldTagSeeName,
		})

		enumValue.Elem().Field(i).SetString(singleFieldTagValue)
	}
}

// Names 名字
func (ess *StructString) Names() map[string]string {
	return ess.names
}

// Entrys 条目
func (ess *StructString) Entrys() map[string]string {
	return ess.names
}

// Datas 数据
func (ess *StructString) Datas() []DataString {
	return ess.datas
}

// Keys 键
func (ess *StructString) Keys() []string {
	result := []string{}
	for _, singleEnum := range ess.datas {
		result = append(
			result,
			singleEnum.ID,
		)
	}
	return result
}

// Values 值
func (ess *StructString) Values() []string {
	result := []string{}
	for _, singleEnum := range ess.datas {
		result = append(
			result,
			singleEnum.Name,
		)
	}
	return result
}
