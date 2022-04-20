package testdata

import (
	"github.com/fishedee/tools/query"
	"time"
)

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

	query.CombineMacroRegister([]ContentType{}, []int{}, (func(ContentType, int) ContentType)(nil), queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0)

	query.CombineMacroRegister([]ContentType{}, []ContentType{}, (func(ContentType, ContentType) ContentType)(nil), queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae)

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
