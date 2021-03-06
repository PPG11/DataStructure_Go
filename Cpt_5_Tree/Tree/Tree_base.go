package Tree

import (
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Queue"
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Stack"
)

type RBColor = bool

const (
	RBBLACK RBColor = true
	RBRED   RBColor = false
)

type BinNodePosi = *BinNode

type BinNode struct {
	Parent, LChild, RChild *BinNode
	Data                   interface{}
	Height                 int
	Color                  RBColor
	Npl                    int
}

type BinTree struct {
	_size int
	_root BinNodePosi
}

/* --------- BinNode 方法 --------- */
//后代规模数
func (T *BinNode) Size() int {
	s := 1
	if T.LChild != nil {
		s += T.LChild.Size()
	}
	if T.RChild != nil {
		s += T.RChild.Size()
	}
	return s
}

func (T *BinTree) SizeAdd(i int) {
	T._size += i
}

func (T *BinNode) GetData() interface{} {
	return T.Data
}

func (T *BinNode) InsertAsLC(e interface{}) BinNodePosi {
	newNode := BinNode{Parent: T, Data: e}
	T.LChild = &newNode
	return &newNode
}

func (T *BinNode) InsertAsRC(e interface{}) BinNodePosi {
	newNode := BinNode{Parent: T, Data: e}
	T.RChild = &newNode
	return &newNode
}

//(中序遍历意义下)当前节点的直接后继
func (T *BinNode) Succ() BinNodePosi {
	s := T
	if T.RChild != nil {
		s = T.RChild
		for s.LChild != nil {
			s = s.LChild
		}
	} else {
		for s.Parent.RChild == s {
			s = s.Parent
		}
		s = s.Parent
	}
	return s
}

/* --------- BinTree 方法 --------- */
//virtual
//更新x的高度
func (T *BinTree) UpdateHeight(x BinNodePosi) int {
	x.Height = 1 + max(x.LChild.Stature(), x.RChild.Stature())
	return x.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (T *BinNode) Stature() int {
	if T == nil {
		return -1
	}
	return T.Height
}

//更新x及祖先的高度
func (T *BinTree) UpdateHeightAbove(x BinNodePosi) {
	for x != nil { //可以优化: 高度未变即可终止
		T.UpdateHeight(x)
		x = x.Parent
	}
}

/* -- 三个基本方法 -- */
func (T *BinTree) Size() int {
	return T._size
}

func (T *BinTree) Empty() bool {
	if T._root == nil {
		return true
	}
	return false
}

func (T *BinTree) Root() BinNodePosi {
	return T._root
}

func (T *BinTree) SetRoot(x BinNodePosi) {
	T._root = x
}

/* -- 子树接入删除和分离 -- */
func (T *BinTree) InsertAsRC(x BinNodePosi, e interface{}) BinNodePosi {
	T._size++
	x.InsertAsRC(e)
	T.UpdateHeightAbove(x)
	return x.RChild
}

func (T *BinTree) InsertAsLC(x BinNodePosi, e interface{}) BinNodePosi {
	T._size++
	x.InsertAsLC(e)
	T.UpdateHeightAbove(x)
	return x.LChild
}

/* -- 遍历 -- */
//先序 递归形式
func (T *BinTree) TraversePre(x BinNodePosi, visit func(interface{})) {
	if x == nil {
		return
	}
	visit(x.Data)
	T.TraversePre(x.LChild, visit)
	T.TraversePre(x.RChild, visit)
}

//先序 迭代形式1
//但是不易推广到中序和后序
func (T *BinTree) TravPreI1(x BinNodePosi, visit func(interface{})) {
	var S Stack.Stack
	if x != nil {
		S.Push(x)
	}
	for !S.Empty() {
		x = S.Pop().(BinNodePosi)
		visit(x.Data)
		//右孩子先入后出
		if x.RChild != nil {
			S.Push(x.RChild)
		}
		//左孩子后入先出
		if x.LChild != nil {
			S.Push(x.LChild)
		}
		//注意上面两个次序
	}
}

//先序 迭代形式2 最终形式
func (T *BinTree) visitAlongLeftBranch(x BinNodePosi, visit func(interface{}), S *Stack.Stack) {
	for x != nil {
		visit(x.Data)
		S.Push(x.RChild)
		x = x.LChild
	}
}

func (T *BinTree) TravPreI2(x BinNodePosi, visit func(interface{})) {
	var S Stack.Stack
	for {
		T.visitAlongLeftBranch(x, visit, &S)
		if S.Empty() {
			break
		}
		x = S.Pop().(BinNodePosi)
	}
}

//中序
//中序 递归形式
func (T *BinTree) TraverseIn(x BinNodePosi, visit func(interface{})) {
	if x == nil {
		return
	}
	T.TraverseIn(x.LChild, visit)
	visit(x.Data)
	T.TraverseIn(x.RChild, visit)
}

//中序 迭代形式
func (T *BinTree) goAlongLeftBranch(x BinNodePosi, S *Stack.Stack) {
	for x != nil {
		S.Push(x)
		x = x.LChild
	}
}

func (T *BinTree) TravInI(x BinNodePosi, visit func(interface{})) {
	var S Stack.Stack
	for {
		T.goAlongLeftBranch(x, &S)
		if S.Empty() {
			break
		}
		x = S.Pop().(BinNodePosi)
		visit(x.Data)
		x = x.RChild
	}
}

//后序
//后序 递归形式
func (T *BinTree) TraversePost(x BinNodePosi, visit func(interface{})) {
	if x == nil {
		return
	}
	T.TraversePost(x.LChild, visit)
	T.TraversePost(x.RChild, visit)
	visit(x.Data)
}

//后序 迭代形式
func (T *BinTree) gotoHLVFL(S *Stack.Stack) {
	for x := S.Top().(BinNodePosi); x != nil; {
		switch {
		case (x.LChild != nil) && (x.RChild != nil):
			S.Push(x.RChild)
			S.Push(x.LChild)
		case (x.LChild != nil) && (x.RChild == nil):
			S.Push(x.LChild)
		default:
			S.Push(x.RChild)
		}
		x = S.Top().(BinNodePosi)
	}
	S.Pop()
}

func (T *BinTree) TravPostI(x BinNodePosi, visit func(interface{})) {
	var S Stack.Stack
	if x != nil {
		S.Push(x)
	}
	for !S.Empty() {
		if S.Top().(BinNodePosi) != x.Parent {
			T.gotoHLVFL(&S)
		}
		x = S.Pop().(BinNodePosi)
		visit(x.Data)
	}
}

//层次
func (T *BinTree) TravLevel(visit func(interface{})) {
	var Q Queue.Queue
	var x BinNodePosi
	Q.Enqueue(T._root)
	for !Q.Empty() {
		x = Q.Dequeue().(BinNodePosi)
		visit(x.Data)
		if x.LChild != nil {
			Q.Enqueue(x.LChild)
		}
		if x.RChild != nil {
			Q.Enqueue(x.RChild)
		}
	}
}

//一些基础方法
func (T *BinTree) IsRoot(x BinNodePosi) bool {
	if x.Parent == nil {
		return true
	}
	return false
}

func (T *BinTree) IsLChild(x BinNodePosi) bool {
	if (!T.IsRoot(x)) && (x == x.Parent.LChild) {
		return true
	}
	return false
}

func (T *BinTree) IsRChild(x BinNodePosi) bool {
	if (!T.IsRoot(x)) && (x == x.Parent.RChild) {
		return true
	}
	return false
}

func (T *BinTree) AttachAsLC(x BinNodePosi, S BinNodePosi) BinNodePosi {
	x.LChild = S
	if x.LChild != nil {
		x.LChild.Parent = x
	}
	T.SizeAdd(S.Size())
	T.UpdateHeight(x)
	return x
}

func (T *BinTree) AttachAsRC(x BinNodePosi, S BinNodePosi) BinNodePosi {
	x.RChild = S
	if x.RChild != nil {
		x.RChild.Parent = x
	}
	T.SizeAdd(S.Size())
	T.UpdateHeight(x)
	return x
}

//来自父亲的引用
func (T *BinTree) FromParentTo(x BinNodePosi) BinNodePosi {
	switch {
	case T.IsRoot(x):
		return T._root
	case T.IsLChild(x):
		return x.Parent.LChild
	default:
		return x.Parent.RChild
	}
}

func (T *BinTree) Uncle(x BinNodePosi) BinNodePosi {
	if p := x.Parent; T.IsLChild(p) {
		return p.Parent.RChild
	} else {
		return p.Parent.LChild
	}
}

func (T *BinTree) Sibling(x BinNodePosi) BinNodePosi {
	if T.IsLChild(x) {
		return x.Parent.RChild
	} else {
		return x.Parent.LChild
	}
}
