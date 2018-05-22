package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {

	/*

		Let's convert the first say method to a go-routine
		Append 'go' before the function. thats it. whaaaa??

		Go-routine is a light-weight thread

	*/

	go say("I am groot")
	say("I am Steve Rodgers")
}
