package main

import "fmt"

// ===== Struct for Demonstrating Pointer vs Value Receivers =====

// Counter struct holds a count value
type Counter struct {
	Count int
}

// ===== Pointer Receiver Method =====

// Increment modifies the counter using a pointer receiver
// (c *Counter) means this method receives a pointer to Counter
// Changes made here affect the original struct
func (c *Counter) Increment() {
	c.Count++ // Modifies the original Counter
}

// ===== Value Receiver Method =====

// IncrementValue tries to modify counter using value receiver
// (c Counter) means this method receives a COPY of Counter
// Changes made here do NOT affect the original struct
func (c Counter) IncrementValue() {
	c.Count++ // Only modifies the copy, not the original
}

// ===== Pass by Value =====

// modifyValue receives a copy of the integer
// Changes here don't affect the original variable
func modifyValue(x int) {
	x = 100 // Only modifies the local copy
}

// ===== Pass by Pointer =====

// modifyPointer receives a pointer to an integer
// Can modify the original value through the pointer
// *int means "pointer to int"
func modifyPointer(x *int) {
	*x = 100 // Dereference pointer and modify original value
}

// swap exchanges values of two integers using pointers
// Parameters: a, b (*int) - pointers to integers to swap
func swap(a, b *int) {
	*a, *b = *b, *a // Swap values at the addresses
}

// ===== Returning Pointers =====

// createPerson creates a Counter and returns a pointer to it
// It's safe to return pointer to local variable in Go
// Returns: *Counter - pointer to the created Counter
func createPerson(name string, age int) *Counter {
	c := Counter{Count: age}
	return &c // Return address of local variable (safe in Go)
}

func main() {
	// ===== Pointer Basics =====
	// & = address-of operator (gets memory address)
	// * = dereference operator (gets value at address)
	
	var x int = 42
	var p *int = &x // p is a pointer to int, stores address of x

	fmt.Println("Value of x:", x)           // Print value
	fmt.Println("Address of x:", &x)        // Print memory address
	fmt.Println("Value of p (address):", p) // p holds the address
	fmt.Println("Value at address p:", *p)  // Dereference p to get value

	// ===== Modifying Through Pointer =====
	*p = 100 // Change value at address p points to
	fmt.Println("\nAfter *p = 100:")
	fmt.Println("Value of x:", x)          // x is now 100
	fmt.Println("Value at address p:", *p) // Also 100

	// ===== Nil Pointers =====
	// Uninitialized pointers have nil value
	var ptr *int
	fmt.Println("\nNil pointer:", ptr)
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	// Assign address to pointer
	ptr = &x
	fmt.Println("After assignment, ptr points to:", *ptr)

	// ===== Pass by Value vs Pass by Pointer =====
	num := 50
	fmt.Println("\nBefore modifyValue:", num)
	modifyValue(num) // Passes copy - original unchanged
	fmt.Println("After modifyValue:", num) // Still 50

	fmt.Println("\nBefore modifyPointer:", num)
	modifyPointer(&num) // Passes address - can modify original
	fmt.Println("After modifyPointer:", num) // Now 100

	// ===== Swapping Values Using Pointers =====
	a, b := 10, 20
	fmt.Printf("\nBefore swap: a=%d, b=%d\n", a, b)
	swap(&a, &b) // Pass addresses to swap function
	fmt.Printf("After swap: a=%d, b=%d\n", a, b) // Values swapped

	// ===== Pointer Receiver vs Value Receiver =====
	counter := Counter{Count: 0}
	fmt.Println("\nCounter value:", counter.Count)

	// Pointer receiver - modifies original
	counter.Increment() // Go automatically passes &counter
	fmt.Println("After Increment() with pointer receiver:", counter.Count) // 1

	// Value receiver - doesn't modify original
	counter.IncrementValue() // Passes copy of counter
	fmt.Println("After IncrementValue() with value receiver:", counter.Count) // Still 1

	// ===== Pointer to Struct =====
	ptrCounter := &Counter{Count: 5} // Create pointer to Counter
	fmt.Println("\nPointer to Counter:", ptrCounter.Count)
	ptrCounter.Increment() // Can call methods on pointer
	fmt.Println("After Increment():", ptrCounter.Count) // 6

	// ===== Returning Pointers from Functions =====
	newCounter := createPerson("Alice", 10)
	fmt.Println("\nReturned pointer to Counter:", newCounter.Count)

	// ===== Pointers to Arrays =====
	arr := [3]int{1, 2, 3}
	arrPtr := &arr // Pointer to array
	fmt.Println("\nArray:", arr)
	fmt.Println("Array pointer:", arrPtr)
	arrPtr[0] = 100 // Go auto-dereferences for indexing
	fmt.Println("After modifying via pointer:", arr) // [100 2 3]

	// ===== Pointers to Strings =====
	str := "Hello"
	strPtr := &str // Pointer to string
	fmt.Println("\nString:", str)
	fmt.Println("String pointer:", *strPtr)
	*strPtr = "World" // Modify string through pointer
	fmt.Println("After modification:", str) // "World"

	// ===== Pointer to Pointer =====
	// **int means "pointer to pointer to int"
	var y int = 42
	var p1 *int = &y    // p1 points to y
	var p2 **int = &p1  // p2 points to p1 (which points to y)
	fmt.Println("\nPointer to pointer:")
	fmt.Println("Value of y:", y)
	fmt.Println("Value via p1:", *p1)   // Dereference once
	fmt.Println("Value via p2:", **p2)  // Dereference twice
	**p2 = 99 // Modify y through double pointer
	fmt.Println("After **p2 = 99, y =", y) // y is now 99

	// ===== Slices and Pointers =====
	// Slices are reference types - they already contain a pointer
	// No need to pass slice pointer to modify elements
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("\nSlice before:", slice)
	modifySlice(slice) // Pass slice (contains pointer internally)
	fmt.Println("Slice after modifySlice:", slice) // [999 2 3 4 5]
}

// modifySlice modifies slice elements
// Slices are reference types, so changes affect the original
// Parameters: s ([]int) - slice to modify
func modifySlice(s []int) {
	s[0] = 999 // Modifies original slice
}
