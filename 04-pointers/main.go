package main

import "fmt"

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}

func (c Counter) IncrementValue() {
	c.Count++
}

func modifyValue(x int) {
	x = 100
}

func modifyPointer(x *int) {
	*x = 100
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func createPerson(name string, age int) *Counter {
	c := Counter{Count: age}
	return &c
}

func main() {
	var x int = 42
	var p *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address):", p)
	fmt.Println("Value at address p:", *p)

	*p = 100
	fmt.Println("\nAfter *p = 100:")
	fmt.Println("Value of x:", x)
	fmt.Println("Value at address p:", *p)

	var ptr *int
	fmt.Println("\nNil pointer:", ptr)
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	ptr = &x
	fmt.Println("After assignment, ptr points to:", *ptr)

	num := 50
	fmt.Println("\nBefore modifyValue:", num)
	modifyValue(num)
	fmt.Println("After modifyValue:", num)

	fmt.Println("\nBefore modifyPointer:", num)
	modifyPointer(&num)
	fmt.Println("After modifyPointer:", num)

	a, b := 10, 20
	fmt.Printf("\nBefore swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	counter := Counter{Count: 0}
	fmt.Println("\nCounter value:", counter.Count)

	counter.Increment()
	fmt.Println("After Increment() with pointer receiver:", counter.Count)

	counter.IncrementValue()
	fmt.Println("After IncrementValue() with value receiver:", counter.Count)

	ptrCounter := &Counter{Count: 5}
	fmt.Println("\nPointer to Counter:", ptrCounter.Count)
	ptrCounter.Increment()
	fmt.Println("After Increment():", ptrCounter.Count)

	newCounter := createPerson("Alice", 10)
	fmt.Println("\nReturned pointer to Counter:", newCounter.Count)

	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Println("\nArray:", arr)
	fmt.Println("Array pointer:", arrPtr)
	arrPtr[0] = 100
	fmt.Println("After modifying via pointer:", arr)

	str := "Hello"
	strPtr := &str
	fmt.Println("\nString:", str)
	fmt.Println("String pointer:", *strPtr)
	*strPtr = "World"
	fmt.Println("After modification:", str)

	var y int = 42
	var p1 *int = &y
	var p2 **int = &p1
	fmt.Println("\nPointer to pointer:")
	fmt.Println("Value of y:", y)
	fmt.Println("Value via p1:", *p1)
	fmt.Println("Value via p2:", **p2)
	**p2 = 99
	fmt.Println("After **p2 = 99, y =", y)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("\nSlice before:", slice)
	modifySlice(slice)
	fmt.Println("Slice after modifySlice:", slice)
}

func modifySlice(s []int) {
	s[0] = 999
}
