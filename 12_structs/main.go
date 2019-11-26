package main

import (
	"fmt"
	"strconv"
)

// Person struct defined
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am a " + strconv.Itoa(p.age) + " " + p.gender
}

// hasBirthday method (pointed receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// changeGender (pointer receiver)
func (p *Person) changeGender(newGender string) {
	p.gender = newGender
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Rodrigo",
		lastName:  "Marcondes",
		city:      "Campinas",
		gender:    "male",
		age:       30,
	}

	// Alternative short form based on the order defined
	person1Alternative := Person{"Rodrigo", "Marcondes", "Campinas", "male", 30}

	fmt.Println(person1)
	fmt.Println(person1Alternative)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.changeGender("vegan")
	fmt.Println(person1.greet())
}
