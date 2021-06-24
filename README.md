# LRU Cache Implementation

A LRU (Least Recently Used) cache is a type of caching system that allows the most recently used elements to be tracked and access quickly.
The same is true for the least frequenlty used element, allowing us to remove the least important element on the system , in order to make room for new elements.

LRU caching consists of 2 primary data structures Hash Table, LinkedList

# Hash table
The hash table allows us to hash a string to a node, the node has 4 properties
1) Key
2) Value
3) Prev *Node <- pointer to previous node
4) Next *Node <- pointer to next node
# Linked List
Linked List allows us to store our values in an efficient structure,
the list consists of a Head , Tail *Node <- both pointers to the respective nodes.

Both of these structures make up a LRUCache

# Cache
To create an instance of the Cach use 
cache := NewCache(size int)

The var cache can then be used to insert key(string) Value(int) pairs such as
cache.InsertKeyValuePair("aapl",255)

cache.GetValueFromKey("aapl") can then be used to return the value 255.

Everytime a value gets read or modified, the cache.ListOfMostRecent *LinkedList <- (pointer to list) property gets updated and the node is set as the head of the list.

This same list is used to determine the least recently used, being the tail of the list.



