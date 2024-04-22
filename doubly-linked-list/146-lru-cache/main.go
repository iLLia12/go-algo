package main

type NodeList struct {
	key int
	value int
	next *NodeList
	prev *NodeList
}

type LRUCache struct {
	capacity int
	cache map[int]*NodeList
	head *NodeList
	tail *NodeList
}

func Constructor(capacity int) LRUCache {
	head := &NodeList{0,0, nil, nil}
	tail := &NodeList{0,0, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cache: make(map[int]*NodeList),
		capacity: capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Insert(node *NodeList) {
	prev := this.tail.prev
	prev.next = node
	node.prev = prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) Delete(node *NodeList) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if ok {
		this.Delete(node)
		this.Insert(node)
		return this.cache[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
			node.value = value
			this.Delete(node)
			this.Insert(node)
			return
	}

	newNode := &NodeList{key: key, value: value}
	this.cache[key] = newNode
	this.Insert(newNode)

	if len(this.cache) > this.capacity {
			lru := this.head.next
			this.Delete(lru)
			delete(this.cache, lru.key)
	}
}

func main() {
	
}