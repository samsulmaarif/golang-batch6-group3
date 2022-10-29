package service

import (
	"database/sql"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
	"log"
	"time"

	"github.com/google/uuid"
)

type ProductServices struct {
	repo repository.ProductRepo
}

func NewProductServices(repo repository.ProductRepo) *ProductServices {
	return &ProductServices{
		repo: repo,
	}
}

func (p *ProductServices) GetProducts() *view.Response {
	products, err := p.repo.GetProducts()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_PRODUCT_FAIL", err.Error())
	}

	return view.SuccessFindAll(view.NewProductFindAllResponse(products))
}

func (p *ProductServices) AddProduct(req *params.ProductCreate) *view.Response {
	product := req.ParseToModel()

	product.Id = uuid.NewString()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := p.repo.AddProduct(product)
	if err != nil {
		log.Printf("get error register user with error %v\n", "")
		return view.ErrInternalServer("CREATE_PRODUCT_FAIL", err.Error())
	}

	return view.SuccessCreated(product)
}

func (p *ProductServices) DeleteProductById(id string) *view.Response {
	product, err := p.repo.FindProductById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_PRODUCT_FAIL", err.Error())
	}
	if product != nil {
		delete, err := p.repo.DeleteProductById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return view.ErrNotFound()
			}
			return view.ErrInternalServer("DELETE_PRODUCT_FAIL", err.Error())
		}

		return view.SuccessFindAll(delete)
	}
	return nil
}

func (p *ProductServices) UpdateProductById(id string, req *params.ProductUpdate) *view.Response {
	product := req.ParseToModel()

	err := p.repo.UpdateProductById(id, product)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("UPDATE_PRODUCT_FAIL", err.Error())
	}

	return view.SuccessFindAll("UPDATE_PRODUCT_SUCCESS")
}

func (p *ProductServices) FindProductById(id string) *view.Response {
	product, err := p.repo.FindProductById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_PRODUCT_FAIL", err.Error())
	}
	return view.SuccessFindAll(product)
}
