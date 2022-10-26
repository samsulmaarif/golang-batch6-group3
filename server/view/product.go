package view

import "golang-batch6-group3/server/model"

type ProductCreateResponse struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
	ImgUrl      string `json:"img_url"`
}

func NewProductCreateResponse(product *model.Product) *ProductCreateResponse {
	return &ProductCreateResponse{
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImgUrl:      product.ImgUrl,
	}
}

type ProductFindAllResponse struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
	ImgUrl      string `json:"img_url"`
}

func NewProductFindAllResponse(products *[]model.Product) []ProductFindAllResponse {
	var productsFind []ProductFindAllResponse
	for _, product := range *products {
		productsFind = append(productsFind, *parseModelToProductFind(&product))
	}
	return productsFind
}

func parseModelToProductFind(product *model.Product) *ProductFindAllResponse {
	return &ProductFindAllResponse{
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImgUrl:      product.ImgUrl,
	}
}
