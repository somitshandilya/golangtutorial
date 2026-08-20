package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ===== Custom Error Types =====

// ValidationError is a custom error type
// Implements the error interface by having an Error() method
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface for ValidationError
// This method is required to satisfy the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// ===== Basic Error Handling =====

// divide performs division with error handling
// Returns result and error (nil if successful)
func divide(a, b float64) (float64, error) {
	// Check for division by zero
	if b == 0 {
		// Return zero value and error
		return 0, errors.New("division by zero")
	}
	// Return result and nil error (no error)
	return a / b, nil
}

// ===== Error with fmt.Errorf =====

// sqrt calculates square root with formatted error
// fmt.Errorf allows creating errors with formatted strings
func sqrt(x float64) (float64, error) {
	if x < 0 {
		// Create formatted error message
		return 0, fmt.Errorf("cannot calculate square root of negative number: %f", x)
	}
	// Simplified sqrt (not accurate, just for demo)
	return x / 2, nil
}

// ===== Multiple Error Checks =====

// validateAge checks if age is valid
// Returns custom ValidationError if invalid
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "age cannot be greater than 150",
		}
	}
	// Return nil if no error
	return nil
}

// ===== Error Wrapping (Go 1.13+) =====

// processData demonstrates error wrapping
// Wrapping preserves the original error while adding context
func processData(input string) error {
	// Try to convert string to int
	_, err := strconv.Atoi(input)
	if err != nil {
		// Wrap the original error with additional context
		// %w verb wraps the error
		return fmt.Errorf("failed to process data '%s': %w", input, err)
	}
	return nil
}

// ===== Sentinel Errors =====

// Predefined errors that can be compared with errors.Is()
var (
	ErrNotFound     = errors.New("item not found")
	ErrUnauthorized = errors.New("unauthorized access")
	ErrInvalidInput = errors.New("invalid input")
)

// findItem simulates finding an item by ID
// Returns sentinel error if not found
func findItem(id int) (string, error) {
	if id < 0 {
		return "", ErrInvalidInput
	}
	if id > 100 {
		return "", ErrNotFound
	}
	return fmt.Sprintf("Item-%d", id), nil
}

// ===== Panic and Recover =====

// riskyOperation demonstrates panic and recover
// Panic stops normal execution, recover catches panics
func riskyOperation(shouldPanic bool) (result string) {
	// defer runs when function returns (even after panic)
	defer func() {
		// recover() catches panic and returns the panic value
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = "recovered"
		}
	}()

	if shouldPanic {
		// panic stops execution and starts unwinding the stack
		panic("something went wrong!")
	}

	return "success"
}

// ===== Error Checking Patterns =====

// readConfig demonstrates multiple error checks
// Early return pattern - check error and return immediately
func readConfig(filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}

	// Simulate file reading
	if filename == "invalid.txt" {
		return fmt.Errorf("failed to read config file: %s", filename)
	}

	fmt.Println("Config loaded successfully")
	return nil
}

func main() {
	// ===== Basic Error Handling =====
	fmt.Println("===== Basic Error Handling =====")
	
	// Successful division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Division by zero - triggers error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Will print error
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// ===== Formatted Errors =====
	fmt.Println("\n===== Formatted Errors =====")
	
	sqrtResult, err := sqrt(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("sqrt(16) ≈ %.2f\n", sqrtResult)
	}

	// Negative number - triggers error
	_, err = sqrt(-5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ===== Custom Error Types =====
	fmt.Println("\n===== Custom Error Types =====")
	
	// Valid age
	err = validateAge(25)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age 25 is valid")
	}

	// Invalid age - negative
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Validation error:", err)
		// Type assertion to access custom error fields
		if valErr, ok := err.(*ValidationError); ok {
			fmt.Printf("  Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	// Invalid age - too high
	err = validateAge(200)
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	// ===== Error Wrapping =====
	fmt.Println("\n===== Error Wrapping =====")
	
	err = processData("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data processed successfully")
	}

	// Invalid input - triggers wrapped error
	err = processData("abc")
	if err != nil {
		fmt.Println("Error:", err)
		// errors.Unwrap() extracts the wrapped error
		fmt.Printf("Unwrapped error: %v\n", errors.Unwrap(err))
	}

	// ===== Sentinel Errors =====
	fmt.Println("\n===== Sentinel Errors =====")
	
	// Valid ID
	item, err := findItem(50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", item)
	}

	// Not found error
	_, err = findItem(150)
	// errors.Is() checks if error matches sentinel error
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found (expected)")
	}

	// Invalid input error
	_, err = findItem(-1)
	if errors.Is(err, ErrInvalidInput) {
		fmt.Println("Invalid input (expected)")
	}

	// ===== Panic and Recover =====
	fmt.Println("\n===== Panic and Recover =====")
	
	// Normal execution
	result1 := riskyOperation(false)
	fmt.Println("Result:", result1)

	// Panic and recover
	result2 := riskyOperation(true)
	fmt.Println("Result after panic:", result2)

	// ===== Error Checking Patterns =====
	fmt.Println("\n===== Error Checking Patterns =====")
	
	// Successful config load
	err = readConfig("config.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Failed config load
	err = readConfig("invalid.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Empty filename
	err = readConfig("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ===== Multiple Error Returns =====
	fmt.Println("\n===== Multiple Error Returns =====")
	
	// Function that returns multiple values with error
	value, count, err := processMultiple("test", 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Value: %s, Count: %d\n", value, count)
	}

	// Error case
	_, _, err = processMultiple("", 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// processMultiple demonstrates multiple return values with error
// Returns two values and an error
func processMultiple(input string, count int) (string, int, error) {
	if input == "" {
		return "", 0, errors.New("input cannot be empty")
	}
	if count <= 0 {
		return "", 0, errors.New("count must be positive")
	}
	return input, count * 2, nil
}
