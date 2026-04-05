package main

import "fmt"

// Node represents a doubly linked list node
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache structure
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // dummy head (MRU side)
	tail     *Node // dummy tail (LRU side)
}

// Constructor initializes the cache
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Get returns value and marks it as recently used
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

// Put inserts/updates key-value pair
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// Update value and move to front
		node.value = value
		this.moveToFront(node)
		return
	}

	// Create new node
	newNode := &Node{
		key:   key,
		value: value,
	}

	this.cache[key] = newNode
	this.addToFront(newNode)

	// Evict if over capacity
	if len(this.cache) > this.capacity {
		lru := this.removeLRU()
		delete(this.cache, lru.key)
	}
}

// --------------------
// Internal Helpers
// --------------------

// addToFront adds node right after head (MRU position)
func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

// remove removes a node from the list
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// moveToFront moves a node to MRU position
func (this *LRUCache) moveToFront(node *Node) {
	this.remove(node)
	this.addToFront(node)
}

// removeLRU removes the least recently used node (before tail)
func (this *LRUCache) removeLRU() *Node {
	lru := this.tail.prev
	this.remove(lru)
	return lru
}

// --------------------
// Testing
// --------------------

func lrucache() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1)) // 1

	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // -1

	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}
