package main

import "fmt"

type person struct {
	name string
	age  int
}

func funcStruct() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "yang", age: 18})
	fmt.Println(person{name: "zhang"})

	fmt.Println(&person{name: "ann", age: 20})

	s := person{name: "sean", age: 20}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 22
	fmt.Println(sp.age)
}
