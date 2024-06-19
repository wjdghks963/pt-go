package functions

import (
	"errors"
	"fmt"
)

// formatter라는 func(string, string) string 이라는 함수를 인자로 받고
// func(stirng, string) 이라는 함수를 반환함
func getLogger(formatter func(string, string) string) func(string, string) {

	return func(a string, b string) {
		fmt.Println(formatter(a,b))
	}

}


func curryFuncTest(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func curryFunc() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	curryFuncTest("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	curryFuncTest("Error on mail server", mailErrors, commaDelimit)
}