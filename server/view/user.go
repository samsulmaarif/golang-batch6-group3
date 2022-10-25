package view

import "golang-batch6-group3/server/model"

type UserCreateResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func NewUserCreateResponse(user *model.User) *UserCreateResponse {
	return &UserCreateResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
}

type UserFindAllResponse struct {
	Id         string `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Street     string `json:"street"`
	CityId     int    `json:"city_id"`
	ProvinceId int    `json:"province_id"`
}

func NewUserFindAllResponse(users *[]model.User) []UserFindAllResponse {
	var usersFind []UserFindAllResponse
	for _, user := range *users {
		usersFind = append(usersFind, *parseModelToUserFind(&user))
	}
	return usersFind
}

func parseModelToUserFind(user *model.User) *UserFindAllResponse {
	return &UserFindAllResponse{
		Id:         user.Id,
		Fullname:   user.Fullname,
		Email:      user.Email,
		Gender:     user.Gender,
		Street:     user.Street,
		CityId:     user.CityId,
		ProvinceId: user.ProvinceId,
	}
}
