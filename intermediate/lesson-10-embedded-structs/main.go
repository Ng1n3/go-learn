package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	//	emplyeeInfo Person
	Person
	empId  string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old. \n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
}

func main() {

	emp := Employee{
		Person: Person{name: "Sam", age: 25},
		empId:  "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Emp ID", emp.empId)
	fmt.Println("Salary", emp.salary)

	emp.introduce()
}
