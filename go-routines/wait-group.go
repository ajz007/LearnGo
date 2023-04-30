package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting the main go routine..")

	var waitGrp sync.WaitGroup

	//set initial value
	waitGrp.Add(100)
	for i := 0; i < 100; i++ {
		go calculate(i, &waitGrp)
	}

	waitGrp.Wait()
	fmt.Println("Ending the main go routine..")
}

func calculate(i int, waitGrp *sync.WaitGroup) {
	fmt.Printf("go routine i = %d and calculated value = %d \n", i, i+i)
	waitGrp.Done()
}
