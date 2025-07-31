package main

import (
	"fmt"
	"slices"
)

// Golang Practice Set 3: Arrays & Slices

func ArraySlice() {
	// 1. Create and print an array
	var arr = []int{2, 3, 4}
	fmt.Println(arr)

	// 2. Iterate over array
	for _, x := range arr {
		fmt.Println(x)
	}

	// 3. Sum elements of array
	sum := 0
	for _, x := range arr {
		sum += x
	}

	// 4. Slice creation from array
	slc := arr[:2]
	fmt.Println(slc)

	// 5. Create and append to slice
	slice := make([]int, 3)
	slice = append(slice, slc...)
	fmt.Println(slice)

	// 6. Copy slice
	cpslice := make([]int, len(slice))
	copy(cpslice, slice)
	fmt.Println(cpslice)

	// 7. Slice length and capacity
	newslice := make([]int, 3)
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))
	newslice = append(newslice, 3)
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	// 8. Delete element from slice
	newslice = append(newslice[:1], newslice[2:]...)
	fmt.Println(newslice)

	// 9. Reverse a slice
	index := len(newslice)
	fmt.Println(index)
	slices.Reverse(newslice)
	fmt.Println(newslice)

	// 10. Find max in slice
	fmt.Println(slices.Max(newslice))

	// 11. Slice filter even numbers
	fil1 := []int{2, 4, 5, 5, 6, 0, 7}
	for i := len(fil1) - 1; i >= 0; i-- {
		if fil1[i]%2 != 0 {
			fil1 = append(fil1[:i], fil1[i+1:]...)
		}
	}
	fmt.Println(fil1)

	// 12. Slice map square
	var res []int
	for _, x := range fil1 {
		res = append(res, x*x)
	}
	fmt.Println(res)

	// 13. Compare two slices
	s1 := []int{1, 2, 3}
	s2 := []int{5, 3, 4}

	for i, x := range s1 {
		if x != s2[i] {
			fmt.Println(false)
		}
	}
	fmt.Println(true)

	// 14. Slice of slices (2D)
	sliceSlice := [][]int{{1, 2}, {2, 3}}
	for _, x := range sliceSlice {
		for _, y := range x {
			fmt.Println(y)
		}
	}

	// 15. Remove duplicates from slice
	fil3 := []int{1, 2, 3, 3, 4, 4, 5, 6}
	mapd := map[int]bool{}
	resnew := []int{}
	for _, x := range fil3 {
		if !mapd[x] {
			mapd[x] = true
			resnew = append(resnew, x)
		}
	}
	fmt.Println(resnew)

	// 16. Insert element in slice
	s3 := []int{3, 4, 5, 7}
	index1 := 2
	elem := 10
	s3 = append(s3[:index1+1], s3[index1:]...)
	s3[index1] = elem
	fmt.Println(s3)

	// 17. Merge two slices
	s1 = append(s1, s2...)
	fmt.Println(s1)

	// 18. Sort a slice
	slices.Sort(s1)
	fmt.Println(s1)

	// 19. Find index of element

	for i, x := range s3 {
		if x == elem {
			fmt.Println(i)
			break
		}
	}

	// 20. Split slice in two
	len := len(fil3) / 2

	f1 := fil3[:len]
	f2 := fil3[len:]
	fmt.Println(f1, f2)
}
