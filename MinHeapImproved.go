package main

// MinHeap is a binary tree based on slice of ints
type MinHeap struct {
	Values []int
}

type Element struct {
	Value interface{}
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
	if h.Len() > 0 {
		output = h.Values[0]
	}
	return
}

// ExtractMin returns smallest value in heap
// and repairs heap
func (h *MinHeap) ExtractMin() (output int) {
	output = -1
	n := h.Len() - 1
	if n >= 0 {
		h.Swap(0, n)
		output = h.Values[n]
		h.Values = h.Values[:n]
	}
	if n >= 1 {
		h.siftdown(0)
	}
	return
}

// Push adds element to heap preserving its structure
func (h *MinHeap) Push(element int) {
	h.Values = append(h.Values, element)
	h.siftup(h.Len() - 1)
}

// Remove deletes element with requested index from heap
// and repairs heap
func (h *MinHeap) Remove(index int) {
	if index < 0 || index >= h.Len() {
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
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.siftdown(i)
	}
}

func (h *MinHeap) siftup(child int) {
	for {
		parent := (child - 1) / 2
		if parent == child || !h.Less(child, parent) {
			break
		}
		h.Swap(child, parent)
		child = parent
	}
}

func (h *MinHeap) siftdown(parent int) {
	n := h.Len()
	for {
		lchild := parent*2 + 1
		rchild := lchild + 1
		if lchild >= n {
			break
		}
		child := lchild
		if rchild < n && h.Less(rchild, lchild) {
			child = rchild
		}
		if !h.Less(child, parent) {
			break
		}
		h.Swap(child, parent)
		parent = child
	}
}
