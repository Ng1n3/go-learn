package main

import "fmt"

func main() {

	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 20
	numbers[1] = 25
	numbers[2] = 26
	numbers[3] = 30
	numbers[4] = 40

	fmt.Println(numbers)

	fruits := [5]string{"Apple", "Banana", "Orange", "Pear"}
	fmt.Println("Fruits Array: ", fruits)

	originalArray := [3]int{1, 3, 4}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("copied Array: ", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	a, b := someFunction()
	fmt.Println(a)
	fmt.Println(b)

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array1 is equal to Array2: ", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	newOriginalArray := [3]int{1, 2, 3}
	var newCopiedArray *[3]int

	newCopiedArray = &newOriginalArray
	newCopiedArray[0] = 100

	fmt.Println("original array: ", newOriginalArray)
	fmt.Println("Copied array: ", newCopiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
