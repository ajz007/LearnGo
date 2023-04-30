package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(4)

	ch := make(chan string, 2)
	ch2 := make(chan string, 2)
	ch3 := make(chan string, 2)

	go send(ch, &wg, 1)
	go receive(ch, &wg)

	go send(ch2, &wg, 2)
	go send(ch3, &wg, 3)

	select {
	case val1 := <-ch:
		fmt.Println("val --> ", val1)
		break
	case val2 := <-ch2:
		fmt.Println("val --> ", val2)
	case val3 := <-ch3:
		fmt.Println("val --> ", val3)
		break
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
		//default:
		//	fmt.Println("val --> Default!!") // if there is no default block
		//then this entire select block becomes blocking as it is waiting to receive data from channel
	}

	wg.Wait()
}

func send(ch chan string, wg *sync.WaitGroup, i int) {
	time.Sleep(3 * time.Second)
	ch <- "Data in Channel! " + fmt.Sprint(i)
	wg.Done()
}

func receive(ch chan string, wg *sync.WaitGroup) {
	v := <-ch
	fmt.Println("Printing value ---> ", v)
	wg.Done()
}
