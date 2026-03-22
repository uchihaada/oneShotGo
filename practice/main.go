package main

import (
	"fmt"
)

// package level variable
// var ada int = 15

// type Person struct {
// 	name  string
// 	age   int
// 	email string
// }

// func (p *Person) fmtstring() (a string) {
// 	a = fmt.Sprintf("%s (%d) - %s", p.name, p.age, p.email)
// 	return a
// }

// // constructor function
// func NewPerson(name string, age int, email string) *Person {
// 	return &Person{
// 		name:  name,
// 		age:   age,
// 		email: email,
// 	}
// }

// type Rectangle struct {
// 	width  float64
// 	height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.width * r.height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.width + r.height)
// }

// type ColoredRectangle struct {
// 	Rectangle // embedded struct
// 	color     string
// }

// type Shape interface {
// 	Area() float64
// }

// func shapeArea(s Shape) {
// 	fmt.Println(s.Area())
// }

// type Circle struct {
// 	r float64
// }

// func (c *Circle) Area() float64 {
// 	return 3.14 * c.r * c.r
// }

// type Square struct {
// 	a float64
// }

// func (s *Square) Area() float64 {
// 	return s.a * s.a
// }

// type CustomError struct {
// 	a   int
// 	msg string
// }

// func (c *CustomError) Error() string {
// 	return fmt.Sprintf("val :%d , error :%s ", c.a, c.msg)
// }

func main() {
	fmt.Println("golu")

	// Question 1: Create a program that declares variables of different types (int, float64, string, bool) and prints them with their zero values. Then reassign them with actual values and print again.

	// var a int
	// var b float64
	// var c string
	// var d bool
	// fmt.Println(a, b, c, d)
	// a = 2
	// b = 3.45
	// c = "ada"
	// d = true
	// fmt.Println(a, b, c, d)

	// Write a program that demonstrates the difference between var, short declaration (:=), and constants. Create examples of each and show a scenario where you MUST use var instead of :=.

	// var a int = 32
	// b := 45.00
	// const c = 34

	// Write a program that takes a number and uses if/else statements to determine if it's positive, negative, or zero. Then use a for loop to print all even numbers from 2 to 20.

	// check(3)

	// Create a program that uses a switch statement to convert a number (1-7) to the corresponding day of the week. Also, write a for loop that breaks when it finds the first number divisible by both 3 and 5 between 1 and 100.

	// for day := 1; day <= 7; day++ {

	// 	switch day {
	// 	case 1:
	// 		fmt.Println("Monday")
	// 	case 2:
	// 		fmt.Println("Tuesday")
	// 	case 3:
	// 		fmt.Println("Wednesday")
	// 	case 4:
	// 		fmt.Println("Thursday")
	// 	case 5:
	// 		fmt.Println("Friday")
	// 	case 6:
	// 		fmt.Println("Saturday")
	// 	default:
	// 		fmt.Println("Sunday")
	// 	}
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%15 == 0 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	// Write a function that takes two integers and returns their sum, difference, and product as multiple return values. Also create a variadic function that calculates the average of any number of float64 values.

	// fmt.Println(calculator(2, 3))
	// fmt.Println(variadic(2.5, 3.4, 4.4))

	// Create a function that demonstrates named return values. Write a function called divide that takes two integers and returns the quotient and remainder. If the divisor is zero, return 0 for both values. Use named returns and show how they work.

	// fmt.Println(divide(1, 2))
	// fmt.Println(quotient(2, 4))

	//  Create an array of 5 integers, fill it with values, then create a slice from the middle 3 elements. Demonstrate appending to the slice and show how it affects (or doesn't affect) the original array.

	// var arr = [5]int{7, 1, 2, 5, 3}
	// // arr = append(arr, 7)
	// // arr = append(arr, 1)
	// // arr = append(arr, 0)
	// // arr = append(arr, 5)
	// // arr = append(arr, 2)
	// fmt.Println(arr)
	// mid := len(arr) / 2
	// arr1 := arr[mid-1 : mid+2]
	// arr1 = append(arr1, 48)
	// fmt.Println("Original array after append:", arr)
	// fmt.Println("Modified slice after append:", arr1)

	// Write a function that takes a slice of integers and returns a new slice containing only the even numbers. Also demonstrate slice operations: make(), copy(), and show how to delete an element from a slice by index.
	// fmt.Println(evenSlice(arr[:]))
	// fmt.Println(deleteByIndex(arr[:], 1))

	// Create a map that stores student names as keys and their grades as values. Write functions to: add a student, get a student's grade, and check if a student exists. Also demonstrate the "comma ok" idiom when accessing map values.

	// mapStudent := make(map[string]rune)
	// var name string = "ada"
	// var grade rune = 'a'
	// addStudent(mapStudent, name, grade)
	// fmt.Println(getStudent(mapStudent,name))

	// Write a program that counts the frequency of words in a string. Create a map where keys are words and values are their counts. Also show how to iterate over a map and delete entries where the count is less than 2.

	// var st string = "golu ada golu ada car bike bike ada golu car tree phone"
	// a := strings.Split(st, " ")
	// mapstr := make(map[string]int)

	// for _, item := range a {
	// 	mapstr[item]++
	// }

	// for i, x := range mapstr {
	// 	if x < 2 {
	// 		delete(mapstr, i)
	// 	}
	// }
	// fmt.Println(mapstr)

	// Create a Person struct with fields: Name (string), Age (int), and Email (string). Write a method for the struct that returns a formatted string like "John (25) - john@email.com". Also create a constructor function that returns a new Person.

	// p1 := NewPerson("golu", 25, "john@email.com")
	// fmt.Println(p1.fmtstring())

	// Create a Rectangle struct with Width and Height fields. Add methods to calculate Area and Perimeter. Then create an embedded struct ColoredRectangle that embeds Rectangle and adds a Color field. Show how embedded methods work.

	// rect := Rectangle{width: 5, height: 10}
	// fmt.Println("Area of Rectangle:", rect.Area())
	// fmt.Println("Perimeter of Rectangle:", rect.Perimeter())
	// coloredRect := ColoredRectangle{
	// 	Rectangle: Rectangle{width: 4, height: 8},
	// 	color:     "red",
	// }
	// fmt.Println("Area of Colored Rectangle:", coloredRect.Area())
	// fmt.Println("Perimeter of Colored Rectangle:", coloredRect.Perimeter())
	// fmt.Println("Color of Colored Rectangle:", coloredRect.color)

	// Write a function that takes a pointer to an integer and doubles its value. Then create another function that swaps two integers using pointers. Demonstrate the difference between passing by value vs passing by pointer.

	// var single int = 2
	// double(&single)
	// fmt.Println(single)
	// doubleEx(single)
	// fmt.Println(single)

	// a := 1
	// b := 2
	// swap(&a, &b)
	// fmt.Println(a, b)

	// Create a function that returns a pointer to a newly created integer. Show how to work with nil pointers safely - write a function that takes a pointer to an integer and safely prints its value (checking for nil first).

	// ptr := createInt(42)
	// printPointer(ptr)

	//  Create an interface called Shape with a method Area() float64. Then create two structs (Circle and Square) that implement this interface. Write a function that takes a Shape interface and prints its area.

	// c := Circle{r: 12.4}
	// s := Square{a: 23.3}
	// shapeArea(&s)
	// shapeArea(&c)

	// Create an empty interface (interface{}) function that can accept any type and uses type assertion to handle different types (int, string, bool) differently. Also demonstrate type switches to handle multiple types in one function.

	// assertioncall("hello")
	// assertioncall(34)
	// assertioncall(true)
	// switchcall("hello")
	// switchcall(34)
	// switchcall(true)

	// Create a function that divides two numbers and returns both the result and an error (if division by zero). Use Go's idiomatic error handling pattern. Then write code that calls this function and properly handles the error.

	// val, err := divide(3.0, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(val)
	// }

	//  Create a custom error type by implementing the error interface. Make a ValidationError struct that holds both an error message and a field name. Write a function that validates user input and returns this custom error when validation fails.

	// val, err := divideCustomError(3.0, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(val)
	// }

	// Create a goroutine that calculates the sum of numbers from 1 to 1000 and sends the result through a channel. In your main function, start the goroutine and receive the result from the channel.

	// ch := make(chan int)

	// go func() {
	// 	sum := 0
	// 	for i := 1; i <= 1000; i++ {
	// 		sum += i
	// 	}
	// 	ch <- sum
	// 	close(ch)
	// }()
	// res := <-ch
	// fmt.Println(res)

	// Create a program that uses multiple goroutines and channels. Launch 3 goroutines that each calculate squares of numbers (1-5, 6-10, 11-15) and send results to a shared channel. Use a buffered channel and demonstrate channel direction (send-only, receive-only parameters).

	// ch := make(chan int, 15)
	// go calculateSquare(ch, 1, 5)
	// go calculateSquare(ch, 6, 10)
	// go calculateSquare(ch, 11, 15)

	// for i := 0; i < 15; i++ {
	// 	fmt.Println(<-ch)
	// }

}

// func calculateSquare(ch chan<- int, a, b int) {
// 	for i := a; i <= b; i++ {
// 		ch <- i * i
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func divideCustomError(a, b int) (int, error) {
// 	if b == 0 {t
// 		return 0, &CustomError{
// 			a:   b,
// 			msg: "divided by zero",
// 		}
// 	}

// 	return a / b, nil
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("divided by zero")
// 	} else {
// 		return (a / b), nil
// 	}
// }

// func assertioncall(data interface{}) {

// 	if str, ok := data.(string); ok {
// 		fmt.Println(str)
// 	} else if intiger, ok := data.(int); ok {
// 		fmt.Println(intiger)
// 	} else if b, ok := data.(bool); ok {
// 		fmt.Println(b)
// 	}
// }

// func switchcall(data interface{}) {
// 	switch v := data.(type) {
// 	case int:
// 		fmt.Println(v)
// 	case string:
// 		fmt.Println(v)
// 	case bool:
// 		fmt.Println(v)
// 	}
// }

// func check(a int) {
// 	if a > 0 {
// 		fmt.Println("Positive")
// 	} else if a < 0 {
// 		fmt.Println("Negative")
// 	} else {
// 		fmt.Println("Zero")
// 	}
// 	for i := 2; i <= 20; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func calculator(a, b int) (int, int, int) {
// 	var sub int
// 	if a > b {
// 		sub = a - b
// 	} else {
// 		sub = b - a
// 	}
// 	return (a + b), (sub), (a * b)
// }

// func variadic(args ...float64) float64 {
// 	var count int
// 	var avg float64
// 	// for _, arg := range args {
// 	// 	v, _ := arg.(int)
// 	// 	// fmt.Println(v)
// 	// 	avg += float64(v)
// 	// 	count++
// 	// }

// 	for _, arg := range args {
// 		avg += arg
// 		count++
// 	}
// 	return (avg / float64(count))

// }

// func divide(a, b int) (int, int) {
// 	return a / b, a % b
// }

// func quotient(a, b int) (r1, r2 int) {
// 	if b == 0 {
// 		return r1, r2
// 	}
// 	r1 = a / b
// 	r2 = a % b
// 	return r1, r2

// }

// func evenSlice(a []int) []int {
// 	var b []int
// 	for _, x := range a {
// 		if x%2 == 0 {
// 			b = append(b, x)
// 		}
// 	}
// 	res := make([]int, len(b))
// 	copy(res, b)
// 	return res
// }

// func deleteByIndex(b []int, a int) []int {
// 	b = append(b[:a], b[a+1:]...)
// 	return b

// }

// func addStudent(a map[string]rune, b string, c rune) {
// 	a[b] = c
// }

// func getStudent(b map[string]rune,a string) (c rune) {

// 	c,ok := b[a]
// 	if ok {
// 		return c
// 	}
// 	return 0
// }

// func double(single *int) {
// 	*single *= 2
// }

// func swap(a, b *int) {

// 	temp := *a
// 	*a = *b
// 	*b = temp
// }

// func doubleEx(a int) {
// 	a *= 2
// }

// func createInt(a int) *int {
// 	var new int = a
// 	return &new
// }

// func printPointer(a *int) {
// 	if a == nil {
// 		fmt.Println("Pointer is nil")
// 		return
// 	}
// 	fmt.Println("Pointer value:", *a)

// }
