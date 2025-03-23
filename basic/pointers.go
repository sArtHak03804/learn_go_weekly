package main

import "fmt"

func updateValue(ptr *int) {
	*ptr = 20
}

func pointers() {
	value := 10
	fmt.Println("Original Value:", value)

	updateValue(&value)
	fmt.Println("Updated Value:", value)
}
