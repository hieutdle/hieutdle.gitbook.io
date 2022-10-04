# 📔 Learning Go: An Idiomatic Approach to Real-World Go Programming

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

**Variadic parameters** allows any number of input parameters

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

Functions declared inside of functions are closures.

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

### Go Is Call By Value
It means that when you supply a variable for a parameter to a function, Go
always makes a copy of the value of the variable

```go
type person struct {
  age int
  name string
}

func modifyFails(i int, s string, p person) {
  i = i * 2
  s = "Goodbye"
  p.name = "Bob"
}

func main() {
  p := person{}
  i := 2
  s := "Hello"
  modifyFails(i, s, p)
  fmt.Println(i, s, p)
  // won't change: 2 Hello {0 }
}
```
Write a function to modify a map parameter and a function to modify
a slice parameter.  Any changes made to a map parameter
are reflected in the variable passed into the function (includes delete). For a slice, you can modify any element in the slice, but you can’t lengthen the slice. It’s because maps and slices are both implemented with pointers.

## Pointer

A **pointer type**
```go
x := 10
var pointerToX *int // *int is just a type like: var i int
pointerToX = &x // store address of a variable
fmt.Println(pointerToX) // prints a memory address
fmt.Println(*pointerToX) // prints 10
z := 5 + *pointerToX
fmt.Println(z) // prints 15
```

While different types of variables can take up different numbers of memory locations (boolean 1 byte, int32 4 byte),
every pointer, no matter what type it is pointing to, is always the same size: a number
that holds the location in memory where the data is stored. 

```go
func failedUpdate(px *int) {
  x2 := 20
  px = &x2
}
func update(px *int) {
  *px = 20
}
func main() {
  x := 10
  failedUpdate(&x)
  fmt.Println(x) // prints 10
  update(&x)
  fmt.Println(x) // prints 20
}
```
we start with `x` in `main` set to `10`. When we call `failedUpdate`, we
copy the address of `x` into the parameter `px`. Next, we declare `x2` in `failedUpdate`, set
to `20`. We then point `px` in `failedUpdate` to the address of `x2`. When we return to
main, the value of `x` is unchanged. When we call update, we copy the address of `x` into
`px` again. However, this time we change the value of what `px` in update points to, the
variable `x` in main. When we return to main, `x` has been changed.

### Pointer Passing Performance

The time to pass a pointer
into a function is constant for all data sizes, roughly one nanosecond. This makes
sense, as the size of a pointer is the same for all data types. It takes about a millisecond once the value
gets to be around 10 megabytes of data.

The behavior for returning a pointer versus returning a value is more interesting. For
data structures that are smaller than a megabyte, it is actually slower to return a
pointer type than a value type. For example, a 100-byte data structure takes around 10
nanoseconds to be returned, but a pointer to that data structure takes about 30 nano‐
seconds. Once your data structures are larger than a megabyte, the performance
advantage flips. It takes nearly 2 milliseconds to return 10 megabytes of data, but a
little more than half a millisecond to returnq a pointer to it.

### The Difference Between Maps and Slices

You should avoid using maps for input parameters or return values,
especially on public APIs. You should avoid using maps for input parameters or return values,
especially on public APIs.

Changing the values in the slice changes the memory that the pointer points to, so the
changes are seen in both the copy and the original. Changes to the length and capacity are not reflected back in the original, because they
are only in the copy. Changing the capacity means that the pointer is now pointing to
a new, bigger block of memory.

The reason you can pass a slice of any size to a function is that the
data that’s passed to the function is the same for any size slice: two
int values and a pointer. The reason that you can’t write a function
that takes an array of any size is because the entire array is passed
to the function, not just a pointer to the data.

## Types, Methods, and Interfaces

Go allows you to declare a type at any block level, from the package block down.

```go
type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```

### Methods
```go
type Person struct {
  FirstName string
  LastName string
  Age int
}
func (p Person) String() string {
  return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```

Method declarations look just like function declarations, with one addition: the
receiver specification. The receiver appears between the keyword func and the name
of the method. The receiver name `p` appears before the type `Person`.

### Pointer Receivers and Value Receivers

* If your method modifies the receiver or needs to handle nil instances, you must use a pointer receiver.
* If your method doesn’t modify the receiver, you can use a value receiver
```go
type Counter struct {
  total int
  lastUpdated time.Time
}
func (c *Counter) Increment() {
  c.total++
  c.lastUpdated = time.Now()
}
func (c Counter) String() string {
  return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
```
The rules for passing values to functions still apply. If you pass
a value type to a function and call a pointer receiver method on the passed value, you
are invoking the method on a copy.
```go
func doUpdateWrong(c Counter) {
  c.Increment()
  fmt.Println("in doUpdateWrong:", c.String())
}
func doUpdateRight(c *Counter) {
  c.Increment()
  fmt.Println("in doUpdateRight:", c.String())
}
func main() {
  var c Counter
  doUpdateWrong(c)
  fmt.Println("in main:", c.String())
  doUpdateRight(&c)
  fmt.Println("in main:", c.String())
}
```
### Methods Are Functions Too

```go
type Adder struct {
  start int
}
func (a Adder) AddTo(val int) int {
  return a.start + val
}
```

* We create an instance of the type in the usual way and invoke its method:
```go
myAdder := Adder{start: 10}
fmt.Println(myAdder.AddTo(5)) // prints 15
```

* We can also assign the method to a variable or pass it to a parameter of type
`func(int)int`. This is called a method value:
```go
f1 := myAdder.AddTo
fmt.Println(f1(10))
```

* You can also create a function from the type itself. This is called a method expression:
```go
f2 := Adder.AddTo
fmt.Println(f2(myAdder, 15)) // prints 25
```

In the case of a method expression, the first parameter is the receiver for the method;
our function signature is func(Adder, int) int.

### iota

his means that it
assigns 0 to the first constant , 1 to the second constant ,
and so on. When a new const block is created, iota is set back to 0.

```go
type BitField int
const (
  Field1 BitField = 1 << iota // assigned 1
  Field2 // assigned 2
  Field3 // assigned 4
  Field4 // assigned 8
)
```

### Use Embedding for Composition

```go
type Employee struct {
  Name string
  ID string
}
func (e Employee) Description() string {
  return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}
type Manager struct {
  Employee
  Reports []Employee
  }
func (m Manager) FindNewEmployees() []Employee {
// do business logic
}

m := Manager{
  Employee: Employee{
  Name: "Bob Bobson",
  ID: "12345",
  },
  Reports: []Employee{},
}
  fmt.Println(m.ID) // prints 12345
  fmt.Println(m.Description()) // prints Bob Bobson (12345)
```

### Embedding Is Not Inheritance

```go
var eFail Employee = m // compilation error!
var eOK Employee = m.Employee // ok!
```

While embedding one concrete type inside another won’t allow you to treat the outer
type as the inner type, the methods on an embedded field do count toward the
method set of the containing struct. This means they can make the containing struct
implement an interface
```go
type Inner struct {
  A int
}
func (i Inner) IntPrinter(val int) string {
  return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) Double() string {
  return i.IntPrinter(i.A * 2)
}
type Outer struct {
  Inner
  S string
}
func (o Outer) IntPrinter(val int) string {
  return fmt.Sprintf("Outer: %d", val)
}
func main() {
    o := Outer{
    Inner: Inner{
    A: 10,
    },
    S: "Hello",
  }
  fmt.Println(o.Double())
}
Running this code produces the output:
Inner: 20
```

### Interface

Interfaces specify what callers need. The client code defines the
interface to specify what functionality it requires.

```go
type LogicProvider struct {}
func (lp LogicProvider) Process(data string) string {
  // business logic
}
type Logic interface {
  Process(data string) string
}
type Client struct{
  L Logic
}
func(c Client) Program() {
  // get data from somewhere
  c.L.Process(data)
}
main() {
  c := Client{
  L: LogicProvider{},
  } 
  c.Program()
}
```

Interfaces are named collections of method signatures.

You have a square and a circle, and want to get the area for each.

```go
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
```


```go
// Now you need functions to display the area and perim. You might write code like this.
func (r rect) area() float64 {
  return r.width * r.height
}
func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}
```
Both functions do exactly the same thing. 
Because Go is strongly typed we need to have two functions. 
One to display the area of rect, and one to display the area of circles.
We can write a single function that takes both.
```go
type geometry interface {
  area() float64
  perim() float64
}
```

Now we can throw away our two functions and write a single function that takes an instance of 
`geometry`.


```go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    measure(r) // {3 4} 12 14
    measure(c) // 5 78.53981633974483 31.41592653589793
}
```

### Accept Interfaces, Return Structs

You’ll often hear experienced Go developers say that your code should “Accept inter‐
faces, return structs.” What this means is that the business logic invoked by your
functions should be invoked via interfaces, but the output of your functions should
be a concrete type.

### The Empty Interface Says Nothing

Sometimes in a statically typed language, you need a way to say that a variable could
store a value of any type. Go uses interface{} to represent this:
```go
  var i interface{}
  i = 20
  i = "hello"
  i = struct {
  FirstName string
  LastName string
  } {"Fred", "Fredson"}
```

You should note that interface{} isn’t special case syntax. An empty interface type
simply states that the variable can store any value whose type implements zero or
more methods. This just happens to match every type in Go. Because an empty inter‐
face doesn’t tell you anything about the value it represents, there isn’t a lot you can do
with it. One common use of the empty interface is as a placeholder for data of uncer‐
tain schema that’s read from an external source, like a JSON file:
```go
  // one set of braces for the interface{} type,
  // the other to instantiate an instance of the map
  data := map[string]interface{}{}
  contents, err := ioutil.ReadFile("testdata/sample.json")
  if err != nil {
  return err
  }
  defer contents.Close()
  json.Unmarshal(contents, &data)
// the contents are now in the data map
```
Another use of interface{} is as a way to store a value in a user-created data structure. 
```go
type LinkedList struct {
  Value interface{}
  Next *LinkedList
}
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
    if ll == nil || pos == 0 {
    return &LinkedList{
      Value: val,
      Next: ll,
    }
  }
  ll.Next = ll.Next.Insert(pos-1, val)
  return ll
}
```

### Type Switches

```go
func doThings(i interface{}) {
 switch j := i.(type) {
 case nil:
 // i is nil, type of j is interface{}
 case int:
 // j is of type int
 case MyInt:
 // j is of type MyInt
 case io.Reader:
 // j is of type io.Reader
 case string:
 // j is a string
 case bool, rune:
 // i is either a bool or rune, so j is of type interface{}
 default:
 // no idea what i is, so j is of type interface{}
 }
}
```

## Errors

When a function executes
as expected, `nil` is returned for the error parameter. If something goes wrong, an
error value is returned instead. 

```go
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
 if denominator == 0 {
   return 0, 0, errors.New("denominator is 0")
 }
 return numerator / denominator, numerator % denominator, nil
}
```

Unlike languages with exceptions, Go doesn’t have special constructs to detect if an
error was returned. Whenever a function returns, use an if statement to check the
error variable to see if it is non-nil:

```go
func main() {
 numerator := 20
 denominator := 3
 remainder, mod, err := calcRemainderAndMod(numerator, denominator)
 if err != nil {
   fmt.Println(err)
   os.Exit(1)
 }
 fmt.Println(remainder, mod)
}
// error is a built-in interface that defines a single method:
type error interface {
 Error() string
}
```

### Use Strings for Simple Errors

Go’s standard library provides two ways to create an error from a string. `errors.New()` 
and `fmt.Errorf()`

```go
func doubleEven(i int) (int, error) {
 if i % 2 != 0 {
   return 0, errors.New("only even numbers are processed")
 }
 return i * 2, nil
}

func doubleEven(i int) (int, error) {
if i % 2 != 0 {
  return 0, fmt.Errorf("%d isn't an even number", i)
}
return i * 2, nil
}

```

## Modules, Packages, and Imports

A module is the root of a Go
library or application, stored in a repository. Modules consist of one or more pack‐
ages, which give the module organization and structure.

Package names should be descriptive. Rather
than have a package called `util`, create a package name that describes 
the functionality provided by the package. For example, say you have two helper 
functions: one to
extract all names from a string and another to format names properly. Don’t create
two functions in a `util` package called `ExtractNames` and `FormatNames`. If you do,
every time you use these functions, they will be referred to as `util.ExtractNames`
and `util.FormatNames`, and that util package tells you nothing about what the func‐
tions do.

It’s better to create one function called `Names` in a package called `extract` 
and a second function called `Names` in a package called `format`. 
It’s OK for these two functions
to have the same name, because they will always be disambiguated by their package
names. The first will be referred to as `extract.Names` when imported, and the second
will be referred to as `format.Names`.

# References

* [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://www.amazon.de/-/en/Jon-Bodner/dp/1492077216)