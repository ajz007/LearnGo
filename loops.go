package main

import "fmt"

func main() {

	i := 10

	for i < 100 {
		fmt.Println("Printing the number = ", i)
		i++
	}

	//short method
	for i := 0; i < 21; i++ {
		fmt.Println("Printing the number in loop 2= ", i)
	}

	//Fizz Buzz
	for j := 1; j < 100; j++ {
		if j%5 == 0 && j%3 == 0 {
			fmt.Printf("Fizz Buzz \n")
		} else if j%3 == 0 {
			fmt.Printf("Fizz \n")
		} else if j%5 == 0 {
			fmt.Printf("Buzz \n")
		}

	}
}
