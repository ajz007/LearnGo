package main

import "fmt"

func main() {
	fmt.Println("Starting the main go routine..")

	for i := 0; i < 100; i++ {
		go calculate(i)
	}

	fmt.Println("Ending the main go routine..")
}

func calculate(i int) {
	fmt.Printf("go routine i = %d and calculated value = %d \n", i, i+i)
}
