package Tree

import (
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Queue"
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Stack"
)

type BinNodePosi = *BinNode

type BinNode struct {
	parent, lChild, rChild *BinNode
	data                   interface{}
	height                 int
}

type BinTree struct {
	_size int
	_root BinNodePosi
}

/* --------- BinNode 方法 --------- */
//后代规模数
func (T *BinNode) Size() int {
	s := 1
	if T.lChild != nil {
		s += T.lChild.Size()
	}
	if T.rChild != nil {
		s += T.rChild.Size()
	}
	return s
}

func (T *BinNode) InsertAsLC(e interface{}) BinNodePosi {
	newNode := BinNode{parent: T, data: e}
	T.lChild = &newNode
	return &newNode
}

func (T *BinNode) InsertAsRC(e interface{}) BinNodePosi {
	newNode := BinNode{parent: T, data: e}
	T.rChild = &newNode
	return &newNode
}

//(中序遍历意义下)当前节点的直接后继
//func (T *BinNode) Succ() BinNodePosi {}

/* --------- BinTree 方法 --------- */
//virtual
//更新x的高度
func (T *BinTree) updateHeight(x BinNodePosi) int {
	x.height = 1 + max(x.lChild.stature(), x.rChild.stature())
	return x.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (T *BinNode) stature() int {
	if T == nil {
		return -1
	}
	return T.height
}

//更新x及祖先的高度
func (T *BinTree) updateHeightAbove(x BinNodePosi) {
	for x != nil { //可以优化: 高度未变即可终止
		T.updateHeight(x)
		x = x.parent
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

/* -- 子树接入删除和分离 -- */
func (T *BinTree) InsertAsRC(x BinNodePosi, e interface{}) BinNodePosi {
	T._size++
	x.InsertAsRC(e)
	T.updateHeightAbove(x)
	return x.rChild
}

func (T *BinTree) InsertAsLC(x BinNodePosi, e interface{}) BinNodePosi {
	T._size++
	x.InsertAsLC(e)
	T.updateHeightAbove(x)
	return x.lChild
}

/* -- 遍历 -- */
//先序 递归形式
func (T *BinTree) TraversePre(x BinNodePosi, visit func(interface{})) {
	if x == nil {
		return
	}
	visit(x.data)
	T.TraversePre(x.lChild, visit)
	T.TraversePre(x.rChild, visit)
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
		visit(x.data)
		//右孩子先入后出
		if x.rChild != nil {
			S.Push(x.rChild)
		}
		//左孩子后入先出
		if x.lChild != nil {
			S.Push(x.lChild)
		}
		//注意上面两个次序
	}
}

//先序 迭代形式2 最终形式
func (T *BinTree) visitAlongLeftBranch(x BinNodePosi, visit func(interface{}), S *Stack.Stack) {
	for x != nil {
		visit(x.data)
		S.Push(x.rChild)
		x = x.lChild
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
	T.TraverseIn(x.lChild, visit)
	visit(x.data)
	T.TraverseIn(x.rChild, visit)
}

//中序 迭代形式
func (T *BinTree) goAlongLeftBranch(x BinNodePosi, S *Stack.Stack) {
	for x != nil {
		S.Push(x)
		x = x.lChild
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
		visit(x.data)
		x = x.rChild
	}
}

//后序
//后序 递归形式
func (T *BinTree) TraversePost(x BinNodePosi, visit func(interface{})) {
	if x == nil {
		return
	}
	T.TraversePost(x.lChild, visit)
	T.TraversePost(x.rChild, visit)
	visit(x.data)
}

//后序 迭代形式
func (T *BinTree) gotoHLVFL(S *Stack.Stack) {
	for x := S.Top().(BinNodePosi); x != nil; {
		switch {
		case (x.lChild != nil) && (x.rChild != nil):
			S.Push(x.rChild)
			S.Push(x.lChild)
		case (x.lChild != nil) && (x.rChild == nil):
			S.Push(x.lChild)
		default:
			S.Push(x.rChild)
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
		if S.Top().(BinNodePosi) != x.parent {
			T.gotoHLVFL(&S)
		}
		x = S.Pop().(BinNodePosi)
		visit(x.data)
	}
}

//层次
func (T *BinTree) TravLevel(visit func(interface{})) {
	var Q Queue.Queue
	var x BinNodePosi
	Q.Enqueue(T._root)
	for !Q.Empty() {
		x = Q.Dequeue().(BinNodePosi)
		visit(x.data)
		if x.lChild != nil {
			Q.Enqueue(x.lChild)
		}
		if x.rChild != nil {
			Q.Enqueue(x.rChild)
		}
	}
}
