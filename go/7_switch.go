package main

import (
	"fmt"
	"time"
)

func main() {
	var no = 10

	switch {
	case no < 0:
		fmt.Println(no, " is negative")

	default:
		fmt.Println(no, " is positive")
	}

	no = -11

	switch no % 2 {
	case 1:
		fmt.Println(no, "  is odd")
	case 0:
		fmt.Println(no, " is even")
	default:
		fmt.Println(no, " is negative and odd")
	}

	var currentTime = time.Now()
	fmt.Println("\ncurrent time is ", currentTime.Hour())
	switch {
	case currentTime.Hour() > 10 && currentTime.Hour() < 18:
		fmt.Println("that means you are working")
	default:
		fmt.Println("that means you are not working")
	}
}
