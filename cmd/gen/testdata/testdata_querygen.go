package testdata

import (
	"github.com/fishedee/tools/cmd/gen/testdata/subtest"
	"github.com/fishedee/tools/query"
	"time"
)

func queryColumnMap_904b262f8e2329ec73c320ca0e5ca82f14165586(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make(map[int]int, len(dataIn))

	for _, single := range dataIn {
		result[single] = single
	}
	return result
}

func queryColumnMap_ac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make(map[int]User, len(dataIn))

	for _, single := range dataIn {
		result[single.UserID] = single
	}
	return result
}

func queryColumn_904b262f8e2329ec73c320ca0e5ca82f14165586(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumn_ac46a6e2d4d6d4f163cc177eb335bc2bb166d92b(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.UserID
	}
	return result
}

func queryColumn_b6a4a6f17a7bc9f4857b953563c0a001e04b0df4(data interface{}, column string) interface{} {
	dataIn := data.([]subtest.Address)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.City
	}
	return result
}

func queryColumn_c6bebe695f9ff26a9409d88809a85fd9cceda86d(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryCombine_38f41d1ea9151d195cb01ed01c28e94b7fbd938b(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]int)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(int, User) User)
	newData := make([]User, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombine_dd9cf383efe9adb9dedf293cf43f875133066c23(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]Admin)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(Admin, User) AdminUser)
	newData := make([]AdminUser, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroup_7c15a02754e8e39a158cd3d3e8088258012c6f55(data interface{}, groupType string, groupFunctor interface{}) interface{} {
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

func queryGroup_b6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41(data interface{}, groupType string, groupFunctor interface{}) interface{} {
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

func queryGroup_ec4d3b2c280c9fc398a0f8a554bc1c2ec7010257(data interface{}, groupType string, groupFunctor interface{}) interface{} {
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

func queryJoin_29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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

func queryJoin_3b9eb26d0b915e0c6d26c912e94883240a196b12(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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

func queryJoin_4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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

func querySelect_6ee296b569bde42f94e969ef503e43ca5827eda6(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]int)
	selectFunctorIn := selectFunctor.(func(int) User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_d595e8831f2025c62e11d6273f64a4d65d32f4c2(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]User)
	selectFunctorIn := selectFunctor.(func(User) Sex)
	result := make([]Sex, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySort_24a673fdaa750030447c9f4189cf7c07434d6efc(data interface{}, sortType string) interface{} {
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

func querySort_af891d058d5a2e0a3ac4b4b291ae9bb959364795(data interface{}, sortType string) interface{} {
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

func querySort_c863de209dea4542794c518ccbd6908305b5d21e(data interface{}, sortType string) interface{} {
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

func queryWhere_ae59d599ddb6a7c038d3dc70deaebd78a39febff(data interface{}, whereFunctor interface{}) interface{} {
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

func queryWhere_df6742e675632943121cefdb3ad29ba75c08eaac(data interface{}, whereFunctor interface{}) interface{} {
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

	query.ColumnMapMacroRegister([]int{}, ".", queryColumnMap_904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMapMacroRegister([]User{}, "UserID", queryColumnMap_ac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMacroRegister([]int{}, ".", queryColumn_904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMacroRegister([]User{}, "UserID", queryColumn_ac46a6e2d4d6d4f163cc177eb335bc2bb166d92b)

	query.ColumnMacroRegister([]subtest.Address{}, "City", queryColumn_b6a4a6f17a7bc9f4857b953563c0a001e04b0df4)

	query.ColumnMacroRegister([]User{}, ".", queryColumn_c6bebe695f9ff26a9409d88809a85fd9cceda86d)

	query.CombineMacroRegister([]int{}, []User{}, (func(int, User) User)(nil), queryCombine_38f41d1ea9151d195cb01ed01c28e94b7fbd938b)

	query.CombineMacroRegister([]Admin{}, []User{}, (func(Admin, User) AdminUser)(nil), queryCombine_dd9cf383efe9adb9dedf293cf43f875133066c23)

	query.GroupMacroRegister([]User{}, "UserID", (func([]User) Department)(nil), queryGroup_7c15a02754e8e39a158cd3d3e8088258012c6f55)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) Department)(nil), queryGroup_b6f1c2171838a5cc1ac9e63d0ecaada8ee2f1e41)

	query.GroupMacroRegister([]int{}, ".", (func([]int) Department)(nil), queryGroup_ec4d3b2c280c9fc398a0f8a554bc1c2ec7010257)

	query.JoinMacroRegister([]Admin{}, []User{}, "left", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoin_29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf)

	query.JoinMacroRegister([]User{}, []int{}, "right", "UserID = .", (func(User, int) User)(nil), queryJoin_3b9eb26d0b915e0c6d26c912e94883240a196b12)

	query.JoinMacroRegister([]Admin{}, []User{}, "inner", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoin_4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f)

	query.SelectMacroRegister([]int{}, (func(int) User)(nil), querySelect_6ee296b569bde42f94e969ef503e43ca5827eda6)

	query.SelectMacroRegister([]User{}, (func(User) Sex)(nil), querySelect_d595e8831f2025c62e11d6273f64a4d65d32f4c2)

	query.SortMacroRegister([]User{}, "UserID asc", querySort_24a673fdaa750030447c9f4189cf7c07434d6efc)

	query.SortMacroRegister([]int{}, ". desc", querySort_af891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]User{}, "UserID desc,Name asc,CreateTime asc", querySort_c863de209dea4542794c518ccbd6908305b5d21e)

	query.WhereMacroRegister([]User{}, (func(User) bool)(nil), queryWhere_ae59d599ddb6a7c038d3dc70deaebd78a39febff)

	query.WhereMacroRegister([]int{}, (func(int) bool)(nil), queryWhere_df6742e675632943121cefdb3ad29ba75c08eaac)

}
