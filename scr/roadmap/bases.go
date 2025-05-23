package roadmap

import "fmt"

func Base() {
	s := "some name"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		result := oddOrEven(Number{number: i}) // imitation of position argument. equal with odd_or_even(i) without struct
		fmt.Printf("(oddOrEven) number %d is %s \n", i, result)
	}

	fmt.Println("=======================")
	for index := 1; index <= 6; index++ {
		result := numberSwitchCase(index)
		fmt.Printf("(numberSwitchCase) number %d is %s \n", index, result)
	}

	fmt.Println("=======================")
	numberPrintForCycle(10)

	fmt.Println("=======================")
	numberPrintWhileCycle(15)

	Test(1, 2, "From main.go")

	testListValue := []int{101, 102, 103, 104, 105}
	printArrayValues(testListValue)

	testMapValue := map[string]int{"one": 1, "two": 2, "three": 3}
	printMapValuesAndKeys(testMapValue)
}

func Test(one int, two int, message string) int {
	//TIP <p> some tip </p>
	var result int = one + two
	fmt.Println("result =", result)
	fmt.Printf("message = %s\n", message)
	return result
}

type Number struct {
	number int
}

func oddOrEven(numberStruct Number) string {

	//if numberStruct.number  %2 == 0 {
	//	return "even"
	//} else if numberStruct.number % 2 == 1 {
	//	return "odd"
	//} else {
	//	return "error"
	//}

	if numberStruct.number%2 == 0 {
		return "even"
	}
	return "odd"
}

func numberSwitchCase(number int) string {
	switch number {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	default:
		return "I don't know what to do"
	}
}

func numberPrintForCycle(numberCycle int) {
	for index := 1; index <= numberCycle; index++ {
		fmt.Printf("(numberPrintForCycle) Current number is: %d \n", index)
	}
}

func numberPrintWhileCycle(numberWhile int) {
	for number := 1; number <= numberWhile; {
		fmt.Printf("(numberPrintWhileCycle) Current number is: %d \n", number)
		number += 1
	}
}

func printArrayValues(list []int) {
	for number := 0; number < len(list); number++ {
		fmt.Printf("printArrayValues: %d \n", list[number])
	}
}

func printMapValuesAndKeys(nestedMap map[string]int) {
	for key, value := range nestedMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)

	}
}
