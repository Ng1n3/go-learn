package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCall
}

type PhoneHomeCall struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

type Address struct {
	city    string
	country string
}

func main() {

	p1 := Person{
		firstName: "Sam",
		lastName:  "Kehrr",
		age:       40,
		address: Address{
			city:    "Lokoja",
			country: "Nigeria",
		},
		PhoneHomeCall: PhoneHomeCall{
			home: "2348037787921",
			cell: "454545454545",
		},
	}

	fmt.Println(p1)

	p2 := Person{
		firstName: "Djon",
		age:       20,
	}

	p3 := Person{firstName: "Djon", age: 20}

	p2.address.city = "Ikeja"
	p2.address.country = "Nigeria"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.address.country)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)

	fmt.Println("Are p1 and p2 equal?: ", p3 == p2)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "Doue",
		email:    "Psuedoemail@mail.com",
	}

	// Assigning methods to Structs
	fmt.Println(p1.fullName())
	fmt.Println(user.username)

	// Using Pointers
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)
}
