package model

var Json struct {
	Name string `json:"name" binding:"required"`
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Users []User
