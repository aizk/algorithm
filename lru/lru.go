package lru

// lru with hashmap and list
type LRU struct {
	m        map[interface{}]*Node
	capacity int
	head     *Node
	tail     *Node
}

func NewLRU(cap int) *LRU {
	return &LRU{
		m:        make(map[interface{}]*Node, cap),
		capacity: cap,
	}
}

func (l *LRU) Get(key interface{}) interface{} {
	return l.m[key].value
}

func (l *LRU) Keys() (keys []interface{}) {
	for key := range l.m {
		keys = append(keys, key)
	}
	return
}

// set node
func (l *LRU) Put(key, value interface{}) {
	oldNode, ok := l.m[key]
	if ok {
		// exist
		// set new value
		oldNode.value = value
		// remove old node
		l.Remove(oldNode)
		// set to head
		l.setHead(oldNode)
		// return value
	} else {
		node := NewNode(key, value)
		if len(l.m) >= l.capacity {
			// need remove tail
			tailKey := l.tail.key
			l.Remove(l.m[tailKey])
			delete(l.m, tailKey)
			l.setHead(node)
		} else {
			l.setHead(node)
		}
		// set key node
		l.m[key] = node
	}
}

func (l *LRU) setHead(node *Node) {
	node.next = l.head
	node.pre = nil

	if l.head != nil {
		l.head.pre = node
	}

	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

// remove node
func (l *LRU) Remove(node *Node) {
	// set next
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		l.head = node.next
	}

	// set pre
	if node.next != nil {
		node.next.pre = node.pre
	} else {
		l.tail = node.pre
	}
}

type Node struct {
	key interface{}
	value interface{}
	pre *Node
	next *Node
}

func NewNode(key, value interface{}) *Node {
	return &Node{
		key: key,
		value: value,
	}
}

