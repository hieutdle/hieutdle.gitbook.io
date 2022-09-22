# 📔 Learning Go: An Idiomatic Approach to Real-World Go Programming

**[You can find all the code for the notes of this book here](https://github.com/ledinhtrunghieu/learning-go)**

## Primitive Types and Declarations

### Literals

**Integer literals** are normally base ten, but different
prefixes are used to indicate other bases: `0b` for binary (base two), `0o` or leading 0 with no letter follow for octal (base
eight), or `0x` for hexadecimal (base sixteen). 

**Rune literals** represent characters and are surrounded by single quotes. In Go, Single quotes and double quotes 
are **not interchangeable**. Rune literals can be written as single Unicode characters `'a'`, 
8-bit octal numbers `'\141'`, 8-bit hexadecimal numbers `'\x61'`, etc.

A **byte** is an alias for **uint8**. Note: **Uint8** value range is **0 to 255** and **int8** range is **-128 to 127**
(-128: 10000000 (1 is a sign bit), 127: 01111111).

An **int** (or **unit**). On a 32-bit or 64-bit CPU, `int` is a 32-bit or 64-bit signed integer like
an int32 or int64.

**Float**: float32 and float64. While Go lets you use `==` and `!=` to compare floats, don’t do it.
Instead, define a maximum allowed variance and see if the difference 
between two floats is less than that.

**Strings** in Go are immutable; you can reassign the value of a string variable, but you
cannot change the value of the string that is assigned to it.

The **rune** type is an alias for the **int32** type.

### var Versus :=

```go
var x int = 10
var x = 10
var x int
x:= 10
```

Situations within functions where you should avoid `:=`:
* When initializing a variable to its zero value, use `var x int`. This makes it clear
  that the zero value is intended.
* specify
  the type of the value: `x := byte(20)`, it is idiomatic to write
  `var x byte = 20`

### Typed and Untyped Constants
Constants can be typed or untyped. 
All of the following assignments are legal:

````go
const x = 10
var y float = x
var z byte = x
````

### Naming Variables and Constants

The names `k` and `v` (short for key and value) are used as
the variable names in a `for-range` loop, standard for loop: `i` and `j`
are common names for the index variable. People will use the first
letter of a type (for example, `i` for integers, `f` for floats, `b` for
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


When a slice grows via append, the Go runtime usually increases a slice
by more than one each time it runs out of capacity. The rules as of Go 1.14 are to
double the size of the slice when the capacity is less than 1,024 and then grow by at
least 25% afterward.

### Declaring slices

```go
var data []int // a slice that stays nil
var x = []int{} // empty slice zero-length, non-nil
data := []int{2, 4, 6, 8} // slice with default values
x := make([]int, 0, 5)
```

If you have a good idea of how large your slice needs to be, use `make` with a zero length and a specified capacity. 

Append can make overlapping slices. The full slice expression includes a third part, which
indicates the last position in the parent slice’s capacity that’s available for the subslice.

```go
y := x[:2:2]
z := x[2:4:4]
```

### copy

Create a slice that’s independent of the original, use the built-in
`copy` function.

```go
x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
copy(y, x[2:]) // first is destination and second is source
```

### Converting strings to slices

```go
var s string = "Hello"
var bs []byte = []byte(s)
var rs []rune = []rune(s)

[72 101 108 108 111 44 32 240 159 140 158]
[72 101 108 108 111 44 32 127774]
```

### Map

`map[keyType]valueType`
```go
var nilMap map[string]int // A nil map is equivalent to an empty map 
// except that elements can’t be added.
totalWins := map[string]int{} // empty map
m := map[string]int{
    "hello": 5,
    "world": 10,
}
ages := make(map[int][]string, 10) // know length

totalWins["Orcas"] = 1 // write
fmt.Println(totalWins["Orcas"]) // read
delete(m, "hello") // delete
```

#### The comma ok Idiom

```go
m := map[string]int{
  "hello": 5,
  "world": 0,
} 
v, ok := m["hello"]
```

#### Using Maps as Sets

```go
intSet := map[int]bool{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
    intSet[v] = true
}
if intSet[100] {
    fmt.Println("100 is in the set")
}
```

### Comparing and Converting Structs
Go does allow you to perform a type conversion from one struct type to
another if the fields of both structs have the same names, order, and types.

## Blocks, Shadows, and Control Structures

### for, Four Ways

```go
// A complete, C-style for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
// A condition-only for
i := 1
for i < 100 {
	fmt.Println(i)
	i = i * 2
}
// An infinite for
for {
	fmt.Println("Hello")
}
// for-range
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
	fmt.Println(v)
}

uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
for k := range uniqueNames {
    fmt.Println(k)
}	
```
The most common reason for iterating over the key is when a map is being used as a
set. 

## Struct

**Variadic parameters**

```go
func addTo(base int, vals ...int) []int {
		out := make([]int, 0, len(vals))
		for _, v := range vals {
		out = append(out, base+v)
	}
		return out
}
fmt.Println(addTo(3))
fmt.Println(addTo(3, 2))
fmt.Println(addTo(3, 2, 4, 6, 8))
a := []int{4, 3}
fmt.Println(addTo(3, a...))
fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
```

While named return values can sometimes help clarify your code, they do have some
potential corner cases - problem of shadowing

```go
func divAndRemainder(numerator, denominator int) (result int, remainder int, 
	err error) {
  // assign some values
  result, remainder = 20, 30
  if denominator == 0 {
  return 0, 0, errors.New("cannot divide by zero")
  }
  return numerator / denominator, numerator % denominator, nil
}
```
Pass 5  and 2 to this function. Return 2 1.

**Blank Returns—Never Use These!**

When there’s invalid input, we return immediately. At first, 
you might find blank returns handy since they allow you to avoid some typing. 
However, most experienced Go developers consider blank returns a bad idea
because they make it harder to understand data flow.

**Functions Are Values**
```go
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
var opMap = map[string]func(int, int) int{
  "+": add,
  "-": sub,
}
```

### Closures


