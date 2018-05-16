package main 

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas uint16
	brake uint16
	steering int16
	top_speed float64
	car_type string
}


/* Value receiver methods */
func (c car) kmph() float64 {
	return float64(c.gas) * (c.top_speed/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas) * (c.top_speed/usixteenbitmax/kmh_multiple)
}

/* Pointer Receiver methods */
func (c *car) changeSpeed(newSpeed float64) {
	c.top_speed = newSpeed
}

func (c *car) changeCarType(newType string) {
	c.car_type = newType
}

func main() {
	
	/* We have two types of methods: Value receivers and Pointer receiver */
	a_car := car{gas: 22341, brake: 0, steering: 12561, top_speed: 125.64, car_type: "Mazda"}
	
	fmt.Println("Details for a_car:")
	fmt.Printf("a_car type: %s\n", a_car.car_type)
	fmt.Printf("a_car top speed: %f\n", a_car.top_speed)
	fmt.Printf("Gas for a_car in kmph: %f\n", a_car.kmph())
	fmt.Printf("Gas for a_car in mph: %f\n", a_car.mph())
	
	a_car.changeCarType("Audi")
	a_car.changeSpeed(160.8)
	
	fmt.Println("Details for a_car after change:\n")
	fmt.Printf("a_car type: %s\n", a_car.car_type)
	fmt.Printf("a_car top speed: %f\n", a_car.top_speed)
	

}

