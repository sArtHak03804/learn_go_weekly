package main

import "fmt"

func slices() {
	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)

	slice = append(slice, 4, 5)
	fmt.Println("Appended Slice:", slice)

	newSlice := slice[1:4]
	fmt.Println("Sliced Slice:", newSlice)
}
