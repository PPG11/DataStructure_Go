package Vector

import (
	"datastructure/DataStructure_Go/Cpt_10_PQ/PriorityQueue"
	"math/rand"
	"reflect"
	"time"
)

type Rank = int

const DefaultCapacity int = 3

type VectorInterface interface {
	Size() Rank
	Get(r Rank) interface{}
	Put(r Rank, e interface{})
	Insert(r Rank, e interface{})
	Remove(r Rank) interface{}
	Disordered() int
	Sort()
	Find(e interface{}) Rank
	Search(e interface{}, lo Rank, hi Rank) Rank
	Deduplicate() int
	Uniquify() int
	Traverse(fun1 func(interface{}))
}

type Vector struct {
	_size     Rank
	_capacity int
	_elem     []interface{}
}

/* ----- Constructor Function ----- */
// capability, scale, v = initial element
func (T *Vector) Vector(c int, s int, v interface{}) {
	T._elem = make([]interface{}, s, c)
	if c < DefaultCapacity {
		T._capacity = DefaultCapacity
	} else {
		T._capacity = c
	}
	T._size = 0
	for ; T._size < s; T._size++ {
		T._elem[T._size] = v
	}
}

//func (T *Vector) VectorCopy(A Vector, n Rank) {
//	T.CopyFrom(A._elem, 0, n)
//}

func (T *Vector) VectorCopyN(A Vector, n Rank) {
	T.CopyFrom(A._elem, 0, n)
}

func (T *Vector) VectorCopyLH(A Vector, lo Rank, hi Rank) {
	T.CopyFrom(A._elem, lo, hi)
}

func (T *Vector) CopyWholeVector(V Vector) {
	T.CopyFrom(V._elem, 0, V._size)
}

func (T *Vector) CopyVectorLH(V Vector, lo Rank, hi Rank) {
	T.CopyFrom(V._elem, lo, hi)
}

/* ----- Read Only Interface ----- */
func (T *Vector) Size() Rank {
	return T._size
}

func (T *Vector) Empty() bool {
	if T._size == 0 {
		return true
	} else {
		return false
	}
}

func (T *Vector) Disordered() int {
	if reflect.TypeOf(T._elem[0]).String() == "Slice" || reflect.TypeOf(T._elem[0]).String() == "Map" || reflect.TypeOf(T._elem[0]).String() == "Func" {
		return -1
	}
	n := 0
	for i := 1; i < T._size; i++ {
		if T._elem[i-1].(float64) > T._elem[i].(float64) {
			n++
		}
	}
	return n
}

func (T *Vector) Find(e interface{}) Rank {
	return T.findLH(e, 0, T._size)
}

func (T *Vector) findLH(e interface{}, lo Rank, hi Rank) Rank {
	hi--
	for lo <= hi {
		if T._elem[hi] == e {
			break
		}
		hi--
	}
	return hi
	// hi < lo 则意味着失败
	// 交给上层算法判断
}

func (T *Vector) Search(e interface{}) Rank {
	return T.SearchLH(e, 0, T._size)
}

func (T *Vector) SearchLH(e interface{}, lo Rank, hi Rank) Rank {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return T.binSearch(e, lo, hi)
	} else {
		return T.fibSearch(e, lo, hi)
	}
}

func (T *Vector) binSearch(e interface{}, lo Rank, hi Rank) Rank {
	for lo < hi {
		switch mi := (lo + hi) >> 1; {
		case e.(float64) < T._elem[mi].(float64):
			hi = mi
		case T._elem[mi].(float64) < e.(float64):
			lo = mi + 1
		default:
			return mi
		}
	}
	return -1
}

func (T *Vector) binSearch2(e interface{}, lo Rank, hi Rank) Rank {
	for lo < hi {
		if mi := (lo + hi) >> 1; e.(float64) < T._elem[mi].(float64) {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo - 1
}

func (T *Vector) fibSearch(e interface{}, lo Rank, hi Rank) Rank {
	fib := Fib{f: 0, g: 1}
	for hi-lo > fib.g {
		fib.prev()
	}
	switch mi := lo + fib.f - 1; {
	case e.(float64) < T._elem[mi].(float64):
		hi = mi
	case T._elem[mi].(float64) < e.(float64):
		lo = mi + 1
	default:
		return mi
	}
	return -1
}

type Fib struct {
	f int
	g int
}

func (F Fib) prev() {
	F.g = F.g + F.f
	F.f = F.g - F.f
}

/* ----- Accessible Interface ----- */
func (T *Vector) Remove(r Rank) interface{} {
	e := T._elem[r]
	_ = T.removeLH(r, r+1)
	return e
}

func (T *Vector) removeLH(lo Rank, hi Rank) int {
	if lo == hi {
		return 0
	}
	T._elem = append(T._elem[:lo], T._elem[hi:]...)
	T._size -= hi - lo
	T.shrink()
	return hi - lo
}

func (T *Vector) Insert(r Rank, e interface{}) {
	T._size++
	T.expand()
	behindElem := append([]interface{}{e}, T._elem[r:]...)
	T._elem = append(T._elem[:r], behindElem...)
}

func (T *Vector) InsertEnd(e interface{}) {
	T._size++
	T.expand()
	T._elem = append(T._elem, e)
}

func (T *Vector) sortLH(lo Rank, hi Rank) {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(4) {
	case 0:
		T.bubbleSort(lo, hi)
	case 1:
		T.mergeSort(lo, hi)
	case 2:
		T.heapSort(lo, hi)
	case 3:
		T.quickSort(lo, hi)
		//default:
		//T.selectionSort(lo, hi)
	}
}

func (T *Vector) Sort() {
	T.sortLH(0, T._size)
}

//func (T *Vector) unsortLH(lo Rank, hi Rank) {}

//func (T *Vector) unsort() {
//	T.unsortLH(0, T._size)
//}

func (T *Vector) Deduplicate() int {
	// disorder vector uniquify
	oldSize := T._size
	i := 1
	for i < T._size {
		if T.findLH(T._elem[i], 0, i) < 0 {
			i++
		} else {
			T.Remove(i)
		}
	}
	return oldSize - T._size
}

func (T *Vector) uniquifyBad() int {
	oldSize := T._size
	i := 0
	for i < T._size {
		if T._elem[i] == T._elem[i+1] {
			T.Remove(i + 1)
		} else {
			i++
		}
	}
	return oldSize - T._size
}

func (T *Vector) Uniquify() int {
	var i, j int
	for i, j = 0, 1; j < T._size; j++ {
		if T._elem[i] != T._elem[j] {
			i++
			T._elem[i] = T._elem[j]
		}
	}
	i++
	T._size = i
	T.shrink()
	return j - i
}

/* ----- Protected Interface ----- */
func (T *Vector) CopyFrom(A []interface{}, lo Rank, hi Rank) {
	T._elem = make([]interface{}, 2*(hi-lo))
	T._size = 0
	for lo < hi {
		T._elem[T._size] = A[lo]
		lo++
		T._size++
	}
}

func (T *Vector) expand() {
	if T._size < T._capacity {
		return
	}
	if T._capacity < DefaultCapacity {
		T._capacity = DefaultCapacity
	}
	T._capacity = T._capacity << 1
}

func (T *Vector) shrink() {
	if T._size >= T._capacity/2 {
		return
	}
	if T._capacity < DefaultCapacity {
		T._capacity = DefaultCapacity
	}
	T._capacity = T._capacity >> 1
}

func (T *Vector) bubble(lo Rank, hi Rank) bool {
	sorted := true
	lo++
	for lo < hi {
		if T._elem[lo-1].(float64) > T._elem[lo].(float64) {
			sorted = false
			T._elem[lo-1], T._elem[lo] = T._elem[lo], T._elem[lo-1]
		}
	}
	return sorted
}

func (T *Vector) bubble2(lo Rank, hi Rank) Rank {
	last := lo
	for lo < hi {
		if T._elem[lo-1].(float64) > T._elem[lo].(float64) {
			last = lo
			T._elem[lo-1], T._elem[lo] = T._elem[lo], T._elem[lo-1]
		}
	}
	return last
}

func (T *Vector) bubbleSort(lo Rank, hi Rank) {
	// type 1
	//for !T.bubble(lo, hi) { // bubble one by one
	//	hi--
	//}

	for lo < hi {
		hi = T.bubble2(lo, hi)
	}
}

//func (T *Vector) max(lo Rank, hi Rank) Rank {}

//func (T *Vector) selectionSort(lo Rank, hi Rank) {}

func (T *Vector) merge(lo Rank, mi Rank, hi Rank) {
	lb, lc := mi-lo, hi-mi
	A := T._elem[lo:hi]
	B := make([]interface{}, lb)
	_ = copy(B, T._elem[lo:mi])
	C := T._elem[mi:hi]

	for i, j, k := 0, 0, 0; j < lb; {
		if lc <= k || B[j].(float64) <= C[k].(float64) {
			A[i] = B[j]
			i++
			j++
		}
		if k < lc && C[k].(float64) < B[j].(float64) {
			A[i] = C[k]
			i++
			k++
		}
	}
}

func (T *Vector) mergeSort(lo Rank, hi Rank) {
	if hi-lo < 2 {
		return
	}
	mi := (hi + lo) >> 1
	T.mergeSort(lo, mi)
	T.mergeSort(mi, hi)
	T.merge(lo, mi, hi)
}

func (T *Vector) partition1(lo Rank, hi Rank) Rank {
	pivot := T._elem[lo]
	for lo+1 < hi {
		for pivot.(int) <= T._elem[hi-1].(int) {
			hi--
		}
		T._elem[lo] = T._elem[hi-1]
		for T._elem[lo].(int) <= pivot.(int) {
			lo++
		}
		T._elem[hi-1] = T._elem[lo]
	}
	T._elem[lo] = pivot
	return lo
}

func (T *Vector) partition(lo Rank, hi Rank) Rank {
	rand.Seed(time.Now().UnixNano())
	T.Swap(lo, lo+rand.Intn(hi-lo)) //随即交换
	pivot := T._elem[lo]
	mi := lo
	for k := lo + 1; k <= hi; k++ {
		if T._elem[k].(int) < pivot.(int) {
			mi++
			T.Swap(mi, k)
		}
	}
	T.Swap(lo, mi)
	return mi
}

func (T *Vector) quickSort(lo Rank, hi Rank) {
	if hi-lo < 2 {
		return
	}
	mi := T.partition(lo, hi-1)
	T.quickSort(lo, mi)
	T.quickSort(mi+1, hi)
}

func (T *Vector) heapSort(lo Rank, hi Rank) {
	var H PriorityQueue.PQComplHeap
	H.CopyFrom(T._elem, lo, hi)
	for !H.Empty() {
		hi--
		T._elem[hi] = H.DelMax()
	}
}

func (T *Vector) majority(maj interface{}) bool {
	maj = T.majEleCandidate()
	return T.majEleCheck(maj)
}

func (T *Vector) majEleCandidate() interface{} {
	var maj interface{}
	for c, i := 0, 0; i < T.Size(); i++ {
		if c == 0 {
			maj = T._elem[i]
			c = 1
		} else {
			if maj == T._elem[i] {
				c++
			} else {
				c--
			}
		}
	}
	return maj
}

func (T *Vector) majEleCheck(maj interface{}) bool {
	var c int = 0
	for i := 0; i < T.Size(); i++ {
		if T._elem[i] == maj {
			c++
		} else {
			c--
		}
	}
	if c > 0 {
		return true
	} else {
		return false
	}
}

func (T *Vector) quickSelect(k Rank) {
	for lo, hi := 0, T.Size()-1; lo < hi; {
		i, j := lo, hi
		pivot := T._elem[lo]
		for i < j {
			for i < j && pivot.(int) <= T._elem[j].(int) {
				j--
			}
			T._elem[i] = T._elem[j]
			for i < j && T._elem[j].(int) <= pivot.(int) {
				i++
			}
			T._elem[j] = T._elem[i]
		}
		T._elem[i] = pivot
		if k <= i {
			hi = i - 1
		}
		if i <= k {
			lo = i + 1
		}
	}
}

/* ----- traverse ----- */

func (T *Vector) Traverse(fun1 func(interface{})) {
	for item := range T._elem {
		fun1(item)
	}
}

func (T *Vector) Get(r Rank) interface{} {
	return T._elem[r]
}

func (T *Vector) Put(r Rank, e interface{}) {
	T._elem[r] = e
}

func (T *Vector) Swap(i, j Rank) {
	T._elem[i], T._elem[j] = T._elem[j], T._elem[i]
}
