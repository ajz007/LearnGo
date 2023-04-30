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
	//ch <- "The PayLoad3 !!"
	close(ch)

	fmt.Printf("Sent")
	waitGrp.Done()
}

func receive(ch chan string, waitGrp *sync.WaitGroup) string {
	fmt.Printf("Receiving data from channel...")
	v, ok := <-ch
	if ok {
		fmt.Println("Data received from channel! ", v)
	}

	v2, ok := <-ch

	if ok {
		fmt.Println("Data received from channel2 ! ", v2)
	}

	v4, ok := <-ch
	if ok {
		fmt.Println("Data received from channel4 ! ", v4)
	}
	fmt.Println("Channel open ? ", ok)

	waitGrp.Done()
	return v
}
