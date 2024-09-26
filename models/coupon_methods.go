package models

func (cd *CartWiseDetails) IsApplicable(cart *Cart) (bool, float32) {
	cart.Total = calculateTotal(cart)
	return cart.Total >= cd.Threshold, cd.calculateDiscount(cart)
}

func (cd *CartWiseDetails) ApplyCoupon(cart *Cart) {
    cart.Total = calculateTotal(cart)
    cart.TotalDiscount = cd.calculateDiscount(cart)
    cart.FinalPrice = cart.Total - cart.TotalDiscount
}

func (pd *ProductWiseDetails) IsApplicable(cart *Cart) (bool, float32) {
	for _, item := range cart.Items {
		if item.Product_id == pd.ProductID {
			return true, pd.calculateDiscount(item)
		}
	}
	return false, 0
}


func (pd *ProductWiseDetails) ApplyCoupon(cart *Cart) {
    discount := float32(0)
    cart.Total = calculateTotal(cart)
    for i := range cart.Items {
        if cart.Items[i].Product_id == pd.ProductID {
            cart.Items[i].TotalDiscount = pd.calculateDiscount(cart.Items[i])
            discount += cart.Items[i].TotalDiscount
        }
    }

    cart.TotalDiscount = discount
    cart.FinalPrice = calculateTotal(cart) - discount
}

func (bx *BxGyDetails) IsApplicable(cart *Cart) (bool, float32) {
	cart.Total = calculateTotal(cart)
	buyProductMap := make(map[int]int)

	for _, item := range cart.Items {
		for _, buyProduct := range bx.BuyProducts {
			if item.Product_id == buyProduct.ProductID {
				buyProductMap[buyProduct.ProductID] += item.Quantity
			}
		}
	}

	totalBuyProducts := 0
	for _, buyProduct := range bx.BuyProducts {
		totalBuyProducts += buyProductMap[buyProduct.ProductID]
	}

	requiredBuyProducts := bx.BuyProducts[0].Quantity
	if totalBuyProducts < requiredBuyProducts {
		return false, 0
	}

	for _, getProduct := range bx.GetProducts {
		getProductInCart := false
		for _, cartItem := range cart.Items {
			if cartItem.Product_id == getProduct.ProductID && cartItem.Quantity >= getProduct.Quantity {
				getProductInCart = true
				break
			}
		}
		if !getProductInCart {
			return false, 0
		}
	}

	return true, bx.calculateDiscount(cart)
}

func (bx *BxGyDetails) ApplyCoupon(cart *Cart)  {
    discount := bx.calculateDiscount(cart)
    cart.TotalDiscount = discount
    cart.FinalPrice = calculateTotal(cart) - discount
}

func calculateTotal(cart *Cart) float32 {
	cartTotal := float32(0)
	for _, item := range cart.Items {
		cartTotal += item.Price * float32(item.Quantity)
	}
	return cartTotal
}

func (cd *CartWiseDetails) calculateDiscount(cart *Cart) float32 {
    return cart.Total * float32(cd.Discount) / 100
}

func (pd *ProductWiseDetails) calculateDiscount(item Product) float32 {
    return item.Price * float32(pd.Discount) / 100
}

func (bx *BxGyDetails) calculateDiscount(cart *Cart) float32 {
	buyProductMap := make(map[int]int)

	for _, item := range cart.Items {
		for _, buyProduct := range bx.BuyProducts {
			if item.Product_id == buyProduct.ProductID {
				buyProductMap[buyProduct.ProductID] += item.Quantity
			}
		}
	}

	totalBuyProducts := 0
	for _, buyProduct := range bx.BuyProducts {
		totalBuyProducts += buyProductMap[buyProduct.ProductID]
	}

	maxRepetitions := totalBuyProducts / bx.BuyProducts[0].Quantity
	if maxRepetitions > bx.RepitionLimit {
		maxRepetitions = bx.RepitionLimit
	}

	var totalDiscount float32
	for _, getProduct := range bx.GetProducts {
		for _, cartItem := range cart.Items {
			if cartItem.Product_id == getProduct.ProductID {
				discount := float32(getProduct.Quantity) * cartItem.Price * float32(maxRepetitions)
                cartItem.TotalDiscount = discount
				totalDiscount += discount
			}
		}
	}
	return totalDiscount
}
