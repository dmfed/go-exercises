package main

import "fmt"

type MaxHeap []int

func NewMaxHeap(arr []int) MaxHeap {
	var mh MaxHeap
	for _, val := range arr {
		mh.Insert(val)
	}
	return mh
}

func (mh *MaxHeap) Insert(n int) {
	(*mh) = append((*mh), n)
	mh.siftUp(len((*mh)) - 1)
}

func (mh *MaxHeap) GetMax() int {
	return (*mh)[0]
}

func (mh *MaxHeap) ExtractMax() int {
	n := (*mh)[0]
	(*mh)[0] = (*mh)[len((*mh))-1]
	(*mh) = (*mh)[:len((*mh))-1]
	if len((*mh)) > 1 {
		mh.siftDown(0)
	}
	return n
}

func (mh *MaxHeap) siftUp(i int) {
	p := mh.parent(i)
	for (*mh)[i] > (*mh)[p] {
		(*mh)[i], (*mh)[p] = (*mh)[p], (*mh)[i]
		i = p
		p = mh.parent(i)
	}
}

func (mh *MaxHeap) siftDown(p int) {
	ch := mh.maxChild(p)
	for (*mh)[p] < (*mh)[ch] {
		(*mh)[p], (*mh)[ch] = (*mh)[ch], (*mh)[p]
		p = ch
		ch = mh.maxChild(p)
	}
}

func (mh *MaxHeap) parent(ch int) int {
	if ch >= 1 {
		return (ch - 1) / 2
	}
	return 0
}

func (mh *MaxHeap) maxChild(p int) int {
	ch1, ch2, maxchild := p, p, p
	if p*2+1 < len((*mh)) {
		ch1 = p*2 + 1
	}
	if p*2+2 < len((*mh)) {
		ch2 = p*2 + 2
	}
	if (*mh)[ch2] > (*mh)[ch1] {
		maxchild = ch2
	} else {
		maxchild = ch1
	}
	return maxchild
}

func main() {
	a := []int{2, 3, 1, 4, 6, 5, 8, 10, 7, 9, 12, 11}

	b := NewMaxHeap(a)
	for len(b) > 0 {
		fmt.Println(b)
		fmt.Println(b.ExtractMax())
	}
}
