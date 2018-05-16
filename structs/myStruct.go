package main

import (
	"fmt"
)

type car struct {
	car_type string
	gas uint16
}

func main() {
	fmt.Println("Whoa, this is my go project")
	a_car := car{car_type: "Mazda", gas: 12345}
	fmt.Println(a_car.gas)
	
}