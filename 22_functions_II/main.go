/* package main
import "fmt"
func main() {
	hello("Baran", 22) // Argument Function Run
}
func hello(name string, age int) { // Parameter Function Write
	fmt.Printf("My name is %s, I am %d years old.", name, age)
} */

/* package main
import "fmt"
func main() {
	fmt.Println(result(40))
}
func result(grade int) string {
	if grade >= 50 {
		return "You passed."
	}
	return "You failed."
	fmt.Println("I AM IN THE FUNCTION.")
} */

/* package main
import "fmt"
func main() {
	merhaba("Arin", 6)
	name := "Elis"
	age := 4
	fmt.Printf("My name is  %s, I am  %d years old.", name, age)
}
func merhaba(name string, age int) {
	fmt.Printf("My name is %s, I am  %d years old.\n", name, age)
} */

/* package main
import (
	"fmt"
	"time"
)
func main() {
	var x int = 10
	fmt.Println(x)
	var moment time.Time = time.Now() // Now ---> method
	fmt.Println(moment)
} */

/* package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	fmt.Print("Please Enter Your Exam Result:")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
} */

package main

import "fmt"

func main() {

	division, remainder := divider(104, 5)
	fmt.Println(division, remainder)

}

// 104 / 5 =====> 20 - 4

func divider(dividend, divider int) (division, remainder int) {
	division = dividend / divider
	remainder = dividend % divider

	return division, remainder
}
