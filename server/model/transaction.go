package model

type Transaction struct {
	BaseModel
	Quantity  int    `json:"quantity"`
	Ongkir    int    `json:"ongkir"`
	Ekspedisi string `json:"ekspedisi"`
	Estimasi  string `json:"estimasi"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Status    string `json:"status"`

	User    User    `json:"user"`
	Product Product `json:"product"`
}
