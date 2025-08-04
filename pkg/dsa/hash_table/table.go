package hashtable

type HashTable[T any] struct {
	size    int
	dataMap []*Node[T]
}

func (h *HashTable[T]) Hash(key string) int {
	hash := 0
	var chars []rune
	for _, s := range key {
		chars = append(chars, s)
	}

	for i := 0; i < len(chars); i++ {
		ascii := int(chars[i])
		hash = (hash + ascii*23) % len(h.dataMap)
	}
	return hash
}

func (h *HashTable[T]) Set(key string, value T) {
	idx := h.Hash(key)
	n := NewNode(key, value)
	if h.dataMap[idx] == nil {
		h.dataMap[idx] = n
	} else {
		tmp := h.dataMap[idx]
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = n
	}
}

func (h *HashTable[T]) Get(key string) *T {
	idx := h.Hash(key)
	tmp := h.dataMap[idx]
	for tmp != nil {
		if tmp.key == key {
			return &tmp.value
		}
		tmp = tmp.next
	}
	return nil
}

func (h *HashTable[T]) Keys() []string {
	var keys []string
	for i := 0; i < len(h.dataMap); i++ {
		tmp := h.dataMap[i]
		for tmp != nil {
			keys = append(keys, tmp.key)
			tmp = tmp.next
		}
	}
	return keys
}
