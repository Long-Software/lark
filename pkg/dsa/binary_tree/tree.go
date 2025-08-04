package binarytree

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) New(value int) *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(value int) bool {
	n := &Node{value: value}
	if b.root == nil {
		b.root = n
		return true
	}

	tmp := b.root
	for tmp != nil {
		if tmp.value == n.value {
			return false
		}
		if tmp.value > n.value {
			if tmp.left == nil {
				tmp.left = n
				return true
			}
			tmp = tmp.left
		} else {
			if tmp.right == nil {
				tmp.right = n
				return true
			}
			tmp = tmp.right
		}
	}
	return false
}

func (b *BinaryTree) Contains(value int) bool {
	if b.root == nil {
		return false
	}
	tmp := b.root
	for tmp != nil {
		if value < tmp.value {
			tmp = tmp.left
		} else if value > tmp.value {
			tmp = tmp.right
		} else {
			return true
		}
	}
	return false
}
