package main

import "fmt"

func divide(a, b int) (result int, err string) {
	if b == 0 {
		err = "Cannot divide by zero"
		return
	}
	result = a / b
	return
}

func namedReturns() {
	result, err := divide(10, 2)
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
