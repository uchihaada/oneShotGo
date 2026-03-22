package main

import (
	"fmt"
)

func reversestring() {

	var s string
	s = "saawadawd"
	n := len(s)
	fmt.Println(s)
	// var res string

	// for i := len(s) - 1; i >= 0; i-- {
	// 	res += string(s[i])
	// }
	// fmt.Print(res)

	runes := []rune(s)
	fmt.Println(runes)

	i := 0

	for i = 0; i < n/2; i++ {
		j := n - i - 1
		fmt.Println(i, j)
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(runes)
	fmt.Println(string(runes))

}
