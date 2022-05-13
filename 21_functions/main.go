package main

import "fmt"

func main() {

	/* 	var x, y, sum int
	   	x = 5
	   	y = 10
	   	sum = x + y
	   	fmt.Printf("%d + %d is equal to %d\n", x, y, sum)
	   	x = 7
	   	y = 11
	   	sum = x + y
	   	fmt.Printf("%d + %d is equal to %d\n", x, y, sum) */

	// Moduler Programming
	// Readable code
	// From Complex To Simple

	/* 	fmt.Println(sum(5, 10)) */

	// func funcName(params) return type { code  }

	/* 	hello()
	   	fmt.Println(sum(5, 10))
	   	fmt.Println(sum(3, 5))
	   	fmt.Println(sum(2, 7))
	   	hello()
		   hello() */

	// fmt.Println(x)

	// return vs print

	/* 	z := sum(5, 10)
	   	fmt.Println(z)
		   sum2(6, 11) */

	hello2("Elis", 4)

	// Function Naming
	// first character must be a letter
	// camel Case -- mySum, myBestFunction
	// out of package --> first letter uppercase

}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func hello() {
	fmt.Println("My name is Baran.")
}

func hello2(name string, age int) {
	fmt.Printf("My name is %s, I am %d years old.", name, age)
}
