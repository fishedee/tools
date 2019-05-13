package main

import "time"

// User User
type User struct {
	UserID     int
	Age        int
	Name       string
	CreateTime time.Time
}

// Admin Admin
type Admin struct {
	AdminID int
	Level   int
}

// AdminUser AdminUser
type AdminUser struct {
	AdminID    int
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
