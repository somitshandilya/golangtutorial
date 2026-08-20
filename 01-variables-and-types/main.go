package main

import "fmt"

func main() {
	// Declare variables using different methods
	var name string = "Alice"
	age := 30
	var height float64 = 5.6
	isStudent := false

	// Print variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	// Constants
	const Pi = 3.14159
	fmt.Println("Pi:", Pi)

	// Type conversion
	var x int = 10
	var y float64 = float64(x)
	fmt.Printf("x (int): %d, y (float64): %f\n", x, y)

	// Zero values demonstration
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	fmt.Println("\nZero values:")
	fmt.Printf("int: %d, float64: %f, bool: %t, string: '%s'\n", zeroInt, zeroFloat, zeroBool, zeroString)

	// Multiple variable declaration
	var a, b, c = 1, "hello", true
	fmt.Println("\nMultiple variables:", a, b, c)
}
