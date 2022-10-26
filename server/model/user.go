package model

import "time"

type BaseModel struct {
	Id        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Contact    string `json:"contact"`
	Gender     string `json:"gender"`
	Street     string `json:"street"`
	CityId     int    `json:"city_id"`
	ProvinceId int    `json:"province_id"`
}

var Users = []User{}
