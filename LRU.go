package main


import "fmt"

// define struct for LRU cache
// index maps string to Node pointer
// define max size of Cache
// 
type LRUCache struct {
	index            map[string]*DoublyLinkedListNode
	maxSize          int
	currentSize      int
	listOfMostRecent *DoublyLinkedList
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		index:            map[string]*DoublyLinkedListNode{},
		maxSize:          size,
		currentSize:      0,
		listOfMostRecent: &DoublyLinkedList{},
	}
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	if _, found := cache.index[key]; found {
		cache.replaceKey(key, value)
	} else {
		if cache.currentSize == cache.maxSize {
			cache.evictLeastRecent()
		} else {
			cache.currentSize += 1
		}
		cache.index[key] = &DoublyLinkedListNode{
			key:   key,
			value: value,
		}
	}
	cache.updateMostRecent(cache.index[key])
}

func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	if node, found := cache.index[key]; !found {
		return 0, false
	} else {
		cache.updateMostRecent(node)
		return node.value, true
	}
}

func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if cache.listOfMostRecent.head == nil {
		return "", false
	}
	return cache.listOfMostRecent.head.key, true
}

func (cache *LRUCache) evictLeastRecent() {
	key := cache.listOfMostRecent.tail.key
	cache.listOfMostRecent.removeTail()
	delete(cache.index, key)
}

func (cache *LRUCache) updateMostRecent(node *DoublyLinkedListNode) {
	cache.listOfMostRecent.setHeadTo(node)
}

func (cache *LRUCache) replaceKey(key string, value int) {
	if node, found := cache.index[key]; !found {
		panic("The provided key isn't in the cache!")
	} else {
		node.value = value
	}
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (list *DoublyLinkedList) setHeadTo(node *DoublyLinkedListNode) {
	if list.head == node {
		return
	}
	if list.head == nil {
		list.head, list.tail = node, node
		return
	}
	if list.head == list.tail {
		list.tail.prev = node
		list.head = node
		list.head.next = list.tail
		return
	}
	if list.tail == node {
		list.removeTail()
	}
	list.head.prev = node
	node.next = list.head
	list.head = node
}

func (list *DoublyLinkedList) removeTail() {
	if list.tail == nil {
		return
	}
	if list.tail == list.head {
		list.head, list.tail = nil, nil
		return
	}
	list.tail = list.tail.prev
	list.tail.next.removeBindings()
}

type DoublyLinkedListNode struct {
	key   string
	value int
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

func (node *DoublyLinkedListNode) removeBindings() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev, node.next = nil, nil
}


func main(){
	cache := NewLRUCache(5)
	fmt.Println("LRU cache max size = 5 current length: ", cache.currentSize)
	fmt.Println("Insert value: 1")
	cache.InsertKeyValuePair("1",1)

	fmt.Println("Insert value: 2")
	cache.InsertKeyValuePair("2",2)

	fmt.Println("Insert value: 3")
	cache.InsertKeyValuePair("3",3)

	fmt.Println("Insert value: 4")
	cache.InsertKeyValuePair("4",4)

	fmt.Println("Insert value: 5")
	cache.InsertKeyValuePair("5",5)

	fmt.Println("Inserting value 6 when cache is full")

	cache.InsertKeyValuePair("6", 6 )

	fmt.Println("Get Most recent key:")
	
	fmt.Println(cache.GetMostRecentKey())

	fmt.Println("Current size: ")
	fmt.Println(cache.currentSize)

}