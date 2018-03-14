package main

import "fmt"

func main() {
	fmt.Println("--------------Go By Example--------------")

	fmt.Println("--------------Values--------------")

	values()

	fmt.Println("--------------Variables--------------")
	variables()

	fmt.Println("--------------constants--------------")
	constants()

	fmt.Println("--------------for--------------")
	funcFor()

	fmt.Println("--------------else if--------------")
	ifelse()

	fmt.Println("--------------switch--------------")
	funcSwitch()

	fmt.Println("--------------array--------------")
	funcArr()
	fmt.Println("--------------slice--------------")
	funcSlice()
	fmt.Println("--------------map--------------")
	funcMap()
	fmt.Println("--------------range--------------")
	funcRange()

	fmt.Println("--------------func--------------")

	plus(1, 2)

	fibo := fibonaci(10)
	fmt.Println("fibo:", fibo)
	fmt.Println("--------------multireturn--------------")
	multiReturn()

	fmt.Println("--------------variadicfunctions--------------")
	variadicfunctions()

	fmt.Println("--------------clousers--------------")
	closures()

	fmt.Println("--------------recurision--------------")
	recurision()

	fmt.Println("--------------pointer--------------")
	funcPointer()

	fmt.Println("--------------struct--------------")
	funcStruct()

	fmt.Println("--------------struct-methods--------------")
	structMethods()

	fmt.Println("--------------struct-interface--------------")
	funcInterface()

	fmt.Println("--------------error--------------")
	funcError()

	fmt.Println("--------------goroutine--------------")
	//funcGoroutines()

	fmt.Println("--------------channel--------------")
	funcChannel()

	fmt.Println("--------------channel buffer--------------")
	channelbuffer()

	fmt.Println("--------------channel async--------------")
	channelAsync()

	fmt.Println("--------------channel directions--------------")
	channeldirection()

	fmt.Println("--------------select--------------")
	funcSelect()

	fmt.Println("--------------sorting--------------")
	sorting()

	fmt.Println("--------------json--------------")
	funcJson()
}
