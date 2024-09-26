package main

import (
	"monkcommerce/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/coupons", couponsHandler)
	http.HandleFunc("/coupons/", couponByIDHandler)
	http.HandleFunc("/applicable-coupons", handler.HandleGetApplicableCoupons)
	http.HandleFunc("/apply-coupon/", handler.HandleApplyCoupon)
	http.ListenAndServe(":8080", nil)
}

func couponsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.HandleCreateCoupon(w, r)
	case http.MethodGet:
		handler.HandleGetCoupon(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func couponByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.HandleGetCouponByID(w, r)

	case http.MethodPut:
		handler.HandleUpdateCoupon(w, r)
	case http.MethodDelete:
		handler.HandleDeleteCoupon(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
