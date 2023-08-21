package multi_provider_configs

import (
	"sync"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
	mu     sync.Mutex
}

// NewQueue initializes and returns a new Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds a value to the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newNode := &Node[T]{value: value}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.length++
}

// Dequeue removes and returns the value from the front of the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.head == nil {
		var zero T
		return zero, false
	}

	value := q.head.value
	nextNode := q.head.next
	q.head.next = nil // Ensure memory can be garbage collected
	q.head = nextNode

	if q.head == nil {
		q.tail = nil
	}

	q.length--
	return value, true
}

// Length returns the number of elements in the queue.
func (q *Queue[T]) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.length
}

func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.length == 0
}
