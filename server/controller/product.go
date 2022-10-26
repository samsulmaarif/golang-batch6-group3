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
