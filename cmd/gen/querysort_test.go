package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/donnol/tools/assert"
	"github.com/donnol/tools/query"
)

func TestQuerySort(t *testing.T) {
	data := []User{
		User{UserId: 3, Name: "a"},
		User{UserId: 3, Name: "c"},
		User{UserId: 23, Name: "d"},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		User{UserId: 23, Name: "a"},
		User{UserId: 1},
		User{UserId: 1},
	}
	assert.Equal(t, query.Sort(data, "UserId desc,Name asc,CreateTime asc"), []User{
		User{UserId: 23, Name: "a"},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(1, 0)},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(29, 0)},
		User{UserId: 23, Name: "c", CreateTime: time.Unix(33, 0)},
		User{UserId: 23, Name: "d"},
		User{UserId: 3, Name: "a"},
		User{UserId: 3, Name: "c"},
		User{UserId: 1},
		User{UserId: 1},
	})
	assert.Equal(t, query.Sort([]int{3, 2, 1, 7, -8}, ". desc"), []int{7, 3, 2, 1, -8})
}

func initQuerySortData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserId = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQuerySortHand(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make([]User, len(data), len(data))
		copy(newData, data)
		sort.SliceStable(newData, func(i int, j int) bool {
			return newData[i].UserId < newData[j].UserId
		})
	}
}

func BenchmarkQuerySortMacro(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Sort(data, "UserId asc")
	}
}

func BenchmarkQuerySortReflect(b *testing.B) {
	data := initQuerySortData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.Sort(data, "Age asc")
	}
}
