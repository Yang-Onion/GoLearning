package main

import "fmt"

func main() {

	//fmt.Println("--------go-变量--------")
	//variable()

	//fmt.Println("--------go-常量--------")
	//constant()

	//fmt.Println("--------go-常量--------")
	//referance()

	fmt.Println("--------go-控制流--------")
	logicFlow()
}

//定义全局变量,即使在方法内没有使用它,也不会报错
var aa string

//定义在方法内部的局部变量,如果没有使用,会报错
func variable() {
	var a int32

	a = 12
	//只能用在方法内容，定义局部变量
	b := 13

	c, d := "hello", 2018

	var e float32 = 3.1415

	var f, g string = "mr", "yang"
	h := false

	var i = true

	aa = "全局变量"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%s,%d \n", c, d)
	fmt.Println(e)
	fmt.Printf("%s.%s \n", f, g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(aa)
}

//常量定义了不使用也不会报错
func constant() {
	//多常量初始化
	const x, y int = 1, 2
	//常量,类型推断
	const s = "hello "
	//常 组
	const (
		a, b      = 10, 100
		c    bool = false
	)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Friday)
}

//引用类型
func referance() {
	a := []int{1, 2, 3}
	a[1] = 11

	b := make([]int, 3)
	b[1] = 10

	c := new([]int)

	fmt.Printf("%d,%d,%d", a, b[1], c)
}

//控制流
func logicFlow() {

	x := 0
	if n := "abc"; x > 0 {
		println(n[0])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[2])
	}

}
