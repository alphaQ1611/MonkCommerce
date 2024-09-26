package service

import (
	"encoding/json"
	"fmt"
	"monkcommerce/models"
	"monkcommerce/storage"

	"github.com/google/uuid"
)

// Temporary structure to hold raw `details`
type CouponTemp struct {
	Id         string          `json:"Id"`
	CouponType string          `json:"type"`
	Details    json.RawMessage `json:"details"` // RawMessage for unmarshalling without knowing the Details struct
}

func CreateCoupon(body []byte) (models.Coupon, error) {
	var tempCoupon CouponTemp
	var coupon models.Coupon

	err := json.Unmarshal(body, &tempCoupon)
	if err != nil {
		return coupon, fmt.Errorf("Error while unmarshalling main coupon: %v", err)
	}

	coupon.CouponType = tempCoupon.CouponType

	switch tempCoupon.CouponType {
	case "cart-wise":
		var details models.CartWiseDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling cart-wise details: %v", err)
		}
		coupon.CouponDetails = details

	case "product-wise":
		var details models.ProductWiseDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling product-wise details: %v", err)
		}
		coupon.CouponDetails = details

	case "bxgy":
		var details models.BxGyDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling bxgy details: %v", err)
		}
		coupon.CouponDetails = details

	default:
		return coupon, fmt.Errorf("Invalid coupon type")
	}

	// Set a new coupon ID and add to storage, simulating an id here from 0 to 100
	coupon.Id = int(uuid.New().ID()) % 100
	storage.AddCoupon(coupon)

	// Print storage for debugging
	for k, v := range storage.Coupons {
		fmt.Println(k, v)
	}
	return coupon, nil
}

func GetCoupons() ([]models.Coupon, error) {
	return storage.GetCoupons()
}

func GetCouponByID(id int) (models.Coupon, error) {
	return storage.GetCouponByID(id)
}

func UpdateCoupon(id int, body []byte) (models.Coupon, error) {
	coupon, err := storage.GetCouponByID(id)
	if err != nil {
		return coupon, err
	}
	var tempCoupon CouponTemp

	err = json.Unmarshal(body, &tempCoupon)
	if err != nil {
		return coupon, fmt.Errorf("Error while unmarshalling main coupon: %v", err)
	}

	coupon.CouponType = tempCoupon.CouponType

	switch tempCoupon.CouponType {
	case "cart-wise":
		var details models.CartWiseDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling cart-wise details: %v", err)
		}
		coupon.CouponDetails = details

	case "product-wise":
		var details models.ProductWiseDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling product-wise details: %v", err)
		}
		coupon.CouponDetails = details

	case "bxgy":
		var details models.BxGyDetails
		if err := json.Unmarshal(tempCoupon.Details, &details); err != nil {
			return coupon, fmt.Errorf("Error while unmarshalling bxgy details: %v", err)
		}
		coupon.CouponDetails = details

	default:
		return coupon, fmt.Errorf("Invalid coupon type")
	}

	storage.UpdateCoupon(id, coupon)

	return coupon, nil
}

func DeleteCoupon(id int) error {
	return storage.DeleteCoupon(id)
}

type CartRequest struct {
	Cart models.Cart `json:"cart"`
}

func GetApplicableCoupons(body []byte) ([]models.ApplicableCoupon, error) {
	var cart models.Cart
	var cartRequest CartRequest

	err := json.Unmarshal(body, &cartRequest)
	cart = cartRequest.Cart
	if err != nil {
		fmt.Println("Here")
		return nil, fmt.Errorf("Invalid cart data")
	}

	allCoupons, _ := storage.GetCoupons()
	if len(allCoupons) == 0 {
		return nil, nil
	}

	coupons := []models.ApplicableCoupon{}

    for _, coupon := range allCoupons {
        ok, discount := coupon.CouponDetails.(models.CouponDetails).IsApplicable(&cart)
        if ok {
            applicableCoupon := models.ApplicableCoupon{
                CouponID: coupon.Id,
                Type:     coupon.CouponType,
                Discount: discount,
            }
            coupons = append(coupons, applicableCoupon)
        }
    }

	return coupons, nil
}

func ApplyCoupon(id int, body []byte) (models.Cart, error) {
	var cart models.Cart
	var cartRequest CartRequest

	err := json.Unmarshal(body, &cartRequest)
	cart = cartRequest.Cart
	if err != nil {
		return models.Cart{}, fmt.Errorf("Invalid cart data")
	}
	coupon, err := storage.GetCouponByID(id)
	if err != nil {
		return models.Cart{}, fmt.Errorf("Coupon not found")
	}
	coupon.CouponDetails.(models.CouponDetails).ApplyCoupon(&cart)

	return cart, nil
}
