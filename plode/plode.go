package plode

import (
	"strconv"
	"strings"
)

// Explode Explode
func Explode(input string, separator string) []string {
	dataResult := strings.Split(input, separator)
	dataResultNew := make([]string, 0, len(dataResult))
	for _, singleResult := range dataResult {
		singleResult = strings.Trim(singleResult, " ")
		if len(singleResult) == 0 {
			continue
		}
		dataResultNew = append(dataResultNew, singleResult)
	}
	return dataResultNew
}

// Implode Implode
func Implode(data []string, separator string) string {
	return strings.Join(data, separator)
}

// ExplodeInt ExplodeInt
func ExplodeInt(input string, separator string) []int {
	dataResult := strings.Split(input, separator)
	dataResultNew := make([]int, 0, len(dataResult))
	for _, singleResult := range dataResult {
		singleResult = strings.Trim(singleResult, " ")
		if len(singleResult) == 0 {
			continue
		}
		singleResultInt, err := strconv.Atoi(singleResult)
		if err != nil {
			panic(err)
		}
		dataResultNew = append(dataResultNew, singleResultInt)
	}
	return dataResultNew
}

// ImplodeInt ImplodeInt
func ImplodeInt(data []int, separator string) string {
	result := make([]string, 0, len(data))
	for _, singleData := range data {
		result = append(result, strconv.Itoa(singleData))
	}
	return strings.Join(result, separator)
}
