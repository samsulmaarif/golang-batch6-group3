package model

import "time"

type BaseModel struct {
	Id        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	Fullname string
	Email    string
	Password string
	Role     string
}

var Users = []User{}
