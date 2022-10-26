package params

import (
	"golang-batch6-group3/server/model"
)

type ProductCreate struct {
	Name        string `validate:"required"`
	Category    string `validate:"required"`
	Description string `validate:"required"`
	Price       string `validate:"required"`
	Stock       string `validate:"required"`
	ImgUrl      string `validate:"required"`
}

type ProductUpdate struct {
	Name          string
	Category      string
	Description   string
	CatPriceegory string
	Stock         string
	ImgUrl        string
}

func (u *ProductCreate) ParseToModel() *model.Product {
	return &model.Product{
		Name:        u.Name,
		Category:    u.Category,
		Description: u.Description,
		Price:       u.Price,
		Stock:       u.Stock,
		ImgUrl:      u.ImgUrl,
	}
}
