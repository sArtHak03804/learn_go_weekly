// calculator performs basic arithmetic operations using static values.
package main

import "fmt"

func calculator() {
	// Static values
	number1 := 10
	number2 := 5

	// Perform all operations
	sum := number1 + number2
	difference := number1 - number2
	product := number1 * number2

	// Handle division by zero
	if number2 != 0 {
		quotient := number1 / number2
		fmt.Println("Division:", quotient)
	} else {
		fmt.Println("Division: Error (division by zero)")
	}

	// Print results
	fmt.Println("Addition:", sum)
	fmt.Println("Subtraction:", difference)
	fmt.Println("Multiplication:", product)
}


