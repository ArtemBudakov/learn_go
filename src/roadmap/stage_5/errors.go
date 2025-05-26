package stage_5

import (
	"errors"
	"fmt"
)

func ErrorTest(panic bool) {
	resultFirst, err := SafeDivide(1, 0)
	fmt.Println(resultFirst, err)
	resultSecond, err := SafeDivide(44, 22)
	fmt.Println(resultSecond, err)

	if panic {
		UnsafeDevice(100, 0)
	}
	CustomPanicDeferred()
	DeferTestEmpty()
	DeferTest(CustomFunctionForTestDeferAndPanic)

	err = CallPanic()
	fmt.Println(err)
}

func SafeDivide(value, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("never divided by zero")
	}
	return value / divider, nil
}

func UnsafeDevice(value, divider int) {
	// if divider is 0 we will get an error. type is panic
	result := value / divider
	fmt.Println("you will never get this message")
	fmt.Println(result)
}

func CustomPanicDeferred() {
	defer func() {
		if panicError := recover(); panicError != nil {
			fmt.Println("Recovered from panic:", panicError)
		}
	}()
	panic("custom panic")
}

func DeferTestEmpty() {
	defer func() {
		if panicError := recover(); panicError != nil {
			fmt.Println("Recovered from panic:", panicError)
		}
	}()
	fmt.Println("DeferTestEmpty without panic")
}

func DeferTest(function func()) {
	defer func() {
		if panicError := recover(); panicError != nil {
			fmt.Println("DeferTest with panic. Recovered from panic:", panicError)
			FirstFunction()
		}
	}()
	function()

	fmt.Println("DeferTest without panic")
	SecondFunction()
}

func CustomFunctionForTestDeferAndPanic() {
	panic("custom panic from CustomFunctionForTestDeferAndPanic")
}

func FirstFunction() {
	fmt.Println("FirstFunction")
}

func SecondFunction() {
	fmt.Println("SecondFunction")
}

func CallPanic() (err error) {
	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic error: %v", panicError)
		}
	}()
	panic("custom panic from CallPanic")
}
