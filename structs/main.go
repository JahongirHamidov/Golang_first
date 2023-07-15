package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + "years old"
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "male" {
		return
	} else {
		person.lastName = spouseLastName
	}

}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Jakhongir",
		lastName:  "Kahmidov",
		city:      "NYC",
		gender:    "male",
		age:       23,
	}
	//Alternative
	name := "Samuel"
	person2 := Person{name, "Altman", "California", "female", 40}

	fmt.Println(person1)
	fmt.Println(person2)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())

}
