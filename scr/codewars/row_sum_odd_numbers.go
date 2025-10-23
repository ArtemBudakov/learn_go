package codewars

func RowSumOddNumbers(number int) int {
	//Calculate the sum of the numbers in the nth row of this
	//triangle (starting at index 1) e.g.: (Input --> Output)
	// number is deep line of triangle
	//             1
	//          3     5
	//       7     9    11
	//   13    15    17    19
	//21    23    25    27    29

	// elements in line == number || 4 elements for 4 line
	// first element in line == (1 + 2 + 3) * 2 + 1 ==  13

	if number == 1 {
		return number
	}

	start_value := 0
	for i := number - 1; i >= 1; i-- {
		start_value += i
		//fmt.Printf("index: %d\n", i)
		//fmt.Printf("start_value: %d\n", start_value)
	}
	first_value := start_value*2 + 1

	result := 0
	for i := number - 1; i >= 0; i-- {
		result += first_value + (2 * i)
		//fmt.Printf("result: %d\n", result)
	}
	return result
}

func CleverRowSumOddNumbers(number int) int {
	return number * number * number
}
