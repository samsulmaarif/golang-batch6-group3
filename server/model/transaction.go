package model

type Transaction struct {
	BaseModel
	Quantity  int    `json:"quantity"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`

	User    User    `json:"user"`
	Product Product `json:"product"`
}
