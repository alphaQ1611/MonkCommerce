package storage

import (
	"fmt"
	"monkcommerce/models"
)

var Coupons = make(map[int]models.Coupon)

func AddCoupon(coupon models.Coupon) error {
	Coupons[coupon.Id] = coupon
	return nil
}

func GetCoupons() ([]models.Coupon, error) {
	coupons := make([]models.Coupon, 0)

	for _, coupon := range Coupons {
		coupons = append(coupons, coupon)
	}
	return coupons, nil
}

func GetCouponByID(id int) (models.Coupon, error) {
	coupon, ok := Coupons[id]
	if !ok {
		return coupon, fmt.Errorf("Coupon not found")
	}
	return coupon, nil
}

func UpdateCoupon(id int, coupon models.Coupon) (models.Coupon, error) {
	Coupons[id] = coupon
	return coupon, nil
}

func DeleteCoupon(id int) error{
    _, ok := Coupons[id]
    if !ok {
        return fmt.Errorf("Coupon not found")
    }
     delete(Coupons, id)
    return nil

}
