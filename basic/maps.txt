package main

import "fmt"

func maps() {
	myMap := map[string]int{"apple": 2, "banana": 3}
	fmt.Println("Map:", myMap)

	myMap["cherry"] = 5
	fmt.Println("Updated Map:", myMap)

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
