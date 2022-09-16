# 📔 Learning Go: An Idiomatic Approach to Real-World Go Programming

**[You can find all the code for the notes of this book here](https://github.com/ledinhtrunghieu/learning-go)**

## Primitive Types and Declarations

### Literals

**Integer literals** are sequences of numbers; they are normally base ten, but different
prefixes are used to indicate other bases: `0b` for binary (base two), `0o` for octal (base
eight), or `0x` for hexadecimal (base sixteen). You can use either or upper- or lowercase
letters for the prefix. A leading 0 with no letter after it is another way to represent an
**octal literal**. Do not use it, as it is very confusing.

**Rune literals** represent characters and are surrounded by single quotes. In Go single quotes and double quotes 
are **not interchangeable**. Rune literals can be written as single Unicode characters `('a')`, 
8-bit octal numbers `('\141')`, 8-bit hexadecimal numbers `('\x61')`, 16-bit hexadecimal numbers
`('\u0061')`, or 32-bit Unicode numbers `('\U00000061')`. There are also several back‐
slash escaped rune literals, with the most useful ones being newline `('\n')`, tab
`('\t')`, single quote `('\'')`, double quote `('\"')`, and backslash `('\\')`.

A **byte** is an alias for **uint8**; it is legal to assign, compare, or perform mathematical operations between 
a byte and a uint8. Note: **Uint8** value range is **0 to 255** and **int8** range is **-128 to 127**
(-128: 10000000 (1 is a sign bit), 127: 01111111).

The second special name is **int** and **unit**. On a 32-bit CPU, int is a 32-bit signed integer like
an int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.

**Float**: float32 and float64. While Go lets you use == and != to compare floats, don’t do it. Due to the inexact
nature of floats, two floating point values might not be equal when you think they
should be. Instead, define a maximum allowed variance and see if the difference 
between two floats is less than that.

**Strings** in Go are immutable; you can reassign the value of a string variable, but you
cannot change the value of the string that is assigned to it.

The **rune** type is an alias for the **int32** type, just like byte is an alias for uint8.

### var Versus :=

The most verbose way to declare a variable in Go uses the var keyword, an explicit
type, and an assignment. It looks like this:
```go
var x int = 10
	
var x = 10
	
var x int
```

Go also supports a short declaration format.

``` go
x, y := 30, "hello"
```

Situations within functions where you should avoid :=:
* When initializing a variable to its zero value, use var x int. This makes it clear
  that the zero value is intended.
* When assigning an untyped constant or a literal to a variable and the default type
  for the constant or literal isn’t the type you want for the variable, use the long var
  form with the type specified. While it is legal to use a type conversion to specify
  the type of the value and use := to write x := byte(20), it is idiomatic to write
  `var x byte = 20`

### Typed and Untyped Constants
Constants can be typed or untyped. An untyped constant works exactly like a literal;
it has no type of its own, but does have a default type that is used when no other type
can be inferred.
All of the following assignments are legal:

````go
const x = 10

var y int = x

var z float64 = x

var d byte = x
````

In many languages, constants are always written in all uppercase letters, with words
separated by underscores (names like `INDEX_COUNTER` or `NUMBER_TRIES`). Go
does not follow this pattern. This is because Go uses the case of the first letter in the
name of a package-level declaration to determine if the item is accessible outside the
package. 

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


## Slices

When a slice grows via append, it takes time for the Go runtime to allocate new mem‐
ory and copy the existing data from the old memory to the new. The old memory also
needs to be garbage collected. For this reason, the Go runtime usually increases a slice
by more than one each time it runs out of capacity. The rules as of Go 1.14 are to
double the size of the slice when the capacity is less than 1,024 and then grow by at
least 25% afterward.


