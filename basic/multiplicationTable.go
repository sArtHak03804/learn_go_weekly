package main

import "fmt"

// Generates the multiplication table for a given number up to a specified range
func generateTable(number, limit int) {
	fmt.Printf("Multiplication Table for %d up to %d:\n", number, limit)
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

func multiplicationTable() {
	var number, limit int
	fmt.Print("Enter the number for the multiplication table: ")
	fmt.Scanln(&number)
	fmt.Print("Enter the range for the table: ")
	fmt.Scanln(&limit)

	generateTable(number, limit)
}
