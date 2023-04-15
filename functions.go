package main

import "fmt"

func main() {
	fmt.Println(getSum(3, 4))
	fmt.Println(getSum2(17, 4))
}

func getSum(i int, j int) int {

	return i + j
}

func getSum2(i, j int) int {

	return i + j
}
