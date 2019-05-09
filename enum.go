package query

import (
	"reflect"
	"strconv"
	"strings"
)

// EnumData 枚举数据
type EnumData struct {
	ID   int
	Name string
}

// EnumStruct 枚举结构体
type EnumStruct struct {
	names map[string]string
	datas []EnumData
}

// InitEnumStruct 初始化
func InitEnumStruct(this interface{}) {
	enumInfo := reflect.TypeOf(this).Elem()
	enumValue := reflect.ValueOf(this)
	result := enumValue.Elem().FieldByName("EnumStruct").Addr().Interface().(*EnumStruct)
	result.names = map[string]string{}
	result.datas = []EnumData{}

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
		result.datas = append(result.datas, EnumData{
			ID:   singleFieldTagValue,
			Name: singleFieldTagSeeName,
		})
		enumValue.Elem().Field(i).SetInt(int64(singleFieldTagValue))
	}
}

// Names 名字
func (es *EnumStruct) Names() map[string]string {
	return es.names
}

// Entrys 条目
func (es *EnumStruct) Entrys() map[int]string {
	result := map[int]string{}
	for key, value := range es.names {
		singleKey, _ := strconv.Atoi(key)
		result[singleKey] = value
	}
	return result
}

// Datas 数据
func (es *EnumStruct) Datas() []EnumData {
	return es.datas
}

// Keys 键
func (es *EnumStruct) Keys() []int {
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
func (es *EnumStruct) Values() []string {
	result := []string{}
	for _, singleEnum := range es.datas {
		result = append(
			result,
			singleEnum.Name,
		)
	}
	return result
}

// EnumDataString 字符串类型
type EnumDataString struct {
	ID   string
	Name string
}

// EnumStructString 字符串类型
type EnumStructString struct {
	names map[string]string
	datas []EnumDataString
}

// InitEnumStructString 初始化
func InitEnumStructString(this interface{}) {
	enumInfo := reflect.TypeOf(this).Elem()
	enumValue := reflect.ValueOf(this)
	result := enumValue.Elem().FieldByName("EnumStructString").Addr().Interface().(*EnumStructString)
	result.names = map[string]string{}
	result.datas = []EnumDataString{}

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
		result.datas = append(result.datas, EnumDataString{
			ID:   singleFieldTagValue,
			Name: singleFieldTagSeeName,
		})

		enumValue.Elem().Field(i).SetString(singleFieldTagValue)
	}
}

// Names 名字
func (ess *EnumStructString) Names() map[string]string {
	return ess.names
}

// Entrys 条目
func (ess *EnumStructString) Entrys() map[string]string {
	return ess.names
}

// Datas 数据
func (ess *EnumStructString) Datas() []EnumDataString {
	return ess.datas
}

// Keys 键
func (ess *EnumStructString) Keys() []string {
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
func (ess *EnumStructString) Values() []string {
	result := []string{}
	for _, singleEnum := range ess.datas {
		result = append(
			result,
			singleEnum.Name,
		)
	}
	return result
}
