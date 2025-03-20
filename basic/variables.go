// variables demonstrates different types of variable declarations in Go.
// We can use the import statement in two ways:
//  1. Single-line import: import "fmt"
//  2. Multi-line import (preferred when importing multiple packages):
//     import (
//     "fmt"
//     "math"
//     "time"
//     )
package main

import (
	"fmt"
)

func variables() {
	// 1. Explicit Type Declaration
	var integerVariable int = 10
	var stringVariable string = "Welcome to Go"

	// 2. Implicit Type Declaration (type inferred by the compiler)
	var floatVariable = 3.14
	var booleanVariable = true

	// 3. Short Variable Declaration (commonly used within functions)
	shortDeclaration := "Efficient and concise"
	shortNumber := 42

	// 4. Multiple Variable Declaration (single line and block format)
	var x, y int = 100, 200
	var (
		blockString string  = "Block style declaration"
		blockFloat  float64 = 9.81
	)

	// 5. Constants (immutable values)
	const appName = "GoLang Variable Demo"
	const maxUsers int = 1000

	// 6. Blank Identifier (used to ignore unwanted values)
	_, ignoredValue := 1, "Ignored"

	// Output all variables with spacing for clarity
	fmt.Println("Explicit Type Declaration:", integerVariable, stringVariable)
	fmt.Println()
	// Note ---> this "	fmt.Println()" blank print statement is used for making some space between outputs so cont get confused.

	fmt.Println("Implicit Type Declaration:", floatVariable, booleanVariable)
	fmt.Println()

	fmt.Println("Short Declaration:", shortDeclaration, shortNumber)
	fmt.Println()

	fmt.Println("Multiple Declaration:", x, y)
	fmt.Println()

	fmt.Println("Block Declaration:", blockString, blockFloat)
	fmt.Println()

	fmt.Println("Constants:", appName, maxUsers)
	fmt.Println()

	fmt.Println("Ignored Value Placeholder:", ignoredValue)
	fmt.Println()
}
