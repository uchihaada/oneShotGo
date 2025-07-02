package main

// import "fmt"

// func pointers() {

// 	x := 5
// 	fmt.Println(x)
// 	y := &x // y is a pointer to x
// 	fmt.Println(y)
// 	fmt.Println(*y)
// 	*y = 23
// 	fmt.Println(x)
// }

// Pointers are variables that store the memory address of another variable.
// They allow you to directly manipulate the value stored in that memory address.
// Pointers are useful for passing large data structures to functions without copying them, and for modifying the original data in place.
// In Go, pointers are declared using the `*` operator, and you can get the address of a variable using the `&` operator.
// When you dereference a pointer using the `*` operator, you can access or modify the value stored at that memory address.
// Pointers can be nil, meaning they do not point to any valid memory address.
// Pointers are commonly used in Go for performance optimization, especially when dealing with large structs or arrays.
