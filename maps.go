package main

import "fmt"

func main() {

	//declare
	emails := make(map[string]string)
	//declare 2
	myAddr := map[string]string{"1": "One", "2": "Two"}

	//assign
	emails["aji"] = "aji.shu@gmail.com"
	emails["aju"] = "aju.shu@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(myAddr)

	//delete
	delete(emails, "aji")
	fmt.Println(emails)
}
