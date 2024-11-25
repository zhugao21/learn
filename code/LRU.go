package code

//type LRUCache struct {
//	capacity int
//	cache    map[int]*lruCacheNode
//	head     *lruCacheNode
//	tail     *lruCacheNode
//}
//
//type lruCacheNode struct {
//	pre  *lruCacheNode
//	next *lruCacheNode
//	key  int
//	val  int
//}
//
//func Constructor(capacity int) LRUCache {
//	head := &lruCacheNode{key: -1, val: -1}
//	tail := &lruCacheNode{key: -1, val: -1}
//	head.next = tail
//	tail.pre = head
//	return LRUCache{
//		capacity: capacity,
//		cache:    make(map[int]*lruCacheNode, capacity),
//		head:     head,
//		tail:     tail,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	if node, ok := this.cache[key]; ok {
//		this.Put(key, node.val)
//		return node.val
//	}
//	return -1
//}
//
//func (this *LRUCache) deleteNode(node *lruCacheNode) {
//	pre := node.pre
//	next := node.next
//	pre.next = next
//	next.pre = pre
//	delete(this.cache, node.key)
//}
//
//func (this *LRUCache) addNodeBeforeTail(node *lruCacheNode) {
//	pre := this.tail.pre
//	pre.next = node
//	node.pre = pre
//	node.next = this.tail
//	this.tail.pre = node
//	this.cache[node.key] = node
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if node, ok := this.cache[key]; ok {
//		this.deleteNode(node)
//		node.val = value
//		this.addNodeBeforeTail(node)
//	} else {
//		if len(this.cache) >= this.capacity {
//			this.deleteNode(this.head.next)
//		}
//		this.addNodeBeforeTail(&lruCacheNode{key: key, val: value})
//	}
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	Capacity   int
	Cache      map[int]*LRUCacheNode
	Head, Tail *LRUCacheNode
}

type LRUCacheNode struct {
	Pre, Next *LRUCacheNode
	key, val  int
}

func Constructor(capacity int) LRUCache {
	head := &LRUCacheNode{}
	tail := &LRUCacheNode{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*LRUCacheNode, capacity),
		Head:     head,
		Tail:     tail,
	}

}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.deleteNode(node)
		this.Put(node.key, node.val)
		return node.val
	}
	return -1
}

func (this *LRUCache) deleteNode(node *LRUCacheNode) {
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
	delete(this.Cache, node.key)
}

func (this *LRUCache) addNode(node *LRUCacheNode) {
	pre := this.Tail.Pre
	pre.Next = node
	node.Pre = pre
	node.Next = this.Tail
	this.Tail.Pre = node
	this.Cache[node.key] = node
}

func (this *LRUCache) Put(key int, value int) {
	newNode := &LRUCacheNode{key: key, val: value}
	if node, ok := this.Cache[key]; ok {
		this.deleteNode(node)
	} else {
		if len(this.Cache) >= this.Capacity {
			this.deleteNode(this.Head.Next)
		}
	}
	this.addNode(newNode)
}
