// package main

// // Example of implementing constraints with Go generics

// import "fmt"

// // Custom constraint for types that can calculate area and perimeter
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Rectangle implements Shape
// type Rectangle struct {
// 	Width, Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// // Circle implements Shape
// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14159 * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * 3.14159 * c.Radius
// }

// // Generic function using custom constraint
// func PrintShapeInfo[T Shape](s T) {
// 	fmt.Printf("Area: %.2f\n", s.Area())
// 	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
// }

// func example() {
// 	r := Rectangle{Width: 5, Height: 3}
// 	c := Circle{Radius: 2.5}

// 	PrintShapeInfo(r)
// 	PrintShapeInfo(c)
// }

package main

import (
	"fmt"
)

// Numeric is a parametric constraint that allows int, float64, or int64
type Numeric interface {
	int | float64 | int64
}

// Sum adds up all elements in a slice of Numeric type
func Sum[T Numeric](slice []T) T {
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

func example() {
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.1, 2.2, 3.3}
	int64s := []int64{10, 20, 30}

	fmt.Println("Sum of ints:", Sum(ints))     // Output: 10
	fmt.Println("Sum of floats:", Sum(floats)) // Output: 6.6
	fmt.Println("Sum of int64s:", Sum(int64s)) // Output: 60
}
