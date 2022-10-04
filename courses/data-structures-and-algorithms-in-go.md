# Data Structures, Algorithms and Pointers in Go

## Data Structures 

This chapter is based on [Data Structures in Golang by Junmin Lee](https://www.youtube.com/playlist?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6)

**[You can find all the code of this chapter here](https://github.com/ledinhtrunghieu/learning-go)**

### Linked List

Linked List put the values in nodes. Nodes are linked to each others by containing address
of the next nodes.

**Benefits:**
* Adding or removing values at the beginning of the list: **O(1)**. 
 Array need to shift when adding new values.


**Disadvantages**
* Travel and replace the value of the node: **0(n)**, arrays cost just **O(1)**


#### Doubly linkedin list
Contains the address of the next and also the previous node

#### Implement

```go
package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
func main() {
	mylist := linkedList{}
	node1 := &node{data: 6}
	node2 := &node{data: 5}
	node3 := &node{data: 4}
	node12 := &node{data: 3}
	node22 := &node{data: 2}
	node32 := &node{data: 1}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node12)
	mylist.prepend(node22)
	mylist.prepend(node32)
	fmt.Println(mylist)
	mylist.printListData()
	mylist.deleteWithValue(6)
	mylist.deleteWithValue(7)
	mylist.deleteWithValue(1)
	mylist.printListData()
	emptylist := linkedList{}
	emptylist.deleteWithValue(10)
}
```

### Stacks
* Last in first out (**LIFO**)
* Add: **Push**. Remove: **Pop**


#### Implement

```go
package main

import "fmt"

type Stack struct {
	items []int
}

// Push will add value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// End return the removed value
func (s *Stack) Pop() int {

	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
```

### Queues
* First in first out
* Add : **Enqueue**. Remove: **Dequeue**

#### Implement

```go
package main

import "fmt"

type Queue struct {
	items []int
}

// Push will add value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Pop will remove a value at the end
// End return the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
```
## Pointers



### Algorithms

[Let's Learn Algorithms by Jon Calhoun](https://www.calhoun.io/lets-learn-algorithms)


