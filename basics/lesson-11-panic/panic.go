package main

import "fmt"

func main() {
	process(10)

	process(-1)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		//    fmt.Println("After Panic")
	}

	fmt.Println("Processing input:", input)
}
