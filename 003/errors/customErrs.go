package errors

import "fmt"


type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf(
		"can't divide %v by zero",
		de.dividend,
	)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, divideError{dividend: a}
	}
	return a / b, nil
}

func customErr() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
