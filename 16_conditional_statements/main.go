package main

import "fmt"

func main() {

	// if <boolean expression> { code }  x%2 == 0 (false)

	/* 	x := 27
	   	if false {
	   		fmt.Println(x, "is an odd number")
	   	} */

	/* 	if !true {
		fmt.Println("Message will be displayed")
	} */

	/* 	if 5 > 3 {
		fmt.Println("Will the message be shown?")
	} */

	/* 	x := 4
	   	if x%2 == 0 {
	   		fmt.Println(x, "is an even number.")
	   	} else {
	   		fmt.Println(x, "is an odd number.")
		   } */

	// if <boolean expression> { code } else { code }

	/* 	if false {
		fmt.Println("Message will be displayed")
	} */

	x := -5

	if x < 0 {
		fmt.Println(x, "is a negative number.")
	} else if x%2 == 0 {
		fmt.Println(x, "is an even number.")
	} else {
		fmt.Println(x, "is an odd number.")
	}
	// if <boolean expression> { code } else if <boolean expression> else { code }
}
