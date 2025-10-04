package main

import (
	"fmt"
	"math"
)


func main() {

  var a, b int = 10, 3
  var result int

  result = a + b
  fmt.Println("Addition:",result)

  result = a - b
  fmt.Println("Subtraction:", result)

  result = a * b
  fmt.Println("Multiplication:", result)

  result = a / b
  fmt.Println("Division:", result)

  result = a % b
  fmt.Println("Remainder:", result)

  const p float64 = 22 / 7.0
  fmt.Println(p)

  // Overflow with signed Integer
  var maxInt int64 = 922337203636854775807 // max value for int64 can hold
  fmt.Println(maxInt)

  maxInt = maxInt + 1
  fmt.Println(maxInt)

  // Overflow with unsigned integers
  var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
  fmt.Println(uMaxInt) // error (overflow)

  uMaxInt = uMaxInt + 1
  fmt.Println(uMaxInt) // 0

  // Undeflow with floating point numbers
  var smallFloat float64 = 1.0e-323
  fmt.Println(smallFloat)

  smallFloat = smallFloat / math.MaxFloat64
  fmt.Println(smallFloat)
}
