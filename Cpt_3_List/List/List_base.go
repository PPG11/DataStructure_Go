package List

type ListNode struct {
	data interface{}
	pred *ListNode
	succ *ListNode
}

type List struct {
	_size   int
	header  ListNode
	trailer ListNode
}

func (L List) init() {
	L.header.pred = nil
	L.header.succ = &L.trailer

	L.trailer.pred = &L.header
	L.trailer.succ = nil

	L._size = 0
}
