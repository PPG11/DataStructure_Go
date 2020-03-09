package List

type ListNode struct {
	data interface{}
	pred *ListNode
	succ *ListNode
}

type List struct {
	_size  int
	header ListNode
	tailer ListNode
}

func (L List) init() {
	L.header.pred = nil
	L.header.succ = &L.tailer

	L.tailer.pred = &L.header
	L.tailer.succ = nil

	L._size = 0
}

func NewList() List {
	L := List{}
	L.init()
	return L
}

/****** Basic ******/
func (L List) get(r int) interface{} {
	p := &L.header
	for ; 0 < r; r-- {
		p = p.succ
	}
	return p.data
}

func (L List) first() ListNode {
	return *L.header.succ
}

func (L List) last() ListNode {
	return *L.tailer.pred
}

/****** Find ******/
func (L List) findBefore(e interface{}, n int, p ListNode) *ListNode {
	for ; 0 < n; n-- {
		p = *p.pred
		if e == p.data {
			return &p
		}
	}
	return nil
}

func (L List) findAfter(e interface{}, p ListNode, n int) *ListNode {
	for ; 0 < n; n-- {
		p = *p.succ
		if e == p.data {
			return &p
		}
	}
	return nil
}

/****** Insert ******/
func (L List) insertBefore(p ListNode, e interface{}) ListNode {
	L._size++
	return p.insertAsPred(e)
}

func (p ListNode) insertAsPred(e interface{}) ListNode {
	x := ListNode{e, p.pred, &p}
	p.pred.succ = &x
	p.pred = &x
	return x
}

func (L List) insertAfter(p ListNode, e interface{}) ListNode {
	L._size++
	return p.insertAsSucc(e)
}

func (p ListNode) insertAsSucc(e interface{}) ListNode {
	x := ListNode{e, &p, p.succ}
	p.succ.pred = &x
	p.succ = &x
	return x
}

func (L List) insertAsLast(e interface{}) {
	L.insertBefore(L.tailer, e)
}

func (L List) insertAsFirst(e interface{}) {
	L.insertAfter(L.header, e)
}

/****** Copy part of list ******/
func (L List) copyNodes(p ListNode, n int) {
	L.init()
	for ; n != 0; n-- {
		L.insertAsLast(p.data)
		p = *p.succ
	}
}

/****** Remove ******/
func (L List) remove(p ListNode) interface{} {
	e := p.data
	p.pred.succ = p.succ
	p.succ.pred = p.pred
	L._size--
	return e
}

/****** Deduplicate ******/
func (L List) deduplicate() int {
	if L._size < 2 {
		return 0
	}
	oldSize := L._size
	p := L.first()
	r := 1
	for p = *p.succ; L.tailer != p; p = *p.succ {
		q := L.findBefore(p.data, r, p)
		if q != nil {
			L.remove(*q)
		} else {
			r++
		}
	}
	return oldSize - L._size
}

/****** For Sort List: uniquify ******/
func (L List) uniquify() int {
	if L._size < 2 {
		return 0
	}
	oldSize := L._size
	p := L.first()
	for q := *p.succ; L.tailer != q; q = *p.succ {
		if q.data == p.data {
			L.remove(q)
		} else {
			p = q
		}
	}
	return oldSize - L._size
}

/****** For Sort List: search ******/
func (L List) search(e interface{}, n int, p ListNode) ListNode {
	for n--; 0 < n; n-- {
		p = *p.pred
		if p.data.(float64) <= e.(float64) {
			break
		}
	}
	return p
}

/****** Selection Sort ******/
func (L List) selectionSort(p ListNode, n int) {
	head := *p.pred
	tail := p

	// put tail into the right place
	for i := 0; i < n; i++ {
		tail = *p.succ
	}

	for 1 < n {
		L.insertBefore(tail, L.remove(L.selectMax(*head.succ, n)))
		tail = *tail.pred
		n--
	}
}

func (L List) selectMax(p ListNode, n int) ListNode { // 1 < n
	max := p
	for cur := p; 1 < n; n-- {
		cur = *cur.succ
		if max.data.(float64) <= cur.data.(float64) {
			max = cur
		}
	}
	return max
}

/****** Selection Sort ******/
func (L List) insertionSort(p ListNode, n int) {
	for r := 0; r < n; r++ {
		L.insertAfter(L.search(p.data, r, p), p.data)
		p = *p.succ
		L.remove(*p.pred)
	}
}
