package List

import (
	"math/rand"
	"time"
)

type ListInterface interface {
	Size() int
	First() *ListNode
	Last() *ListNode
	InsertAsFirst(e interface{})
	InsertAsLast(e interface{})
	InsertAfter(p ListNode, e interface{}) ListNode
	InsertBefore(p ListNode, e interface{}) ListNode
	Remove(p *ListNode) interface{}
	//disordered()
	Sort(p ListNode, n int)
	//find()
	Search(e interface{}, n int, p ListNode) ListNode
	Deduplicate() int
	Uniquify() int
	//traverse()
}

type ListNode struct {
	Data interface{}
	pred *ListNode
	succ *ListNode
}

type List struct {
	_size  int
	header ListNode
	tailer ListNode
}

func (L *List) init() {
	L.header.pred = nil
	L.header.succ = &L.tailer

	L.tailer.pred = &L.header
	L.tailer.succ = nil

	L._size = 0
}

//func NewList() List {
//	L := List{}
//	L.init()
//	return L
//}

/****** Basic ******/
func (L *List) Get(r int) interface{} {
	p := &L.header
	for ; 0 < r; r-- {
		p = p.succ
	}
	return p.Data
}

func (L *List) First() *ListNode {
	return L.header.succ
}

func (L *List) Last() *ListNode {
	return L.tailer.pred
}

func (L *List) Size() int {
	return L._size
}

func (L *List) Empty() bool {
	if L._size == 0 {
		return true
	} else {
		return false
	}
}

/****** Find ******/
func (L *List) findBefore(e interface{}, n int, p ListNode) *ListNode {
	for ; 0 < n; n-- {
		p = *p.pred
		if e == p.Data {
			return &p
		}
	}
	return nil
}

func (L *List) findAfter(e interface{}, p ListNode, n int) *ListNode {
	for ; 0 < n; n-- {
		p = *p.succ
		if e == p.Data {
			return &p
		}
	}
	return nil
}

/****** Insert ******/
func (L *List) InsertBefore(p ListNode, e interface{}) ListNode {
	L._size++
	return p.insertAsPred(e)
}

func (p *ListNode) insertAsPred(e interface{}) ListNode {
	x := ListNode{e, p.pred, p}
	p.pred.succ = &x
	p.pred = &x
	return x
}

func (L *List) InsertAfter(p ListNode, e interface{}) ListNode {
	L._size++
	return p.insertAsSucc(e)
}

func (p *ListNode) insertAsSucc(e interface{}) ListNode {
	x := ListNode{e, p, p.succ}
	p.succ.pred = &x
	p.succ = &x
	return x
}

func (L *List) InsertAsLast(e interface{}) {
	L.InsertBefore(L.tailer, e)
}

func (L *List) InsertAsFirst(e interface{}) {
	L.InsertAfter(L.header, e)
}

/****** Copy part of list ******/
func (L *List) copyNodes(p ListNode, n int) {
	L.init()
	for ; n != 0; n-- {
		L.InsertAsLast(p.Data)
		p = *p.succ
	}
}

/****** Remove ******/
func (L *List) Remove(p *ListNode) interface{} {
	e := p.Data
	p.pred.succ = p.succ
	p.succ.pred = p.pred
	L._size--
	return e
}

/****** Deduplicate ******/
func (L *List) Deduplicate() int {
	if L._size < 2 {
		return 0
	}
	oldSize := L._size
	p := *L.First()
	r := 1
	for p = *p.succ; L.tailer != p; p = *p.succ {
		q := L.findBefore(p.Data, r, p)
		if q != nil {
			L.Remove(q)
		} else {
			r++
		}
	}
	return oldSize - L._size
}

/****** For Sort List: uniquify ******/
func (L *List) Uniquify() int {
	if L._size < 2 {
		return 0
	}
	oldSize := L._size
	p := *L.First()
	for q := *p.succ; L.tailer != q; q = *p.succ {
		if q.Data == p.Data {
			L.Remove(&q)
		} else {
			p = q
		}
	}
	return oldSize - L._size
}

/****** For Sort List: search ******/
func (L *List) Search(e interface{}, n int, p ListNode) ListNode {
	for n--; 0 < n; n-- {
		p = *p.pred
		if p.Data.(float64) <= e.(float64) {
			break
		}
	}
	return p
}

/****** Selection Sort ******/
func (L *List) selectionSort(p ListNode, n int) {
	head := *p.pred
	tail := p

	// put tail into the right place
	for i := 0; i < n; i++ {
		tail = *p.succ
	}

	for 1 < n {
		L.InsertBefore(tail, L.Remove(L.selectMax(*head.succ, n)))
		tail = *tail.pred
		n--
	}
}

func (L *List) selectMax(p ListNode, n int) *ListNode { // 1 < n
	max := p
	for cur := p; 1 < n; n-- {
		cur = *cur.succ
		if max.Data.(float64) <= cur.Data.(float64) {
			max = cur
		}
	}
	return &max
}

/****** Selection Sort ******/
func (L *List) insertionSort(p ListNode, n int) {
	for r := 0; r < n; r++ {
		L.InsertAfter(L.Search(p.Data, r, p), p.Data)
		p = *p.succ
		L.Remove(p.pred)
	}
}

/****** Sort ******/
func (L *List) Sort(p ListNode, n int) {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(2) {
	case 0:
		L.selectionSort(p, n)
	case 1:
		L.insertionSort(p, n)
	}
}
