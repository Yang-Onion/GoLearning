package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

//只要struct拥有interface中的所有方法，那么这个type就自动实现了这个interface

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func funcInterface() {
	r := rectangle{width: 5, height: 6}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
