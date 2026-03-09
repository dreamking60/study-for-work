# day2
## Multiple Package import
Every Go program is made up of packages.

Programs start running in package main.

```go
package main

import (
    "fmt"
    "time"
)
```
## Exported names
In Go, a name is exported if it begins with a capital letter.
For example, Pizza is an exported name, as is Pi, which is exported
from the math package. 
pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. 
Any “unexported" names are not accessible from outside the pacakge.

## Functions
A function can take zero or more arguments.
In this example, add take two parameters of type int.
Notice that the type comes after the variable name.

```go
func add(x int, y int) int {
	return x + y
}
```
When two or more consecutive named function parameters share a type,
you can omit the type from all but the last.

from `x int, y int` to `x, y int`


A function can return any number of results.
```go
func swap(x, y string) (string, string) {
	return y, x
}
```

Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

## Variable
The var statement declares a list of variables; 
as in function argument lists, the type is last.

A var statement can be at package or function level. 
We see both in this example.

```go
var a, b, c int
```

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; 
the variable will take the type of the initializer.

```go
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
Inside a function, the := short assignment statement can be used 
in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword 
(var, func, and so on) and so the := construct is not available.

```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
## Basic types
Go's basic types are
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
The example shows variables of several types, 
and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems 
and 64 bits wide on 64-bit systems. When you need an integer value you should use int 
unless you have a specific reason to use a sized or unsigned integer type.


## Zero Values
Variables declared without an explicit initial value are given their zero value.

The zero value is:
- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

## Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

## Type inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

var i int
j := i // j is an int
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
Try changing the initial value of v in the example code and observe how its type is affected.

## Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using := syntax.


Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```
## Struct
A struct.
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

## Error
The go import multiple package require () but not {}.
```
cmd/app/main.go:2:8: missing import path
```

