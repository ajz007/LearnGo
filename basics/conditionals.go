package main

import "fmt"

func main() {

	x := 25
	y := 35
	if x < y {
		fmt.Printf("%d is less than %d \n ", x, y)
	} else if y < x {
		fmt.Printf("%d is less than %d \n ", y, x)
	} else {
		fmt.Printf("default")
	}

	color := "red"

	switch color {
	case "red":
		{
			fmt.Printf("red")
		}
	case "green":
		{
			fmt.Printf("green")
		}
	default:
		fmt.Printf("default")
	}

}
