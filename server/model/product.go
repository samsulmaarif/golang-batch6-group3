package model

type Product struct {
	BaseModel
	Name        string `json:"name"`
	Category    string `json:"category"`
	Weight      int    `json:"weight"`
	Description string `json:"decription"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	ImgUrl      string `json:"img_url"`
}
