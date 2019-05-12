package testdata

import (
	"github.com/donnol/tools/cmd/gen/testdata/subtest"
	"github.com/donnol/tools/query"
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

func queryColumnMap_bda61613d581d4e6502573a96c698a03a7031fb8(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make(map[int]User, len(dataIn))

	for _, single := range dataIn {
		result[single.UserId] = single
	}
	return result
}

func queryColumn_07eb99a7704735990f03c0468f629ef16515d104(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumn_08e5ce68309eb818810f3ae5380a7481c31cd332(data interface{}, column string) interface{} {
	dataIn := data.([]subtest.Address)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.City
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

func queryColumn_bda61613d581d4e6502573a96c698a03a7031fb8(data interface{}, column string) interface{} {
	dataIn := data.([]User)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.UserId
	}
	return result
}

func queryCombine_31206c0e8c8b497d4b835869326c201a5cf97694(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]int)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(int, User) User)
	newData := make([]User, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombine_9f09f990590f2b42c871c5d7bbe24b0aff03948f(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]Admin)
	rightDataIn := rightData.([]User)
	combineFunctorIn := combineFunctor.(func(Admin, User) AdminUser)
	newData := make([]AdminUser, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroup_aaca024b7640468b608ed8755af609636323ce8b(data interface{}, groupType string, groupFunctor interface{}) interface{} {
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

func queryGroup_f1f1ce9226ecaea3ea6fc82777419ed52fd0c36f(data interface{}, groupType string, groupFunctor interface{}) interface{} {
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

func queryGroup_fdfad344ace2df32db1ae62e4e232f157c763992(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]User)
	groupFunctorIn := groupFunctor.(func([]User) Department)
	bufferData := make([]User, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make([]Department, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].UserId
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

func queryJoin_18bde861c12acb1e0c3173fc7d84b195077b32db(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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
		fieldValue := leftValue.UserId
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

func queryJoin_54262cadc51750499fb0541ee5c55b7425830e0f(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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
		fieldValue := rightDataIn[i].UserId
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
		fieldValue := leftValue.AdminId
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

func queryJoin_d4e9bfd7d31916dde8a632646c5ccc753c1c09fd(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
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
		fieldValue := rightDataIn[i].UserId
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
		fieldValue := leftValue.AdminId
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

func querySelect_19f1028c0a8b4408b1ad4bf02088582f0dace753(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]User)
	selectFunctorIn := selectFunctor.(func(User) Sex)
	result := make([]Sex, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_dff3049e104cedc6dd907972a1a8914f5648f3b2(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]int)
	selectFunctorIn := selectFunctor.(func(int) User)
	result := make([]User, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySort_7156c5177f9a3291d16da5de11a32d3872285009(data interface{}, sortType string) interface{} {
	dataIn := data.([]User)
	newData := make([]User, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].UserId < newData[j].UserId {
			return -1
		} else if newData[i].UserId > newData[j].UserId {
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

	language.QuerySortInternal(len(newData), func(i int, j int) int {
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

func querySort_b27281fbf5dcbd12639fa2e2405ea432ced93be3(data interface{}, sortType string) interface{} {
	dataIn := data.([]User)
	newData := make([]User, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].UserId < newData[j].UserId {
			return 1
		} else if newData[i].UserId > newData[j].UserId {
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

func queryWhere_3c489ad1034aa03a1c14eeb5109005fc9a7a97c1(data interface{}, whereFunctor interface{}) interface{} {
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

	query.ColumnMapMacroRegister([]User{}, "UserId", queryColumnMap_bda61613d581d4e6502573a96c698a03a7031fb8)

	query.ColumnMacroRegister([]User{}, ".", queryColumn_07eb99a7704735990f03c0468f629ef16515d104)

	query.ColumnMacroRegister([]subtest.Address{}, "City", queryColumn_08e5ce68309eb818810f3ae5380a7481c31cd332)

	query.ColumnMacroRegister([]int{}, ".", queryColumn_904b262f8e2329ec73c320ca0e5ca82f14165586)

	query.ColumnMacroRegister([]User{}, "UserId", queryColumn_bda61613d581d4e6502573a96c698a03a7031fb8)

	query.CombineMacroRegister([]int{}, []User{}, (func(int, User) User)(nil), queryCombine_31206c0e8c8b497d4b835869326c201a5cf97694)

	query.CombineMacroRegister([]Admin{}, []User{}, (func(Admin, User) AdminUser)(nil), queryCombine_9f09f990590f2b42c871c5d7bbe24b0aff03948f)

	query.GroupMacroRegister([]int{}, ".", (func([]int) Department)(nil), queryGroup_aaca024b7640468b608ed8755af609636323ce8b)

	query.GroupMacroRegister([]User{}, "CreateTime", (func([]User) Department)(nil), queryGroup_f1f1ce9226ecaea3ea6fc82777419ed52fd0c36f)

	query.GroupMacroRegister([]User{}, "UserId", (func([]User) Department)(nil), queryGroup_fdfad344ace2df32db1ae62e4e232f157c763992)

	query.JoinMacroRegister([]User{}, []int{}, "right", "UserId = .", (func(User, int) User)(nil), queryJoin_18bde861c12acb1e0c3173fc7d84b195077b32db)

	query.JoinMacroRegister([]Admin{}, []User{}, "inner", "AdminId = UserId", (func(Admin, User) AdminUser)(nil), queryJoin_54262cadc51750499fb0541ee5c55b7425830e0f)

	query.JoinMacroRegister([]Admin{}, []User{}, "left", "AdminId = UserId", (func(Admin, User) AdminUser)(nil), queryJoin_d4e9bfd7d31916dde8a632646c5ccc753c1c09fd)

	query.SelectMacroRegister([]User{}, (func(User) Sex)(nil), querySelect_19f1028c0a8b4408b1ad4bf02088582f0dace753)

	query.SelectMacroRegister([]int{}, (func(int) User)(nil), querySelect_dff3049e104cedc6dd907972a1a8914f5648f3b2)

	query.SortMacroRegister([]User{}, "UserId asc", querySort_7156c5177f9a3291d16da5de11a32d3872285009)

	query.SortMacroRegister([]int{}, ". desc", querySort_af891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]User{}, "UserId desc,Name asc,CreateTime asc", querySort_b27281fbf5dcbd12639fa2e2405ea432ced93be3)

	query.WhereMacroRegister([]User{}, (func(User) bool)(nil), queryWhere_3c489ad1034aa03a1c14eeb5109005fc9a7a97c1)

	query.WhereMacroRegister([]int{}, (func(int) bool)(nil), queryWhere_df6742e675632943121cefdb3ad29ba75c08eaac)

}
