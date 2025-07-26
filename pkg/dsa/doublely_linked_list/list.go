package doublelylinkedlist

import (
	"fmt"
	"strings"
)

type DoublelyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T any](val T) *DoublelyLinkedList[T] {
	node := newNode(val)
	return &DoublelyLinkedList[T]{
		head: node,
		tail: node,
		len:  1,
	}
}

func (d *DoublelyLinkedList[T]) Append(val T) {
	node := newNode(val)
	if d.IsEmpty() {
		d.head, d.tail = node, node
	} else {
		d.tail.next = node
		node.prev = d.tail
		d.tail = node
	}
	d.len++
}

func (d *DoublelyLinkedList[T]) Pop() *Node[T] {
	if d.IsEmpty() {
		return nil
	}

	tmp := d.tail
	d.tail = d.tail.prev
	d.tail.next = nil
	tmp.prev = nil
	d.len--
	if d.IsEmpty() {
		d.head, d.tail = nil, nil
	}
	return tmp
}

func (d *DoublelyLinkedList[T]) Prepend(val T) {
	node := newNode(val)
	if d.IsEmpty() {
		d.head, d.tail = node, node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.len++
}

func (d *DoublelyLinkedList[T]) RemoveFirst() *Node[T] {
	if d.IsEmpty() {
		return nil
	}
	tmp := d.head
	if d.len == 1 {
		d.head, d.tail = nil, nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
		tmp.next = nil
	}
	d.len--
	return tmp
}
func (d *DoublelyLinkedList[T]) Get(index int) *Node[T] {
	if d.isOutOfBound(index) {
		return nil
	}
	var tmp *Node[T]
	if index < d.len/2 {
		tmp = d.head
		for i := 0; i < index; i++ {
			tmp = tmp.next
		}
	} else {
		tmp = d.tail
		for i := d.len - 1; i > index; i-- {
			tmp = tmp.prev
		}
	}
	return tmp
}

func (d *DoublelyLinkedList[T]) Set(index int, val T) bool {
	tmp := d.Get(index)
	if tmp != nil {
		tmp.SetValue(val)
		return true
	}
	return false
}

func (d *DoublelyLinkedList[T]) Insert(index int, val T) bool {
	if d.isOutOfBound(index) {
		return false
	}
	if index == 0 {
		d.Prepend(val)
		return true
	}
	if index == d.len {
		d.Append(val)
		return true
	}

	node := newNode(val)
	prev := d.Get(index - 1)
	next := prev.next

	node.prev = prev
	node.next = next
	prev.next = node
	next.prev = node
	d.len++
	return true
}

func (d *DoublelyLinkedList[T]) Remove(index int) *Node[T] {
	if d.isOutOfBound(index) {
		return nil
	}
	if index == 0 {
		return d.RemoveFirst()
	}
	if index == d.len {
		return d.Pop()
	}

	tmp := d.Get(index)
	tmp.next.prev = tmp.prev
	tmp.prev.next = tmp.next

	tmp.prev = nil
	tmp.next = nil
	d.len--
	return tmp
}

func (d *DoublelyLinkedList[T]) Head() *Node[T] {
	return d.head
}
func (d *DoublelyLinkedList[T]) Tail() *Node[T] {
	return d.tail
}
func (d *DoublelyLinkedList[T]) Size() int {
	return d.len
}
func (d *DoublelyLinkedList[T]) IsEmpty() bool {
	return d.len == 0
}

func (l *DoublelyLinkedList[T]) isOutOfBound(index int) bool {
	return index < 0 || index >= l.len
}

/*
*
Returns a string representation of the linked list
*/
func (l *DoublelyLinkedList[T]) String() string {
	var bd strings.Builder
	bd.WriteString("[")

	for tmp := l.head; tmp != nil; tmp = tmp.next {
		bd.WriteString(fmt.Sprintf("%+v", tmp.value))
		if tmp.next != nil {
			bd.WriteString(", ")
		}
	}

	bd.WriteString("]")
	return bd.String()
}
