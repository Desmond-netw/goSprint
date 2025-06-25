package main

import (
	"fmt"
	"math"
)

/*
 1. initilize variables
 - base_price float
 -quantity float
 -tax_rate float
 -DISCOUNT_QTY 5
 -DISCOUNT_rate 0.1
 2. let subtotal = base_price * quantity
 3. if subtotal >= DISCOUNT_QTY
 - dicounted = subtotal * DISCOUNT_RATE
 -subtotal -= discounted
 4. final-total = subtotal + (subtotal * tax_rate)
 5.Print final_total
*/

const (
	DISCOUNT_QTY  = 5
	DISCOUNT_RATE = 0.1
)

func calculate_order_total(base_price float64, quantity int, tax_rate float64) float64 {
	// initilize subtotal
	subtotal := base_price * float64(quantity)

	// APPLY discount to subtotal
	if subtotal >= DISCOUNT_QTY {
		discount := subtotal * DISCOUNT_RATE
		// apply discount subtotal
		subtotal -= discount
	}
	// apply tax
	tax := subtotal * tax_rate
	final_total := subtotal + tax

	return math.Round(final_total*100) / 100
}
func main() {
	price_total := calculate_order_total(100, 5, 0.2)
	fmt.Printf("Total Price : %.2f\n", price_total)
}
