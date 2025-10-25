package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Rect1 struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (r Rect1) area() float64 {
	return r.height * r.width
}

func (r Rect1) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := Rect{width: 4, height: 5}
	c := Circle{radius: 4}
	r1 := Rect1{width: 4, height: 5}

	measure(r)
	measure(c)
	measure(r1)

	myPrinter(1, "John", 46.9, true)

	printType(9)
	printType("John")
	printType(false)
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")
	}
}

func myPrinter(i, ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
