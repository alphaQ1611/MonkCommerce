Coupon Types

1. Cart-wise Coupons

	•	Functionality: A discount is applied to the entire cart if the total amount exceeds a specified threshold.
	•	Use Cases:
	•	Valid cart total exceeds threshold: Apply discount.
	•	Cart total below threshold: No discount applied.

2. Product-wise Coupons

	•	Functionality: A discount is applied to specific products identified by their IDs.
	•	Use Cases:
	•	Discount applied to eligible products in the cart.
	•	Non-eligible products remain at full price.

3. Buy X, Get Y (BxGy) Coupons

	•	Functionality: Customers can buy a specified quantity of products from one array and get specified products from another array for free, with a repetition limit.
	•	Use Cases:
	•	Valid buy products present in the cart: Apply free products based on defined conditions.
	•	Not enough buy products: Coupon is not applicable.
	•	Exceeds repetition limit: Apply discount according to the limit.

Assumptions
	•	The system assumes that all product IDs are unique.
	•	The cart and coupon structures are well-defined and adhere to the expected formats.
    •	The system assumes that the coupon IDs are unique.

Limitations
	•	The system currently does not handle expired coupons. Adding expiration dates for coupons could enhance usability.
	

Implemented Cases

	1.	Cart-wise Coupons: The system correctly checks if the cart total exceeds the threshold and applies the discount.
	2.	Product-wise Coupons: Discounts are applied correctly to the specific products in the cart.
	3.	BxGy Coupons: Discounts are applied correctly to products in the cart based on the repetition limit.

Unimplemented Cases

	•	Expiration Handling: Implementing expiration dates for coupons would add complexity to the system but is essential for real-world applications.
    

Suggestions for Improvement

	•	Implement unit tests to ensure all functionality works as expected and to prevent regressions.
	•	Consider integrating a database to persist coupon information and support more complex queries.
	•	Extend the API to handle coupon expiration and user management for a more robust system.
