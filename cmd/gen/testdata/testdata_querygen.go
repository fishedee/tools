package testdata

import (
	"github.com/fishedee/tools/query"
	testdata "github.com/fishedee/tools/query/test_data"
)

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

	query.CombineMacroRegister([]testdata.ContentType{}, []int{}, (func(testdata.ContentType, int) testdata.ContentType)(nil), queryCombineV32ceb64b78fbf30e491600c88b60e25966b3d0c0)

	query.CombineMacroRegister([]int{}, []User{}, (func(int, User) User)(nil), queryCombineV38f41d1ea9151d195cb01ed01c28e94b7fbd938b)

	query.CombineMacroRegister([]testdata.ContentType{}, []testdata.ContentType{}, (func(testdata.ContentType, testdata.ContentType) testdata.ContentType)(nil), queryCombineV67e4a61d96d7ecbc2c0ef31db8c2bb9496b45dae)

	query.CombineMacroRegister([]Admin{}, []User{}, (func(Admin, User) AdminUser)(nil), queryCombineVdd9cf383efe9adb9dedf293cf43f875133066c23)

	query.JoinMacroRegister([]Admin{}, []User{}, "left", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoinV29a4e5efaa3321b67511a20dc57f29e0ec5aa0cf)

	query.JoinMacroRegister([]User{}, []int{}, "right", "UserID = .", (func(User, int) User)(nil), queryJoinV3b9eb26d0b915e0c6d26c912e94883240a196b12)

	query.JoinMacroRegister([]Admin{}, []User{}, "inner", "AdminID = UserID", (func(Admin, User) AdminUser)(nil), queryJoinV4f860e3bcc925b8d2d6fae8fab6a2a5cc4fba96f)

	query.SelectMacroRegister([]User{}, (func(User) Sex)(nil), querySelectVd595e8831f2025c62e11d6273f64a4d65d32f4c2)

	query.SortMacroRegister([]User{}, "UserID asc", querySortV24a673fdaa750030447c9f4189cf7c07434d6efc)

	query.SortMacroRegister([]Admin{}, "IsMale asc", querySortV3fd878965c2e3841c633e35be48e5a35315692a9)

	query.SortMacroRegister([]int{}, ". desc", querySortVaf891d058d5a2e0a3ac4b4b291ae9bb959364795)

	query.SortMacroRegister([]User{}, "UserID desc,Name asc,CreateTime asc", querySortVc863de209dea4542794c518ccbd6908305b5d21e)

	query.WhereMacroRegister([]User{}, (func(User) bool)(nil), queryWhereVae59d599ddb6a7c038d3dc70deaebd78a39febff)

	query.WhereMacroRegister([]int{}, (func(int) bool)(nil), queryWhereVdf6742e675632943121cefdb3ad29ba75c08eaac)

}
