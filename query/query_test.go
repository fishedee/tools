package query_test

import (
	"testing"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
	"github.com/fishedee/tools/query/testdata"
)

func TestQuerySelect(t *testing.T) {
	testCase := testdata.GetQuerySelectTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Select(singleTestCase.Origin, singleTestCase.Function)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryWhere(t *testing.T) {
	testCase := testdata.GetQueryWhereTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Where(singleTestCase.Origin, singleTestCase.Function)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryReduce(t *testing.T) {
	testCase := testdata.GetQueryReduceTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Reduce(singleTestCase.Origin, singleTestCase.Function, singleTestCase.InitNum)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQuerySort(t *testing.T) {
	testCase := testdata.GetQuerySortTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Sort(singleTestCase.Origin, singleTestCase.SortName)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQueryJoin(t *testing.T) {
	testCase := testdata.GetQueryJoinTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Join(singleTestCase.LeftData, singleTestCase.RightData, singleTestCase.JoinPlace, singleTestCase.JoinType, singleTestCase.JoinFuctor)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryGroup(t *testing.T) {
	testCase := testdata.GetQueryGroupTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Group(singleTestCase.Data, singleTestCase.GroupType, singleTestCase.GroupFuctor)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}

}

func TestQueryColumn(t *testing.T) {
	testCase := testdata.GetQueryColumnTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Column(singleTestCase.Data, singleTestCase.Column)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryColumnMap(t *testing.T) {
	testCase := testdata.GetQueryColumnMapTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.ColumnMap(singleTestCase.Data, singleTestCase.Column)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryReverse(t *testing.T) {
	testCase := testdata.GetQueryReverseTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Reverse(singleTestCase.Data)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}

func TestQueryCombine(t *testing.T) {
	testCase := testdata.GetQueryCombineTestCase()

	for _, singleTestCase := range testCase {
		result := query.Combine(singleTestCase.Origin, singleTestCase.Origin2, singleTestCase.Functor)
		assert.Equal(t, result, singleTestCase.Target)
	}
}

func TestQueryDistinct(t *testing.T) {
	testCase := testdata.GetQueryDistinctTestCase()

	for singleTestCaseIndex, singleTestCase := range testCase {

		result := query.Distinct(singleTestCase.Origin, singleTestCase.UniqueName)
		assert.Equal(t, result, singleTestCase.Target, singleTestCaseIndex)

	}
}
