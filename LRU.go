package lru

import "errors"

/* This package contains definition of LRU (Least Recently Used) Cache, DoublyLinkedList, and DoublyLinkedListNode
Functions for manipulation LRU include:
ctor: NewLRUCache(size int): LRUCache, essentially a constructor to initialize a new LRU with a max size.
*/

// Cache is a struct that has fields index(maps a string to a DoublyLinkedListNode)
// index:  is a map, key(string) value Node,
// maxSize int:  is the max size of cache before it starts to evict node on LRU principles.
// currentSize int : holds current size of cache, needed to be able to know if we need to call evictLeastRecent().
// listOfMostRecent &DoublyLinkedList: is the most recently used LinkedList, important to keep track of for implementiong LRU
type Cache struct {
	index            map[string]*DoublyLinkedListNode
	maxSize          int
	currentSize      int
	listOfMostRecent *DoublyLinkedList
}

// NewCache is a constructor/ initializer , takes an int for maxSize to initialize Cache with.
func NewCache(size int) *Cache {
	return &Cache{
		index:            map[string]*DoublyLinkedListNode{},
		maxSize:          size,
		currentSize:      0,
		listOfMostRecent: &DoublyLinkedList{},
	}
}

// InsertKeyValuePair takes in a key(string) and value(int) and inserts the pair into the cache, updating MostRecent()
func (cache *Cache) InsertKeyValuePair(key string, value int) {
	// _ is for throwaway var (node that we dont need in this fn())
	_, found := cache.index[key]

	if found {
		cache.ReplaceKey(key, value)
	} else {
		if cache.currentSize == cache.maxSize {
			cache.EvictLeastRecent()
		} else {
			cache.currentSize++
		}
		cache.index[key] = &DoublyLinkedListNode{
			key:   key,
			value: value,
		}
	}
	cache.UpdateMostRecent(cache.index[key])
}

// GetValueFromKey takes in a key(string) if found returns value,true and updatesMostRecent() ,if !found returns 0, false
func (cache *Cache) GetValueFromKey(key string) (int, bool) {
	node, found := cache.index[key]
	if !found {
		return 0, false
	}
	cache.UpdateMostRecent(node)
	return node.value, true
}

// GetMostRecentKey returns they key(string), true , if Cache is empty, will return empty string "", false
func (cache *Cache) GetMostRecentKey() (string, bool) {
	if cache.listOfMostRecent.head == nil {
		return "", false
	}
	return cache.listOfMostRecent.head.key, true
}

// EvictLeastRecent removes cache.list of mostRecent tail (LeastRecentlyUsed) and removed the index key
func (cache *Cache) EvictLeastRecent() error {
	if cache.currentSize <1 {
		return errors.New("Empty cache cannot EvictLeastRecent")
	} else{
		key := cache.listOfMostRecent.tail.key
		cache.listOfMostRecent.RemoveTail()
		delete(cache.index, key)
		cache.currentSize --
		return nil
	}
}

// UpdateMostRecent updates the listOfMostRecent.head(MostRecent) to input node
func (cache *Cache) UpdateMostRecent(node *DoublyLinkedListNode) {
	cache.listOfMostRecent.SetHeadTo(node)
}

// ReplaceKey takes in a key(string) value(int) , if key !exists : throw error, else update value with key
func (cache *Cache) ReplaceKey(key string, value int) (bool, error) {
	node, found := cache.index[key]
	if !found {
		return false, errors.New("Key not found in cache")
	}
	node.value = value
	return true, nil

}

// DoublyLinkedList defines struct which has 2 fields, head and tail of type *DoublyLinkedListNode
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

// SetHeadTo takes input node and makes it head of LinkedList
func (list *DoublyLinkedList) SetHeadTo(node *DoublyLinkedListNode) {
	// checks if current head already is input node, does nothing
	if list.head == node {
		return
	}
	// if the list head is nil , list is empty and sets head and tail to input node
	if list.head == nil {
		list.head, list.tail = node, node
		return
	}
	// if list only contains 1 element, list.tail.prev is set to input node
	// list head set to input node
	// links.list.head next to list.tail
	if list.head == list.tail {
		list.tail.prev = node
		list.head = node
		list.head.next = list.tail
		return
	}
	// if the tail is the input node
	// remove the current tail
	// list.head.prev is set to input node
	// input node.next is set to list.head
	// head is set to input node
	if list.tail == node {
		list.RemoveTail()
	}
	list.head.prev = node
	node.next = list.head
	list.head = node
}

// RemoveTail will remove current tail and sets tail.prev.next = nil
func (list *DoublyLinkedList) RemoveTail() {
	// if current tail is nil , do nothing
	if list.tail == nil {
		return
	}
	// if list only has 1 element, set list.head and list.tail = nil
	if list.tail == list.head {
		list.head, list.tail = nil, nil
		return
	}
	// set tail to current tail's previous node
	list.tail = list.tail.prev
	// remove pointers from previous tail
	list.tail.next.RemoveBindings()
}

// DoublyLinkedListNode structure has key(string), value(int), prev,next *Node (pointer to previous and next nodes)
type DoublyLinkedListNode struct {
	key   string
	value int
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

// RemoveBindings removes pointers(prev, next) from calling node
func (node *DoublyLinkedListNode) RemoveBindings() {
	// check if node is != head
	if node.prev != nil {
		node.prev.next = node.next
	}
	// check if node is !tail
	if node.next != nil {
		node.next.prev = node.prev
	}
	// remove prev and next pointers
	node.prev, node.next = nil, nil
}
