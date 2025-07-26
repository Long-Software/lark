package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	top    *Node[T]
	height int
}

func New[T any](val T) *Stack[T] {
	node := newNode(val)
	return &Stack[T]{
		top:    node,
		height: 1,
	}
}

func (s *Stack[T]) Push(val T) {
	node := newNode(val)
	if s.IsEmpty() {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}
	s.height++
}

func (s *Stack[T]) Pop(val T) *Node[T] {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.top
	s.top = s.top.next
	tmp.next = nil
	s.height--
	return tmp
}
func (s *Stack[T]) IsEmpty() bool {
	return s.height == 0
}
func (s *Stack[T]) Top() *Node[T] {
	return s.top
}
func (s *Stack[T]) Height() int {
	return s.height
}

/*
*
Returns a string representation of the linked list
*/
func (s *Stack[T]) String() string {
	var bd strings.Builder
	bd.WriteString("[")

	for tmp := s.top; tmp != nil; tmp = tmp.next {
		bd.WriteString(fmt.Sprintf("%+v", tmp.value))
		if tmp.next != nil {
			bd.WriteString(", ")
		}
	}

	bd.WriteString("]")
	return bd.String()
}
