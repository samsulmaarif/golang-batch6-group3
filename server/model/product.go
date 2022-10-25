package model

type Product struct {
	BaseModel
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"decription"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
	ImgUrl      string `json:"img_url"`
}
