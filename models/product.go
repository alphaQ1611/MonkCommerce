package models

type Product struct {
	Product_id    int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float32 `json:"price"`
	TotalDiscount float32 `json:"total_discount"`
}
