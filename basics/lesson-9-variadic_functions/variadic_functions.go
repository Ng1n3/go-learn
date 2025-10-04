package main

import "fmt"

func main() {
	// ... Ellipsis
	// fmt.Println("Sum of 1, 2, 3: ", sum(1,2 3))
	statement, total := sum("The sum of 1, 2, 3, is", 1, 2, 3, 4)
	fmt.Println(statement, total)

	sequence, total := sum2(1, 20, 30, 40, 50, 60)
	fmt.Println("sequence: ", sequence, "Total", total)

	sequence2, total2 := sum2(2, 40, 36, 40, 50, 60)
	fmt.Println("sequence: ", sequence2, "Total", total2)

	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence3, total3 := sum2(3, numbers...)
	fmt.Println("sequence: ", sequence3, "Total", total3)

	integers3 := []int{1, 2, 3, 4}
	sequence4, total4 := sum2(6, integers3...)
	fmt.Println("sequence: ", sequence4, "Total: ", total4)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
func sum2(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
