package main

import (
	"fmt"
	"sync"
)

func main() {
	//var chnl chan string
	chnl := make(chan string, 3)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go send(chnl, &waitGrp)
	go receive(chnl, &waitGrp)
	waitGrp.Wait()
}

func send(ch chan string, waitGrp *sync.WaitGroup) {
	fmt.Printf("Sending data to channel...")
	ch <- "The PayLoad1 !!"
	ch <- "The PayLoad2 !!"
	ch <- "The PayLoad3 !!"
	//1,2,3 would be recieved and next 3 would be buffered
	ch <- "The PayLoad4 !!"
	ch <- "The PayLoad5 !!"
	ch <- "The PayLoad6 !!"

	fmt.Printf("Sent")
	waitGrp.Done()
}

func receive(ch chan string, waitGrp *sync.WaitGroup) string {
	fmt.Printf("Receiving data from channel...")
	v := <-ch
	fmt.Println("Data received from channel! ", v)
	v2 := <-ch
	fmt.Println("Data received from channel2 ! ", v2)
	v3 := <-ch
	fmt.Println("Data received from channel3 ! ", v3)

	waitGrp.Done()
	return v
}
