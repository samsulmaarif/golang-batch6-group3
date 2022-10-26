package controller

import (
	"fmt"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/service"
	"golang-batch6-group3/server/view"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserServices
}

func NewUserHandler(svc *service.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) GinGetUsers(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := u.svc.GetUsers()

	WriteJsonResponseGin(c, resp)

}

func (u *UserHandler) GinRegister(c *gin.Context) {
	var req params.UserCreate
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

	// if len(req.Fullname) < 5 {

	// 	resp := view.ErrBadRequest("user name length must be greater than 4")
	// 	WriteJsonResponseGin(c, resp)
	// 	return
	// }

	resp := u.svc.CreateUser(&req)
	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) GinLogin(c *gin.Context) {
	var req params.UserLogin
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

	resp := u.svc.Login(&req)
	WriteJsonResponseGin(c, resp)
}
