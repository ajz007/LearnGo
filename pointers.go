package main

import "fmt"

func main() {

	a := 5
	b := &a
	fmt.Println("Printing data and pointers a =", a)
	fmt.Println("Printing data and pointers b =", b)

	fmt.Printf("%T \n", b)
	//use * to read the value from the pointer
	fmt.Printf("The data for b is %d", *b)

	//update value of pointer
	*b = 25

	fmt.Printf("The data for a,b is %d,%d", a, *b)

}
