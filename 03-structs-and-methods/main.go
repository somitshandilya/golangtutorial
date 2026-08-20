package main

import "fmt"

// ===== Struct Definitions =====

// Person is a struct that groups related data together
// Structs are like classes in other languages but without inheritance
type Person struct {
	Name string // Field: person's name
	Age  int    // Field: person's age
	City string // Field: person's city
}

// Rectangle struct represents a rectangle shape
type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}

// ===== Methods with Value Receiver =====

// Area calculates the area of a rectangle
// (r Rectangle) is a value receiver - receives a copy of the struct
// The method doesn't modify the original struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
// Value receiver - doesn't modify the original
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ===== Methods with Pointer Receiver =====

// Scale modifies the rectangle dimensions by a factor
// (r *Rectangle) is a pointer receiver - can modify the original struct
// Use pointer receivers when you need to modify the struct
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor  // Modifies the original struct
	r.Height *= factor // Modifies the original struct
}

// Circle struct represents a circle shape
type Circle struct {
	Radius float64 // Radius of the circle
}

// Area calculates the area of a circle (πr²)
// Value receiver - doesn't modify the circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the circumference of a circle (2πr)
// Value receiver - doesn't modify the circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// ===== Struct Embedding (Composition) =====

// Employee embeds Person struct (composition, not inheritance)
// Employee "has a" Person, gaining all Person fields
type Employee struct {
	Person              // Embedded struct - fields promoted to Employee
	EmployeeID int      // Additional field specific to Employee
	Department string   // Additional field specific to Employee
}

// DisplayInfo prints employee information
// Can access embedded Person fields directly (e.Name, e.Age, e.City)
func (e Employee) DisplayInfo() {
	fmt.Printf("Employee: %s, Age: %d, City: %s, ID: %d, Dept: %s\n",
		e.Name, e.Age, e.City, e.EmployeeID, e.Department)
}

func main() {
	// ===== Struct Initialization Method 1: Named Fields =====
	// Most readable way - explicitly name each field
	p1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Println("Person:", p1)
	// Access struct fields using dot notation
	fmt.Printf("Name: %s, Age: %d, City: %s\n", p1.Name, p1.Age, p1.City)

	// ===== Struct Initialization Method 2: Positional =====
	// Values must be in the exact order as struct definition
	p2 := Person{"Bob", 25, "London"}
	fmt.Println("\nPerson 2:", p2)

	// ===== Struct Initialization Method 3: Zero Value + Assignment =====
	// Declare struct with zero values, then assign fields
	var p3 Person
	p3.Name = "Charlie"
	p3.Age = 35
	p3.City = "Paris"
	fmt.Println("Person 3:", p3)

	// ===== Calling Methods on Structs =====
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("\nRectangle: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
	// Call value receiver methods
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// ===== Pointer Receiver Method =====
	// Scale modifies the original struct because it uses pointer receiver
	fmt.Println("\nScaling rectangle by 2...")
	rect.Scale(2) // Go automatically takes address (&rect)
	fmt.Printf("New dimensions: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Area: %.2f\n", rect.Area())

	// ===== Another Struct Example =====
	circle := Circle{Radius: 7}
	fmt.Printf("\nCircle: Radius=%.2f\n", circle.Radius)
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())

	// ===== Embedded Struct (Composition) =====
	// Employee contains Person - all Person fields are accessible
	emp := Employee{
		Person: Person{ // Initialize embedded struct
			Name: "David",
			Age:  28,
			City: "Tokyo",
		},
		EmployeeID: 1001,
		Department: "Engineering",
	}
	fmt.Println("\nEmployee Info:")
	emp.DisplayInfo()
	// Access embedded Person fields directly (promoted fields)
	fmt.Printf("Accessing embedded fields: %s from %s\n", emp.Name, emp.City)

	// ===== Pointer to Struct =====
	// Create pointer to struct using &
	ptrPerson := &Person{Name: "Eve", Age: 32, City: "Berlin"}
	fmt.Println("\nPointer to struct:", ptrPerson)
	// Go automatically dereferences pointer for field access
	fmt.Printf("Name via pointer: %s\n", ptrPerson.Name)
}
