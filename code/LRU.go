package code

type LRUCache struct {
	capacity int
	cache    map[int]*lruCacheNode
	head     *lruCacheNode
	tail     *lruCacheNode
}

type lruCacheNode struct {
	pre  *lruCacheNode
	next *lruCacheNode
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {
	head := &lruCacheNode{key: -1, val: -1}
	tail := &lruCacheNode{key: -1, val: -1}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*lruCacheNode, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.Put(key, node.val)
		return node.val
	}
	return -1
}

func (this *LRUCache) deleteNode(node *lruCacheNode) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
	delete(this.cache, node.key)
}

func (this *LRUCache) addNodeBeforeTail(node *lruCacheNode) {
	pre := this.tail.pre
	pre.next = node
	node.pre = pre
	node.next = this.tail
	this.tail.pre = node
	this.cache[node.key] = node
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.deleteNode(node)
		node.val = value
		this.addNodeBeforeTail(node)
	} else {
		if len(this.cache) >= this.capacity {
			this.deleteNode(this.head.next)
		}
		this.addNodeBeforeTail(&lruCacheNode{key: key, val: value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
