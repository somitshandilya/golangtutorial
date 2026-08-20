package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
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

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

type Employee struct {
	Person
	EmployeeID int
	Department string
}

func (e Employee) DisplayInfo() {
	fmt.Printf("Employee: %s, Age: %d, City: %s, ID: %d, Dept: %s\n",
		e.Name, e.Age, e.City, e.EmployeeID, e.Department)
}

func main() {
	p1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Println("Person:", p1)
	fmt.Printf("Name: %s, Age: %d, City: %s\n", p1.Name, p1.Age, p1.City)

	p2 := Person{"Bob", 25, "London"}
	fmt.Println("\nPerson 2:", p2)

	var p3 Person
	p3.Name = "Charlie"
	p3.Age = 35
	p3.City = "Paris"
	fmt.Println("Person 3:", p3)

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("\nRectangle: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	fmt.Println("\nScaling rectangle by 2...")
	rect.Scale(2)
	fmt.Printf("New dimensions: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Area: %.2f\n", rect.Area())

	circle := Circle{Radius: 7}
	fmt.Printf("\nCircle: Radius=%.2f\n", circle.Radius)
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())

	emp := Employee{
		Person: Person{
			Name: "David",
			Age:  28,
			City: "Tokyo",
		},
		EmployeeID: 1001,
		Department: "Engineering",
	}
	fmt.Println("\nEmployee Info:")
	emp.DisplayInfo()
	fmt.Printf("Accessing embedded fields: %s from %s\n", emp.Name, emp.City)

	ptrPerson := &Person{Name: "Eve", Age: 32, City: "Berlin"}
	fmt.Println("\nPointer to struct:", ptrPerson)
	fmt.Printf("Name via pointer: %s\n", ptrPerson.Name)
}
