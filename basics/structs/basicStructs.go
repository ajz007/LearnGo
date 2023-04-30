package main

import (
	"fmt"
	"strconv"
)

func main() {

	person := Person{firstName: "Ajit", lastName: "Bhakalu", age: 31, gender: "M"}

	person2 := Person{"Ajit", "Bhakalu", 31, "M"}
	fmt.Printf("Person \n", person)
	fmt.Printf("Person2 \n", person2)

	//update
	person2.firstName = "Sameer"
	fmt.Println("Person 1 First name", person.firstName)
	fmt.Println("Person 2 First name", person2.firstName)

	//calling methods of structs
	fmt.Println("The greeting ", person.greet())
	person.hasBirthDay()
	fmt.Println("The greeting ", person.greet())
	person.hasBirthDay()
	fmt.Println("The greeting ", person.greet())

}

//Please note that the structs dont have any methods defined inside them. These are defined outside and objects are passed to it

// value receivers methods-- which return values ( like getters in java)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + "and my age is " + strconv.Itoa(p.age)
}

// Pointer Receivers
func (p *Person) hasBirthDay() {
	p.age++
}

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}
