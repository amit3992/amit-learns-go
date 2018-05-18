package main 

import (
	"fmt"
)

func main() {
	
	/* Some more info about maps: 
	
	- Maps are not safe for concurrent use. One way to make maps safe from concurrent goroutines is with 'sync.RWMutex'
	
	*/
	
	/* Map definition. This map is a refernce so it's value is 'nil'. We need to initialize a map to use it */
	var grades map[string]float32
	
	if grades == nil {
		fmt.Println("Map is empty right now. Let's initialize it on the next line.\n")
	}
	
	/* Map initialization */
	grades = make(map[string]float32) // Key -> string, value -> float32
	
	if grades != nil {
		fmt.Println("Yay! Map has been initialized. We can now add values to it.\n")
	} else {
		fmt.Println("Womp womp. Map not initialized yet.\n")
	}
	grades["Alexander"] = 42
	grades["Aaron"] = 40
	grades["Eliza"] = 67
	grades["Marquis"] = 53
	grades["Angelica"] = 78
	
	AlexGrades := grades["Alexander"]
	fmt.Printf("Alexander's grades: %f\n", AlexGrades)
	fmt.Println(grades)
	
	fmt.Println("===============================================================================\n")
	
	
	/* ================================= MAP METHODS ====================================== */
	fmt.Println("Let's take a look at some Map methods:\n")
	
	/* A] Delete a key from the map*/
	delete(grades, "Aaron")
	fmt.Println("Deleted Aaron from the map. Sorry Aaron.")
	fmt.Println(grades)
	
	/* B] Test for existence of key in the map */
	a_val, a_contains := grades["Alexander"]
	b_val, b_contains := grades["Aaron"]
	
	if a_contains == true {
		fmt.Printf("\nAlexander exists in the map and the value is: %f\n", a_val)
	} else {
		fmt.Println("Womp womp. Alexander does not exist in the map.\n")
	}
	
	if b_contains == true {
		fmt.Printf("Aaron exists in the map and the value is: %f\n", b_val)
	} else {
		fmt.Println("Womp womp. Aaron does not exist in the map.\n")
	}
	
	// Test for existence without retrieving the value
	_, ok := grades["John"]
	fmt.Printf("Does John exist in the map? -> %t\n", ok)
	
	
	/* C] Iterate maps */
	fmt.Println("\nIterating a map:\n")
	for key, value := range grades {
	    fmt.Println("Key:", key, "Value:", value)
	}
	
	
	/* ============================ SOME MORE MAP DEFINITIONS ========================================= */
	 
	 // Map of a map. Key -> String, Value -> Map[string]int
	 var hits map[string]map[string]int
	 
	 hits = make(map[string]map[string]int)
}

