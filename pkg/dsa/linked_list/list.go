package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T any](val T) *LinkedList[T] {
	node := newNode(val)
	return &LinkedList[T]{
		head: node,
		tail: node,
		len:  1,
	}
}

/*
*
Add a new value to the end of the linked list
*/
func (l *LinkedList[T]) Append(val T) {
	node := newNode(val)
	if l.IsEmpty() { // if the list is empty
		l.head, l.tail = node, node
	} else { // if there are items in the list
		l.tail.next = node // set the current tail link to the next node
		l.tail = node      // change the current tail to the next node
	}

	l.len++
}

/*
*	Remove the last item of the linked list
 */
func (l *LinkedList[T]) Pop() *Node[T] {
	if l.IsEmpty() {
		return &Node[T]{}
	}
	var tmp, prev *Node[T]
	for tmp = l.head; tmp != nil; tmp = tmp.next {
		if tmp.next != nil {
			prev = tmp
		}
	}
	l.tail = prev
	l.tail.next = nil
	l.len--
	if l.IsEmpty() {
		l.head, l.tail = nil, nil
	}
	return tmp
}

/*
*
Add an item to the first of the linked list
*/

func (l *LinkedList[T]) Prepend(val T) {
	node := newNode(val)
	if l.IsEmpty() {
		l.head, l.tail = node, node
	} else {
		node.next = l.head
		l.head = node
	}
	l.len++
}

func (l *LinkedList[T]) RemoveFirst() *Node[T] {
	tmp := l.head
	l.head = l.head.next
	tmp.next = nil
	l.len--
	if l.IsEmpty() {
		l.tail = nil
	}
	return tmp
}

func (l *LinkedList[T]) Get(index int) *Node[T] {
	if l.isOutOfBound(index) {
		return nil
	}
	tmp := l.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	return tmp
}

func (l *LinkedList[T]) Set(index int, val T) bool {
	tmp := l.Get(index)
	if tmp != nil {
		tmp.SetValue(val)
		return true
	}
	return false
}

func (l *LinkedList[T]) Insert(index int, val T) bool {
	if l.isOutOfBound(index) {
		return false
	}
	if index == 0 {
		l.Prepend(val)
		return true
	}

	if index == l.len {
		l.Append(val)
		return true
	}
	node := newNode(val)
	tmp := l.Get(index - 1)
	node.next = tmp.next
	tmp.next = node
	l.len++
	return true
}

func (l *LinkedList[T]) Remove(index int) *Node[T] {
	if l.isOutOfBound(index) {
		return nil
	}
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.len-1 {
		return l.Pop()
	}

	prev := l.Get(index - 1)
	tmp := prev.next
	prev.next = tmp.next
	tmp.next = nil
	l.len--
	return tmp
}

func (l *LinkedList[T]) Reverse() {
	tmp := l.head
	l.head = l.tail
	l.tail = tmp

	prev, next := &Node[T]{}, tmp.next

	for i := 0; i < l.len; i++ {
		next = tmp.next
		tmp.next = prev
		prev = tmp
		tmp = next
	}
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}
func (l *LinkedList[T]) Tail() *Node[T] {
	return l.tail
}
func (l *LinkedList[T]) Size() int {
	return l.len
}
func (l *LinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList[T]) isOutOfBound(index int) bool {
	return index < 0 || index >= l.len
}

/*
*
Returns a string representation of the linked list
*/
func (l *LinkedList[T]) String() string {
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
