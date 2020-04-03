package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestNewCache(t *testing.T) {
	// arrange
	maxSize := 10
	c := NewCache(maxSize)
	// act
	result := c.maxSize
	// assert
	assert.Equal(t, result, 10, "maxSize Expected: 10")
}
func TestCache_InsertKeyValuePair(t *testing.T) {
	// arrange
	c := NewCache(10)
	newKey := "aapl"
	newValue := 1
	// act
	c.InsertKeyValuePair(newKey,newValue)
	result :=  c.index[newKey]
	resultKey , resultValue := result.key, result.value
	//assert
	assert.Equal(t, resultKey, newKey, "input key should be == resultKey")
	assert.Equal(t, resultValue, newValue, "input value should == resultValue")
}
func TestCache_InsertKeyValuePairInsertSamePairThreeTimes(t *testing.T) {
	// arrange
	c := NewCache(10)
	k := "aapl"
	v := 1
	// act
	c.InsertKeyValuePair(k,v)
	c.InsertKeyValuePair(k,v)
	c.InsertKeyValuePair(k,v)
	//assert
	assert.Equal(t, 1 , c.currentSize, "CurrentSize should still == 1")
}
func TestCache_GetValueFromKey(t *testing.T) {
	// arrange
	c := NewCache(5)
	// act
	c.InsertKeyValuePair("aaple",1)
	c.InsertKeyValuePair("msft",2)
	c.InsertKeyValuePair("amd", 3)

	resultInt, resultBool := c.GetValueFromKey("msft")
	//assert
	assert.Equal(t,2 ,resultInt, "Should return value from key 'msft' , Expected: 2")
	assert.Equal(t, true, resultBool, "Should return bool true, since key exists")

}
func TestCache_GetValueFromKeyWhereKeyDoesntExist(t *testing.T) {
	// arrange
	c := NewCache(5)
	//act
	c.InsertKeyValuePair("aaple",1)
	c.InsertKeyValuePair("msft",2)
	c.InsertKeyValuePair("sq",3)
	resultInt, resultBool := c.GetValueFromKey("roku")
	// assert
	assert.Equal(t,false,resultBool, "Key doesnt exist; expected: false")
	assert.Equal(t,0, resultInt,"key doesnt exist; expected int: 0")

}
func TestCache_GetMostRecentKeyEmptyCache(t *testing.T) {
	// arrange
	c := NewCache(5)
	// act
	resultString, resultBool := c.GetMostRecentKey()
	// assert
	assert.Equal(t,"", resultString, "empty cache should return '' empty string ")
	assert.Equal(t, false, resultBool, "empty cache should return false")
}
func TestCache_GetMostRecentKey(t *testing.T) {
	// arrange
	c := NewCache(5)
	//act
	c.InsertKeyValuePair("aapl",1)
	c.InsertKeyValuePair("msft",2)
	c.InsertKeyValuePair("sq",3)
	c.InsertKeyValuePair("amd",4)
	resultValue , _  := c.GetValueFromKey("aapl")
	expectedValue := 1
	//assert
	assert.Equal(t,expectedValue,resultValue,"After several inserts, after calling GetValueFromKey('aapl') should update to MRU")
}
func TestCache_EvictLeastRecent(t *testing.T) {
	// arrange
	c := NewCache(5)
	c.InsertKeyValuePair("aapl",1)
	c.InsertKeyValuePair("msft",2)
	c.InsertKeyValuePair("sq",3)
	c.InsertKeyValuePair("roku", 4)
	// act
	c.EvictLeastRecent()

	resultValue, resultBool := c.GetValueFromKey("aapl")
	//assert
	assert.Equal(t,0,resultValue,"After ELR aapl should not exist in cache")
	assert.Equal(t, false, resultBool,"GetValueFromKey(aapl) should return false after removed from cache")
	assert.Equal(t,3 , c.currentSize, "After 4 inserts and 1 evict, current size should be 3")
}
func TestCache_EvictLeastRecentEmptyCache(t *testing.T) {
	// arrange
	c := NewCache(5)
	// act
	c.InsertKeyValuePair("aapl", 1)
	c.InsertKeyValuePair("msft", 1)

	c.EvictLeastRecent()
	c.EvictLeastRecent()
	result := c.EvictLeastRecent()
	//assert
	assert.Error(t, result, "")
}
func TestCache_UpdateMostRecent(t *testing.T) {
	// arrange
	c := NewCache(2)
	node := DoublyLinkedListNode{"roku",4,nil, nil}
	// act
	c.InsertKeyValuePair("aaple",1)
	c.InsertKeyValuePair("msft",2)
	c.InsertKeyValuePair("sq",3)
	c.UpdateMostRecent(&node)
	resultStr, _  := c.GetMostRecentKey()
	// assert
	assert.Equal(t,"roku",resultStr,"Most recent key should be updated to node.key")
}
func TestCache_ReplaceKey(t *testing.T) {
	// arrange
	c := NewCache(5)
	// act
	c.InsertKeyValuePair("aapl", 1)
	resultBool, err := c.ReplaceKey("aapl", 2)
	// assert
	assert.Equal(t, resultBool, true, "Expect true")
	assert.Equal(t, err, nil, "Error = nil when successful replacement")
}
func TestCache_ReplaceKeyExpectError(t *testing.T){
	// arrange
	c := NewCache(1)
	//act
	c.InsertKeyValuePair("aapl",1)
	resultBool, err := c.ReplaceKey("msft",2)
	//assert
	assert.Error(t, err,"Key doesnt exist in cache, can not replace")
	assert.Equal(t,false, resultBool,"Key doesnt exists, expect false")
}
func TestDoublyLinkedList_SetHeadTo(t *testing.T) {
	// arrange
	list := DoublyLinkedList{head: nil, tail:nil}
	a := DoublyLinkedListNode{"aapl",1,nil, nil}
	// act
	list.SetHeadTo(&a)
	// assert
	assert.Equal(t,list.head.key ,a.key,"List head.key should == a.key")
	assert.Equal(t,list.tail.key, a.key, "List tail.key should be == a.key after inserting first value into empty linked list")
}
func TestDoublyLinkedList_SetHeadToInsertTail(t *testing.T) {
	// arrange
	list := DoublyLinkedList{head:nil, tail:nil}
	a := DoublyLinkedListNode{key:"aapl",value: 1,prev: nil,next: nil}
	b := DoublyLinkedListNode{key:"msft", value: 2, prev: &a, next: nil}
	// act
	a.next = &b
	b.prev = &a
	list.SetHeadTo(&a)
	list.tail = &b
	list.SetHeadTo(&b)
	// assert
	assert.Equal(t,list.head.key, b.key, "Setting head of list == tail, removes tail from list and sets as head")
}
func TestDoublyLinkedList_RemoveTail(t *testing.T) {
	// arrange
	list:= DoublyLinkedList{head:nil, tail: nil}
	a := DoublyLinkedListNode{key:"aapl", value: 1, prev:nil, next:nil}
	b := DoublyLinkedListNode{key:"msft", value: 2, prev:nil, next:nil}
	c := DoublyLinkedListNode{key: "roku", value:3, prev:nil, next:nil}
	// act
	list.SetHeadTo(&a)
	a.next = &b
	b.prev = &a
	b.next = &c
	c.prev = &b
	list.tail = &c
	list.RemoveTail()
	//assert
	assert.Equal(t,b.key, list.tail.key)
}
func TestDoublyLinkedListNode_RemoveBindings(t *testing.T) {
	// arrange
	a := DoublyLinkedListNode{key:"aapl", value: 1, prev: nil, next: nil}
	b := DoublyLinkedListNode{key:"msft", value: 2, prev: &a, next: nil}
	c := DoublyLinkedListNode{key:"roku", value: 3, prev: &b, next: nil}
	// act
	b.next = &c
	a.RemoveBindings()
	b.RemoveBindings()
	c.RemoveBindings()
	// assert
	assert.Nil(t,a.next)
	assert.Nil(t,b.prev)
	assert.Nil(t,c.prev)
}
func TestDoublyLinkedListNode_RemoveBindingsRemoveOneElement(t *testing.T) {
	// arrange
	a := DoublyLinkedListNode{key:"aapl", value: 1, prev: nil, next: nil}
	b := DoublyLinkedListNode{key:"msft", value: 2, prev: &a, next: nil}
	c := DoublyLinkedListNode{key:"roku", value: 3, prev: &b, next: nil}
	// act
	a.next = &b
	b.next = &c
	b.RemoveBindings()
	// assert
	assert.True(t, a.next.key ==  c.key,"")
}