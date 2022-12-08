---
coverY: 0
---

# ☯ gods

This post is inspired by Problem Solving with Algorithms 
and Data Structures using Python and C++

## 1. Basic Data Structures

### Stack

**LIFO, Last-in first-out**

![Stack](images/gods/stack.png)

#### Operations
* `Size()` returns the number of items on the stack.
* `Push(item)` adds a new item to the top of the stack.
* `Pop()` removes the top item from the stack and returns it.
* `Peek()` returns the top item from the stack.
* `IsEmpty()` tests to see whether the stack is empty and returns a boolean value.

#### Implementation

```go
type Stack struct {
    items []int
}

// Size() returns the number of items on the stack.
func (s *Stack) Size() int {
    return len(s.items)
}

// Push(item) adds a new item to the top of the stack.
func (s *Stack) Push(i int) {
    s.items = append(s.items, i)
}

// Pop() will remove a value at the end and return it.
func (s *Stack) Pop() int {
    l := s.Size()
    toRemove := s.items[l-1]
    s.items = s.items[:l-1]
    return toRemove
}

// Peek() returns the top item from the stack.
func (s *Stack) Peek() int {
    l := s.Size()
    return s.items[l-1]
}

// IsEmpty() tests to see whether the stack is empty and returns a boolean value.
func (s *Stack) IsEmpty() bool {
    return s.Size() == 0
}
```

#### Decimal to any Base
![Decimal-to-Binary Conversion](images/gods/dectobin.png)

```go
func baseConverter(decNumber int, base int) string {
	remStack := Stack{}
	for decNumber > 0 {
		rem := decNumber % base
		remStack.Push(rem)
		decNumber = decNumber / base
	}
	newString := ""
	for !(remStack.IsEmpty()) {
		newString = newString + strconv.Itoa(remStack.Pop())
	}
	return newString
}
```

### Queue

