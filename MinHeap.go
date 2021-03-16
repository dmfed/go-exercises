package main

// MinHeap is a binary tree based on slice of ints
type MinHeap struct {
	Values []int
}

// NewMinHeap is basically a heapify method for slice which returns instance of MinHeap
// It modifies incoming slice in place
func NewMinHeap(slice []int) *MinHeap {
	var heap MinHeap
	heap.Values = slice
	heap.Heapify()
	return &heap
}

// GetMin returns smalles value in heap
func (h *MinHeap) GetMin() (output int) {
	output = -1
	if len(h.Values) > 0 {
		output = h.Values[0]
	}
	return
}

// ExtractMin returns smallest value in heap
// and repairs heap
func (h *MinHeap) ExtractMin() (output int) {
	output = -1
	if len(h.Values) > 0 {
		output = h.Values[0]
		h.Values[0] = h.Values[len(h.Values)-1]
		h.Values = h.Values[:len(h.Values)-1]
	}
	if len(h.Values) > 1 {
		h.siftdown(0)
	}
	return
}

// Push adds element to heap preserving its structure
func (h *MinHeap) Push(element int) {
	h.Values = append(h.Values, element)
	h.siftup(len(h.Values) - 1)
}

// Remove deletes element with requested index from heap
// and repairs heap
func (h *MinHeap) Remove(index int) {
	if index < 0 || index >= len(h.Values) {
		return
	}
	h.Values[index] = -1 //We assume that values in the MinHeap are within [0, 10^9]
	h.siftup(index)
	h.ExtractMin()
}

// Len returns current number of elements in the heap
func (h *MinHeap) Len() int {
	return len(h.Values)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.Values[i] < h.Values[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.Values[i], h.Values[j] = h.Values[j], h.Values[i]
}

// Return returns inderlying slice
func (h *MinHeap) Return() []int {
	return h.Values
}

// Heapify tranforms undelying slice to usable
// heap.
func (h *MinHeap) Heapify() {
	for i := len(h.Values) / 2; i >= 0; i-- {
		h.siftdown(i)
	}
}

func (h *MinHeap) siftup(index int) {
	parent := (index - 1) / 2
	if parent < 0 {
		parent = 0
	}
	if h.Less(index, parent) {
		h.Swap(index, parent)
		h.siftup(parent)
	}
}

func (h *MinHeap) siftdown(index int) {
	child := index
	lchild := index*2 + 1
	rchild := index*2 + 2
	if lchild < len(h.Values) {
		child = lchild
	}
	if rchild < len(h.Values) && h.Less(rchild, lchild) {
		child = rchild
	}
	if h.Less(child, index) {
		h.Swap(child, index)
		h.siftdown(child)
	}
}

func (h *MinHeap) indexIsValid(index int) (result bool) {
	if index >= 0 && index < len(h.Values) {
		result = true
	}
	return
}
