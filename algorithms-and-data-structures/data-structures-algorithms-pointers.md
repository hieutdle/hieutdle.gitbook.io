# Data Structures in Go

## Golang Tutorials

This chapter is based on [Golang Tutorial 2 - Learn how go works with the hello world code by Junmin Lee](https://www.youtube.com/watch?v=7PMEbo5Ryeg\&t=129s)

C++, Go, Java, Python: High Level Languages, which means human readable

Human -> Machine: Compile (Convert human to machine readable code)

For example, when you finishing a program, **compiler** will transform it to a executable file (file.exe)

A package is a collection of source files in the same directory that are compiled together.

A module is a collection of related go packages that are released together

## Data Structures

This chapter is based on [Data Structures in Golang by Junmin Lee](https://www.youtube.com/playlist?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6)

[**You can find all the code of this chapter here**](https://github.com/ledinhtrunghieu/learning-go)

### Linked List

Linked List put the values in nodes. Nodes are linked to each others by containing address of the next nodes.

**Benefits:**

* Adding or removing values at the beginning of the list: **O(1)**. Array need to shift when adding new values.

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

Return false immediately if the words not in the tree (the first character of the words not in the first parents)

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

Search for the name in the array => You have to go through all the array to search for it.

But if you know the index (**hash code**) of the name, for example, 82 for Andy => You can immediately access the name.

When Inserting, but Andy into the **Hash function**, get the hash code, and go to that index to find Andy.

#### Simple Hash Algorithm

RANDY => 82 + 65 + 78 + 68 + 89 (ASCII Code) = 382

382 % 100 = 82.

Choose 100 because is the size of the array we want to store. So the index always between 0 and 99

Eric -> 91.

When we search for STAN, we put into the hash function => get the index and find it in the array

#### Collision Handling

Two name have the same hash code. There are two ways for collision handling

First is **Open Addressing**. We store Andy in index 4 and Eric in index 5. When we search for Eric, we go to the original location, which is 4. If 4 is not Eric, we go to the next index, which is 5. This method still faster than searching each element in the array.

Drawbacks: more and more name in the same address => loose benefit of the hash table.

Second is **Separate Chaining**: Storing multiple names in one slot by using Linked List. Each index will hold a pointer point to the head of a linked list. That has a list of name. The linked list will be called bucket.

#### Insert/Delete/Search

Hash table has the best of array and linked list.

Best case: 0(1) for Insert Delete and Search

Worst Case: 0(n) like a linked list when every element in the same linked list.

#### Implement

```go
package main

import (
	"fmt"
)

// ArraySize is the size of the hash table array

const ArraySize = 7

// HashTable will hold an array

type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the array

type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and add it to the hash table array

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Init will create a bucket in each slot of the hash table

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// insert will take in a key, create a node with a key
// and insert the node in the bucket

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search will take in a key and return true if the bucket has that key

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete the node from the bucket

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	testHashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		testHashTable.Insert(v)
	}

	testHashTable.Delete("STAN")
	fmt.Println(testHashTable.Search("STAN"))
	fmt.Println(testHashTable.Search("KYLE"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	testBucket.insert("RANDY")
	testBucket.delete("RANDY")
	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("ERIC"))
}
```

### Heap

The parent node is larger than a child.

* parent index \* 2 + 1 = left child index.
* parent index \* 2 + 2 = right child index

Very fast when getting the biggest or lowest value (highest node).

#### Insert

When ever insert: add the node to the bottom right of the tree.

#### Heapify

We need to re-arrange the tree by swapping the parent and child nodes if the child node is larger than the parent node. The process of rearranging the indices as **Heapify**.

#### Extract

Extract means remove the highest key of the tree. Right after taking out the highest nodes, we will take the last node of the tree to the root position. Then we swap with its larger child.

#### Time complexity

Heapify up or down is depending on the height of the tree => 0(h) (extract or insert)

If you want to replace to it the number of elements of the array => O(logn) because the height and the number of indices have a logarithmic relation.

#### Implement

```go
package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap.
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Print("cannot extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will be heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	// loop while index at least one child
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			// when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			//when left child is larger
			childToCompare = l
		} else {
			// when right child is larger
			childToCompare = r
		}

		// compare array value of current index
		// to larger child and swap if smaller

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			// it means it find the right place
			return
		}

	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 2
}

// get the right child index
func right(i int) int {
	return 2*i + 1
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
```

## Pointers

\*:

* `*int`: this whole thing is a type(pointer type, `int` is a base)
* `*p`: \* in front of variable => operator returns what p is pointing to (print value the pointer pointing to - dereferencing)
* value of p is now the address of the value

```go
package main

import "fmt"

func main() {
 i, j := 42, 2701

 p := &i
 fmt.Println(*p) // 42
 fmt.Printf("%T\n", p)
 *p = 21        // change value of i, still keep the address of i
 fmt.Println(i) // print 21

 p = &j         // now point to j
 *p = *p / 37   // change value of j, still keep the address of j
 fmt.Println(j) // 73
}
```

Why you pointer? Because you can access the variable through pointer from different part of the program rather than copy it each time you use or want to manipulate.

Goroutines: independent path of execution - stack of memory

```go
func main(){
	a :=4
	squareVal(a)
}

func squareVal(v int)  {
    v *= v
	fmt.Println(&v,v)
}
```

![](../images/images-courses/pointer.png)

They are isolated frames.

After squareVal create a copy of `a` => `v` with value of 16, when the active frame return to main, `a` still has the value of 4.

```go
func main(){
    a :=4
    squareAdd(&a)
}

func squareAdd(p *int){
	*p *= *p
	fmt.Println(p,*p)
}
```

![](../images/images-courses/pointer2.png)

#### return `m` (value) and `&m` (pointer)

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() person {
	m:= person{name:"noname", age:50}
	return m
}

func main()  {
	fmt.Println(initPerson())
}
```

When we call initPerson, we create `m`. Then you change the value of `m`, because of the isolation characteristic. We can not send `m` to the `main()` function. Instead, we make a copy of `m`.

![](../images/images-courses/pointer3.png)

Let's return the address of `m`

```go
func initPerson() *person {
       m:= person{name:"noname", age:50}
	   fmt.Printf("initPerson -> %p\n",&m)
       return &m
}

func main()  {
	fmt.Println(initPerson())
    fmt.Printf("main -> %p\n",initPerson())
}
```

![](../images/images-courses/pointer4.png)

We have an address pointing to `m`, but when the `initPerson()` function finish. That frame is become invalid so the address we copied into the active frame is useless. That where heaps come in so heaps is going to solve this problem for us. The name heaps is different from data structure heaps.

The compiler will analyze what's going on and figures out that this may cause the problem and copy `m` to the heap. Then the `initPerson()` function will return the address of `m` in the heap. After the return when the address of `m` is copied to the frame of the `main()` function. We would be able to access 'm\` with that address.

![](../images/images-courses/pointer5.png)

#### Garbage Collector

We are doing this in the cost of heap allocation. Which can be a burden for the garbage collector and it can cost us performance.

Stacks don't need garbage collector because it is self-cleaning. When function is called and finish, it will discards the frame and everything inside it. When another function is called, the space will be used by other frames.

If we put something in the heap, that will create job for the garbage collector.

There is a specific algorithms for the garbage collector automatically sets the memory free for ones that we don't use and just keep the ones that we need.

![](../images/images-courses/pointer6.png)

### Algorithms

[Let's Learn Algorithms by Jon Calhoun](https://www.calhoun.io/lets-learn-algorithms)
