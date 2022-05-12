package main

import "fmt"

func main() {

	/* 	x := 10
	   	if x := -5; x < 0 {
	   		fmt.Println(x, "is an negative number.")
	   	} else if x%2 == 0 {
	   		fmt.Println(x, "is an even number.")
	   	} else {
	   		fmt.Println(x, "is an odd number.")
	   	}
		   fmt.Println(x) */

	if x := -25; x < 0 {
		fmt.Println(x, "is a negative number.")
		fmt.Println("My name is Baran.")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "is an even number.")
		} else {
			fmt.Println(x, "is an odd number.")
		}
	}

}
