package view

import "golang-batch6-group3/server/model"

type TransactionAddResponse struct {
	Quantity  int    `json:"quantity"`
	Ongkir    int    `json:"ongkir"`
	Ekspedisi string `json:"ekspedisi"`
	Estimasi  string `json:"estimasi"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Status    string `json:"status"`
}

func NewTransactionAddResponse(transaction *model.Transaction) *TransactionAddResponse {
	return &TransactionAddResponse{
		Quantity:  transaction.Quantity,
		Ongkir:    transaction.Ongkir,
		Ekspedisi: transaction.Ekspedisi,
		Estimasi:  transaction.Estimasi,
		UserId:    transaction.UserId,
		ProductId: transaction.ProductId,
		Status:    transaction.Status,
	}
}

type TransactionFindAllResponse struct {
	Quantity  int    `json:"quantity"`
	Ongkir    int    `json:"ongkir"`
	Ekspedisi string `json:"ekspedisi"`
	Estimasi  string `json:"estimasi"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Status    string `json:"status"`
}

func NewTransactionFindAllResponse(transactions *[]model.Transaction) []TransactionFindAllResponse {
	var transactionsFind []TransactionFindAllResponse
	for _, transaction := range *transactions {
		transactionsFind = append(transactionsFind, *parseModelToTransactionFind(&transaction))
	}
	return transactionsFind
}

func parseModelToTransactionFind(transaction *model.Transaction) *TransactionFindAllResponse {
	return &TransactionFindAllResponse{
		Quantity:  transaction.Quantity,
		Ongkir:    transaction.Ongkir,
		Ekspedisi: transaction.Ekspedisi,
		Estimasi:  transaction.Estimasi,
		UserId:    transaction.UserId,
		ProductId: transaction.ProductId,
		Status:    transaction.Status,
	}
}
