package main

import "fmt"

type MaxHeap struct {
	nodes []int
}

func NewMaxHeap(arr []int) MaxHeap {
	var mh MaxHeap
	mh.nodes = []int{}
	for _, val := range arr {
		mh.Insert(val)
	}
	return mh
}

func (mh *MaxHeap) Insert(n int) {
	mh.nodes = append(mh.nodes, n)
	mh.siftUp(len(mh.nodes) - 1)
}

func (mh *MaxHeap) GetMax() int {
	return mh.nodes[0]
}

func (mh *MaxHeap) ExtractMax() int {
	n := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]
	if len(mh.nodes) > 1 {
		mh.siftDown(0)
	}
	return n
}

func (mh *MaxHeap) siftUp(i int) {
	p := mh.parent(i)
	for mh.nodes[i] > mh.nodes[p] {
		mh.nodes[i], mh.nodes[p] = mh.nodes[p], mh.nodes[i]
		i = p
		p = mh.parent(i)
	}

}

func (mh *MaxHeap) siftDown(p int) {
	ch := mh.maxChild(p)
	for mh.nodes[p] < mh.nodes[ch] {
		mh.nodes[p], mh.nodes[ch] = mh.nodes[ch], mh.nodes[p]
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
	if p*2+1 < len(mh.nodes)-1 {
		ch1 = p*2 + 1
	}
	if p*2+2 < len(mh.nodes)-1 {
		ch2 = p*2 + 2
	}
	if mh.nodes[ch2] > mh.nodes[ch1] {
		maxchild = ch2
	} else {
		maxchild = ch1
	}
	return maxchild
}

func main() {
	a := []int{2, 9, 3, 6, 7, 10, 5}
	b := NewMaxHeap(a)
	fmt.Println(b)
	fmt.Println(b.GetMax())
	b.Insert(11)
	for len(b.nodes) > 1 {
		fmt.Println(b.ExtractMax())
	}
	fmt.Println(b.nodes)
	fmt.Println(b)
	c := NewMaxHeap([]int{1})
	fmt.Println(c)
	c.ExtractMax()
	fmt.Println(c)
}
