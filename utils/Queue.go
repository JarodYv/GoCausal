package utils

type elem struct {
	value interface{}
	prev  *elem
	next  *elem
}

type LinkedQueue struct {
	head *elem
	tail *elem
	size int
}

func (queue *LinkedQueue) Size() int {
	return queue.size
}

func (queue *LinkedQueue) Peek() interface{} {
	if queue.head == nil {
		panic("Empty queue.")
	}
	return queue.head.value
}

func (queue *LinkedQueue) Append(value interface{}) {
	newElem := &elem{value, queue.tail, nil}
	if queue.tail == nil {
		queue.head = newElem
		queue.tail = newElem
	} else {
		queue.tail.next = newElem
		queue.tail = newElem
	}
	queue.size++
	newElem = nil
}

func (queue *LinkedQueue) Pop() interface{} {
	if queue.head == nil {
		panic("Empty queue.")
	}
	firstElem := queue.head
	queue.head = firstElem.next
	queue.size--
	return firstElem
}
