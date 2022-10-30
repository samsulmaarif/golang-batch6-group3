package controller

import (
	"fmt"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/service"
	"golang-batch6-group3/server/view"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *service.ProductServices
}

func NewProductHandler(svc *service.ProductServices) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

func (p *ProductHandler) GinGetProducts(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := p.svc.GetProducts()

	WriteJsonResponseGin(c, resp)

}

func (p *ProductHandler) GinAddProduct(c *gin.Context) {
	var req params.ProductCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := p.svc.AddProduct(&req)
	WriteJsonResponseGin(c, resp)
}

func (p *ProductHandler) GinDeleteProduct(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("product Id not found in params"))
		return
	}
	p.svc.DeleteProductById(Id)
	WriteJsonResponseGin(c, view.SuccessDeleted("DELETE_PRODUCT_SUCCESS", Id))
}

func (p *ProductHandler) GinUpdateProduct(c *gin.Context) {
	var req params.ProductUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("product Id not found in params"))
		return
	}

	p.svc.UpdateProductById(Id, &req)
	WriteJsonResponseGin(c, view.SuccessUpdated("UPDATE_PRODUCT_SUCCESS", req))
}

func (p *ProductHandler) GetProductById(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("product Id not found in params"))
		return
	}

	resp := p.svc.FindProductById(Id)
	WriteJsonResponseGin(c, resp)
}
