package main

import "fmt"

func main() {

	var arr [2]string
	veggiesArr := [2]string{"Aalu", "Kachalu"}
	fruitsArr := []string{"Graves", "Banana", "Mango"}

	arr[0] = "Apples"
	arr[1] = "Orange"
	fmt.Println(arr)
	fmt.Println(fruitsArr)
	fmt.Println(veggiesArr)
	fmt.Println(len(fruitsArr))
	fmt.Println(fruitsArr[1:2])

}
