package main

import "fmt"

func counter() {

	s := "swiss"

	mp := make(map[rune]int)
	runes := []rune(s)

	for _, x := range runes {
		mp[x]++
	}
	fmt.Println(mp)
}
