package main

import (
 "fmt"
)

func main() {

	x := 10
	y := 10

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than x")
	}

	// else if 

	if x > y {
		fmt.Println("x is greater than y")
	} else if y > x {
		fmt.Println("y is greater than x")
	} else {
		fmt.Println("x and y are equal")
	}

	// nested 

	if x > y {
		fmt.Println("x is greater than y")
	  	 if y > x {
			fmt.Println("y is greater than x")
			if x == y {
				fmt.Println("x and y are equal")
			}
		} 
	} else {
		fmt.Println("error")
	}


}