package main

import "fmt"

func foo(c chan int, someVal int) {
	c <- someVal * 5
}

func main() {

	/* Declaring a channel */
	fooChan := make(chan int)

	/* Below go routines are sending value to the channel */
	go foo(fooChan, 10)
	go foo(fooChan, 20)

	/* Receiving values from the channel */
	val_1 := <-fooChan
	val_2 := <-fooChan

	fmt.Println(val_1, val_2)
}
