package lrucache146

//DListNode 链表的结点
type DListNode struct {
	key  int
	val  int
	prio *DListNode
	next *DListNode
}

//LRUCache 双向链表
type LRUCache struct {
	cache    map[int]*DListNode
	capacity int
	head     *DListNode
	tail     *DListNode
}

//Constructor 构造LRUCache 双向链表
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    map[int]*DListNode{},
		capacity: capacity,
		head:     &DListNode{key: 0, val: 0},
		tail:     &DListNode{key: 0, val: 0},
	}
	lru.head.next = lru.tail
	lru.tail.prio = lru.head
	return lru
}

//RemoveNode 删除制定结点
func (l *LRUCache) RemoveNode(node *DListNode) {
	node.prio.next = node.next
	node.next.prio = node.prio
}

//RemoveNode 删除最后一个结点
func (l *LRUCache) RemoveNodeLast() *DListNode {
	lastNode := l.tail.prio
	l.RemoveNode(lastNode)
	return lastNode
}

//AddNodeFirst 添加结点到第一位
func (l *LRUCache) AddNodeFirst(node *DListNode) {
	node.prio = l.head
	node.next = l.head.next
	l.head.next.prio = node
	l.head.next = node
}

//MoveNodeFirst 移动结点到第一位
func (l *LRUCache) MoveNodeFirst(node *DListNode) {
	l.RemoveNode(node)
	l.AddNodeFirst(node)
}

//Get
// 1. 如果获取的到 则将缓存放在第一位
// 2.  获取不到 返回-1
func (l *LRUCache) Get(key int) int {
	if v, ok := l.cache[key]; ok {
		l.MoveNodeFirst(v)
		return v.val
	}
	return -1
}

//Put
// 1. key存在 则将缓存放在第一位，并且更新value
// 2. key不存在 则创建新的链表结点放到第一位，并且添加到cache中
// 3. 缓存满了 删除最后一位 并且删除cache中的key
func (l *LRUCache) Put(key int, value int) {
	node := l.cache[key]
	if node != nil {
		node.val = value
		l.MoveNodeFirst(node)
		return
	}
	node = &DListNode{key: key, val: value}
	l.AddNodeFirst(node)
	l.cache[key] = node

	if len(l.cache) > l.capacity {
		rmNode := l.RemoveNodeLast()
		delete(l.cache, rmNode.key)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
