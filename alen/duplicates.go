package main

func duplicates() {

	a := []int{1, 2, 4, 4, 34, 4, 4, 4, 4, 3, 3, 5}

	mapp := make(map[int]int)

	for _, v := range a {
		mapp[v] += 1
		if mapp[v] == 2 {
			println(v)
		}
	}

}
