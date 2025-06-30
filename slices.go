package main

import "fmt"

// func slicesExample() {
// 	s := []int{1, 2, 3, 4}
// 	fmt.Println(s)

// 	// get slice of the array
// 	subSlice := s[1:3]
// 	fmt.Println(subSlice)

// 	// append new element
// 	s = append(s, 56)
// 	fmt.Println(s)

// 	// creating dynamic slices
// 	dynamicSlice := make([]int, len(s))
// 	fmt.Println(dynamicSlice)
// 	fmt.Println(len(dynamicSlice), cap(dynamicSlice))
// 	dynamicSlice = append(dynamicSlice, 10, 20, 30)
// 	fmt.Println(dynamicSlice)
// 	fmt.Println(len(dynamicSlice), cap(dynamicSlice))

// 	// length and capacity are different.
// 	// Length is the number of elements in the slice.
// 	// Capacity is the number of elements the slice can hold before it needs to be resized.

// 	// copying slices
// 	copySlices := make([]int, len(s))
// 	copy(copySlices, s)

// 	// Removing elements from a slice
// 	fmt.Println("s:", s)
// 	s1 := append(s[:1], s[3:]...)
// 	fmt.Println(s1)

// 	// Iterating over a slice

// 	for i, v := range s {
// 		fmt.Printf("%d element is %d\n", i, v)
// 	}

// }

// slices references arrays, so if you modify the slice, the original array is also modified.
// slices are reference types, so they point to the same underlying array.
// slices are more flexible than arrays, as they can grow and shrink in size dynamically.
// slices are more efficient than arrays, as they do not require copying the entire array when passing it to a function.

// Multidimensional slices can point to same arrays.

// a function that takes a slice as an argument can modify the original slice, but it cannot change the length or capacity of the slice.

// cap() returns the capacity of the slice, which is the number of elements the slice can hold before it needs to be resized.

// len() returns the length of the slice, which is the number of elements in the slice.

// len() and cap() are safe function so if they are empty // they will return 0 and not panic.

// Variadic functions can take a variable number of arguments, which are passed as a slice.

// sum is a variadic function that takes any number of int arguments and returns their sum.
// The ...int syntax means the function can accept zero or more ints, which are available as a slice inside the function.
// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

// Example usage:
// result := sum(1, 2, 3, 4) // result will be 10
// result2 := sum() // result2 will be 0, as no arguments are passed

// Spread operator can be used to pass a slice as individual arguments to a function.

// func printString(strs ...string) {
// 	for _, str := range strs {
// 		fmt.Println(str)
// 	}
// }

// Example usage:
// strings := []string{"Hello", "World", "Go"}
// printString(strings...) // This will print each string in the slice on a new line

// type Cost struct {
// 	day   int
// 	value float64
// }

// func practiceSlice(cost []Cost) []float64 {

// 	var result []float64

// 	// Use a slice to accumulate values by day (assuming days are 0 to 5)
// 	dayTotals := make([]float64, 6)
// 	for _, c := range cost {
// 		dayTotals[c.day] += c.value
// 	}
// 	result = append(result, dayTotals...)
// 	return result

// }

// func testPractice() {
// 	var cost []Cost
// 	cost = append(cost,
// 		Cost{day: 0, value: 5.0},
// 		Cost{day: 1, value: 10.5},
// 		Cost{day: 1, value: 12.0},
// 		Cost{day: 5, value: 20.0},
// 	)
// 	fmt.Println(practiceSlice(cost))

// }

// Slices of slices

func matrix() {
	matrix := make([][]int, 0)

	for i := 0; i <= 10; i++ {
		row := make([]int, 0)
		for j := 0; j <= 10; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	for i, row := range matrix {
		fmt.Println("Row", i, ":", row)
	}
}

// when using append i should always assign the result back to the slice variable.
// This is because append may return a new slice if the underlying array needs to be resized.
