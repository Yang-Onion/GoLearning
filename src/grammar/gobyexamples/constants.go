package main

import (
	"fmt"
	"math"
)

const hello string = "Hello,World"

func constants() {
	fmt.Println(hello)

	const n = 5000
	const PI = 3.1415
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
