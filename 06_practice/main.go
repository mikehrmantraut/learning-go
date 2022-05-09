// 1-) studentName --> John Doe, grade --> 77, isPassed --> true
// Create these variables in 3 different ways

/* package main
import "fmt"
func main() { */

/* 	var studentName string = "John Doe"
   	var grade int = 77
   	var isPassed bool = true */

/* 	var studentName = "John Doe"
   	var grade = 77
	   var isPassed = true */

/* 	studentName := "John Doe"
	grade := 77
	isPassed := true
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}
*/

// 2-) Define the above-mentioned variables in a single line.
/*
package main
import "fmt"
func main() {
	/* 	var studentName string = "John Doe"
	   	var grade int = 77
	   	var isPassed bool = true */

/* var studentName, grade, isPassed = "John Doe", 77, true */

/* 	studentName, grade, isPassed := "John Doe", 77, true
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}  */

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" --> Explain these terms.

/* package main
import "fmt"
func main() {
	var studentName string = "John Doe"
	studentName = "Michael Scott"
	fmt.Println(studentName)
} */


// 4-) ":=" vs "=" show difference between these two, double declaration

package main

import "fmt"

func main() {

	/* 	var studentName string = "John Doe"
	   	studentName = "Michael Scott" */

	studentName := "John Doe"
	studentName = "Michael Scott"

	fmt.Println(studentName)

}
