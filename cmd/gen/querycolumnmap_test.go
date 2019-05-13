package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fishedee/tools/assert"
	"github.com/fishedee/tools/query"
)

// User User
type User struct {
	UserId     int
	Age        int
	Name       string
	CreateTime time.Time
}

// Admin Admin
type Admin struct {
	AdminId int
	Level   int
}

// AdminUser AdminUser
type AdminUser struct {
	AdminId    int
	Level      int
	Name       string
	CreateTime time.Time
}

// Department Department
type Department struct {
	Employees []User
}

// Sex Sex
type Sex struct {
	IsMale bool
}

func TestQueryColumnMap(t *testing.T) {
	data := []User{
		User{UserId: 1},
		User{UserId: -2},
		User{UserId: 3},
	}
	assert.Equal(t, query.ColumnMap(data, "UserId"), map[int]User{
		1:  User{UserId: 1},
		-2: User{UserId: -2},
		3:  User{UserId: 3},
	})
	assert.Equal(t, query.ColumnMap([]int{5, 6, 8, 8, 0, 6}, "."), map[int]int{
		5: 5,
		6: 6,
		8: 8,
		0: 0,
	})
}

func initQueryColumnMapData() []User {
	data := make([]User, 1000, 1000)
	for i := range data {
		data[i].UserId = rand.Int()
		data[i].Age = rand.Int()
	}
	return data
}

func BenchmarkQueryColumnMapHand(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		newData := make(map[int]User, len(data))
		for _, single := range data {
			newData[single.UserId] = single
		}
	}
}

func BenchmarkQueryColumnMapMacro(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap(data, "UserId")
	}
}

func BenchmarkQueryColumnMapReflect(b *testing.B) {
	data := initQueryColumnMapData()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		query.ColumnMap(data, "Age")
	}
}
