package main

import "fmt"

func main() {
	var i = 10

	for i >= 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println("\n\n\nnew loop")
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n\n\nnew loop")
	var k = 1
	for {
		if k == 5 {
			break
		}
		k++
		if k != 5 {
			fmt.Println(k)
			continue
		}
		fmt.Println("statment executes even after break")
		fmt.Println("HOW THE F###")
	}

	fmt.Println("\n\n\nnew loop")
	for {
		fmt.Println("break")
		break
	}

	//ERRORS:-

	//ERROR: Cannot creaeting variable inside for loop using var keyword
	// for var j := 0; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	//ERROR: need all conditions inside for loop else only 1 condition of checking
	// for j := 0; j < 5 {
	// 	fmt.Println(j)
	// 	j++
	// }
}
