package main 

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Time now is %s\n", start)
	
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("GO was launched at %s\n", t)
	
	fmt.Println("Date stats: \n")
	fmt.Printf("The month is: %s\n", start.Month())
}

