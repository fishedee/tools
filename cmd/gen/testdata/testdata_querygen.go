package testdata

import (
	"github.com/fishedee/tools/cmd/gen/testdata/subtest"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/test_data"
	"time"
)

func queryColumnMapV1634476ac5d81ffe821151e2cd007944bb75387e(data []testdata.ContentType, column string) map[string]testdata.ContentType {
	dataIn := data
	result := make(map[string]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV28f674724a19ea93b16beb429196eaa567b10f28(data []User, column string) map[int][]User {
	dataIn := data
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make(map[int][]User, len(dataIn))

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

		result[bufferData[kbegin].UserID] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41(data []string, column string) map[string]string {
	dataIn := data
	result := make(map[string]string, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapV3c9c70c40c63019848d44a57a05b107340215e80(data []testdata.ContentType, column string) map[float32]testdata.ContentType {
	dataIn := data
	result := make(map[float32]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Money] = dataIn[i]
	}
	return result
}

func queryColumnMapV3f666f2858a8335d8dedbc686b034302e2f4fc0f(data []testdata.ContentType, column string) map[float64]testdata.ContentType {
	dataIn := data
	result := make(map[float64]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].CardMoney] = dataIn[i]
	}
	return result
}

func queryColumnMapV74382f13397675a69df627f3964ab4362f62b343(data []User, column string) map[string]User {
	dataIn := data
	result := make(map[string]User, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV84f666378aff13a68045428e41ef52e9bf17a800(data []testdata.QueryInnerStruct2, column string) map[int]testdata.QueryInnerStruct2 {
	dataIn := data
	result := make(map[int]testdata.QueryInnerStruct2, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].QueryInnerStruct.MM] = dataIn[i]
	}
	return result
}

func queryColumnMapV890fbb2d61811a21c543c5d899f2e91b964343eb(data []testdata.ContentType, column string) map[bool]testdata.ContentType {
	dataIn := data
	result := make(map[bool]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Ok] = dataIn[i]
	}
	return result
}

func queryColumnMapV8be389735876e433a498564bd3f63c8f1232d915(data []testdata.ContentType, column string) map[string]testdata.ContentType {
	dataIn := data
	result := make(map[string]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV904b262f8e2329ec73c320ca0e5ca82f14165586(data []int, column string) map[int]int {
	dataIn := data
	result := make(map[int]int, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2(data []int, column string) map[int]int {
	dataIn := data
	result := make(map[int]int, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i]] = dataIn[i]
	}
	return result
}

func queryColumnMapV9379e410e09e784ed64aa5e8311734b853ca1260(data []testdata.ContentType, column string) map[int]testdata.ContentType {
	dataIn := data
	result := make(map[int]testdata.ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Age] = dataIn[i]
	}
	return result
}

func queryColumnMapVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data []User, column string) map[int]User {
	dataIn := data
	result := make(map[int]User, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].UserID] = dataIn[i]
	}
	return result
}

func queryColumnV1326c8079077532a2df04e5705513b21ec7ccfde(data []testdata.ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV1634476ac5d81ffe821151e2cd007944bb75387e(data []testdata.ContentType, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV3923b792e276005e09637544ecb3aec8be870f41(data []string, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV3c9c70c40c63019848d44a57a05b107340215e80(data []testdata.ContentType, column string) []float32 {
	dataIn := data
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnV3f666f2858a8335d8dedbc686b034302e2f4fc0f(data []testdata.ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV67d1dbaa7a84019e2b2b100fc5d35aa18ad7d0ce(data []testdata.ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV74382f13397675a69df627f3964ab4362f62b343(data []User, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV84f666378aff13a68045428e41ef52e9bf17a800(data []testdata.QueryInnerStruct2, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.QueryInnerStruct.MM
	}
	return result
}

func queryColumnV890fbb2d61811a21c543c5d899f2e91b964343eb(data []testdata.ContentType, column string) []bool {
	dataIn := data
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumnV8be389735876e433a498564bd3f63c8f1232d915(data []testdata.ContentType, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV904b262f8e2329ec73c320ca0e5ca82f14165586(data []int, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2(data []int, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnV9379e410e09e784ed64aa5e8311734b853ca1260(data []testdata.ContentType, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumnVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data []User, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.UserID
	}
	return result
}

func queryColumnVb6a4a6f17a7bc9f4857b953563c0a001e04b0df4(data []subtest.Address, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.City
	}
	return result
}

func queryColumnVba800328deafbc2a4fbda5fd19e9eaca7bd4f9a1(data []testdata.ContentType, column string) []float32 {
	dataIn := data
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnVc2be6eac9135e9e0a63d2c9155bea08a4f3c2b5e(data []testdata.QueryInnerStruct2, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.MM
	}
	return result
}

func queryColumnVc6bebe695f9ff26a9409d88809a85fd9cceda86d(data []User, column string) []User {
	dataIn := data
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumnVd8a8ba7093b7cf4407a40f064340077d6ad3be2f(data []testdata.ContentType, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0(leftData []testdata.ContentType, rightData []int, combineFunctor func(testdata.ContentType, int) testdata.ContentType) []testdata.ContentType {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]testdata.ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV38f41d1ea9151d195cb01ed01c28e94b7fbd938b(leftData []int, rightData []User, combineFunctor func(int, User) User) []User {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]User, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae(leftData []testdata.ContentType, rightData []testdata.ContentType, combineFunctor func(testdata.ContentType, testdata.ContentType) testdata.ContentType) []testdata.ContentType {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]testdata.ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineVdd9cf383efe9adb9dedf293cf43f875133066c23(leftData []Admin, rightData []User, combineFunctor func(Admin, User) AdminUser) []AdminUser {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]AdminUser, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroupV289d77ccfff4ea272a58267e3b9c9de43411f84b(data []string, groupType string, groupFunctor func([]string) testdata.ContentType) []testdata.ContentType {
	dataIn := data
	bufferData := make([]string, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupV5e87b2e0994d2f6b51e103cd03f0e20bc2b60c2a(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) int) []int {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d(data []int, groupType string, groupFunctor func([]int) int) []int {
	dataIn := data
	bufferData := make([]int, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupV7c15a02754e8e39a158cd3d3e8088258012c6f55(data []User, groupType string, groupFunctor func([]User) Department) []Department {
	dataIn := data
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupV80b407696d3665d0652382da3e9fda32506f530f(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []float64) *[]float64 {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVb0f32b70a6b484b0da2fb2a791874fe6a632fbf5(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []testdata.ContentType) *[]testdata.ContentType {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVb6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41(data []User, groupType string, groupFunctor func([]User) Department) []Department {
	dataIn := data
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupVb7a74b6b9fc71210fab248a70af4f4beccf9c794(data []testdata.QueryInnerStruct2, groupType string, groupFunctor func([]testdata.QueryInnerStruct2) []testdata.QueryInnerStruct2) *[]testdata.QueryInnerStruct2 {
	dataIn := data
	bufferData := make([]testdata.QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVd527b6188165d052cd11b364e72e6bdfe4fd76d6(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []testdata.ContentType) *[]testdata.ContentType {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVd9d416d8cc0214bf9f6454fc0f9361ce6d224172(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []testdata.ContentType) *[]testdata.ContentType {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVe43c990928ddf5d988a6464ebd9aa14669069506(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []testdata.ContentType) *[]testdata.ContentType {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVe7d5282fd7785eff7f4ab513493be85a01ef834b(data []User, groupType string, groupFunctor func([]User) []Department) *[]Department {
	dataIn := data
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVec4d3b2c280c9fc398a0f8a554bc1c2ec7010257(data []int, groupType string, groupFunctor func([]int) Department) []Department {
	dataIn := data
	bufferData := make([]int, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupVed5e4b40611c7106ad6cc07c33aef9aa052ff6f1(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) float64) []float64 {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryGroupVef320495b62cbaeae4ebd6f3805c6b2d56041b01(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) []testdata.ContentType) *[]testdata.ContentType {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

	return &result
}

func queryGroupVf7bec03379fe55d2997b3d5ac7360bbc400c27f3(data []testdata.ContentType, groupType string, groupFunctor func([]testdata.ContentType) float32) []float32 {
	dataIn := data
	bufferData := make([]testdata.ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor
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

func queryJoinV29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf(leftData []Admin, rightData []User, joinPlace, joinType string, joinFunctor func(Admin, User) AdminUser) []AdminUser {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
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

func queryJoinV3b9eb26d0b915e0c6d26c912e94883240a196b12(leftData []User, rightData []int, joinPlace, joinType string, joinFunctor func(User, int) User) []User {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
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

func queryJoinV4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f(leftData []Admin, rightData []User, joinPlace, joinType string, joinFunctor func(Admin, User) AdminUser) []AdminUser {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
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

func querySelectV6ee296b569bde42f94e969ef503e43ca5827eda6(data []int, selectFunctor func(a int) User) []User {
	result := make([]User, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVd595e8831f2025c62e11d6273f64a4d65d32f4c2(data []User, selectFunctor func(a User) Sex) []Sex {
	result := make([]Sex, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySortV24a673fdaa750030447c9f4189cf7c07434d6efc(data []User, sortType string) []User {
	newData := make([]User, len(data), len(data))
	copy(newData, data)

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

func querySortV3fd878965c2e3841c633e35be48e5a35315692a9(data []Admin, sortType string) []Admin {
	newData := make([]Admin, len(data), len(data))
	copy(newData, data)

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

func querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795(data []int, sortType string) []int {
	newData := make([]int, len(data), len(data))
	copy(newData, data)

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

func querySortVc863de209dea4542794c518ccbd6908305b5d21e(data []User, sortType string) []User {
	newData := make([]User, len(data), len(data))
	copy(newData, data)

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

func queryWhereVae59d599ddb6a7c038d3dc70deaebd78a39febff(data []User, whereFunctor func(User) bool) []User {
	result := make([]User, 0, len(data))

	for _, single := range data {
		shouldStay := whereFunctor(single)
		if shouldStay {
			result = append(result, single)
		}
	}
	return result
}

func queryWhereVdf6742e675632943121cefdb3ad29ba75c08eaac(data []int, whereFunctor func(int) bool) []int {
	result := make([]int, 0, len(data))

	for _, single := range data {
		shouldStay := whereFunctor(single)
		if shouldStay {
			result = append(result, single)
		}
	}
	return result
}

func init() {

	query.ColumnMapMacroRegister[testdata.ContentType, string]([]testdata.ContentType{}, "     Name         ", queryColumnMapV1634476ac5d81ffe821151e2cd007944bb75387e)

	query.ColumnMapMacroRegister[User, int]([]User{}, "[]UserID", queryColumnMapV28f674724a19ea93b16beb429196eaa567b10f28)

	query.ColumnMapMacroRegister[string, string]([]string{}, " . ", queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMapMacroRegister[testdata.ContentType, float32]([]testdata.ContentType{}, "    Money  ", queryColumnMapV3c9c70c40c63019848d44a57a05b107340215e80)

	query.ColumnMapMacroRegister[testdata.ContentType, float64]([]testdata.ContentType{}, "    CardMoney", queryColumnMapV3f666f2858a8335d8dedbc686b034302e2f4fc0f)

	query.ColumnMapMacroRegister[User, string]([]User{}, "Name", queryColumnMapV74382f13397675a69df627f3964ab4362f62b343)

	query.ColumnMapMacroRegister[testdata.QueryInnerStruct2, int]([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnMapV84f666378aff13a68045428e41ef52e9bf17a800)

	query.ColumnMapMacroRegister[testdata.ContentType, bool]([]testdata.ContentType{}, "Ok        ", queryColumnMapV890fbb2d61811a21c543c5d899f2e91b964343eb)

	query.ColumnMapMacroRegister[testdata.ContentType, string]([]testdata.ContentType{}, " Name ", queryColumnMapV8be389735876e433a498564bd3f63c8f1232d915)

	query.ColumnMapMacroRegister[int, int]([]int{}, ".", queryColumnMapV904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMapMacroRegister[int, int]([]int{}, " . ", queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMapMacroRegister[testdata.ContentType, int]([]testdata.ContentType{}, "Age        ", queryColumnMapV9379e410e09e784ed64aa5e8311734b853ca1260)

	query.ColumnMapMacroRegister[User, int]([]User{}, "UserID", queryColumnMapVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMacroRegister([]testdata.ContentType{}, "CardMoney  ", queryColumnV1326c8079077532a2df04e5705513b21ec7ccfde)

	query.ColumnMacroRegister([]testdata.ContentType{}, "     Name         ", queryColumnV1634476ac5d81ffe821151e2cd007944bb75387e)

	query.ColumnMacroRegister([]string{}, " . ", queryColumnV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMacroRegister([]testdata.ContentType{}, "    Money  ", queryColumnV3c9c70c40c63019848d44a57a05b107340215e80)

	query.ColumnMacroRegister([]testdata.ContentType{}, "    CardMoney", queryColumnV3f666f2858a8335d8dedbc686b034302e2f4fc0f)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  CardMoney  ", queryColumnV67d1dbaa7a84019e2b2b100fc5d35aa18ad7d0ce)

	query.ColumnMacroRegister([]User{}, "Name", queryColumnV74382f13397675a69df627f3964ab4362f62b343)

	query.ColumnMacroRegister([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnV84f666378aff13a68045428e41ef52e9bf17a800)

	query.ColumnMacroRegister([]testdata.ContentType{}, "Ok        ", queryColumnV890fbb2d61811a21c543c5d899f2e91b964343eb)

	query.ColumnMacroRegister([]testdata.ContentType{}, " Name ", queryColumnV8be389735876e433a498564bd3f63c8f1232d915)

	query.ColumnMacroRegister([]int{}, ".", queryColumnV904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMacroRegister([]int{}, " . ", queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMacroRegister([]testdata.ContentType{}, "Age        ", queryColumnV9379e410e09e784ed64aa5e8311734b853ca1260)

	query.ColumnMacroRegister([]User{}, "UserID", queryColumnVac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMacroRegister([]subtest.Address{}, "City", queryColumnVb6a4a6f17a7bc9f4857b953563c0a001e04b0df4)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  Money  ", queryColumnVba800328deafbc2a4fbda5fd19e9eaca7bd4f9a1)

	query.ColumnMacroRegister([]testdata.QueryInnerStruct2{}, "  MM  ", queryColumnVc2be6eac9135e9e0a63d2c9155bea08a4f3c2b5e)

	query.ColumnMacroRegister([]User{}, ".", queryColumnVc6bebe695f9ff26a9409d88809a85fd9cceda86d)

	query.ColumnMacroRegister([]testdata.ContentType{}, "  Age  ", queryColumnVd8a8ba7093b7cf4407a40f064340077d6ad3be2f)

	query.CombineMacroRegister([]testdata.ContentType{}, []int{}, (func(testdata.ContentType, int) testdata.ContentType)(nil), queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0)

	query.CombineMacroRegister([]int{}, []User{}, (func(int, User) User)(nil), queryCombineV38f41d1ea9151d195cb01ed01c28e94b7fbd938b)

	query.CombineMacroRegister([]testdata.ContentType{}, []testdata.ContentType{}, (func(testdata.ContentType, testdata.ContentType) testdata.ContentType)(nil), queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae)

	query.CombineMacroRegister([]Admin{}, []User{}, (func(Admin, User) AdminUser)(nil), queryCombineVdd9cf383efe9adb9dedf293cf43f875133066c23)

	query.GroupMacroRegister([]string{}, ".", (func([]string) testdata.ContentType)(nil), queryGroupV289d77ccfff4ea272a58267e3b9c9de43411f84b)

	query.GroupMacroRegister([]testdata.ContentType{}, "Register ", (func([]testdata.ContentType) int)(nil), queryGroupV5e87b2e0994d2f6b51e103cd03f0e20bc2b60c2a)

	query.GroupMacroRegister([]int{}, ".", (func([]int) int)(nil), queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d)

	query.GroupMacroRegister([]User{}, "UserID", (func([]User) Department)(nil), queryGroupV7c15a02754e8e39a158cd3d3e8088258012c6f55)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) []float64)(nil), queryGroupV80b407696d3665d0652382da3e9fda32506f530f)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVb0f32b70a6b484b0da2fb2a791874fe6a632fbf5)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) Department)(nil), queryGroupVb6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41)

	query.GroupMacroRegister([]testdata.QueryInnerStruct2{}, "QueryInnerStruct.MM", (func([]testdata.QueryInnerStruct2) []testdata.QueryInnerStruct2)(nil), queryGroupVb7a74b6b9fc71210fab248a70af4f4beccf9c794)

	query.GroupMacroRegister([]testdata.ContentType{}, "Name", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVd527b6188165d052cd11b364e72e6bdfe4fd76d6)

	query.GroupMacroRegister([]testdata.ContentType{}, "Ok", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVd9d416d8cc0214bf9f6454fc0f9361ce6d224172)

	query.GroupMacroRegister([]testdata.ContentType{}, " Ok ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVe43c990928ddf5d988a6464ebd9aa14669069506)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) []Department)(nil), queryGroupVe7d5282fd7785eff7f4ab513493be85a01ef834b)

	query.GroupMacroRegister([]int{}, ".", (func([]int) Department)(nil), queryGroupVec4d3b2c280c9fc398a0f8a554bc1c2ec7010257)

	query.GroupMacroRegister([]testdata.ContentType{}, " Age ", (func([]testdata.ContentType) float64)(nil), queryGroupVed5e4b40611c7106ad6cc07c33aef9aa052ff6f1)

	query.GroupMacroRegister([]testdata.ContentType{}, "Register ", (func([]testdata.ContentType) []testdata.ContentType)(nil), queryGroupVef320495b62cbaeae4ebd6f3805c6b2d56041b01)

	query.GroupMacroRegister([]testdata.ContentType{}, "Name", (func([]testdata.ContentType) float32)(nil), queryGroupVf7bec03379fe55d2997b3d5ac7360bbc400c27f3)

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
