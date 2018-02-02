package main

import "fmt"

//全局变量定义之后即使不使用也不会报错
var n string

func variables() {
	var s string = "initial"
	fmt.Println(s)
	var a, b int = 1, 2
	fmt.Printf("a=%d,b=%d \n", a, b)

	i, j := 3, 3.1415
	fmt.Println(i, j)
	//局部变量,如果没有使用则会报错
	var k = false
	var m string

	if !k {
		fmt.Println("m default value is empty",m)
	}
}
