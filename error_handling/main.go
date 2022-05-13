/* package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your number: ", result)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("ERROR: You did not enter an even number.")
	}

	return num, nil
} */

/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("can not take the square root of negative numbers")
	}

	return math.Sqrt(num), nil
} */

/* package main

import (
	"fmt"
	"math"
)

func main() {

	result := sRoot(-4)
	{
		fmt.Println(result)
	}
}

func sRoot(num float64) float64 {

	return math.Sqrt(num)
} */

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
    fmt.Println("Your file:", file)
	}

}
