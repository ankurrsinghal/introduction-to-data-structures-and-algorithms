package LinkedListQueue

type Node struct {
	data int
	next *Node
}

// MyCircularQueue ...
type MyCircularQueue struct {
	head *Node
	tail *Node
	size int
	cap  int
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	node := Node{data: value}
	node.next = this.tail
	this.tail = &node
	if this.head == nil {
		this.head = this.tail
	}
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
	}
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.head.data
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.tail.data
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

// Constructor ...
func Constructor(cap int) MyCircularQueue {
	return MyCircularQueue{size: 0, cap: cap}
}
