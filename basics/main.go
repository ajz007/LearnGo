package main

import "fmt"

func main() {
	var str = "Hello World!"
	var str2 string = "Hello World!"
	str3 := "Hello world 5"
	name, nameAgain := "Hello Name", "hello Name again"

	const str4 = "Hello world constant!"

	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(name)
	fmt.Println(nameAgain)

	fmt.Printf("%T\n", str3)

	fmt.Printf("%T\n", str4)
}
