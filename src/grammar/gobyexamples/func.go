package main

import "fmt"

func plus(a, b int) {
	fmt.Println(a + b)
}
func fibonaci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return fibonaci(n-1) + fibonaci(n-2)
}
