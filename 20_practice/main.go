// 1-) Print numbers 1 to 10 as odd or even with the if structure.

/* package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "çittir")
		} else {
			fmt.Println(i, "tektir")
		}
	}

} */

// 2-) Give an example of a non-Go while loop using the for construct.

/* package main

import "fmt"

func main() {

	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

} */

// 3-) Explain the expression of switch fallthrough

/* package main

import "fmt"

func main() {

	switch x := 75; {
	case x < 20:
		fmt.Printf("%d is less than 20\n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d is less than 50\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d is less than 100\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d is less than 200\n", x)
	}

} */

// 4-) Make the following if loop more idiomatic.

/* package main

import (
	"fmt"
)

func main() {
	if x := 20; x%2 == 0 {
		fmt.Println(x, "is an even number.")
	} else {
		fmt.Println(x, "is an odd number.")
	}
} */

/* package main

import "fmt"

func main() {

	x := 20

	if x%2 == 0 {
		fmt.Println(x, "is an even number.")
		return
	}

	fmt.Println(x, "is an odd number.")
} */

// 5-) Write a program that displays prime numbers from 1 to 50.

package main

import "fmt"

func main() {

	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf("%d is a prime number. \n", x)
		}
	}
}
