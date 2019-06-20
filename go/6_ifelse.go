package main

import "fmt"

func main() {
	if 2%2 == 1 {
		fmt.Println("two is odd number")
	} else {
		fmt.Println("two is even number")
	}

	// ERRORS:-
	// if{
	//	 	###########
	// }								//else should be written along with "}" of previous "if"
	// else{
	// 		################
	// }

	// ERRORS:-
	// if 'sush'					//cannot create string with single quote
	//
	if "sush" != "sushil" {
		fmt.Println("sush is not present in sushil")
	}

	var no = 10
	if no > 10 {
		fmt.Println("no greater than 10")
	} else if no < 10 {
		fmt.Println("no is less than 10")
	} else {
		fmt.Println("no is equal to 10")
	}
}
