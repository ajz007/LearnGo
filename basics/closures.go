package main

import "fmt"

func main() {

	//passing functions as values is called closures
	a := adder()

	for i := 0; i < 10; i++ {
		fmt.Println("the data for adder", a(i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
