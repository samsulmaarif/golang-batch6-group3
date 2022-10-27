package controller

import (
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/service"
	"golang-batch6-group3/server/view"

	"github.com/gin-gonic/gin"
)

type RajaOngkirHandler struct {
	svc *service.RajaOngkirServices
}

func NewRajaOngkirHandler(svc *service.RajaOngkirServices) *RajaOngkirHandler {
	return &RajaOngkirHandler{
		svc: svc,
	}
}

func (ro *RajaOngkirHandler) GetProvinceById(c *gin.Context) {
	var req params.RajaOngkirQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}
	resp := ro.svc.FindProvinceById(&req)
	WriteJsonResponseGin(c, resp)
}

func (ro *RajaOngkirHandler) GetCityById(c *gin.Context) {
	var req params.RajaOngkirQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}
	resp := ro.svc.FindCityById(&req)
	WriteJsonResponseGin(c, resp)
}
