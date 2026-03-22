package main

import "fmt"

func checkpalindrome() {

	s := "adadawdawdwada"
	fmt.Println(s)
	n := len(s)
	runes := []rune(s)

	for i := 0; i < n/2; i++ {
		j := n - i - 1
		if runes[i] != runes[j] {
			println("false")
			break
		}
	}
	println("true")

}
