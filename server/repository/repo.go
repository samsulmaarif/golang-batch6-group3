package repository

import "golang-batch6-group3/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	UpdateUserByEmail(email string, user *model.User) error
}

type ProductRepo interface {
	GetProducts() (*[]model.Product, error)
	AddProduct(product *model.Product) error
	DeleteProductById(id string) (*model.Product, error)
	UpdateProductById(id string, product *model.Product) error
	FindProductById(id string) (*model.Product, error)
}

type TransactionRepo interface {
	GetTransactions() (*[]model.Transaction, error)
	GetMemberTransactions(id string) (*[]model.Transaction, error)
	UpdateTransactionStatusById(id string, status string) error
	AddTransaction(transaction *model.Transaction) error
	FindTransactionById(id string) (*model.Transaction, error)
}
