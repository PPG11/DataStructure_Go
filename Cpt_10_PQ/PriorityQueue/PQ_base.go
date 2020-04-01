package PriorityQueue

import (
	"datastructure/DataStructure_Go/Cpt_2_Vector/Vector"
	"datastructure/DataStructure_Go/Cpt_5_Tree/Tree"
)

type PQ struct {
	Vector.Vector
}

func (T *PQ) Parent(i Vector.Rank) Vector.Rank {
	return (i - 1) >> 1
}

func (T *PQ) LChild(i Vector.Rank) Vector.Rank {
	return 1 + (i << 1)
}

func (T *PQ) RChild(i Vector.Rank) Vector.Rank {
	return (1 + i) << 1
}

func (T *PQ) InHeap(i Vector.Rank) bool {
	if -1 < i && i < T.Size() {
		return true
	}
	return false
}

func (T *PQ) LastInternal() Vector.Rank {
	return T.Parent(T.Size() - 1)
}

func (T *PQ) ParentValid(i Vector.Rank) bool {
	return 0 < i
}

func (T *PQ) LChildValid(i Vector.Rank) bool {
	return T.InHeap(T.LChild(i))
}

func (T *PQ) RChildValid(i Vector.Rank) bool {
	return T.InHeap(T.RChild(i))
}

func (T *PQ) Bigger(i, j Vector.Rank) Vector.Rank {
	//相等时返回 i (前者)
	if T.Get(i).(int) < T.Get(j).(int) {
		return j
	} else {
		return i
	}
}

func (T *PQ) ProperParent(i Vector.Rank) Vector.Rank { //i lchild rchild 三者中最大者
	switch {
	case T.RChildValid(i):
		return T.Bigger(i, T.Bigger(T.LChild(i), T.RChild(i)))
	case T.LChildValid(i):
		return T.Bigger(i, T.LChild(i))
	default:
		return i
	}
}

type PQComplHeap struct {
	PQ
}

func (T *PQ) GetMax() int {
	return T.Get(0).(int)
}

func (T *PQ) InsertHeap(e interface{}) {
	T.InsertEnd(e)
	T.percolateUp(T.Size() - 1)
}

func (T *PQ) percolateUp(i Vector.Rank) Vector.Rank {
	for T.ParentValid(i) {
		if j := T.Parent(i); T.Get(i).(int) < T.Get(j).(int) {
			break
		} else {
			T.Swap(i, j)
			i = j
		}
	}
	return i
}

func (T *PQ) DelMax() interface{} {
	//maxElem := T.Get(0)
	T.Swap(0, T.Size()-1)
	maxElem := T.Remove(T.Size() - 1)
	T.percolateDown(0)
	return maxElem
}

func (T *PQ) percolateDown(i Vector.Rank) Vector.Rank {
	j := T.ProperParent(i)
	for i != j {
		T.Swap(i, j)
		i = j
		j = T.ProperParent(i)
	}
	return i
}

func (T *PQComplHeap) HeapCopy(A PQComplHeap, n Vector.Rank) {
	T.VectorCopyN(A.Vector, n)
	T.heapify()
}

func (T *PQComplHeap) heapifyBad(n Vector.Rank) {
	//蛮力
	for i := 1; i < n; i++ {
		T.percolateUp(i)
	}
}

func (T *PQComplHeap) heapify() {
	//Floyd
	for i := T.LastInternal(); T.InHeap(i); i-- {
		T.percolateDown(i)
	}
}

//Left Heap
type PQLeftHeap struct {
	PQ
	Tree.BinTree
}

//合并
func (T *PQLeftHeap) Merge(a, b Tree.BinNodePosi) Tree.BinNodePosi {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Data.(int) < b.Data.(int) {
		*a, *b = *b, *a
	}
	a.RChild = T.Merge(a.RChild, b)
	a.RChild.Parent = a
	if a.RChild == nil || a.LChild.Npl < a.RChild.Npl {
		*a.LChild, *a.RChild = *a.RChild, *a.LChild
	}
	if a.RChild != nil {
		a.Npl = a.RChild.Npl + 1
	} else {
		a.Npl = 1
	}
	return a
}

func (T *PQLeftHeap) Insert(e interface{}) {
	v := &Tree.BinNode{Data: e}
	T.SetRoot(T.Merge(T.Root(), v))
	T.SizeAdd(1)
}

func (T *PQLeftHeap) DelMax() interface{} {
	lHeap := T.Root().LChild
	rHeap := T.Root().RChild
	e := T.Root().Data
	T.SizeAdd(-1)
	T.SetRoot(T.Merge(lHeap, rHeap))
	if T.Root() != nil {
		T.Root().Parent = nil
	}
	return e
}
