# Step 1: Variables and Types

## Introduction

In Go, variables are explicitly typed, but the compiler can often infer the type. This step covers the basics of declaring variables and understanding Go's type system.

## Variable Declaration

### Method 1: Using `var` keyword

```go
var name string = "John"
var age int = 30
var isActive bool = true
```

### Method 2: Type inference with `var`

```go
var name = "John"      // Go infers string
var age = 30           // Go infers int
var isActive = true    // Go infers bool
```

### Method 3: Short declaration (`:=`)

This is the most common way to declare variables inside functions:

```go
name := "John"
age := 30
isActive := true
```

**Note:** The `:=` syntax can only be used inside functions, not at package level.

## Basic Types

### Numeric Types

```go
// Integers
var i int = 42           // Platform dependent (32 or 64 bit)
var i8 int8 = 127        // -128 to 127
var i16 int16 = 32767    // -32768 to 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// Unsigned integers
var ui uint = 42         // Platform dependent
var ui8 uint8 = 255      // 0 to 255 (also called byte)
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

### String Type

```go
var message string = "Hello, Go!"
var multiline = `This is a
multiline string
using backticks`
```

### Boolean Type

```go
var isTrue bool = true
var isFalse bool = false
```

## Zero Values

Variables declared without an explicit initial value are given their zero value:

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // "" (empty string)
```

## Constants

Constants are declared using the `const` keyword and cannot be changed:

```go
const Pi = 3.14159
const AppName = "MyApp"
const MaxConnections = 100
```

## Multiple Variable Declaration

```go
// Multiple variables of same type
var x, y, z int = 1, 2, 3

// Multiple variables with type inference
var a, b, c = 1, "hello", true

// Using short declaration
name, age := "Alice", 25
```

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversion
var s string = string(65)  // "A" (ASCII)
```

## Practice Exercise

Create a file `main.go` and try this:

```go
package main

import "fmt"

func main() {
    // Declare variables using different methods
    var name string = "Alice"
    age := 30
    var height float64 = 5.6
    isStudent := false
    
    // Print variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    
    // Constants
    const Pi = 3.14159
    fmt.Println("Pi:", Pi)
    
    // Type conversion
    var x int = 10
    var y float64 = float64(x)
    fmt.Printf("x (int): %d, y (float64): %f\n", x, y)
}
```

## Run the Program

```bash
go run main.go
```

## Expected Output

```
Name: Alice
Age: 30
Height: 5.6
Is Student: false
Pi: 3.14159
x (int): 10, y (float64): 10.000000
```

## Key Takeaways

1. Use `:=` for short variable declaration inside functions
2. Go is statically typed - types are checked at compile time
3. Zero values are automatically assigned to uninitialized variables
4. Type conversion must be explicit
5. Constants are declared with `const` and cannot be changed

## Next Step

Once you've practiced with variables and types, move on to **Step 2: Functions**.
