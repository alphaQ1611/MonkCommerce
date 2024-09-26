package models

type Cart struct {
	Items         []*Product `json:"items"`
	Total         float32   `json:"total_price"`
	TotalDiscount float32   `json:"total_discount"`
	FinalPrice    float32   `json:"final_price"`
}
