package main

import "fmt"

func main() {
	// ===== Simple Function Call =====
	// Calling a function with no parameters and no return value
	sayHello()

	// ===== Function with Parameters =====
	// Passing a string argument to the function
	greet("Alice")

	// ===== Function with Return Value =====
	// Calling a function that returns a single value
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// ===== Multiple Return Values =====
	// Go functions can return multiple values (commonly result + error)
	result, err := divide(10, 2)
	// Always check error before using result
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// ===== Variadic Function =====
	// Function that accepts variable number of arguments
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	// ===== Anonymous Function =====
	// Function without a name, assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(4, 5))

	// ===== Closure =====
	// Function that captures variables from outer scope
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// ===== Defer Statement =====
	// Demonstrates defer execution order
	deferExample()

	// ===== Recursion =====
	// Function calling itself
	fmt.Println("Factorial of 5:", factorial(5))
}

// sayHello is a simple function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// greet takes a string parameter and prints a greeting
// Parameters: name (string) - the name to greet
func greet(name string) {
	fmt.Println("Hello,", name)
}

// add takes two integers and returns their sum
// Parameters: a, b (int) - numbers to add
// Returns: int - the sum of a and b
func add(a, b int) int {
	return a + b
}

// divide performs division and returns result with error handling
// Parameters: a, b (float64) - dividend and divisor
// Returns: (float64, error) - result and error if division by zero
func divide(a, b float64) (float64, error) {
	// Check for division by zero
	if b == 0 {
		// Return zero value and error
		return 0, fmt.Errorf("cannot divide by zero")
	}
	// Return result and nil error (no error)
	return a / b, nil
}

// sumAll is a variadic function that sums any number of integers
// Parameters: numbers (...int) - variable number of integers
// Returns: int - sum of all numbers
func sumAll(numbers ...int) int {
	total := 0
	// range iterates over the slice of numbers
	// _ ignores the index, num is the value
	for _, num := range numbers {
		total += num
	}
	return total
}

// makeCounter returns a closure that maintains state
// Returns: func() int - a function that increments and returns a counter
func makeCounter() func() int {
	// count is captured by the returned function (closure)
	count := 0
	// Return anonymous function that has access to count
	return func() int {
		count++ // Increment the captured variable
		return count
	}
}

// deferExample demonstrates the defer statement
// Deferred functions execute in LIFO order when function returns
func deferExample() {
	// This will execute last (after function completes)
	defer fmt.Println("Deferred: This prints last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
	// When function returns, deferred statement executes
}

// factorial calculates factorial using recursion
// Parameters: n (int) - number to calculate factorial for
// Returns: int - factorial of n
func factorial(n int) int {
	// Base case: stop recursion
	if n <= 1 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)
}
