package main

import "fmt"

func main() {
	
	/*
	var (
		name string = "Baran"
		age int = 22
		isMarried bool = false

		weight float32 = 65.5
		height int = 180
	)
	*/
	
	// var name, age, isMarried, weight, height  = "Baran", 22, false, 65.5, 180 
	
	name, age, isMarried, weight, height := "Baran", 22, false, 65.5, 180

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

}
