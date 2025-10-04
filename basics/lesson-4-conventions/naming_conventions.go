package main

import "fmt"


type Employee struct {
  FirstName string
  LastName string
  Age int
}

func main() {
  // PascaCase
  // E.g. CalculateArea, UserInfo, NewHTTPRequest
  // structs, interfaces, enums

  // snake_case
  // E.g. user_id, first_name, http_request
  
  // UPPERCASE
  // Use case is constants

  // mixedCase
  // E.g. javaScript, htmlDocument, isValid
 
  const MAXRETRIES = 5

  var employeeID = 1001
  fmt.Println("Employee ID: ", employeeID)
}
