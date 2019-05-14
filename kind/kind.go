package kind

import (
	"reflect"
	"sync"

	"github.com/fishedee/tools/enum"
	"github.com/fishedee/tools/plode"
)

// TypeKind 类型
var TypeKind struct {
	enum.Struct
	BOOL      int `enum:"1,布尔"`
	INT       int `enum:"2,有符号整数"`
	UINT      int `enum:"3,无符号整数"`
	FLOAT     int `enum:"4,浮点数"`
	PTR       int `enum:"5,指针"`
	STRING    int `enum:"6,字符串"`
	ARRAY     int `enum:"7,数组"`
	MAP       int `enum:"8,映射"`
	STRUCT    int `enum:"9,结构体"`
	INTERFACE int `enum:"10,接口"`
	FUNC      int `enum:"11,函数"`
	CHAN      int `enum:"12,通道"`
	OTHER     int `enum:"13,其他"`
}

func init() {
	enum.InitStruct(&TypeKind)
}

// GetTypeKind 获取类型
func GetTypeKind(t reflect.Type) int {
	// 检查
	if t == nil {
		return TypeKind.OTHER
	}

	switch t.Kind() {
	case reflect.Bool:
		return TypeKind.BOOL
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return TypeKind.INT
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return TypeKind.UINT
	case reflect.Float32, reflect.Float64:
		return TypeKind.FLOAT
	case reflect.Ptr:
		return TypeKind.PTR
	case reflect.String:
		return TypeKind.STRING
	case reflect.Array, reflect.Slice:
		return TypeKind.ARRAY
	case reflect.Map:
		return TypeKind.MAP
	case reflect.Struct:
		return TypeKind.STRUCT
	case reflect.Interface:
		return TypeKind.INTERFACE
	case reflect.Func:
		return TypeKind.FUNC
	case reflect.Chan:
		return TypeKind.CHAN
	default:
		return TypeKind.OTHER
	}
}

type zeroable interface {
	IsZero() bool
}

// IsEmptyValue 是否0值
func IsEmptyValue(v reflect.Value) bool {
	k := v.Interface()
	switch k.(type) {
	case int:
		return k.(int) == 0
	case int8:
		return k.(int8) == 0
	case int16:
		return k.(int16) == 0
	case int32:
		return k.(int32) == 0
	case int64:
		return k.(int64) == 0
	case uint:
		return k.(uint) == 0
	case uint8:
		return k.(uint8) == 0
	case uint16:
		return k.(uint16) == 0
	case uint32:
		return k.(uint32) == 0
	case uint64:
		return k.(uint64) == 0
	case float32:
		return k.(float32) == 0
	case float64:
		return k.(float64) == 0
	case bool:
		return k.(bool) == false
	case string:
		return k.(string) == ""
	case zeroable:
		return k.(zeroable).IsZero()
	}
	return false
}

type getFieldByNameResult struct {
	structField reflect.StructField
	isExist     bool
}

var (
	getFieldByNameCache = map[reflect.Type]map[string]getFieldByNameResult{}
	getFieldByNameMutex = sync.RWMutex{}
)

func getFieldByNameInner(t reflect.Type, name string) (reflect.StructField, bool) {
	nameArray := plode.Explode(name, ".")
	if len(nameArray) == 0 {
		return reflect.StructField{}, false
	}
	var isExist bool
	var resultStruct reflect.StructField
	resultIndex := []int{}
	for _, singleName := range nameArray {
		resultStruct, isExist = t.FieldByName(singleName)
		if !isExist {
			return reflect.StructField{}, false
		}
		resultIndex = append(resultIndex, resultStruct.Index...)
		t = resultStruct.Type
	}
	resultStruct.Index = resultIndex
	return resultStruct, true
}

// GetFieldByName 获取字段
func GetFieldByName(t reflect.Type, name string) (reflect.StructField, bool) {
	getFieldByNameMutex.RLock()
	result, isExist := getFieldByNameCache[t][name]
	getFieldByNameMutex.RUnlock()

	if isExist {
		return result.structField, result.isExist
	}
	result.structField, result.isExist = getFieldByNameInner(t, name)

	getFieldByNameMutex.Lock()
	typeInfo, isExist := getFieldByNameCache[t]
	if !isExist {
		typeInfo = map[string]getFieldByNameResult{}
	}
	typeInfo[name] = result
	getFieldByNameCache[t] = typeInfo
	getFieldByNameMutex.Unlock()

	return result.structField, result.isExist
}
