package main

import "fmt"

// Basic function that takes two integers as input and returns their sum.
func add(a int, b int) int {
    return a + b
}

// Demonstrating the use of a basic function.
func basicFunctions() {
    sum := add(5, 10)
    fmt.Println("Sum (using add):", sum)
}

// Variadic function: Accepts a variable number of arguments.
// This function calculates the sum of all the numbers passed to it.
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Demonstrating the use of a variadic function.
func variadicFunctions() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println("Sum (using variadic function):", result)
}

// Function with named return values.
// Named return values are declared in the function signature itself.
func multiplyAndDivide(a int, b int) (product int, quotient int) {
    product = a * b
    quotient = a / b
    return // Implicitly returns the named variables.
}

// Demonstrating the use of named return values.
func namedReturnFunctions() {
    product, quotient := multiplyAndDivide(10, 2)
    fmt.Println("Product:", product)
    fmt.Println("Quotient:", quotient)
}

// Anonymous function: A function without a name, often used for short-term tasks.
func anonymousFunctions() {
    // Define and immediately invoke an anonymous function.
    result := func(a int, b int) int {
        return a - b
    }(10, 3)
    fmt.Println("Result (using anonymous function):", result)
}

// Entry point to test all types of functions.
func functions() {
    fmt.Println("Demonstrating different types of functions:")
    
    // Basic function.
    basicFunctions()
    
    // Variadic function.
    variadicFunctions()
    
    // Named return values.
    namedReturnFunctions()
    
    // Anonymous function.
    anonymousFunctions()
}
