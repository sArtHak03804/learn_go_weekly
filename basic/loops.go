package main

import "fmt"
//go language dosen't have while loop and do while loop

// but we can use for loop to achieve the same functionality
// loops function demonstrates different types of loops in Go
//comment out loops that are not needed
func loops() {


	// while loop styled for loop 
	i:=1
	for i<=3{

		fmt.Println(i,"while loop styled for loop iteration")
		i++
	}

	// infiniteLoop 
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// classic for loop
	for i:=0 ; i<10 ; i++{
		fmt.Println(i,"classic for loop iteration")
	}

		// classic for loop with break and continue
	for i:=0 ; i<10 ; i++{

		if i%2==0{
			continue // skip even numbers
		}
		fmt.Println(i,"classic for loop iteration with continue")
		if i==5{
			break // break the loop when i is 5
		}
		fmt.Println(i,"classic for loop iteration with break")
	}

	// range loop(this range feature is used in go version 1.22 or above)
	for i:= range 10{
		fmt.Println(i,"range loop") // prints 0 to 9
	} 

}