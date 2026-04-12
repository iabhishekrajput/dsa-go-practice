package heaps

type Entry[T any] struct {
	Priority int
	Value    T
}

type Heap[T any] struct {
	data    []Entry[T]
	compare func(a, b int) bool
}

func NewMinHeap[T any]() *Heap[T] {
	return &Heap[T]{
		data:    []Entry[T]{},
		compare: func(a, b int) bool { return a < b },
	}
}

func NewMaxHeap[T any]() *Heap[T] {
	return &Heap[T]{
		data:    []Entry[T]{},
		compare: func(a, b int) bool { return a > b },
	}
}

func (h *Heap[T]) Insert(item Entry[T]) {
	h.data = append(h.data, item)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap[T]) Extract() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}

	extracted := h.data[0]

	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	h.heapifyDown(0)

	return extracted.Value, true
}

func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}

	return h.data[0].Value, true
}

func (h *Heap[T]) PeekPriority() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	return h.data[0].Priority, true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) heapifyUp(index int) {
	for index > 0 && h.compare(h.data[index].Priority, h.data[parent(index)].Priority) {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *Heap[T]) heapifyDown(index int) {
	lastIndex := len(h.data) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.compare(h.data[l].Priority, h.data[r].Priority) {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.compare(h.data[childToCompare].Priority, h.data[index].Priority) {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (h *Heap[T]) swap(idx1, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}
