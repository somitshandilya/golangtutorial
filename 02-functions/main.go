package main

import "fmt"

func main() {
	// Simple function call
	sayHello()

	// Function with parameters
	greet("Alice")

	// Function with return value
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// Multiple return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// Variadic function
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	// Anonymous function
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(4, 5))

	// Closure
	counter := makeCounter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

	// Defer
	deferExample()

	// Recursion
	fmt.Println("Factorial of 5:", factorial(5))
}

func sayHello() {
	fmt.Println("Hello, World!")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func deferExample() {
	defer fmt.Println("Deferred: This prints last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
