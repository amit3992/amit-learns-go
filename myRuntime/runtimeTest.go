package main 

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("My OS is: %s\n", runtime.GOOS)
	fmt.Printf("Go-root: %s\n", runtime.GOROOT())
	fmt.Printf("My Go arch: %s\n", runtime.GOARCH)
	
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

}
