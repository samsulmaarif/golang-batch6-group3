package params

import (
	"golang-batch6-group3/server/model"
)

type TransactionCreate struct {
	Quantity  int    `validate:"required"`
	Ongkir    int    `validate:"required"`
	Ekspedisi string `validate:"required"`
	Estimasi  string `validate:"required"`
	UserId    string `validate:"required"`
	ProductId string `validate:"required"`
	Status    string `validate:"required"`
}

func (t *TransactionCreate) ParseToModel() *model.Transaction {
	return &model.Transaction{
		Quantity:  t.Quantity,
		Ongkir:    t.Ongkir,
		Ekspedisi: t.Ekspedisi,
		Estimasi:  t.Estimasi,
		UserId:    t.UserId,
		ProductId: t.ProductId,
		Status:    t.Status,
	}
}
