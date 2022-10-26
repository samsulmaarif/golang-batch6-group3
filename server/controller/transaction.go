package controller

import (
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/service"
	"golang-batch6-group3/server/view"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	svc *service.TransactionServices
}

func NewTransactionHandler(svc *service.TransactionServices) *TransactionHandler {
	return &TransactionHandler{
		svc: svc,
	}
}

func (t *TransactionHandler) GinGetTransactions(c *gin.Context) {
	resp := t.svc.GetTransactions()

	WriteJsonResponseGin(c, resp)

}

func (t *TransactionHandler) GinGetMemberTransactions(c *gin.Context) {
	Id := c.GetString("USER_ID")
	resp := t.svc.GetMemberTransactions(Id)

	WriteJsonResponseGin(c, resp)

}

func (t *TransactionHandler) GinAddTransaction(c *gin.Context) {
	var req params.TransactionCreate
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

	resp := t.svc.CreateTransaction(&req)
	WriteJsonResponseGin(c, resp)
}

func (p *TransactionHandler) GinPackageTransactions(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("transaction Id not found in params"))
		return
	}

	resp := p.svc.UpdateTransactionStatusById(Id, "dikemas")

	WriteJsonResponseGin(c, resp)
}

func (p *TransactionHandler) GinSendTransactions(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("transaction Id not found in params"))
		return
	}

	resp := p.svc.UpdateTransactionStatusById(Id, "dikirim")

	WriteJsonResponseGin(c, resp)
}

func (p *TransactionHandler) GinConfirmTransactions(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("transaction Id not found in params"))
		return
	}

	resp := p.svc.UpdateTransactionStatusById(Id, "terkonfirmasi")

	WriteJsonResponseGin(c, resp)
}

func (t *TransactionHandler) GinGetTransactionById(c *gin.Context) {
	Id, isExist := c.Params.Get("Id")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("transaction Id not found in params"))
		return
	}

	resp := t.svc.FindTransactionById(Id)
	WriteJsonResponseGin(c, resp)
}
