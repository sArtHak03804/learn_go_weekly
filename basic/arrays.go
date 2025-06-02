package main

import "fmt"

func arrays() {
	var arr [5]int = [5]int{1, 2, 3, 78, 5}
	fmt.Println("Array:", arr)

	// Traversing the array
	for i, value := range arr {
		fmt.Printf("Index %d: %d\n", i, value)
	}
}
