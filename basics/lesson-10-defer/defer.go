package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("First defered Hello world!")
	defer fmt.Println("Second defered Hello world!")
	defer fmt.Println("Third defered Hello world!")
	i++
	fmt.Println("Normal Hello world again!")
	fmt.Println("Value of i: ", i)
}
