package main

import "fmt"

func main() {

	sequence := add()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := add()
	fmt.Println(sequence2())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
}

func add() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
