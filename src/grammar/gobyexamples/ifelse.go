package main

import "fmt"

func ifelse() {

	if 0%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	if 3 > 0 {
		fmt.Println("正数")
	}

	if j := 7; j < 0 {
		fmt.Println("负数")
	} else if j > 10 {
		fmt.Println("2位数")
	} else {
		fmt.Println("小于10的非负数")
	}
}
