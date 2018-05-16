package main

import "fmt"

func main() {
	
	var p *int
	
	if p != nil {
		fmt.Printf("The value of p: \n", *p)
		fmt.Printf("The address of p: \n", p)
	} else {
		fmt.Printf("The value of p is nil\n")
	}
	
	var v int = 42
	p = &v
	
	if p != nil {
		fmt.Printf("The value of p: \n", *p)
		fmt.Printf("The address of p: \n", p)
	} else {
		fmt.Printf("The value of p is nil\n")
	}
	
}

