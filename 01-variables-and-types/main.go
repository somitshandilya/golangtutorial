package main

import "fmt"

func main() {
	// ===== Variable Declaration Methods =====
	
	// Method 1: Using 'var' keyword with explicit type
	// Syntax: var variableName type = value
	var name string = "Alice"
	
	// Method 2: Short declaration using := (type inference)
	// This is the most common way inside functions
	age := 30
	
	// Method 3: Using 'var' with explicit type for floating point
	var height float64 = 5.6
	
	// Short declaration for boolean
	isStudent := false

	// ===== Printing Variables =====
	// fmt.Println prints values with automatic spacing
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	// ===== Constants =====
	// Constants are declared with 'const' and cannot be changed
	// They must be assigned at declaration time
	const Pi = 3.14159
	fmt.Println("Pi:", Pi)

	// ===== Type Conversion =====
	// Go requires explicit type conversion (no automatic conversion)
	var x int = 10
	// Convert int to float64 using float64(x)
	var y float64 = float64(x)
	// %d = integer, %f = floating point
	fmt.Printf("x (int): %d, y (float64): %f\n", x, y)

	// ===== Zero Values =====
	// Variables declared without initialization get "zero values"
	var zeroInt int           // 0
	var zeroFloat float64     // 0.0
	var zeroBool bool         // false
	var zeroString string     // "" (empty string)
	fmt.Println("\nZero values:")
	// %t = boolean format
	fmt.Printf("int: %d, float64: %f, bool: %t, string: '%s'\n", zeroInt, zeroFloat, zeroBool, zeroString)

	// ===== Multiple Variable Declaration =====
	// Declare multiple variables of different types in one line
	var a, b, c = 1, "hello", true
	fmt.Println("\nMultiple variables:", a, b, c)
}
