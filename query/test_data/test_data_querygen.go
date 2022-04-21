package testdata

import (
	"github.com/fishedee/tools/query"
	"time"
)

func queryColumnMapV1634476ac5d81ffe821151e2cd007944bb75387e(data []ContentType, column string) map[string]ContentType {
	dataIn := data
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV17413fb3daa269475dd84fbd43de7143d1a1fecd(data []ContentType, column string) map[float32][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[float32]int, len(dataIn))
	result := make(map[float32][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Money
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

		result[bufferData[kbegin].Money] = bufferData[kbegin:k]

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

func queryColumnMapV3c9c70c40c63019848d44a57a05b107340215e80(data []ContentType, column string) map[float32]ContentType {
	dataIn := data
	result := make(map[float32]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Money] = dataIn[i]
	}
	return result
}

func queryColumnMapV3f666f2858a8335d8dedbc686b034302e2f4fc0f(data []ContentType, column string) map[float64]ContentType {
	dataIn := data
	result := make(map[float64]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].CardMoney] = dataIn[i]
	}
	return result
}

func queryColumnMapV53ba67968a45a6ad0844debe5d0894f3bfbf48a2(data []ContentType, column string) map[float64][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[float64]int, len(dataIn))
	result := make(map[float64][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].CardMoney
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

		result[bufferData[kbegin].CardMoney] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapV68400eb3d8154e355447a3652268b3ee45711762(data []ContentType, column string) map[string][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make(map[string][]ContentType, len(dataIn))

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

		result[bufferData[kbegin].Name] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapV7c828f56cce9605a6545677e43765003f2a1e3da(data []ContentType, column string) map[int][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make(map[int][]ContentType, len(dataIn))

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

		result[bufferData[kbegin].Age] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapV8388abeb56ba7226a39979e2f85daf5b77b0ec13(data []ContentType, column string) map[string][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make(map[string][]ContentType, len(dataIn))

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

		result[bufferData[kbegin].Name] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapV84f666378aff13a68045428e41ef52e9bf17a800(data []QueryInnerStruct2, column string) map[int]QueryInnerStruct2 {
	dataIn := data
	result := make(map[int]QueryInnerStruct2, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].QueryInnerStruct.MM] = dataIn[i]
	}
	return result
}

func queryColumnMapV890fbb2d61811a21c543c5d899f2e91b964343eb(data []ContentType, column string) map[bool]ContentType {
	dataIn := data
	result := make(map[bool]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Ok] = dataIn[i]
	}
	return result
}

func queryColumnMapV8be389735876e433a498564bd3f63c8f1232d915(data []ContentType, column string) map[string]ContentType {
	dataIn := data
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
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

func queryColumnMapV9379e410e09e784ed64aa5e8311734b853ca1260(data []ContentType, column string) map[int]ContentType {
	dataIn := data
	result := make(map[int]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Age] = dataIn[i]
	}
	return result
}

func queryColumnMapVa3c39cd2e6c27da10580768a73066e91af5bfb6c(data []ContentType, column string) map[bool][]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	result := make(map[bool][]ContentType, len(dataIn))

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

		result[bufferData[kbegin].Ok] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnMapVc0d612348cc652ffb28125191e9e18b8956d695f(data []QueryInnerStruct2, column string) map[int][]QueryInnerStruct2 {
	dataIn := data
	bufferData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make(map[int][]QueryInnerStruct2, len(dataIn))

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

		result[bufferData[kbegin].QueryInnerStruct.MM] = bufferData[kbegin:k]

	}

	return result
}

func queryColumnV1326c8079077532a2df04e5705513b21ec7ccfde(data []ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV1634476ac5d81ffe821151e2cd007944bb75387e(data []ContentType, column string) []string {
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

func queryColumnV3c9c70c40c63019848d44a57a05b107340215e80(data []ContentType, column string) []float32 {
	dataIn := data
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnV3f666f2858a8335d8dedbc686b034302e2f4fc0f(data []ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV67d1dbaa7a84019e2b2b100fc5d35aa18ad7d0ce(data []ContentType, column string) []float64 {
	dataIn := data
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnV84f666378aff13a68045428e41ef52e9bf17a800(data []QueryInnerStruct2, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.QueryInnerStruct.MM
	}
	return result
}

func queryColumnV890fbb2d61811a21c543c5d899f2e91b964343eb(data []ContentType, column string) []bool {
	dataIn := data
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumnV8be389735876e433a498564bd3f63c8f1232d915(data []ContentType, column string) []string {
	dataIn := data
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
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

func queryColumnV9379e410e09e784ed64aa5e8311734b853ca1260(data []ContentType, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumnVba800328deafbc2a4fbda5fd19e9eaca7bd4f9a1(data []ContentType, column string) []float32 {
	dataIn := data
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnVc2be6eac9135e9e0a63d2c9155bea08a4f3c2b5e(data []QueryInnerStruct2, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.MM
	}
	return result
}

func queryColumnVd8a8ba7093b7cf4407a40f064340077d6ad3be2f(data []ContentType, column string) []int {
	dataIn := data
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0(leftData []ContentType, rightData []int, combineFunctor func(ContentType, int) ContentType) []ContentType {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae(leftData []ContentType, rightData []ContentType, combineFunctor func(ContentType, ContentType) ContentType) []ContentType {
	leftDataIn := leftData
	rightDataIn := rightData
	combineFunctorIn := combineFunctor
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroupV289d77ccfff4ea272a58267e3b9c9de43411f84b(data []string, groupType string, groupFunctor func([]string) ContentType) []ContentType {
	dataIn := data
	bufferData := make([]string, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupV5e87b2e0994d2f6b51e103cd03f0e20bc2b60c2a(data []ContentType, groupType string, groupFunctor func([]ContentType) int) []int {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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

func queryGroupV80b407696d3665d0652382da3e9fda32506f530f(data []ContentType, groupType string, groupFunctor func([]ContentType) []float64) *[]float64 {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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

func queryGroupVb0f32b70a6b484b0da2fb2a791874fe6a632fbf5(data []ContentType, groupType string, groupFunctor func([]ContentType) []ContentType) *[]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupVb7a74b6b9fc71210fab248a70af4f4beccf9c794(data []QueryInnerStruct2, groupType string, groupFunctor func([]QueryInnerStruct2) []QueryInnerStruct2) *[]QueryInnerStruct2 {
	dataIn := data
	bufferData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]QueryInnerStruct2, 0, len(dataIn))

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

func queryGroupVd527b6188165d052cd11b364e72e6bdfe4fd76d6(data []ContentType, groupType string, groupFunctor func([]ContentType) []ContentType) *[]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupVd9d416d8cc0214bf9f6454fc0f9361ce6d224172(data []ContentType, groupType string, groupFunctor func([]ContentType) []ContentType) *[]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupVe43c990928ddf5d988a6464ebd9aa14669069506(data []ContentType, groupType string, groupFunctor func([]ContentType) []ContentType) *[]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupVed5e4b40611c7106ad6cc07c33aef9aa052ff6f1(data []ContentType, groupType string, groupFunctor func([]ContentType) float64) []float64 {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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

func queryGroupVef320495b62cbaeae4ebd6f3805c6b2d56041b01(data []ContentType, groupType string, groupFunctor func([]ContentType) []ContentType) *[]ContentType {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor
	result := make([]ContentType, 0, len(dataIn))

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

func queryGroupVf7bec03379fe55d2997b3d5ac7360bbc400c27f3(data []ContentType, groupType string, groupFunctor func([]ContentType) float32) []float32 {
	dataIn := data
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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

func queryJoinV16717d359669b0a88ee7cc7822a752e5268ac86e(leftData []UserType, rightData []ContentType2, joinPlace, joinType string, joinFunctor func(UserType, ContentType2) resultType) []resultType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "right"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
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
		fieldValue := leftValue.Name
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

func queryJoinV1b68c7344d84699f91e8f8b67ab69c82ebb24396(leftData []UserType, rightData []ContentType2, joinPlace, joinType string, joinFunctor func(UserType, ContentType2) resultType) []resultType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "inner"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
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
		fieldValue := leftValue.Name
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

func queryJoinV3f5255508598e2dd853b2e1a776b591bbc01f661(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "right"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Age
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
		fieldValue := leftValue.Age
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

func queryJoinV594ea381e3a7138b5a71da17e916ffc715a05065(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[float64]int, len(rightDataIn))
	mapDataFirst := make(map[float64]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Money
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
		fieldValue := leftValue.Money
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

func queryJoinV6318a3be49a0b2c65397cbe6095d1fef08f1da1a(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[float64]int, len(rightDataIn))
	mapDataFirst := make(map[float64]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Money
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
		fieldValue := leftValue.CardMoney
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

func queryJoinV65f020a6760fb2879c59487e6b60fb48f5f775eb(leftData []UserType, rightData []ContentType2, joinPlace, joinType string, joinFunctor func(UserType, ContentType2) resultType) []resultType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "outer"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
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
		fieldValue := leftValue.Name
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

func queryJoinV72e7ca024be21c12e5c83b83eadcdbf445435874(leftData []ExtendType, rightData []ExtendType, joinPlace, joinType string, joinFunctor func(ExtendType, ExtendType) ExtendType) []ExtendType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]ExtendType, 0, len(leftDataIn))

	emptyLeftData := ExtendType{}
	emptyRightData := ExtendType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].ContentID
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
		fieldValue := leftValue.ContentID
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

func queryJoinV735feb68194a860dc2ce928b8a1fe0cdd278c5f7(leftData []string, rightData []UserType, joinPlace, joinType string, joinFunctor func(string, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := ""
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Name
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
		fieldValue := leftValue
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

func queryJoinV7ab733e147706522deb4421f1f3d43427bcb47f2(leftData []QueryInnerStruct2, rightData []QueryInnerStruct2, joinPlace, joinType string, joinFunctor func(QueryInnerStruct2, QueryInnerStruct2) QueryInnerStruct2) []QueryInnerStruct2 {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]QueryInnerStruct2, 0, len(leftDataIn))

	emptyLeftData := QueryInnerStruct2{}
	emptyRightData := QueryInnerStruct2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].QueryInnerStruct.MM
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
		fieldValue := leftValue.QueryInnerStruct.MM
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

func queryJoinV8403d43d84a5d29444e02c89eb65b7873baf7d74(leftData []UserType, rightData []ContentType2, joinPlace, joinType string, joinFunctor func(UserType, ContentType2) resultType) []resultType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
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
		fieldValue := leftValue.Name
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

func queryJoinV9bdc75718519891ca0fdf5f7da8ab252229b0998(leftData []string, rightData []ContentType2, joinPlace, joinType string, joinFunctor func(string, ContentType2) ContentType2) []ContentType2 {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]ContentType2, 0, len(leftDataIn))

	emptyLeftData := ""
	emptyRightData := ContentType2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
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
		fieldValue := leftValue
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

func queryJoinVadbb58372ab09b3a7f683ab8ea68ff28ee769a98(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[time.Time]int, len(rightDataIn))
	mapDataFirst := make(map[time.Time]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Register
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
		fieldValue := leftValue.Register
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

func queryJoinVc461ee076fa39ef07ed5e9e12cf9b008ff2cf77d(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[bool]int, len(rightDataIn))
	mapDataFirst := make(map[bool]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Ok
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
		fieldValue := leftValue.Ok
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

func queryJoinVe155a25d203f4a8ba9b8d0fe79617435bbae2020(leftData []UserType, rightData []UserType, joinPlace, joinType string, joinFunctor func(UserType, UserType) UserType) []UserType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Name
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
		fieldValue := leftValue.Name
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

func queryJoinVe642f90bacb1b5af7c628b3bd214fac620e2dc0c(leftData []int, rightData []ExtendType, joinPlace, joinType string, joinFunctor func(int, ExtendType) ExtendType) []ExtendType {
	leftDataIn := leftData
	rightDataIn := rightData
	joinFunctorIn := joinFunctor
	result := make([]ExtendType, 0, len(leftDataIn))

	emptyLeftData := 0
	emptyRightData := ExtendType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].ContentID
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
		fieldValue := leftValue
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

func querySelectV8dd2f22a7d420700a40e4fb90e0ec144ad4ab02a(data []ContentType, selectFunctor func(a ContentType) bool) []bool {
	result := make([]bool, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectV97b3b14d2f425f3e98e2b7107d3cdd7c6c52a735(data []ContentType, selectFunctor func(a ContentType) time.Time) []time.Time {
	result := make([]time.Time, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectV9a85b8e1070a2db925ad62b365386546ff4af9a9(data []ContentType, selectFunctor func(a ContentType) string) []string {
	result := make([]string, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVa8d2c26d0fb3626963157cb810f2deaec4c004b2(data []ContentType, selectFunctor func(a ContentType) map[string]int) []map[string]int {
	result := make([]map[string]int, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVc16afb48d4b6ea5e2998d4dca33159a7844f2b79(data []ContentType, selectFunctor func(a ContentType) ContentType) []ContentType {
	result := make([]ContentType, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVd2b4e246806a4194279f457567ad1af15e1e4693(data []ContentType, selectFunctor func(a ContentType) float64) []float64 {
	result := make([]float64, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVe039fdcac8cace31c190f73c448a8fbb64250c08(data []ContentType, selectFunctor func(a ContentType) int) []int {
	result := make([]int, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySelectVf211ca8299315c42e232800e2496b19f20b10812(data []ContentType, selectFunctor func(a ContentType) float32) []float32 {
	result := make([]float32, len(data), len(data))

	for i, single := range data {
		result[i] = selectFunctor(single)
	}
	return result
}

func querySortV02acdee795c5f46c192fe4d156a7cb6445fbd820(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Name < newData[j].Name {
			return 1
		} else if newData[i].Name > newData[j].Name {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV12a67b856dc43b179cff46187ca40790068008bb(data []QueryInnerStruct2, sortType string) []QueryInnerStruct2 {
	newData := make([]QueryInnerStruct2, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].MM < newData[j].MM {
			return 1
		} else if newData[i].MM > newData[j].MM {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV2246cd31bfe003b82d01c3b267b265ec1e4d4c9c(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return 1
		} else if newData[i].Money > newData[j].Money {
			return -1
		}

		if newData[i].Age < newData[j].Age {
			return -1
		} else if newData[i].Age > newData[j].Age {
			return 1
		}

		if newData[i].Name < newData[j].Name {
			return 1
		} else if newData[i].Name > newData[j].Name {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV3167b008d944f57b21219c7ccd8b7038cb10298e(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Age < newData[j].Age {
			return 1
		} else if newData[i].Age > newData[j].Age {
			return -1
		}

		if newData[i].Ok == false && newData[j].Ok == true {
			return 1
		} else if newData[i].Ok == true && newData[j].Ok == false {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV3a655406a8a554479037c1dbae992af30dc6c515(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].CardMoney < newData[j].CardMoney {
			return -1
		} else if newData[i].CardMoney > newData[j].CardMoney {
			return 1
		}

		if newData[i].Register.Before(newData[j].Register) {
			return 1
		} else if newData[i].Register.After(newData[j].Register) {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV74654e8b45593005ef783b89255269f7c6ecc39b(data []int, sortType string) []int {
	newData := make([]int, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i] < newData[j] {
			return -1
		} else if newData[i] > newData[j] {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortV7b9fcae94950c02af6862470d192f7ab2af36b85(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return -1
		} else if newData[i].Money > newData[j].Money {
			return 1
		}

		if newData[i].Register.Before(newData[j].Register) {
			return 1
		} else if newData[i].Register.After(newData[j].Register) {
			return -1
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

func querySortVbebdee905ede9ba26057790a48fefca7682da6c4(data []QueryInnerStruct2, sortType string) []QueryInnerStruct2 {
	newData := make([]QueryInnerStruct2, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].QueryInnerStruct.MM < newData[j].QueryInnerStruct.MM {
			return -1
		} else if newData[i].QueryInnerStruct.MM > newData[j].QueryInnerStruct.MM {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortVecfb5ac13e272fa79a529c2724ce68ee9d96a3d0(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return 1
		} else if newData[i].Money > newData[j].Money {
			return -1
		}

		if newData[i].Age < newData[j].Age {
			return -1
		} else if newData[i].Age > newData[j].Age {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortVf7ba08f955406257d5dcd7f186e622fe1ed796db(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Ok == false && newData[j].Ok == true {
			return 1
		} else if newData[i].Ok == true && newData[j].Ok == false {
			return -1
		}

		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySortVfdaec6e376e64f78de44a25c5c06271899db26fb(data []ContentType, sortType string) []ContentType {
	newData := make([]ContentType, len(data), len(data))
	copy(newData, data)

	query.SortInternal(len(newData), func(i int, j int) int {
		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func queryWhereV8dd2f22a7d420700a40e4fb90e0ec144ad4ab02a(data []ContentType, whereFunctor func(ContentType) bool) []ContentType {
	result := make([]ContentType, 0, len(data))

	for _, single := range data {
		shouldStay := whereFunctor(single)
		if shouldStay {
			result = append(result, single)
		}
	}
	return result
}

func init() {

	query.ColumnMapMacroRegister[ContentType, string]([]ContentType{}, "     Name         ", queryColumnMapV1634476ac5d81ffe821151e2cd007944bb75387e)

	query.ColumnMapMacroRegister[ContentType, float32]([]ContentType{}, "    []Money  ", queryColumnMapV17413fb3daa269475dd84fbd43de7143d1a1fecd)

	query.ColumnMapMacroRegister[string, string]([]string{}, " . ", queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMapMacroRegister[ContentType, float32]([]ContentType{}, "    Money  ", queryColumnMapV3c9c70c40c63019848d44a57a05b107340215e80)

	query.ColumnMapMacroRegister[ContentType, float64]([]ContentType{}, "    CardMoney", queryColumnMapV3f666f2858a8335d8dedbc686b034302e2f4fc0f)

	query.ColumnMapMacroRegister[ContentType, float64]([]ContentType{}, "    []CardMoney", queryColumnMapV53ba67968a45a6ad0844debe5d0894f3bfbf48a2)

	query.ColumnMapMacroRegister[ContentType, string]([]ContentType{}, " []Name ", queryColumnMapV68400eb3d8154e355447a3652268b3ee45711762)

	query.ColumnMapMacroRegister[ContentType, int]([]ContentType{}, "[]Age        ", queryColumnMapV7c828f56cce9605a6545677e43765003f2a1e3da)

	query.ColumnMapMacroRegister[ContentType, string]([]ContentType{}, "     [] Name         ", queryColumnMapV8388abeb56ba7226a39979e2f85daf5b77b0ec13)

	query.ColumnMapMacroRegister[QueryInnerStruct2, int]([]QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnMapV84f666378aff13a68045428e41ef52e9bf17a800)

	query.ColumnMapMacroRegister[ContentType, bool]([]ContentType{}, "Ok        ", queryColumnMapV890fbb2d61811a21c543c5d899f2e91b964343eb)

	query.ColumnMapMacroRegister[ContentType, string]([]ContentType{}, " Name ", queryColumnMapV8be389735876e433a498564bd3f63c8f1232d915)

	query.ColumnMapMacroRegister[int, int]([]int{}, " . ", queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMapMacroRegister[ContentType, int]([]ContentType{}, "Age        ", queryColumnMapV9379e410e09e784ed64aa5e8311734b853ca1260)

	query.ColumnMapMacroRegister[ContentType, bool]([]ContentType{}, "[]Ok        ", queryColumnMapVa3c39cd2e6c27da10580768a73066e91af5bfb6c)

	query.ColumnMapMacroRegister[QueryInnerStruct2, int]([]QueryInnerStruct2{}, "[]QueryInnerStruct.MM", queryColumnMapVc0d612348cc652ffb28125191e9e18b8956d695f)

	query.ColumnMacroRegister([]ContentType{}, "CardMoney  ", queryColumnV1326c8079077532a2df04e5705513b21ec7ccfde)

	query.ColumnMacroRegister([]ContentType{}, "     Name         ", queryColumnV1634476ac5d81ffe821151e2cd007944bb75387e)

	query.ColumnMacroRegister([]string{}, " . ", queryColumnV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMacroRegister([]ContentType{}, "    Money  ", queryColumnV3c9c70c40c63019848d44a57a05b107340215e80)

	query.ColumnMacroRegister([]ContentType{}, "    CardMoney", queryColumnV3f666f2858a8335d8dedbc686b034302e2f4fc0f)

	query.ColumnMacroRegister([]ContentType{}, "  CardMoney  ", queryColumnV67d1dbaa7a84019e2b2b100fc5d35aa18ad7d0ce)

	query.ColumnMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnV84f666378aff13a68045428e41ef52e9bf17a800)

	query.ColumnMacroRegister([]ContentType{}, "Ok        ", queryColumnV890fbb2d61811a21c543c5d899f2e91b964343eb)

	query.ColumnMacroRegister([]ContentType{}, " Name ", queryColumnV8be389735876e433a498564bd3f63c8f1232d915)

	query.ColumnMacroRegister([]int{}, " . ", queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMacroRegister([]ContentType{}, "Age        ", queryColumnV9379e410e09e784ed64aa5e8311734b853ca1260)

	query.ColumnMacroRegister([]ContentType{}, "  Money  ", queryColumnVba800328deafbc2a4fbda5fd19e9eaca7bd4f9a1)

	query.ColumnMacroRegister([]QueryInnerStruct2{}, "  MM  ", queryColumnVc2be6eac9135e9e0a63d2c9155bea08a4f3c2b5e)

	query.ColumnMacroRegister([]ContentType{}, "  Age  ", queryColumnVd8a8ba7093b7cf4407a40f064340077d6ad3be2f)

	query.CombineMacroRegister([]ContentType{}, []int{}, (func(ContentType, int) ContentType)(nil), queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0)

	query.CombineMacroRegister([]ContentType{}, []ContentType{}, (func(ContentType, ContentType) ContentType)(nil), queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae)

	query.GroupMacroRegister([]string{}, ".", (func([]string) ContentType)(nil), queryGroupV289d77ccfff4ea272a58267e3b9c9de43411f84b)

	query.GroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) int)(nil), queryGroupV5e87b2e0994d2f6b51e103cd03f0e20bc2b60c2a)

	query.GroupMacroRegister([]int{}, ".", (func([]int) int)(nil), queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []float64)(nil), queryGroupV80b407696d3665d0652382da3e9fda32506f530f)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []ContentType)(nil), queryGroupVb0f32b70a6b484b0da2fb2a791874fe6a632fbf5)

	query.GroupMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM", (func([]QueryInnerStruct2) []QueryInnerStruct2)(nil), queryGroupVb7a74b6b9fc71210fab248a70af4f4beccf9c794)

	query.GroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) []ContentType)(nil), queryGroupVd527b6188165d052cd11b364e72e6bdfe4fd76d6)

	query.GroupMacroRegister([]ContentType{}, "Ok", (func([]ContentType) []ContentType)(nil), queryGroupVd9d416d8cc0214bf9f6454fc0f9361ce6d224172)

	query.GroupMacroRegister([]ContentType{}, " Ok ", (func([]ContentType) []ContentType)(nil), queryGroupVe43c990928ddf5d988a6464ebd9aa14669069506)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) float64)(nil), queryGroupVed5e4b40611c7106ad6cc07c33aef9aa052ff6f1)

	query.GroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) []ContentType)(nil), queryGroupVef320495b62cbaeae4ebd6f3805c6b2d56041b01)

	query.GroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) float32)(nil), queryGroupVf7bec03379fe55d2997b3d5ac7360bbc400c27f3)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "right", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV16717d359669b0a88ee7cc7822a752e5268ac86e)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "inner", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV1b68c7344d84699f91e8f8b67ab69c82ebb24396)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "right", "Age=Age", (func(UserType, UserType) UserType)(nil), queryJoinV3f5255508598e2dd853b2e1a776b591bbc01f661)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " Money=Money ", (func(UserType, UserType) UserType)(nil), queryJoinV594ea381e3a7138b5a71da17e916ffc715a05065)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " CardMoney = Money ", (func(UserType, UserType) UserType)(nil), queryJoinV6318a3be49a0b2c65397cbe6095d1fef08f1da1a)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "outer", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV65f020a6760fb2879c59487e6b60fb48f5f775eb)

	query.JoinMacroRegister([]ExtendType{}, []ExtendType{}, " left ", "  ContentID  =  ContentID ", (func(ExtendType, ExtendType) ExtendType)(nil), queryJoinV72e7ca024be21c12e5c83b83eadcdbf445435874)

	query.JoinMacroRegister([]string{}, []UserType{}, "left", " . = Name", (func(string, UserType) UserType)(nil), queryJoinV735feb68194a860dc2ce928b8a1fe0cdd278c5f7)

	query.JoinMacroRegister([]QueryInnerStruct2{}, []QueryInnerStruct2{}, "left", "QueryInnerStruct.MM = QueryInnerStruct.MM", (func(QueryInnerStruct2, QueryInnerStruct2) QueryInnerStruct2)(nil), queryJoinV7ab733e147706522deb4421f1f3d43427bcb47f2)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "left", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV8403d43d84a5d29444e02c89eb65b7873baf7d74)

	query.JoinMacroRegister([]string{}, []ContentType2{}, "left", "  .  =  UserName ", (func(string, ContentType2) ContentType2)(nil), queryJoinV9bdc75718519891ca0fdf5f7da8ab252229b0998)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " Register = Register ", (func(UserType, UserType) UserType)(nil), queryJoinVadbb58372ab09b3a7f683ab8ea68ff28ee769a98)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", "Ok  =  Ok", (func(UserType, UserType) UserType)(nil), queryJoinVc461ee076fa39ef07ed5e9e12cf9b008ff2cf77d)

	query.JoinMacroRegister([]UserType{}, []UserType{}, " left ", "  Name  =  Name ", (func(UserType, UserType) UserType)(nil), queryJoinVe155a25d203f4a8ba9b8d0fe79617435bbae2020)

	query.JoinMacroRegister([]int{}, []ExtendType{}, " left ", "  .  =  ContentID ", (func(int, ExtendType) ExtendType)(nil), queryJoinVe642f90bacb1b5af7c628b3bd214fac620e2dc0c)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), querySelectV8dd2f22a7d420700a40e4fb90e0ec144ad4ab02a)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) time.Time)(nil), querySelectV97b3b14d2f425f3e98e2b7107d3cdd7c6c52a735)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) string)(nil), querySelectV9a85b8e1070a2db925ad62b365386546ff4af9a9)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) map[string]int)(nil), querySelectVa8d2c26d0fb3626963157cb810f2deaec4c004b2)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) ContentType)(nil), querySelectVc16afb48d4b6ea5e2998d4dca33159a7844f2b79)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) float64)(nil), querySelectVd2b4e246806a4194279f457567ad1af15e1e4693)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) int)(nil), querySelectVe039fdcac8cace31c190f73c448a8fbb64250c08)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) float32)(nil), querySelectVf211ca8299315c42e232800e2496b19f20b10812)

	query.SortMacroRegister([]ContentType{}, "Name desc", querySortV02acdee795c5f46c192fe4d156a7cb6445fbd820)

	query.SortMacroRegister([]QueryInnerStruct2{}, "MM desc", querySortV12a67b856dc43b179cff46187ca40790068008bb)

	query.SortMacroRegister([]ContentType{}, " Money desc,Age asc,Name desc", querySortV2246cd31bfe003b82d01c3b267b265ec1e4d4c9c)

	query.SortMacroRegister([]ContentType{}, "Age desc,Ok desc", querySortV3167b008d944f57b21219c7ccd8b7038cb10298e)

	query.SortMacroRegister([]ContentType{}, "CardMoney,Register desc", querySortV3a655406a8a554479037c1dbae992af30dc6c515)

	query.SortMacroRegister([]int{}, ". asc", querySortV74654e8b45593005ef783b89255269f7c6ecc39b)

	query.SortMacroRegister([]ContentType{}, "Money,Register desc", querySortV7b9fcae94950c02af6862470d192f7ab2af36b85)

	query.SortMacroRegister([]int{}, ". desc", querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM asc", querySortVbebdee905ede9ba26057790a48fefca7682da6c4)

	query.SortMacroRegister([]ContentType{}, " Money desc,Age asc", querySortVecfb5ac13e272fa79a529c2724ce68ee9d96a3d0)

	query.SortMacroRegister([]ContentType{}, "Ok desc,Name", querySortVf7ba08f955406257d5dcd7f186e622fe1ed796db)

	query.SortMacroRegister([]ContentType{}, "Name asc", querySortVfdaec6e376e64f78de44a25c5c06271899db26fb)

	query.WhereMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), queryWhereV8dd2f22a7d420700a40e4fb90e0ec144ad4ab02a)

}
