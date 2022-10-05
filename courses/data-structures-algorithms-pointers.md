# Data Structures, Algorithms and Pointers in Go

## Golang Tutorials
This chapter is based on [Golang Tutorial 2 - Learn how go works with the hello world code by Junmin Lee](https://www.youtube.com/watch?v=7PMEbo5Ryeg&t=129s)

C++, Go, Java, Python: High Level Languages, which means human readable

Human -> Machine: Compile (Convert human to machine readable code)

For example, when you finishing a program, **compiler** will transform it 
to a executable file (file.exe)

A package is a collection of source files in the same directory that are compiled together.

A module is a collection of related go packages that are released together
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


### Binary Search Tree

Root -> Parent -> Children -> Leaf

Left Children smaller, Right Children bigger

**Benefits:**
* Speed
* O(logn)
* Worst case still 0(n)

#### Implement

```go
package main

import "fmt"

var count int // Count how many nodes travel

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and Return Trueif there is a node with that value

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		n.Right.Search(k)

	} else if n.Key > k {
		n.Left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(50)
	tree.Insert(202)
	tree.Insert(50)
	tree.Insert(24120)
	tree.Insert(525)
	tree.Insert(15)
	tree.Insert(5213)
	tree.Insert(230)
	tree.Insert(5)
	tree.Insert(8)

	fmt.Println(tree)
	fmt.Println(tree.Search(5213))
	fmt.Println(count)
}
```

### Tries

Trees that store words.

Each nodes hold the array of 26 characters.

Time complexity: 0(m), where m is the length of the words/

Return false immediately if the words not in the tree 
(the first character of the words not in the first parents)

Tries is trade off between Time and Space.

```go
package main

import "fmt"

// AlphabetSize is the numberof possible characters in the trie

const AlphabetSize = 26

type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	results := &Trie{root: &TrieNode{}}
	return results
}

// Insert will take in a word and add it to the tries
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // 'a' = 97. 'b' - 'a' = 1 => a 0 b 1 c 2 ...
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and return TRUE if that word is include in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // 'a' = 97. 'b' - 'a' = 1 => a 0 b 1 c 2 ...
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false // Even will match the word, but the word can continue => not the word we want
}

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie)      //&{0xc000138000}
	fmt.Println(testTrie.root) // &{[<nil> x26 ] false}
	toAdd := []string{
		"test",
		"next",
		"best",
		"aragon",
		"oregon",
	}

	for _, v := range toAdd {
		testTrie.Insert(v)
	}
	fmt.Println(testTrie.Search("aragon"))
}
```

### Hash Tables

#### Introduction
Search for the name in the array 
=> You have to go through all the array to search for it.

But if you know the index (**hash code**) of the name, for example, 82 for Andy 
=> You can immediately access the name.
 
When Inserting, but Andy into the **Hash function**, get the hash code, and go to
that index to find Andy.

#### Simple Hash Algorithm

RANDY => 82 + 65 + 78 + 68 + 89 (ASCII Code) = 382 

382 % 100 = 82.

Choose 100 because is the size of the array we want to store. So the index
always between 0 and 99

Eric -> 91.

When we search for STAN, we put into the hash function => get the index and find it
in the array

#### Collision Handling

Two name have the same hash code. There are two ways for 
collision handling

First is **Open Addressing**. We store Andy in index 4 and Eric in index 5.
When we search for Eric, we go to the original location, which is 4. 
If 4 is not Eric, we go to the next index, which is 5. This method still faster than
searching each element in the array.

Drawbacks: more and more name in the same address => loose benefit of the hash table.

Second is **Separate Chaining**: Storing multiple names in one slot by using 
Linked List. Each index will hold a pointer point to the head of a linked list.
That has a list of name. The linked list will be called bucket.

#### Insert/Delete/Search
Hash table has the best of array and linked list.

Best case: 0(1) for Insert Delete and Search

Worst Case: 0(n) like a linked list when every element in the same linked list.

#### Implement


## Pointers



### Algorithms

[Let's Learn Algorithms by Jon Calhoun](https://www.calhoun.io/lets-learn-algorithms)


