package main

import (
	"fmt"
	"sync"
)

func main() {
	//var chnl chan string
	chnl := make(chan string)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go send(chnl, &waitGrp)
	go receive(chnl, &waitGrp)
	waitGrp.Wait()
}

func send(ch chan string, waitGrp *sync.WaitGroup) {
	fmt.Printf("Sending data to channel...")
	ch <- "The PayLoad!!"
	fmt.Printf("Sent")
	waitGrp.Done()
}

func receive(ch chan string, waitGrp *sync.WaitGroup) string {
	fmt.Printf("Receiving data from channel...")
	v := <-ch
	fmt.Println("Data received from channel!", v)
	waitGrp.Done()
	return v
}
