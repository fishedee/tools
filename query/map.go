package query

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fishedee/tools/decimal"
	"github.com/fishedee/tools/kind"
)

var decimalType = reflect.TypeOf(decimal.Decimal("0"))

func nameMapper(name string) string {
	return strings.ToLower(name[0:1]) + name[1:]
}

func combileMap(result map[string]interface{}, singleResultMap reflect.Value) {
	singleResultMapType := singleResultMap.Type()
	if singleResultMapType.Kind() != reflect.Map {
		return
	}
	singleResultMapKeys := singleResultMap.MapKeys()
	for _, singleKey := range singleResultMapKeys {
		singleResultKey := fmt.Sprintf("%v", singleKey)
		_, isExist := result[singleResultKey]
		if isExist {
			continue
		}
		result[singleResultKey] = singleResultMap.MapIndex(singleKey).Interface()
	}
}

type arrayMappingStructInfo struct {
	name      string
	omitempty bool
	anonymous bool
	canRead   bool
	canWrite  bool
	index     []int
}

type arrayMappingInfo struct {
	kind       int
	isTimeType bool
	field      []arrayMappingStructInfo
}

var arrayMappingInfoMap struct {
	mutex sync.RWMutex
	data  map[string]map[reflect.Type]arrayMappingInfo
}

var interfaceType reflect.Type

func getDataTagInfoInner(dataType reflect.Type, tag string) arrayMappingInfo {
	dataTypeKind := kind.GetTypeKind(dataType)
	result := arrayMappingInfo{}
	result.kind = dataTypeKind
	if dataTypeKind == kind.TypeKind.STRUCT {
		if dataType == reflect.TypeOf(time.Time{}) {
			//时间类型
			result.isTimeType = true
		} else {
			//结构体类型
			result.isTimeType = false
			anonymousField := []arrayMappingStructInfo{}
			noanonymousField := []arrayMappingStructInfo{}
			for i := 0; i != dataType.NumField(); i++ {
				singleDataType := dataType.Field(i)
				if singleDataType.PkgPath != "" && singleDataType.Anonymous == false {
					continue
				}
				var singleName string
				var omitempty bool
				var canRead bool
				var canWrite bool
				singleName = nameMapper(singleDataType.Name)
				canRead = true
				canWrite = true
				omitempty = false

				jsonTag := singleDataType.Tag.Get(tag)
				jsonTagList := strings.Split(jsonTag, ",")
				for singleTagIndex, singleTag := range jsonTagList {
					if singleTag == "-" {
						canRead = false
						canWrite = false
					} else if singleTag == "->" {
						canRead = false
						canWrite = true
					} else if singleTag == "<-" {
						canRead = true
						canWrite = false
					} else if singleTagIndex == 0 && singleTag != "" {
						singleName = singleTag
					} else if singleTagIndex == 1 && singleTag == "omitempty" {
						omitempty = true
					}
				}
				single := arrayMappingStructInfo{}
				single.name = singleName
				single.omitempty = omitempty
				single.canRead = canRead
				single.canWrite = canWrite
				single.index = singleDataType.Index
				single.anonymous = singleDataType.Anonymous
				if singleDataType.Anonymous {
					anonymousField = append(anonymousField, single)
				} else {
					noanonymousField = append(noanonymousField, single)
				}
			}
			result.field = append(noanonymousField, anonymousField...)
		}
	}
	return result
}

func getDataTagInfo(target reflect.Type, tag string) arrayMappingInfo {
	arrayMappingInfoMap.mutex.RLock()
	var result arrayMappingInfo
	var ok bool
	resultArray, okArray := arrayMappingInfoMap.data[tag]
	if okArray {
		result, ok = resultArray[target]
	}
	arrayMappingInfoMap.mutex.RUnlock()

	if ok {
		return result
	}
	result = getDataTagInfoInner(target, tag)

	arrayMappingInfoMap.mutex.Lock()
	if !okArray {
		resultArray = map[reflect.Type]arrayMappingInfo{}
		arrayMappingInfoMap.data[tag] = resultArray
	}
	resultArray[target] = result
	arrayMappingInfoMap.mutex.Unlock()

	return result
}

func arrayToMapInner(dataValue reflect.Value, tag string) (reflect.Value, bool) {
	if dataValue.IsValid() == false {
		return dataValue, true
	}

	var result reflect.Value
	var isEmpty bool
	dataType := getDataTagInfo(dataValue.Type(), tag)
	if dataType.kind == kind.TypeKind.STRUCT && dataType.isTimeType == true {
		timeValue := dataValue.Interface().(time.Time)
		result = reflect.ValueOf(timeValue.Format("2006-01-02 15:04:05"))
		isEmpty = kind.IsEmptyValue(dataValue)
	} else if dataType.kind == kind.TypeKind.STRUCT && dataType.isTimeType == false {
		resultMap := map[string]interface{}{}
		for _, singleType := range dataType.field {
			if singleType.canWrite == false {
				continue
			}
			singleResultMap, isEmptyValue := arrayToMapInner(dataValue.FieldByIndex(singleType.index), tag)
			if singleType.anonymous == false {
				if singleType.omitempty == true && isEmptyValue {
					continue
				}
				if singleResultMap.IsValid() == false {
					continue
				}
				resultMap[singleType.name] = singleResultMap.Interface()
			} else {
				combileMap(resultMap, singleResultMap)
			}
		}
		result = reflect.ValueOf(resultMap)
		isEmpty = (len(resultMap) == 0)
	} else if dataType.kind == kind.TypeKind.ARRAY {
		resultSlice := []interface{}{}
		dataLen := dataValue.Len()
		for i := 0; i != dataLen; i++ {
			singleDataValue := dataValue.Index(i)
			singleDataResult, _ := arrayToMapInner(singleDataValue, tag)
			resultSlice = append(resultSlice, singleDataResult.Interface())
		}
		result = reflect.ValueOf(resultSlice)
		isEmpty = (len(resultSlice) == 0)
	} else if dataType.kind == kind.TypeKind.MAP {
		dataKeyType := dataValue.Type().Key()
		resultMapType := reflect.MapOf(dataKeyType, interfaceType)
		resultMap := reflect.MakeMap(resultMapType)
		dataKeys := dataValue.MapKeys()
		for _, singleDataKey := range dataKeys {
			singleDataValue := dataValue.MapIndex(singleDataKey)
			singleDataResult, _ := arrayToMapInner(singleDataValue, tag)
			resultMap.SetMapIndex(singleDataKey, singleDataResult)
		}
		result = resultMap
		isEmpty = (len(dataKeys) == 0)
	} else if dataType.kind == kind.TypeKind.INTERFACE ||
		dataType.kind == kind.TypeKind.PTR {
		result, isEmpty = arrayToMapInner(dataValue.Elem(), tag)
	} else if dataType.kind == kind.TypeKind.STRING && dataValue.Type() == decimalType {
		decimalString := dataValue.Interface().(decimal.Decimal).String()
		result = reflect.ValueOf(decimalString)
		isEmpty = (decimalString == "0")
	} else {
		result = dataValue
		isEmpty = kind.IsEmptyValue(dataValue)
	}
	return result, isEmpty
}

// ArrayToMap 数组转映射
func ArrayToMap(data interface{}, tag string) interface{} {
	dataValue, _ := arrayToMapInner(reflect.ValueOf(data), tag)
	if dataValue.IsValid() == false {
		return nil
	}

	return dataValue.Interface()
}

func mapToBool(dataValue reflect.Value, target reflect.Value) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind == kind.TypeKind.BOOL {
		target.SetBool(dataValue.Bool())
		return nil
	} else if dataKind == kind.TypeKind.STRING {
		dataBool, err := strconv.ParseBool(dataValue.String())
		if err != nil {
			return fmt.Errorf("不是布尔值，其值为[%s]", dataValue.String())
		}
		target.SetBool(dataBool)
		return nil
	} else {
		return fmt.Errorf("不是布尔值，其类型为[%s]", dataValue.Type().String())
	}
}

func mapToUint(dataValue reflect.Value, target reflect.Value) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind == kind.TypeKind.UINT {
		target.SetUint(dataValue.Uint())
		return nil
	} else if dataKind == kind.TypeKind.INT {
		target.SetUint(uint64(dataValue.Int()))
		return nil
	} else if dataKind == kind.TypeKind.FLOAT {
		target.SetUint(uint64(math.Floor(dataValue.Float() + 0.5)))
		return nil
	} else if dataKind == kind.TypeKind.STRING {
		dataUint, err := strconv.ParseUint(dataValue.String(), 10, 64)
		if err != nil {
			return fmt.Errorf("不是无符号整数，其值为[%s]", dataValue.String())
		}
		target.SetUint(dataUint)
		return nil
	} else {
		return fmt.Errorf("不是无符号整数，其类型为[%s]", dataValue.Type().String())
	}
}

func mapToInt(dataValue reflect.Value, target reflect.Value) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind == kind.TypeKind.INT {
		target.SetInt(dataValue.Int())
		return nil
	} else if dataKind == kind.TypeKind.UINT {
		target.SetInt(int64(dataValue.Uint()))
		return nil
	} else if dataKind == kind.TypeKind.FLOAT {
		target.SetInt(int64(math.Floor(dataValue.Float() + 0.5)))
		return nil
	} else if dataKind == kind.TypeKind.STRING {
		dataInt, err := strconv.ParseInt(dataValue.String(), 10, 64)
		if err != nil {
			return fmt.Errorf("不是整数，其值为[%s]", dataValue.String())
		}
		target.SetInt(dataInt)
		return nil
	} else {
		return fmt.Errorf("不是整数，其类型为[%s]", dataValue.Type().String())
	}
}

func mapToFloat(dataValue reflect.Value, target reflect.Value) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind == kind.TypeKind.FLOAT {
		target.SetFloat(dataValue.Float())
		return nil
	} else if dataKind == kind.TypeKind.INT {
		target.SetFloat(float64(dataValue.Int()))
		return nil
	} else if dataKind == kind.TypeKind.UINT {
		target.SetFloat(float64(dataValue.Uint()))
		return nil
	} else if dataKind == kind.TypeKind.STRING {
		dataFloat, err := strconv.ParseFloat(dataValue.String(), 64)
		if err != nil {
			return fmt.Errorf("不是浮点数，其值为[%s]", dataValue.String())
		}
		target.SetFloat(dataFloat)
		return nil
	} else {
		return fmt.Errorf("不是浮点数，其类型为[%s]", dataValue.Type().String())
	}
}

func mapToString(dataValue reflect.Value, target reflect.Value) error {
	stringValue := fmt.Sprintf("%v", dataValue.Interface())
	if target.Type() == decimalType {
		_, err := decimal.NewDecimal(stringValue)
		if err != nil {
			return fmt.Errorf("不是十进制数字，其值为[%s]", stringValue)
		}
		target.SetString(stringValue)
		return nil
	}

	target.SetString(stringValue)
	return nil
}

func mapToArray(dataValue reflect.Value, target reflect.Value, tag string) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind != kind.TypeKind.ARRAY {
		return fmt.Errorf("不是数组，其类型为[%s]", dataValue.Type().String())
	}
	//增长空间
	dataLen := dataValue.Len()
	targetType := target.Type()
	targetLen := target.Len()
	if targetType.Kind() == reflect.Slice {
		if target.IsNil() == true {
			var newTarget reflect.Value
			newTarget = reflect.MakeSlice(targetType, dataLen, dataLen)
			target.Set(newTarget)
		} else if targetLen != dataLen {
			var newTarget reflect.Value
			newTarget = reflect.MakeSlice(targetType, dataLen, dataLen)
			reflect.Copy(newTarget, target)
			target.Set(newTarget)
		}
		targetLen = dataLen
	}
	//复制数据
	for i := 0; i != targetLen; i++ {
		if i >= dataLen {
			targetElemType := targetType.Elem()
			zeroElemType := reflect.Zero(targetElemType)
			for i := dataLen; i < targetLen; i++ {
				target.Index(i).Set(zeroElemType)
			}
			break
		} else {
			singleData := dataValue.Index(i)
			singleDataTarget := target.Index(i)
			err := mapToArrayInner(singleData, singleDataTarget, tag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func mapToMap(dataValue reflect.Value, target reflect.Value, tag string) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind != kind.TypeKind.MAP {
		return fmt.Errorf("不是映射，其类型为[%s]", dataValue.Type().String())
	}
	dataKeys := dataValue.MapKeys()
	targetType := target.Type()
	targetKeyType := targetType.Key()
	targetValueType := targetType.Elem()
	if target.IsNil() == true {
		var newTarget reflect.Value
		newTarget = reflect.MakeMap(targetType)
		target.Set(newTarget)
	}
	for _, singleDataKey := range dataKeys {
		singleDataTargetKey := reflect.New(targetKeyType).Elem()
		err := mapToArrayInner(singleDataKey, singleDataTargetKey, tag)
		if err != nil {
			return err
		}

		singleDataValue := dataValue.MapIndex(singleDataKey)
		singleDataTargetValue := reflect.New(targetValueType).Elem()
		singleDataTargetValueOld := target.MapIndex(singleDataTargetKey)
		if singleDataTargetValueOld.IsValid() == true {
			singleDataTargetValue.Set(singleDataTargetValueOld)
		}
		err = mapToArrayInner(singleDataValue, singleDataTargetValue, tag)
		if err != nil {
			return fmt.Errorf("参数%s%s", singleDataKey, err.Error())
		}
		target.SetMapIndex(singleDataTargetKey, singleDataTargetValue)
	}
	return nil
}

func mapToTime(dataValue reflect.Value, target reflect.Value) error {
	dataType := dataValue.Type()
	if dataType == reflect.TypeOf(time.Time{}) {
		target.Set(dataValue)
	} else if dataType.Kind() == reflect.String {
		timeValue, err := time.ParseInLocation("2006-01-02 15:04:05", dataValue.String(), time.Now().Local().Location())
		if err != nil {
			return fmt.Errorf("不是时间，其值为[%s]", dataValue.String())
		}
		target.Set(reflect.ValueOf(timeValue))
		return nil
	}
	return fmt.Errorf("不是时间，其类型为[%s]", dataValue.Type().String())
}

func mapToStruct(dataValue reflect.Value, target reflect.Value, targetType arrayMappingInfo, tag string) error {
	dataType := dataValue.Type()
	dataKind := kind.GetTypeKind(dataType)
	if dataKind != kind.TypeKind.MAP {
		return fmt.Errorf("不是映射，其类型为[%s]", dataValue.Type().String())
	}
	dataTypeKey := dataType.Key()
	for _, singleStructInfo := range targetType.field {
		if singleStructInfo.canRead == false {
			continue
		}
		if singleStructInfo.anonymous == true {
			//FIXME 暂不考虑匿名结构体的覆盖问题
			singleDataValue := target.FieldByIndex(singleStructInfo.index)
			err := mapToArrayInner(dataValue, singleDataValue, tag)
			if err != nil {
				return fmt.Errorf("参数%s%s", singleStructInfo.name, err.Error())
			}
		} else {
			singleMapKey := reflect.New(dataTypeKey)
			singleDataKey := reflect.ValueOf(singleStructInfo.name)
			err := mapToArrayInner(singleDataKey, singleMapKey, tag)
			if err != nil {
				return err
			}

			singleDataValue := target.FieldByIndex(singleStructInfo.index)
			singleMapResult := dataValue.MapIndex(singleMapKey.Elem())
			if singleMapResult.IsValid() == false {
				continue
			}
			err = mapToArrayInner(singleMapResult, singleDataValue, tag)
			if err != nil {
				return fmt.Errorf("参数%s%s", singleDataKey, err.Error())
			}
		}
	}
	return nil
}

func mapToPtr(dataValue reflect.Value, target reflect.Value, tag string) error {
	targetElem := target.Elem()
	if targetElem.IsValid() == false {
		targetElem = reflect.New(target.Type().Elem())
		target.Set(targetElem)
	}
	return mapToArrayInner(dataValue, targetElem, tag)
}

func mapToInterface(dataValue reflect.Value, target reflect.Value, tag string) error {
	targetElem := target.Elem()
	if targetElem.IsValid() == false {
		target.Set(dataValue)
		return nil
	}
	newTargetElem := reflect.New(targetElem.Type()).Elem()
	newTargetElem.Set(targetElem)
	err := mapToArrayInner(dataValue, newTargetElem, tag)
	if err != nil {
		return err
	}
	target.Set(newTargetElem)
	return nil
}

func mapToArrayInner(data reflect.Value, target reflect.Value, tag string) error {
	//处理data是个nil的问题
	if data.IsValid() == false {
		target.Set(reflect.Zero(target.Type()))
		return nil
	}
	//处理data是多层嵌套的问题
	dataKind := data.Type().Kind()
	if dataKind == reflect.Interface {
		return mapToArrayInner(data.Elem(), target, tag)
	} else if dataKind == reflect.Ptr {
		return mapToArrayInner(data.Elem(), target, tag)
	}
	//根据target是多层嵌套的问题
	targetType := getDataTagInfo(target.Type(), tag)
	if targetType.kind == kind.TypeKind.PTR {
		return mapToPtr(data, target, tag)
	} else if targetType.kind == kind.TypeKind.INTERFACE {
		return mapToInterface(data, target, tag)
	}
	//处理data是个空字符串
	if dataKind == reflect.String && data.String() == "" {
		target.Set(reflect.Zero(target.Type()))
		return nil
	}
	if targetType.kind == kind.TypeKind.BOOL {
		return mapToBool(data, target)
	} else if targetType.kind == kind.TypeKind.INT {
		return mapToInt(data, target)
	} else if targetType.kind == kind.TypeKind.UINT {
		return mapToUint(data, target)
	} else if targetType.kind == kind.TypeKind.FLOAT {
		return mapToFloat(data, target)
	} else if targetType.kind == kind.TypeKind.STRING {
		return mapToString(data, target)
	} else if targetType.kind == kind.TypeKind.ARRAY {
		return mapToArray(data, target, tag)
	} else if targetType.kind == kind.TypeKind.MAP {
		return mapToMap(data, target, tag)
	} else if targetType.kind == kind.TypeKind.STRUCT {
		if targetType.isTimeType {
			return mapToTime(data, target)
		}
		return mapToStruct(data, target, targetType, tag)
	} else {
		return errors.New("unkown target type " + target.Type().String())
	}
}

// MapToArray 映射转数组
func MapToArray(data interface{}, target interface{}, tag string) error {
	dataValue := reflect.ValueOf(data)
	targetValue := reflect.ValueOf(target)
	if targetValue.IsValid() == false {
		return errors.New("target is nil")
	}
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("invalid target is not ptr")
	}
	return mapToArrayInner(dataValue, targetValue, tag)
}

func init() {
	arrayMappingInfoMap.data = map[string]map[reflect.Type]arrayMappingInfo{}
	var mm struct {
		Test interface{}
	}
	interfaceType = reflect.TypeOf(mm).Field(0).Type
}
