package products

type ProductDTO struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Price int    `json:"price"`
}
