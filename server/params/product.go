package params

import (
	"golang-batch6-group3/server/model"
)

type ProductCreate struct {
	Name        string `validate:"required"`
	Category    string `validate:"required"`
	Weight      int    `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required"`
	Stock       int    `validate:"required"`
	ImgUrl      string `validate:"required"`
}

type ProductUpdate struct {
	Name        string
	Category    string
	Weight      int
	Description string
	Price       int
	Stock       int
	ImgUrl      string
}

func (u *ProductCreate) ParseToModel() *model.Product {
	return &model.Product{
		Name:        u.Name,
		Category:    u.Category,
		Weight:      u.Weight,
		Description: u.Description,
		Price:       u.Price,
		Stock:       u.Stock,
		ImgUrl:      u.ImgUrl,
	}
}

func (u *ProductUpdate) ParseToModel() *model.Product {
	return &model.Product{
		Name:        u.Name,
		Category:    u.Category,
		Weight:      u.Weight,
		Description: u.Description,
		Price:       u.Price,
		Stock:       u.Stock,
		ImgUrl:      u.ImgUrl,
	}
}
