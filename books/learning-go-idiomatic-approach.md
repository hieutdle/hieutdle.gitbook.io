# 📔 Learning Go: An Idiomatic Approach to Real-World Go Programming

**[You can find all the code for the notes of this book here](https://github.com/ledinhtrunghieu/learning-go)**

## Primitive Types and Declarations

### Literals

**Integer literals** are normally base ten, but different
prefixes are used to indicate other bases: `0b` for binary (base two), `0o` for octal (base
eight), or `0x` for hexadecimal (base sixteen). A leading 0 with no letter after it is another way to represent an
**octal literal**. 

**Rune literals** represent characters and are surrounded by single quotes. In Go single quotes and double quotes 
are **not interchangeable**. Rune literals can be written as single Unicode characters `'a'`, 
8-bit octal numbers `'\141'`, 8-bit hexadecimal numbers `'\x61'`, etc.

A **byte** is an alias for **uint8**; it is legal to assign, compare, or perform mathematical operations between 
a byte and a uint8. Note: **Uint8** value range is **0 to 255** and **int8** range is **-128 to 127**
(-128: 10000000 (1 is a sign bit), 127: 01111111).

The second special name is **int** and **unit**. On a 32-bit CPU, int is a 32-bit signed integer like
an int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.

**Float**: float32 and float64. While Go lets you use `==` and `!=` to compare floats, don’t do it.
Instead, define a maximum allowed variance and see if the difference 
between two floats is less than that.

**Strings** in Go are immutable; you can reassign the value of a string variable, but you
cannot change the value of the string that is assigned to it.

The **rune** type is an alias for the **int32** type, just like byte is an alias for uint8.

### var Versus :=

```go
var x int = 10
var x = 10
var x int
x:= 10
```

Situations within functions where you should avoid :=:
* When initializing a variable to its zero value, use `var x int`. This makes it clear
  that the zero value is intended.
* While it is legal to use a type conversion to specify
  the type of the value and use := to write `x := byte(20)`, it is idiomatic to write
  `var x byte = 20`

### Typed and Untyped Constants
Constants can be typed or untyped. 
All of the following assignments are legal:

````go
const x = 10
var y int = x
var z float64 = x
var d byte = x
````

### Naming Variables and Constants

The names `k` and `v` (short for key and value) are used as
the variable names in a `for-range` loop. If you are using a standard for loop, `i` and `j`
are common names for the index variable. People will use the first
letter of a type as the variable name (for example, `i` for integers, `f` for floats, `b` for
booleans).

## Composite Types

### Arrays

```go
var x = [12]int{1, 5: 4, 6, 10: 100, 15}
```

This creates an array of 12 ints with the following values: [1, 0, 0, 0, 0, 4, 6, 0, 0, 0,
100, 15]. `5: 4` means 0 0 0 0 to index 5, then 4.


### Slices


One slice is appended onto another by using the `…` operator to expand the source slice
into individual values.

```go
y := []int{20, 30, 40}
x = append(x, y...)
```


When a slice grows via append, it takes time for the Go runtime to allocate new mem‐
ory and copy the existing data from the old memory to the new. The old memory also
needs to be garbage collected. For this reason, the Go runtime usually increases a slice
by more than one each time it runs out of capacity. The rules as of Go 1.14 are to
double the size of the slice when the capacity is less than 1,024 and then grow by at
least 25% afterward.

### Declaring slices

```go
var data []int // a slice that stays nil
var x = []int{} // empty slice zero-length, non-nil
data := []int{2, 4, 6, 8} // slice with default values
```

If you have a good idea of how large your slice needs to be, but don’t know what those
values will be when you are writing the program, use make with a zero length and a specified capacity. This
allows you to use append to add items to the slice. 

Append can make overlapping slices

```go
x := make([]int, 0, 5)
x = append(x, 1, 2, 3, 4)
y := x[:2]
z := x[2:]
fmt.Println(cap(x), cap(y), cap(z))
y = append(y, 30, 40, 50)
x = append(x, 60)
z = append(z, 70)
fmt.Println("x:", x)
fmt.Println("y:", y)
fmt.Println("z:", z)
```

```go
5 5 3
x: [1 2 30 40 70]
y: [1 2 30 40 70]
z: [30 40 70]
```

The full slice expression includes a third part, which
indicates the last position in the parent slice’s capacity that’s available for the subslice.

```go
y := x[:2:2]
z := x[2:4:4]
```

```go
5 2 2
x: [1 2 3 4 60]
y: [1 2 30 40 50]
z: [3 4 70]
```
