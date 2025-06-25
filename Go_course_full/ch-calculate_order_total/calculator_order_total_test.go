package main

import (
	"math"
	"testing"
)

// helper function
func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

// main testing
func TestCalculeOrderTotal(t *testing.T) {
	// define test struct
	tests := []struct {
		base_price float64
		quantity   int
		tax_rate   float64
		expected   float64
	}{
		{100, 5, 0.1, 495.00},
		{50, 10, 1, 900.00},
		{50, 1, 0.5, 67.50},
	}

	//iterate for test
	for _, test := range tests {
		got := calculate_order_total(test.base_price, test.quantity, test.tax_rate)
		if !floatEquals(got, test.expected) {
			t.Errorf("CalculatedOrderTotal(%v, %v, %v) = %v, want %v", test.base_price, test.quantity, test.tax_rate, got, test.expected)

		}
	}
}
