package main

import "fmt"

var packVar = "Package Scope"

// packVar := "Package Scope" --> error 
// --> syntax error: non-declaration statement outside
// --> function body

func main() {
	
	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	
	fmt.Println(packVar)
	
	myFunc()
}

func myFunc() {
	fmt.Println(packVar)
	//fmt.Println(funcVar) --> error --> undefined: funcVar 
}
