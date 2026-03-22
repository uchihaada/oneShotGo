package main

import (
	"fmt"
)

func twosum() {

	nums := []int{2, 7, 11, 15}
	target := 9

	m := make(map[int]int)

	for i, v := range nums {

		c := target - v

		if idx, ok := m[c]; ok {
			fmt.Println(idx, i)
			break
		}
		m[v] = i
	}
}
