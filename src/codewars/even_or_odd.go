package codewars

func EvenOrOdd(input_value int) string {
	result := ""
	if input_value%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return result
}
