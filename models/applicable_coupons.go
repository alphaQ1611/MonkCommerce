package models

type ApplicableCoupon struct {
	CouponID int     `json:"coupon_id"`
	Type     string  `json:"type"`
	Discount float32 `json:"discount"`
}

type ApplicableCouponsResponse struct {
	ApplicableCoupons []ApplicableCoupon `json:"applicable_coupons"` 
}
