package main

import "fmt"

//type conversions or popularly in go called type assertions are a way to assign the implementations value to a interface variable
//after testing that it is of the required type. Without that it would result in runtime error
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
