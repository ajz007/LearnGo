package main

import "fmt"

func main() {

	ids := []int{12, 23, 234532, 2343, 121, 123, 123}

	emails := map[string]string{"1": "my@gmail.com", "2": "two@gmail.com"}

	//here i is the index and data has the datd from the array
	for i, data := range ids {
		fmt.Printf("%d is the index of %d \n", i, data)
	}

	//if you dont want to use the index
	for _, data := range ids {
		fmt.Printf("is the index of %d \n", data)
	}

	for k, v := range emails {
		fmt.Printf("%sis the key for value %s \n", k, v)
	}

	for key := range emails {
		fmt.Printf("%sis the key for value\n", key)
	}

	for _, value := range emails {
		fmt.Printf("is the key for value %s\n", value)
	}
}
