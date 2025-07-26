package test

import (
	"testing"

	ll "github.com/Long-Software/lark/pkg/dsa/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	arr := ll.New(2)
	assert.Equal(t, arr.Get(0).GetValue(), 2)
}

func TestLinkedListAppend(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)
	assert.Equal(t, arr.Get(1).GetValue(), 5)
	assert.Equal(t, arr.Size(), 2)
}

func TestLinkedListPop(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)
	arr.Pop()
	assert.Equal(t, arr.Size(), 1)
}

func TestLinkedListPrepend(t *testing.T) {
	arr := ll.New(2)
	arr.Prepend(5)
	assert.Equal(t, arr.Get(0).GetValue(), 5)
	assert.Equal(t, arr.Size(), 2)
}

func TestLinkedListRemoveFrist(t *testing.T) {
	arr := ll.New(2)
	node := arr.RemoveFirst()
	assert.Equal(t, arr.Size(), 0)
	assert.Equal(t, node.GetValue(), 2)
}

func TestLinkedListGet(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)
	arr.Append(7)
	arr.Append(9)

	assert.Equal(t, arr.Get(1).GetValue(), 5)
	assert.Equal(t, arr.Get(2).GetValue(), 7)
	assert.Equal(t, arr.Get(3).GetValue(), 9)
}

func TestLinkedListSet(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)
	res := arr.Set(2, 7)
	assert.Equal(t, res, false)

	res = arr.Set(1, 7)
	assert.Equal(t, res, true)
	assert.Equal(t, arr.Get(1).GetValue(), 7)
}

func TestLinkedListInsert(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)

	res := arr.Insert(3, 7)
	assert.Equal(t, res, false)

	res = arr.Insert(1, 7)
	assert.Equal(t, res, true)
	assert.Equal(t, arr.Get(1).GetValue(), 7)
	assert.Equal(t, arr.Get(2).GetValue(), 5)
}

func TestLinkedListRemove(t *testing.T) {
	arr := ll.New(2)

	node := arr.Remove(1)
	assert.Nil(t, node)

	node = arr.Remove(0)
	assert.Equal(t, node.GetValue(), 2)
	assert.Equal(t, arr.Size(), 0)
}

func TestReverse(t *testing.T) {
	arr := ll.New(2)
	arr.Append(5)
	arr.Append(7)
	arr.Append(9)

	arr.Reverse()
	assert.Equal(t, arr.Get(0).GetValue(), 9)
	assert.Equal(t, arr.Get(3).GetValue(), 2)
}
