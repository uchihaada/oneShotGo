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

}

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
