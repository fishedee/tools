package quicktag

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// QuickTag 快速标签
type QuickTag struct {
	tag       string
	cache     sync.Map
	mutex     sync.Mutex
	basicType map[reflect.Kind]bool
}

// NewQuickTag 新建
func NewQuickTag(tag string) *QuickTag {
	basicTypeMap := map[reflect.Kind]bool{}

	basicType := []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
		reflect.Chan,
		reflect.Func,
	}

	for _, kind := range basicType {
		basicTypeMap[kind] = true
	}
	return &QuickTag{
		tag:       tag,
		cache:     sync.Map{},
		mutex:     sync.Mutex{},
		basicType: basicTypeMap,
	}
}

// GetTagType 获取标签类型
func (qt *QuickTag) GetTagType(t reflect.Type) reflect.Type {
	result, isExist := qt.cache.Load(t)
	if isExist {
		return *(result.(*reflect.Type))
	}

	qt.mutex.Lock()
	defer qt.mutex.Unlock()

	hasVisit := map[reflect.Type]bool{}
	return qt.getTagTypeInner(hasVisit, t)
}

func (qt *QuickTag) getTagTypeInner(hasVisit map[reflect.Type]bool, t reflect.Type) reflect.Type {
	timeType := reflect.TypeOf(time.Time{})

	cacheType, isExist := qt.cache.Load(t)
	if isExist {
		return *(cacheType.(*reflect.Type))
	}

	hasVisit[t] = true
	var resultType reflect.Type

	kind := t.Kind()
	if qt.basicType[kind] == true {
		resultType = t
	} else if kind == reflect.Ptr {
		tempType := qt.getTagTypeInner(hasVisit, t.Elem())
		resultType = reflect.PtrTo(tempType)
	} else if kind == reflect.Array {
		tempType := qt.getTagTypeInner(hasVisit, t.Elem())
		resultType = reflect.ArrayOf(t.Len(), tempType)
	} else if kind == reflect.Slice {
		tempType := qt.getTagTypeInner(hasVisit, t.Elem())
		resultType = reflect.SliceOf(tempType)
	} else if kind == reflect.Map {
		tempType := qt.getTagTypeInner(hasVisit, t.Elem())
		resultType = reflect.MapOf(t.Key(), tempType)
	} else if kind == reflect.Interface {
		panic(fmt.Sprintf("quick tag dosnot support interface %v", t))
	} else if t == timeType {
		resultType = reflect.TypeOf(myTime{})
	} else if kind == reflect.Struct {
		resultType = qt.getStructType(hasVisit, t)
	} else {
		panic(fmt.Sprintf("unknown kind %v", kind))
	}

	hasVisit[t] = false
	qt.cache.Store(t, &resultType)

	return resultType
}

func (qt *QuickTag) getStructType(hasVisit map[reflect.Type]bool, t reflect.Type) reflect.Type {
	numField := t.NumField()
	newStructFields := []reflect.StructField{}

	for i := 0; i != numField; i++ {
		field := t.Field(i)
		if hasVisit[field.Type] == true {
			panic(fmt.Sprintf("quick tag can not support circle type %v->%v", t, field.Type))
		}
		fieldType := qt.getTagTypeInner(hasVisit, field.Type)
		fieldName := field.Name
		fieldAnonymous := field.Anonymous
		fieldTag := qt.getTag(field)
		newStructFields = append(newStructFields, reflect.StructField{
			Name:      fieldName,
			Type:      fieldType,
			Tag:       fieldTag,
			Anonymous: fieldAnonymous,
		})
	}
	return reflect.StructOf(newStructFields)
}

func (qt *QuickTag) getTag(field reflect.StructField) reflect.StructTag {
	tag := field.Tag
	firstName := strings.ToLower(field.Name[0:1]) + field.Name[1:]
	secondSet := ""
	originInfo, hasOriginTag := tag.Lookup(qt.tag)
	if hasOriginTag == true {
		originInfoList := strings.Split(originInfo, ",")
		if originInfoList[0] != "" {
			firstName = originInfoList[0]
		}
		if len(originInfoList) >= 2 && originInfoList[1] != "" {
			secondSet = originInfoList[1]
		}
	}
	var result = ""
	if secondSet == "" {
		result = fmt.Sprintf("%s:\"%s\"", qt.tag, firstName)
	} else {
		result = fmt.Sprintf("%s:\"%s,%s\"", qt.tag, firstName, secondSet)
	}
	return reflect.StructTag(result)
}

type emptyInterface struct {
	pt unsafe.Pointer
	pv unsafe.Pointer
}

func (qt *QuickTag) pointerOfType(t reflect.Type) unsafe.Pointer {
	p := *(*emptyInterface)(unsafe.Pointer(&t))
	return p.pv
}

// GetTagInstance 生成有tag实例
func (qt *QuickTag) GetTagInstance(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	srcType := reflect.TypeOf(src)
	eface := *(*emptyInterface)(unsafe.Pointer(&src))
	eface.pt = qt.pointerOfType(qt.GetTagType(srcType))
	dst := *(*interface{})(unsafe.Pointer(&eface))
	return dst
}
