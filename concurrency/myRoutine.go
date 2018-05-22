package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func DeferTest() {

	/* Defer defers execution of a statement after the rest of the function completes
	It works in a LIFO manner as described below.
	*/

	defer fmt.Println("Done! I was deferred first. ")
	defer fmt.Println("Are we done yet?")
	defer fmt.Println("Hmmm so defer works kinda like a stack?")

	fmt.Println("Doing some stuff.")
	fmt.Println("code code, tappity tap, codity code. ")
	fmt.Println("still doing some stuff. Lets wait for 2 seconds. ")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Okay, I'm done now. Lets call the defers. ")
}

func Panic() {

	/*
		panic halts the program and starts to panic. This stops running the program
		and runs all of the deferred statements from the panic function.

		recover function lets us recover from the panicking goroutine
	*/
}

var wg sync.WaitGroup

func main() {

	/*

		Let's convert the first say method to a go-routine
		Append 'go' before the function. thats it. whaaaa??

		Go-routine is a light-weight thread

	*/
	DeferTest()
	Panic()

	wg.Add(1)
	go say("I am groot")
	wg.Add(1)
	go say("I am Steve Rodgers")

	wg.Wait()
}
