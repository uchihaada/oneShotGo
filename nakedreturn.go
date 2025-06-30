package main

// func ret() (a, b, c int) {
// 	a = 1
// 	b = 2
// 	c = 3
// 	return // naked return
// }

// func checknakedreturn() {
// 	a, b, c := ret()
// 	println(a, b, c) // Output: 1 2 3
// }

// For small functions, naked returns can be useful for readability,
// but they can also make the code less clear if overused or used in complex functions. It's generally recommended to use named returns judiciously and only when it enhances clarity.
// Naked returns are often used in Go for short functions where the return values are clear and straightforward.
// They allow the function to return its named return values without explicitly specifying them again in the return statement.
// However, for longer or more complex functions, it's usually better to use explicit return statements to improve code readability and maintainability.

// When there are many values are returned then named returns are useful.

// Guard clauses provides a linear flow of control and makes the code more readable.
// They allow you to handle special cases or errors early in the function, reducing the need for deeply nested code.
// Guard clause is an early return.
