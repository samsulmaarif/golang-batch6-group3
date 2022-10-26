package gorm_postgres

import (
	"golang-batch6-group3/server/model"
	"golang-batch6-group3/server/repository"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepoGormPostgres(db *gorm.DB) repository.ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) GetProducts() (*[]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p *productRepo) AddProduct(product *model.Product) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (p *productRepo) FindProductById(email string) (*model.Product, error) {
	var product model.Product
	err := p.db.Where("email=?", email).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
