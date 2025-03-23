package main

import "fmt"

// data_types demonstrates the use of basic data types in Go, including
// integer, float, boolean, and string, along with type conversion.
func data_types() {
	// Declaring and initializing an integer variable.
	// Integers are whole numbers without a fractional part.
	var integer int = 42
	fmt.Println("Integer:", integer)

	// Declaring and initializing a float variable.
	// Floats are used to represent numbers with a fractional part.
	var floating float64 = 3.14
	fmt.Println("Float:", floating)

	// Declaring and initializing a boolean variable.
	// Booleans can only have two values: true or false.
	var boolean bool = true
	fmt.Println("Boolean:", boolean)

	// Declaring and initializing a string variable.
	// Strings represent sequences of characters enclosed in double quotes.
	var str string = "Hello, Go!"
	fmt.Println("String:", str)

	// Type Conversion:
	// Sometimes, you may need to convert one data type into another.
	// For example, converting an integer to a float.
	converted0 := float64(integer) // Converts the integer '42' to a float (42.0).
	fmt.Println("Converted Integer to Float:", converted0)

	// Converting a float to an integer.
	// Note: This truncates the fractional part, so 3.14 becomes 3.
	converted1 := int(floating)
	fmt.Println("Converted Float to Integer:", converted1)

}

// Key Notes for Learners:
// 1. Go is a statically typed language, meaning you must declare the type of a variable explicitly or let Go infer it.
// 2. Type conversion must be done explicitly; Go does not perform implicit conversions like some other languages.
// 3. When converting from a float to an integer, the fractional part is always truncated (not rounded).
