package main

import "fmt"

func funcFor() {
	i := 1
	for i < 10 {
		fmt.Printf("%d,", i)
		i++
	}
	fmt.Printf("\n")
	for j := 2; j < 5; j++ {
		fmt.Printf("%d,", j)
	}
	fmt.Printf("\n")
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%d,", n)
	}
	fmt.Printf("\n")
}
