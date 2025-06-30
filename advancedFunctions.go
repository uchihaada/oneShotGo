package main

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return abs(x - y)
}

// abs returns the absolute value of an int
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func higherOrderFunc(x, y, z int, method func(int, int) int) int {
	return method(method(x, y), z)

}

func testHigherOrderFunc() {
	println(higherOrderFunc(1, 2, 3, add))  // Output: 6
	println(higherOrderFunc(10, 5, 2, sub)) // Output
}

// higherOrder functions tak other functions as arguments and return the fuctions as return values
// they are used to create more generic and reusable code
// they can be used to create custom logic that can be applied to different data types or operations
// they are often used in functional programming paradigms
