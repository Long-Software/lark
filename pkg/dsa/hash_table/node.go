package hashtable

type Node[T any] struct {
	key   string
	value T
	next  *Node[T]
}

func NewNode[T any](key string, value T) *Node[T] {
	return &Node[T]{
		key:   key,
		value: value,
	}
}
