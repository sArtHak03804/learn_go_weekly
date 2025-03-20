// ifelse demonstrates different types of conditional statements in Go.
package main

import "fmt"


func ifelse() {
	number := 10 // Static value

	// Simple if-else
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// If-else if-else chain
	if number > 0 {
		fmt.Println(number, "is positive")
	} else if number < 0 {
		fmt.Println(number, "is negative")
	} else {
		fmt.Println(number, "is zero")
	}

	// Nested if-else
	if number != 0 {
		if number > 0 {
			if number%2 == 0 {
				fmt.Println(number, "is a positive even number")
			} else {
				fmt.Println(number, "is a positive odd number")
			}
		} else {
			if number%2 == 0 {
				fmt.Println(number, "is a negative even number")
			} else {
				fmt.Println(number, "is a negative odd number")
			}
		}
	} else {
		fmt.Println(number, "is zero")
	}
}
