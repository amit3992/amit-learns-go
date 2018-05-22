package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someVal int) {
	defer wg.Done()
	c <- someVal * 5
}

func main() {

	/* Declaring a channel. Channels by default are blocking on sending and receiving values */
	fooChan := make(chan int, 10)

	/* Below go routines are sending value to the channel */
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooChan, i)
	}

	wg.Wait()
	close(fooChan)

	/* Receiving values from the channel */
	for item := range fooChan {
		fmt.Println(item)
	}

}
