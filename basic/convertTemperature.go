package main

import "fmt"

// Converts temperature between Celsius, Fahrenheit, and Kelvin
func convertTemperature() {
	var choice int
	var temperature float64

	fmt.Println("Temperature Unit Converter:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Celsius to Kelvin")
	fmt.Println("4. Kelvin to Celsius")
	fmt.Print("Choose an option (1-4): ")
	fmt.Scanln(&choice)

	fmt.Print("Enter the temperature: ")
	fmt.Scanln(&temperature)

	switch choice {
	case 1:
		fmt.Printf("%.2f°C = %.2f°F\n", temperature, (temperature*9/5)+32)
	case 2:
		fmt.Printf("%.2f°F = %.2f°C\n", temperature, (temperature-32)*5/9)
	case 3:
		fmt.Printf("%.2f°C = %.2fK\n", temperature, temperature+273.15)
	case 4:
		fmt.Printf("%.2fK = %.2f°C\n", temperature, temperature-273.15)
	default:
		fmt.Println("Invalid choice!")
	}
}
