package testdata

import (
	"github.com/fishedee/tools/query"
	"time"
)

func queryColumnMapV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV1a5b7250371597524e364f0c816390c77a8b3331(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Name] = dataIn[i]
	}
	return result
}

func queryColumnMapV268d58dff08fb0947b9b47bcae328d584ec43d6c(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[bool]ContentType, len(dataIn))

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
	dataIn := data.([]ContentType)
	result := make(map[float32]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Money] = dataIn[i]
	}
	return result
}

func queryColumnMapV904f7e5061ea0a11202b104fcb01960d528c1ccd(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[float64]ContentType, len(dataIn))

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

func queryColumnMapVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[int]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		result[dataIn[i].Age] = dataIn[i]
	}
	return result
}

func queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV1a5b7250371597524e364f0c816390c77a8b3331(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumnV3465d8e8bd06530b67385225226c4b0bf84f1e19(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
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
	dataIn := data.([]ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
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
	dataIn := data.([]ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumnVe4faf04ecdc0c4866eb2ec31e4d6fe4610ea9cae(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumnVfbe7749ddebdcadafcd40309aff9af0ae35a7707(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryCombineV09c7dc794885ed91aba0c8d6332ac0560ddd8c38(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]ContentType)
	rightDataIn := rightData.([]ContentType)
	combineFunctorIn := combineFunctor.(func(ContentType, ContentType) ContentType)
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombineV228612e67e8c710669fd5517896357f50582a609(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]ContentType)
	rightDataIn := rightData.([]int)
	combineFunctorIn := combineFunctor.(func(ContentType, int) ContentType)
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroupV34b1efcd4a92cbf477c338aec5ef9e49e4e25774(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]string)
	groupFunctorIn := groupFunctor.(func([]string) ContentType)
	bufferData := make([]string, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
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

func queryGroupV7c2562e83d5d0523f97388581656549f827c5363(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
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

	return result
}

func queryGroupV7e426b5791161f51e11d28bc0feb633b08a92842(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []float64)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
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

	return result
}

func queryGroupV8738994dbfe455405a3c9c07003ed95fec1dde39(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) int)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
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

	return result
}

func queryGroupVc6ec01b1a8f68f11281a5667b88c5e1967c42c86(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
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

	return result
}

func queryGroupVca76f6ce6b260c880f4be68c0c6f72186fa635b5(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) float32)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) float64)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
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

func queryGroupVff06256d82e26530e6c726fa09cb485f82fa3a55(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
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

	return result
}

func queryJoinV05e7a66203c74f236e031eb75f8bb81cd62b4c5c(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
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

func queryJoinV278136fb3d4e24cb55e3c1a341f21de90dd0c0ef(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinV2d5a6b539e97379decf26a5f6440a1e2b1244f31(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinV5307fcc7a488d36862795265fa06bdf9b785df46(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
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

func queryJoinV6105a55f22b1be36255189a3f5046562cfc2d4a9(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinV7365b213d46b1da711d1e7f819b4f6dd8667ae73(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
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

func queryJoinV7abc7f400d9daa08a82df0f80c5a1ff2961f607b(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinV7bc11e07d61cd1ea2fa514ade01c0b8704e26bd4(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]string)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(string, ContentType2) ContentType2)
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

func queryJoinVa1a3a9e9a8824a018e089e1dd8166ad273eb2929(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinVeb07b54f706db726142724132d348b491f412ff4(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
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

func queryJoinVfcf50f290878012d066ad5c0cfe18fd39296bd0c(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
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

func queryJoinVfdfe281f6058a6868c2e12f32ca13f6c3dabe242(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]string)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(string, UserType) UserType)
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

func querySelectV03b12a1acb9918d749bfa5c88ad41af5d11af1be(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) bool)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV2df5a9ed3c5d3e88750329401c4bf91a6ff879ed(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) string)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV446a51935e0be020e2e4b8c4aefb054faac5e32c(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) map[string]int)
	result := make([]map[string]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV48d789bc642d669f7044381e856125972c5002f7(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) float64)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV49410ca5646fa429d1c1a39d6adbe922a8b398de(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) time.Time)
	result := make([]time.Time, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV4b690bab06e586f19389ddb3bb327d5abd08d9d4(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) ContentType)
	result := make([]ContentType, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV4dfeeca849a875cfddf58f802cd0c7b99c4d6ee5(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) float32)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelectV56ff3e10b6f0a89535dd3a999fd35056dcde6d1b(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySortV10a439a196b4cc9dca0592a40a23aba8392203e4(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortV40a4020cd2e40cf9f18c9e1f6ab38b3954c42a18(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortV621ff803aedd0b5249ba4dfb7425b4776188e65f(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortV74654e8b45593005ef783b89255269f7c6ecc39b(data interface{}, sortType string) interface{} {
	dataIn := data.([]int)
	newData := make([]int, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortVa7adb01486b6faebfa0ad8d4fd2a7e47b68375d0(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortVb0375dbee2cd369a9223aff4984ff75cd4b1379d(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortVb05f1e0866945217824e423729a4f07ffb44848e(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortVc176fe16e99109d1f300a4845f4b46844243aa9b(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func querySortVd0f5de5ff928ef2468012b8408e5088feb13e615(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

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

func queryWhereV03b12a1acb9918d749bfa5c88ad41af5d11af1be(data interface{}, whereFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	whereFunctorIn := whereFunctor.(func(ContentType) bool)
	result := make([]ContentType, 0, len(dataIn))

	for _, single := range dataIn {
		shouldStay := whereFunctorIn(single)
		if shouldStay == true {
			result = append(result, single)
		}
	}
	return result
}

func init() {

	query.ColumnMapMacroRegister([]ContentType{}, "     Name         ", queryColumnMapV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f)

	query.ColumnMapMacroRegister([]ContentType{}, " Name ", queryColumnMapV1a5b7250371597524e364f0c816390c77a8b3331)

	query.ColumnMapMacroRegister([]ContentType{}, "Ok        ", queryColumnMapV268d58dff08fb0947b9b47bcae328d584ec43d6c)

	query.ColumnMapMacroRegister([]string{}, " . ", queryColumnMapV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMapMacroRegister([]ContentType{}, "    Money  ", queryColumnMapV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7)

	query.ColumnMapMacroRegister([]ContentType{}, "    CardMoney", queryColumnMapV904f7e5061ea0a11202b104fcb01960d528c1ccd)

	query.ColumnMapMacroRegister([]int{}, " . ", queryColumnMapV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMapMacroRegister([]ContentType{}, "Age        ", queryColumnMapVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3)

	query.ColumnMacroRegister([]ContentType{}, "     Name         ", queryColumnV0210877b9f45b0e2d7c760cad71c8d1aa3e70a6f)

	query.ColumnMacroRegister([]ContentType{}, " Name ", queryColumnV1a5b7250371597524e364f0c816390c77a8b3331)

	query.ColumnMacroRegister([]ContentType{}, "Ok        ", queryColumnV268d58dff08fb0947b9b47bcae328d584ec43d6c)

	query.ColumnMacroRegister([]ContentType{}, "  Age  ", queryColumnV3465d8e8bd06530b67385225226c4b0bf84f1e19)

	query.ColumnMacroRegister([]string{}, " . ", queryColumnV3923b792e276005e09637544ecb3aec8be870f41)

	query.ColumnMacroRegister([]ContentType{}, "    Money  ", queryColumnV6b4a4fd9e192f5ca29db73c69b9472328b1d4cd7)

	query.ColumnMacroRegister([]ContentType{}, "    CardMoney", queryColumnV904f7e5061ea0a11202b104fcb01960d528c1ccd)

	query.ColumnMacroRegister([]int{}, " . ", queryColumnV91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	query.ColumnMacroRegister([]ContentType{}, "  Money  ", queryColumnVa101312acb15d8adcd8418ad608c0e7fcedba287)

	query.ColumnMacroRegister([]ContentType{}, "  CardMoney  ", queryColumnVe4faf04ecdc0c4866eb2ec31e4d6fe4610ea9cae)

	query.ColumnMacroRegister([]ContentType{}, "Age        ", queryColumnVe56b49b3fa0f6bf953dd89ffa8677a9ed1f2dfe3)

	query.ColumnMacroRegister([]ContentType{}, "CardMoney  ", queryColumnVfbe7749ddebdcadafcd40309aff9af0ae35a7707)

	query.CombineMacroRegister([]ContentType{}, []ContentType{}, (func(ContentType, ContentType) ContentType)(nil), queryCombineV09c7dc794885ed91aba0c8d6332ac0560ddd8c38)

	query.CombineMacroRegister([]ContentType{}, []int{}, (func(ContentType, int) ContentType)(nil), queryCombineV228612e67e8c710669fd5517896357f50582a609)

	query.GroupMacroRegister([]string{}, ".", (func([]string) ContentType)(nil), queryGroupV34b1efcd4a92cbf477c338aec5ef9e49e4e25774)

	query.GroupMacroRegister([]int{}, ".", (func([]int) int)(nil), queryGroupV7959aac2ba701c92b02938af82c21599cbf58c3d)

	query.GroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) []ContentType)(nil), queryGroupV7c2562e83d5d0523f97388581656549f827c5363)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []float64)(nil), queryGroupV7e426b5791161f51e11d28bc0feb633b08a92842)

	query.GroupMacroRegister([]ContentType{}, " Ok ", (func([]ContentType) []ContentType)(nil), queryGroupV8192c48029957133c04b253ae1641389b396caf8)

	query.GroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) int)(nil), queryGroupV8738994dbfe455405a3c9c07003ed95fec1dde39)

	query.GroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) []ContentType)(nil), queryGroupV9f6a97dc8909e876536f87142b5f351e7c589297)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []ContentType)(nil), queryGroupVc6ec01b1a8f68f11281a5667b88c5e1967c42c86)

	query.GroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) float32)(nil), queryGroupVca76f6ce6b260c880f4be68c0c6f72186fa635b5)

	query.GroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) float64)(nil), queryGroupVe2f66d8f0c3b0f6a3ae63ae50e57ab532ca4d858)

	query.GroupMacroRegister([]ContentType{}, "Ok", (func([]ContentType) []ContentType)(nil), queryGroupVff06256d82e26530e6c726fa09cb485f82fa3a55)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "inner", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV05e7a66203c74f236e031eb75f8bb81cd62b4c5c)

	query.JoinMacroRegister([]UserType{}, []UserType{}, " left ", "  Name  =  Name ", (func(UserType, UserType) UserType)(nil), queryJoinV278136fb3d4e24cb55e3c1a341f21de90dd0c0ef)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " Money=Money ", (func(UserType, UserType) UserType)(nil), queryJoinV2d5a6b539e97379decf26a5f6440a1e2b1244f31)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "right", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV5307fcc7a488d36862795265fa06bdf9b785df46)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " CardMoney = Money ", (func(UserType, UserType) UserType)(nil), queryJoinV6105a55f22b1be36255189a3f5046562cfc2d4a9)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "left", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinV7365b213d46b1da711d1e7f819b4f6dd8667ae73)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", "Ok  =  Ok", (func(UserType, UserType) UserType)(nil), queryJoinV7abc7f400d9daa08a82df0f80c5a1ff2961f607b)

	query.JoinMacroRegister([]string{}, []ContentType2{}, "left", "  .  =  UserName ", (func(string, ContentType2) ContentType2)(nil), queryJoinV7bc11e07d61cd1ea2fa514ade01c0b8704e26bd4)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "right", "Age=Age", (func(UserType, UserType) UserType)(nil), queryJoinVa1a3a9e9a8824a018e089e1dd8166ad273eb2929)

	query.JoinMacroRegister([]UserType{}, []UserType{}, "left", " Register = Register ", (func(UserType, UserType) UserType)(nil), queryJoinVeb07b54f706db726142724132d348b491f412ff4)

	query.JoinMacroRegister([]UserType{}, []ContentType2{}, "outer", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoinVfcf50f290878012d066ad5c0cfe18fd39296bd0c)

	query.JoinMacroRegister([]string{}, []UserType{}, "left", " . = Name", (func(string, UserType) UserType)(nil), queryJoinVfdfe281f6058a6868c2e12f32ca13f6c3dabe242)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), querySelectV03b12a1acb9918d749bfa5c88ad41af5d11af1be)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) string)(nil), querySelectV2df5a9ed3c5d3e88750329401c4bf91a6ff879ed)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) map[string]int)(nil), querySelectV446a51935e0be020e2e4b8c4aefb054faac5e32c)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) float64)(nil), querySelectV48d789bc642d669f7044381e856125972c5002f7)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) time.Time)(nil), querySelectV49410ca5646fa429d1c1a39d6adbe922a8b398de)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) ContentType)(nil), querySelectV4b690bab06e586f19389ddb3bb327d5abd08d9d4)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) float32)(nil), querySelectV4dfeeca849a875cfddf58f802cd0c7b99c4d6ee5)

	query.SelectMacroRegister([]ContentType{}, (func(ContentType) int)(nil), querySelectV56ff3e10b6f0a89535dd3a999fd35056dcde6d1b)

	query.SortMacroRegister([]ContentType{}, " Money desc,Age asc", querySortV10a439a196b4cc9dca0592a40a23aba8392203e4)

	query.SortMacroRegister([]ContentType{}, "Age desc,Ok desc", querySortV40a4020cd2e40cf9f18c9e1f6ab38b3954c42a18)

	query.SortMacroRegister([]ContentType{}, "Name desc", querySortV621ff803aedd0b5249ba4dfb7425b4776188e65f)

	query.SortMacroRegister([]int{}, ". asc", querySortV74654e8b45593005ef783b89255269f7c6ecc39b)

	query.SortMacroRegister([]ContentType{}, "CardMoney,Register desc", querySortVa7adb01486b6faebfa0ad8d4fd2a7e47b68375d0)

	query.SortMacroRegister([]int{}, ". desc", querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]ContentType{}, " Money desc,Age asc,Name desc", querySortVb0375dbee2cd369a9223aff4984ff75cd4b1379d)

	query.SortMacroRegister([]ContentType{}, "Name asc", querySortVb05f1e0866945217824e423729a4f07ffb44848e)

	query.SortMacroRegister([]ContentType{}, "Money,Register desc", querySortVc176fe16e99109d1f300a4845f4b46844243aa9b)

	query.SortMacroRegister([]ContentType{}, "Ok desc,Name", querySortVd0f5de5ff928ef2468012b8408e5088feb13e615)

	query.WhereMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), queryWhereV03b12a1acb9918d749bfa5c88ad41af5d11af1be)

}
