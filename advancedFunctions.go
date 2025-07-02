package main

// func add(x, y int) int {
// 	return x + y
// }
// func mul(x, y int) int {
// 	return x * y
// }

// func sub(x, y int) int {
// 	return abs(x - y)
// }

// // abs returns the absolute value of an int
// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// func higherOrderFunc(x, y, z int, method func(int, int) int) int {
// 	return method(method(x, y), z)

// }

// func testHigherOrderFunc() {
// 	println(higherOrderFunc(1, 2, 3, add))  // Output: 6
// 	println(higherOrderFunc(10, 5, 2, sub)) // Output
// }

// higherOrder functions tak other functions as arguments and return the fuctions as return values
// they are used to create more generic and reusable code
// they can be used to create custom logic that can be applied to different data types or operations
// they are often used in functional programming paradigms

// Currying is a technique where a function takes a function or functions as arguments and returns a new function with some of the arguments pre-filled.

// func curry() {
// 	doubleFunc := self(add)
// 	squareFunc := self(mul)

// 	fmt.Println(doubleFunc(5))
// 	fmt.Println(squareFunc(5))

// }

// func self(mathfunc func(int, int) int) func(int) int {
// 	return func(x int) int {
// 		return mathfunc(x, x)
// 	}
// }

// defer keyword is like called when the function returns something so that we do not have to call same code multiple times if we have multiple return statements within the same function

// closures are functions that capture the variables from their surrounding context, allowing them to access and modify those variables even after the context has changed. They are often used to create functions with specific behavior or state.

// func closureExample() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }
// func testClosure() {
// 	counter := closureExample()
// 	fmt.Println(counter()) // Output: 1
// 	fmt.Println(counter()) // Output: 2
// 	fmt.Println(counter()) // Output: 3
// }
