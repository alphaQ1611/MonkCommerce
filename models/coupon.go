package models

type CouponDetails interface {
	IsApplicable(cart *Cart) (bool, float32)
	ApplyCoupon(cart *Cart)
}

type CartWiseDetails struct {
	Threshold float32 `json:"threshold"`
	Discount  int     `json:"discount"`
}

type ProductWiseDetails struct {
	ProductID int `json:"product_id"`
	Discount  int `json:"discount"`
}

type BxGyProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type BxGyDetails struct {
	BuyProducts   []BxGyProduct `json:"buy_products"`
	GetProducts   []BxGyProduct `json:"get_products"`
	RepitionLimit int           `json:"repition_limit"`
}

type Coupon struct {
	Id            int
	CouponType    string        `json:"type"`
	CouponDetails CouponDetails `json:"details"`
}
