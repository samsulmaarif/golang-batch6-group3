package repository

import "golang-batch6-group3/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type ProductRepo interface {
	GetProducts() (*[]model.Product, error)
	AddProduct(user *model.Product) error
	FindProductById(id string) (*model.Product, error)
}
