package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type Writer interface {
	Write(data string) error
}

type Reader interface {
	Read() (string, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type FileHandler struct {
	content string
}

func (f *FileHandler) Write(data string) error {
	f.content = data
	fmt.Println("Writing:", data)
	return nil
}

func (f *FileHandler) Read() (string, error) {
	fmt.Println("Reading:", f.content)
	return f.content, nil
}

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 6, Height: 4, Side1: 5, Side2: 5, Side3: 6}

	fmt.Println("Rectangle:")
	printShapeInfo(rect)

	fmt.Println("\nCircle:")
	printShapeInfo(circle)

	fmt.Println("\nTriangle:")
	printShapeInfo(triangle)

	shapes := []Shape{rect, circle, triangle}
	fmt.Printf("\nTotal area of all shapes: %.2f\n", totalArea(shapes))

	var s Shape
	s = rect
	fmt.Printf("\nShape interface holding Rectangle: Area = %.2f\n", s.Area())

	s = circle
	fmt.Printf("Shape interface holding Circle: Area = %.2f\n", s.Area())

	var fh FileHandler
	fh.Write("Hello, Go!")
	data, _ := fh.Read()
	fmt.Println("Data read:", data)

	var rw ReadWriter = &fh
	rw.Write("Interface composition example")
	content, _ := rw.Read()
	fmt.Println("Content:", content)

	person := Person{Name: "Alice", Age: 30}
	fmt.Println("\nPerson:", person)

	fmt.Println("\nEmpty interface examples:")
	describe(42)
	describe("hello")
	describe(true)
	describe(person)
	describe([]int{1, 2, 3})

	var empty interface{}
	empty = 42
	fmt.Printf("\nEmpty interface: %v (type: %T)\n", empty, empty)

	empty = "now a string"
	fmt.Printf("Empty interface: %v (type: %T)\n", empty, empty)

	value, ok := empty.(string)
	if ok {
		fmt.Printf("Type assertion successful: %s\n", value)
	}

	switch v := empty.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	checkShape(rect)
	checkShape(person)
}

func checkShape(i interface{}) {
	s, ok := i.(Shape)
	if ok {
		fmt.Printf("\nThis is a Shape with area: %.2f\n", s.Area())
	} else {
		fmt.Printf("\nThis is not a Shape: %T\n", i)
	}
}
