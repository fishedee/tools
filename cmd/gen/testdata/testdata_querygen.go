package testdata

import (
	"github.com/fishedee/tools/cmd/gen/testdata/subtest"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
	"time"
)

func queryColumnMapV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[string]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV0278bd21d25efe9d7f53cc03a127d039c73e98a2(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.QueryInnerStruct2)
	result := make(map[int]testdata.QueryInnerStruct2, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].QueryInnerStruct.MM] = dataIn[i]
	}
	return result
}

func queryColumnMapV1a5b7250371597524e364f0c816390c77a8b3331(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[string]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV268d58dff08fb0947b9b47bcae328d584ec43d6c(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[bool]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Ok] = dataIn[i]
	}
	return result
}

func queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41(data interface{}, column string) interface{} {
	dataIn := data.([]string)
	result := make(map[string]string, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[float32]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Money] = dataIn[i]
	}
	return result
}

func queryColumnMapV74382f13397675a69df627f3964ab4362f62b343(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make(map[string]User, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV904b262f8e2329ec73c320ca0e5ca82f14165586(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make(map[int]int, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapV904f7e5061ea0a11202b104fcb01960d528c1ccd(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[float64]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].CardMoney] = dataIn[i]
	}
	return result
}

func queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make(map[int]int, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make(map[int]User, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].UserID] = dataIn[i]
	}
	return result
}

func queryColumnMapVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make(map[int]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Age] = dataIn[i]
	}
	return result
}

func queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV0278bd21d25efe9d7f53cc03a127d039c73e98a2(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.QueryInnerStruct2)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.QueryInnerStruct.MM
	}
	return result
}

func queryColumnV1a5b7250371597524e364f0c816390c77a8b3331(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumnV3465d8e8bd06530b67385225226c4b0bf84f1e19(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumnV3923b792e276005e09637544ecb3aec8be870f41(data interface{}, column string) interface{} {
	dataIn := data.([]string)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnV74382f13397675a69df627f3964ab4362f62b343(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV904b262f8e2329ec73c320ca0e5ca82f14165586(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV90d8f7e69601fe716207d104d716fb6ca0fbfbb9(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.QueryInnerStruct2)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.MM
	}
	return result
}

func queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnVa101312acb15d8adcd8418ad608c0e7fcedba287(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.UserID
	}
	return result
}

func queryColumnVb6a4a6f17a7bc9f4857b953563c0a001e04b0df4(data interface{}, column string) interface{} {
	dataIn := data.([]subtest.Address)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.City
	}
	return result
}

func queryColumnVc6bebe695f9ff26a9409d88809a85fd9cceda86d(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnVe4faf04ecdc0c4866eb2ec31e4d6fe4610ea9cae(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumnVfbe7749ddebdcadafcd40309aff9af0ae35a7707(data interface{}, column string) interface{} {
	dataIn := data.([]testdata.ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryCombineV09c7dc794885ed91aba0c8d6332ac0560ddd8c38(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]testdata.ContentType)
	rightDataIn := rightData.([]testdata.ContentType)
	combineFunctorIn := combineFunctor.(func(testdata.ContentType, testdata.ContentType) testdata.ContentType)
	newData := make([]testdata.ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV228612e67e8c710669fd5517896357f50582a609(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]testdata.ContentType)
	rightDataIn := rightData.([]int)
	combineFunctorIn := combineFunctor.(func(testdata.ContentType, int) testdata.ContentType)
	newData := make([]testdata.ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV38f41d1ea9151d195cb01ed01c28e94b7fbd938b(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]int)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(int, User) User)
	newData := make([]User, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineVdd9cf383efe9adb9dedf293cf43f875133066c23(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]Admin)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(Admin, User) AdminUser)
	newData := make([]AdminUser, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroupV09895bd703c01101d7fdf582f1cc06b6b3ac6c3f(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.QueryInnerStruct2)
	groupFunctorIn := groupFunctor.(func([]testdata.QueryInnerStruct2) []testdata.QueryInnerStruct2)
	bufferData := make([]testdata.QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]testdata.QueryInnerStruct2, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].QueryInnerStruct.MM
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupV34b1efcd4a92cbf477c338aec5ef9e49e4e25774(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]string)
	groupFunctorIn := groupFunctor.(func([]string) testdata.ContentType)
	bufferData := make([]string, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i]
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]int)
	groupFunctorIn := groupFunctor.(func([]int) int)
	bufferData := make([]int, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]int, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i]
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupV7c15a02754e8e39a158cd3d3e8088258012c6f55(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]User)
	groupFunctorIn := groupFunctor.(func([]User) Department)
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]Department, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].UserID
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupV7c2562e83d5d0523f97388581656549f827c5363(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []testdata.ContentType)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupV7e426b5791161f51e11d28bc0feb633b08a92842(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []float64)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]float64, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupV8192c48029957133c04b253ae1641389b396caf8(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []testdata.ContentType)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Ok
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupV8738994dbfe455405a3c9c07003ed95fec1dde39(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) int)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	result := make([]int, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Register
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupV9f6a97dc8909e876536f87142b5f351e7c589297(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []testdata.ContentType)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Register
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupVb6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]User)
	groupFunctorIn := groupFunctor.(func([]User) Department)
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	result := make([]Department, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].CreateTime
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupVc6ec01b1a8f68f11281a5667b88c5e1967c42c86(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []testdata.ContentType)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupVca76f6ce6b260c880f4be68c0c6f72186fa635b5(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) float32)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make([]float32, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupVe2f66d8f0c3b0f6a3ae63ae50e57ab532ca4d858(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) float64)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]float64, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupVe7d5282fd7785eff7f4ab513493be85a01ef834b(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]User)
	groupFunctorIn := groupFunctor.(func([]User) []Department)
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	result := make([]Department, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].CreateTime
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryGroupVec4d3b2c280c9fc398a0f8a554bc1c2ec7010257(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]int)
	groupFunctorIn := groupFunctor.(func([]int) Department)
	bufferData := make([]int, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]Department, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i]
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)
	}

	return result
}

func queryGroupVff06256d82e26530e6c726fa09cb485f82fa3a55(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]testdata.ContentType)
	groupFunctorIn := groupFunctor.(func([]testdata.ContentType) []testdata.ContentType)
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	result := make([]testdata.ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Ok
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0
		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)
	}

	return result
}

func queryJoinV29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]Admin)
	rightDataIn := rightData.([]User)
	joinFunctorIn := joinFunctor.(func(Admin, User) AdminUser)
	result := make([]AdminUser, 0, len(leftDataIn))

	emptyLeftData := Admin{}
	emptyRightData := User{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserID
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.AdminID
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoinV3b9eb26d0b915e0c6d26c912e94883240a196b12(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]User)
	rightDataIn := rightData.([]int)
	joinFunctorIn := joinFunctor.(func(User, int) User)
	result := make([]User, 0, len(leftDataIn))

	emptyLeftData := User{}
	emptyRightData := 0
	joinPlace = "right"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i]
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.UserID
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoinV4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]Admin)
	rightDataIn := rightData.([]User)
	joinFunctorIn := joinFunctor.(func(Admin, User) AdminUser)
	result := make([]AdminUser, 0, len(leftDataIn))

	emptyLeftData := Admin{}
	emptyRightData := User{}
	joinPlace = "inner"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserID
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.AdminID
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func querySelectV6ee296b569bde42f94e969ef503e43ca5827eda6(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]int)
	selectFunctorIn := selectFunctor.(func(int) User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectVd595e8831f2025c62e11d6273f64a4d65d32f4c2(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]User)
	selectFunctorIn := selectFunctor.(func(User) Sex)
	result := make([]Sex, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySortV24a673fdaa750030447c9f4189cf7c07434d6efc(data interface{}, sortType string) interface{} {
	dataIn := data.([]User)
	newData := make([]User, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].UserID < newData[j].UserID {
			return -1
		} else if newData[i].UserID > newData[j].UserID {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV3fd878965c2e3841c633e35be48e5a35315692a9(data interface{}, sortType string) interface{} {
	dataIn := data.([]Admin)
	newData := make([]Admin, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].IsMale == false && newData[j].IsMale == true {
			return -1
		} else if newData[i].IsMale == true && newData[j].IsMale == false {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795(data interface{}, sortType string) interface{} {
	dataIn := data.([]int)
	newData := make([]int, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i] < newData[j] {
			return 1
		} else if newData[i] > newData[j] {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortVc863de209dea4542794c518ccbd6908305b5d21e(data interface{}, sortType string) interface{} {
	dataIn := data.([]User)
	newData := make([]User, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].UserID < newData[j].UserID {
			return 1
		} else if newData[i].UserID > newData[j].UserID {
			return -1
		}

		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		if newData[i].CreateTime.Before(newData[j].CreateTime) {
			return -1
		} else if newData[i].CreateTime.After(newData[j].CreateTime) {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func queryWhereVae59d599ddb6a7c038d3dc70deaebd78a39febff(data interface{}, whereFunctor interface{}) interface{} {
	dataIn := data.([]User)
	whereFunctorIn := whereFunctor.(func(User) bool)
	result := make([]User, 0, len(dataIn))

	for _, single := range dataIn {
		shouldStay := whereFunctorIn(single)
		if shouldStay == true {
			result = append(result, single)
		}
	}
	return result
}

func queryWhereVdf6742e675632943121cefdb3ad29ba75c08eaac(data interface{}, whereFunctor interface{}) interface{} {
	dataIn := data.([]int)
	whereFunctorIn := whereFunctor.(func(int) bool)
	result := make([]int, 0, len(dataIn))

	for _, single := range dataIn {
		shouldStay := whereFunctorIn(single)
		if shouldStay == true {
			result = append(result, single)
		}
	}
	return result
}

func init() {

	query.ColumnMapMacroRegister([]testdata.ContentType{}, "     Name         ", queryColumnMapV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f)

	query.ColumnMapMacroRegister([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnMapV0278bd21d25efe9d7f53cc03a127d039c73e98a2)

	query.ColumnMapMacroRegister([]testdata.ContentType{}, " Name ", queryColumnMapV1a5b7250371597524e364f0c816390c77a8b3331)

	query.ColumnMapMacroRegister([]testdata.ContentType{}, "Ok        ", queryColumnMapV268d58dff08fb0947b9b47bcae328d584ec43d6c)

	query.ColumnMapMacroRegister([]string{}, " . ", queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMapMacroRegister([]testdata.ContentType{}, "    Money  ", queryColumnMapV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7)

	query.ColumnMapMacroRegister([]User{}, "Name", queryColumnMapV74382f13397675a69df627f3964ab4362f62b343)

	query.ColumnMapMacroRegister([]int{}, ".", queryColumnMapV904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMapMacroRegister([]testdata.ContentType{}, "    CardMoney", queryColumnMapV904f7e5061ea0a11202b104fcb01960d528c1ccd)

	query.ColumnMapMacroRegister([]int{}, " . ", queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMapMacroRegister([]User{}, "UserID", queryColumnMapVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMapMacroRegister([]testdata.ContentType{}, "Age        ", queryColumnMapVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3)

	query.ColumnMacroRegister([]testdata.ContentType{}, "     Name         ", queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f)

	query.ColumnMacroRegister([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnV0278bd21d25efe9d7f53cc03a127d039c73e98a2)

	query.ColumnMacroRegister([]testdata.ContentType{}, " Name ", queryColumnV1a5b7250371597524e364f0c816390c77a8b3331)

	query.ColumnMacroRegister([]testdata.ContentType{}, "Ok        ", queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  Age  ", queryColumnV3465d8e8bd06530b67385225226c4b0bf84f1e19)

	query.ColumnMacroRegister([]string{}, " . ", queryColumnV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMacroRegister([]testdata.ContentType{}, "    Money  ", queryColumnV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7)

	query.ColumnMacroRegister([]User{}, "Name", queryColumnV74382f13397675a69df627f3964ab4362f62b343)

	query.ColumnMacroRegister([]int{}, ".", queryColumnV904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMacroRegister([]testdata.ContentType{}, "    CardMoney", queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd)

	query.ColumnMacroRegister([]testdata.QueryInnerStruct2{}, "  MM  ", queryColumnV90d8f7e69601fe716207d104d716fb6ca0fbfbb9)

	query.ColumnMacroRegister([]int{}, " . ", queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  Money  ", queryColumnVa101312acb15d8adcd8418ad608c0e7fcedba287)

	query.ColumnMacroRegister([]User{}, "UserID", queryColumnVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMacroRegister([]subtest.Address{}, "City", queryColumnVb6a4a6f17a7bc9f4857b953563c0a001e04b0df4)

	query.ColumnMacroRegister([]User{}, ".", queryColumnVc6bebe695f9ff26a9409d88809a85fd9cceda86d)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  CardMoney  ", queryColumnVe4faf04ecdc0c4866eb2ec31e4d6fe4610ea9cae)

	query.ColumnMacroRegister([]testdata.ContentType{}, "Age        ", queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3)

	query.ColumnMacroRegister([]testdata.ContentType{}, "CardMoney  ", queryColumnVfbe7749ddebdcadafcd40309aff9af0ae35a7707)

	query.CombineMacroRegister([]testdata.ContentType{}, []testdata.ContentType{}, (func(testdata.ContentType, testdata.ContentType) testdata.ContentType)(nil), queryCombineV09c7dc794885ed91aba0c8d6332ac0560ddd8c38)

	query.CombineMacroRegister([]testdata.ContentType{}, []int{}, (func(testdata.ContentType, int) testdata.ContentType)(nil), queryCombineV228612e67e8c710669fd5517896357f50582a609)

	query.CombineMacroRegister([]int{}, []User{}, (func(int, User) User)(nil), queryCombineV38f41d1ea9151d195cb01ed01c28e94b7fbd938b)

	query.CombineMacroRegister([]Admin{}, []User{}, (func(Admin, User) AdminUser)(nil), queryCombineVdd9cf383efe9adb9dedf293cf43f875133066c23)

	query.GroupMacroRegister([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", (func([]testdata.QueryInnerStruct2) []testdata.QueryInnerStruct2)(nil), queryGroupV09895bd703c01101d7fdf582f1cc06b6b3ac6c3f)

	query.GroupMacroRegister([]string{}, ".", (func([]string) testdata.ContentType)(nil), queryGroupV34b1efcd4a92cbf477c338aec5ef9e49e4e25774)

	query.GroupMacroRegister([]int{}, ".", (func([]int) int)(nil), queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d)

	query.GroupMacroRegister([]User{}, "UserID", (func([]User) Department)(nil), queryGroupV7c15a02754e8e39a158cd3d3e8088258012c6f55)

	query.GroupMacroRegister([]testdata.ContentType{}, "Name", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupV7c2562e83d5d0523f97388581656549f827c5363)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) []float64)(nil), queryGroupV7e426b5791161f51e11d28bc0feb633b08a92842)

	query.GroupMacroRegister([]testdata.ContentType{}, " Ok ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupV8192c48029957133c04b253ae1641389b396caf8)

	query.GroupMacroRegister([]testdata.ContentType{}, "Register ", (func([]testdata.ContentType) int)(nil), queryGroupV8738994dbfe455405a3c9c07003ed95fec1dde39)

	query.GroupMacroRegister([]testdata.ContentType{}, "Register ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupV9f6a97dc8909e876536f87142b5f351e7c589297)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) Department)(nil), queryGroupVb6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVc6ec01b1a8f68f11281a5667b88c5e1967c42c86)

	query.GroupMacroRegister([]testdata.ContentType{}, "Name", (func([]testdata.ContentType) float32)(nil), queryGroupVca76f6ce6b260c880f4be68c0c6f72186fa635b5)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) float64)(nil), queryGroupVe2f66d8f0c3b0f6a3ae63ae50e57ab532ca4d858)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) []Department)(nil), queryGroupVe7d5282fd7785eff7f4ab513493be85a01ef834b)

	query.GroupMacroRegister([]int{}, ".", (func([]int) Department)(nil), queryGroupVec4d3b2c280c9fc398a0f8a554bc1c2ec7010257)

	query.GroupMacroRegister([]testdata.ContentType{}, "Ok", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVff06256d82e26530e6c726fa09cb485f82fa3a55)

	query.JoinMacroRegister([]Admin{}, []User{}, "left", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoinV29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf)

	query.JoinMacroRegister([]User{}, []int{}, "right", "UserID = .", (func(User, int) User)(nil), queryJoinV3b9eb26d0b915e0c6d26c912e94883240a196b12)

	query.JoinMacroRegister([]Admin{}, []User{}, "inner", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoinV4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f)

	query.SelectMacroRegister([]int{}, (func(int) User)(nil), querySelectV6ee296b569bde42f94e969ef503e43ca5827eda6)

	query.SelectMacroRegister([]User{}, (func(User) Sex)(nil), querySelectVd595e8831f2025c62e11d6273f64a4d65d32f4c2)

	query.SortMacroRegister([]User{}, "UserID asc", querySortV24a673fdaa750030447c9f4189cf7c07434d6efc)

	query.SortMacroRegister([]Admin{}, "IsMale asc", querySortV3fd878965c2e3841c633e35be48e5a35315692a9)

	query.SortMacroRegister([]int{}, ". desc", querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]User{}, "UserID desc,Name asc,CreateTime asc", querySortVc863de209dea4542794c518ccbd6908305b5d21e)

	query.WhereMacroRegister([]User{}, (func(User) bool)(nil), queryWhereVae59d599ddb6a7c038d3dc70deaebd78a39febff)

	query.WhereMacroRegister([]int{}, (func(int) bool)(nil), queryWhereVdf6742e675632943121cefdb3ad29ba75c08eaac)

}
