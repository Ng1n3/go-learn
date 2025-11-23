package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	person := Person{FirstName: "John", Email: "example@example.com"}

	//Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("errror marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person2 := Person{FirstName: "Khali", Age: 28, Email: "kalie@example.com", Address: Address{City: "Ibadan", State: "Oyo"}}

	jsonData2, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData2))

	jsonData1 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`

	var employeeFromJSON Employee
	json.Unmarshal([]byte(jsonData1), &employeeFromJSON)

	fmt.Println(employeeFromJSON)
	fmt.Println("Jenny's Age increase by 5 years", employeeFromJSON.Age+5)
	fmt.Println("Jenny's city:", employeeFromJSON.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("Ther was an error marshalling: ", err)
		return
	}

	fmt.Println("JSON List:", string(jsonList))

	// Handling unknown json structures
	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]interface{}

	json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
