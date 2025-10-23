package main

import (
	"first-go/src/codewars"
	"fmt"
)

func main() {
	DEBUG := true
	//roadmap.Base()
	callAllFunctions(DEBUG)

}

func callAllFunctions(debug bool) {
	if !debug {
		codewars.EvenOrOdd(133)
		codewars.RowSumOddNumbers(5)
		codewars.CleverRowSumOddNumbers(5)

		return
	}

	fmt.Printf("codewars.EvenOrOdd(133) %s \n", codewars.EvenOrOdd(133))
	fmt.Printf("codewars.RowSumOddNumbers(5) %d \n", codewars.RowSumOddNumbers(5))
	fmt.Printf("codewars.CleverRowSumOddNumbers(5) %d \n", codewars.CleverRowSumOddNumbers(5))
}
