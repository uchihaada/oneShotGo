package main

func fibbonacci(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibbonacci(n-1) + fibbonacci(n-2)
}
