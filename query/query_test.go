package query

import (
	"testing"

	"github.com/fishedee/tools/assert"
)

func TestQuerySelect(t *testing.T) {
	testCase := GetQuerySelectTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Select(singleTestCase.Origin, singleTestCase.Function)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryWhere(t *testing.T) {
	testCase := GetQueryWhereTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Where(singleTestCase.Origin, singleTestCase.Function)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryReduce(t *testing.T) {
	testCase := GetQueryReduceTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Reduce(singleTestCase.Origin, singleTestCase.Function, singleTestCase.InitNum)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQuerySort(t *testing.T) {
	testCase := GetQuerySortTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Sort(singleTestCase.Origin, singleTestCase.SortName)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQueryJoin(t *testing.T) {
	testCase := GetQueryJoinTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Join(singleTestCase.LeftData, singleTestCase.RightData, singleTestCase.JoinPlace, singleTestCase.JoinType, singleTestCase.JoinFuctor)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryGroup(t *testing.T) {
	testCase := GetQueryGroupTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Group(singleTestCase.Data, singleTestCase.GroupType, singleTestCase.GroupFuctor)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQueryColumn(t *testing.T) {
	testCase := GetQueryColumnTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Column(singleTestCase.Data, singleTestCase.Column)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryColumnMap(t *testing.T) {
	testCase := GetQueryColumnMapTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := ColumnMap(singleTestCase.Data, singleTestCase.Column)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryReverse(t *testing.T) {
	testCase := GetQueryReverseTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Reverse(singleTestCase.Data)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryCombine(t *testing.T) {
	testCase := GetQueryCombineTestCase()

	for _, singleTestCase := range testCase {
		result := Combine(singleTestCase.Origin, singleTestCase.Origin2, singleTestCase.Functor)
		assert.Equal(t, result, singleTestCase.Target)
	}
}

func TestQueryDistinct(t *testing.T) {
	testCase := GetQueryDistinctTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := Distinct(singleTestCase.Origin, singleTestCase.UniqueName)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}
