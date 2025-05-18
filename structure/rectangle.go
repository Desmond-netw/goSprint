// go method to struct

package main

import "fmt"

type rect struct {
	width, height int
}

// method to find area
func (r *rect) area() int {
	rectArea := r.width * r.height

	return rectArea
}

// method to find perimeter of rectangle
func (r rect) peri() int {

	perimeter := 2*r.width + 2*r.height

	return perimeter
}

// funct to print
func main() {
	r := rect{width: 20, height: 10}
	fmt.Println("Area of Rectangle:", r.area())
	fmt.Println("Perimeter of Rectangle:", r.peri())
}
