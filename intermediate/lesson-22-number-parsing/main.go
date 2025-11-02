package main

import (
	"fmt"
	"strconv"
)

func main() {
	numstr := "12345"
	num, err := strconv.Atoi(numstr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num)
	fmt.Println("Parsed Integer: ", num+1)

	numisStr, err := strconv.ParseInt(numstr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("parsed Integer: ", numisStr+1)

	floatstr := "3.14"
	floatVal, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Printf("parsed Float: %.2f\n", floatVal)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing the binary value: ", err)
		return
	}
	fmt.Println("Parsed binary to int", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hexadecimal value: ", err)
		return
	}
	fmt.Println("Parsed hex to decimal: ", hex)
}
