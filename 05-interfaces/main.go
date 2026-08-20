package main

import (
	"fmt"
	"math"
)

// ===== Interface Definition =====

// Shape interface defines behavior for geometric shapes
// Any type that implements Area() and Perimeter() satisfies this interface
// Interfaces are implemented implicitly (no "implements" keyword)
type Shape interface {
	Area() float64      // Method to calculate area
	Perimeter() float64 // Method to calculate perimeter
}

// ===== Types Implementing Shape Interface =====

// Rectangle implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area implements Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter implements Shape interface for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
	Radius float64
}

// Area implements Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter implements Shape interface for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle implements Shape interface
type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

// Area implements Shape interface for Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter implements Shape interface for Triangle
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

// ===== Polymorphism with Interfaces =====

// printShapeInfo accepts any type that implements Shape interface
// This is polymorphism - different types can be passed to same function
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// totalArea calculates total area of multiple shapes
// Slice of Shape interface can hold different concrete types
func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area() // Calls appropriate Area() for each type
	}
	return total
}

// ===== Interface Composition =====

// Writer interface defines write behavior
type Writer interface {
	Write(data string) error
}

// Reader interface defines read behavior
type Reader interface {
	Read() (string, error)
}

// ReadWriter combines Reader and Writer interfaces
// This is interface composition (embedding)
type ReadWriter interface {
	Reader // Embeds Reader interface
	Writer // Embeds Writer interface
}

// FileHandler implements ReadWriter interface
type FileHandler struct {
	content string
}

// Write implements Writer interface
// Uses pointer receiver to modify struct
func (f *FileHandler) Write(data string) error {
	f.content = data
	fmt.Println("Writing:", data)
	return nil
}

// Read implements Reader interface
func (f *FileHandler) Read() (string, error) {
	fmt.Println("Reading:", f.content)
	return f.content, nil
}

// ===== Custom String Representation =====

// Stringer interface (similar to fmt.Stringer in standard library)
type Stringer interface {
	String() string
}

// Person struct
type Person struct {
	Name string
	Age  int
}

// String implements Stringer interface for Person
// When Person is printed, this method is called automatically
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// ===== Empty Interface =====

// describe accepts empty interface (interface{}) - can accept ANY type
// interface{} has no methods, so all types satisfy it
func describe(i interface{}) {
	// %T prints type, %v prints value
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	// ===== Creating Shapes =====
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 6, Height: 4, Side1: 5, Side2: 5, Side3: 6}

	// ===== Polymorphism in Action =====
	// Same function works with different types
	fmt.Println("Rectangle:")
	printShapeInfo(rect) // Rectangle implements Shape

	fmt.Println("\nCircle:")
	printShapeInfo(circle) // Circle implements Shape

	fmt.Println("\nTriangle:")
	printShapeInfo(triangle) // Triangle implements Shape

	// ===== Slice of Interfaces =====
	// Can store different types that implement same interface
	shapes := []Shape{rect, circle, triangle}
	fmt.Printf("\nTotal area of all shapes: %.2f\n", totalArea(shapes))

	// ===== Interface Variables =====
	// Interface variable can hold any type that implements the interface
	var s Shape
	s = rect // Assign Rectangle to Shape interface
	fmt.Printf("\nShape interface holding Rectangle: Area = %.2f\n", s.Area())

	s = circle // Reassign to Circle
	fmt.Printf("Shape interface holding Circle: Area = %.2f\n", s.Area())

	// ===== Interface Composition Example =====
	var fh FileHandler
	fh.Write("Hello, Go!")
	data, _ := fh.Read() // _ ignores error return value
	fmt.Println("Data read:", data)

	// FileHandler implements ReadWriter (has both Read and Write)
	var rw ReadWriter = &fh // Must use pointer (methods have pointer receiver)
	rw.Write("Interface composition example")
	content, _ := rw.Read()
	fmt.Println("Content:", content)

	// ===== Custom String Method =====
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("\nPerson:", person) // Calls String() method

	// ===== Empty Interface (interface{}) =====
	// Can accept any type
	fmt.Println("\nEmpty interface examples:")
	describe(42)              // int
	describe("hello")         // string
	describe(true)            // bool
	describe(person)          // Person
	describe([]int{1, 2, 3})  // slice

	// ===== Dynamic Typing with Empty Interface =====
	// Empty interface can hold any value and change types
	var empty interface{}
	empty = 42 // Holds int
	fmt.Printf("\nEmpty interface: %v (type: %T)\n", empty, empty)

	empty = "now a string" // Now holds string
	fmt.Printf("Empty interface: %v (type: %T)\n", empty, empty)

	// ===== Type Assertion =====
	// Extract concrete type from interface
	// Syntax: value, ok := interfaceVar.(Type)
	value, ok := empty.(string) // Try to assert as string
	if ok {
		fmt.Printf("Type assertion successful: %s\n", value)
	}

	// ===== Type Switch =====
	// Check type of interface value and handle each case
	switch v := empty.(type) { // type switch syntax
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v) // v is string in this case
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// ===== Type Assertion with Interfaces =====
	checkShape(rect)   // Rectangle implements Shape
	checkShape(person) // Person does not implement Shape
}

// checkShape checks if a value implements Shape interface
// Parameters: i (interface{}) - can be any type
func checkShape(i interface{}) {
	// Try to assert i as Shape interface
	s, ok := i.(Shape) // ok is true if i implements Shape
	if ok {
		fmt.Printf("\nThis is a Shape with area: %.2f\n", s.Area())
	} else {
		fmt.Printf("\nThis is not a Shape: %T\n", i)
	}
}
