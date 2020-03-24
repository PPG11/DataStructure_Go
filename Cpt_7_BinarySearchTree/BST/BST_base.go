package BST

import "datastructure/DataStructure_Go/Cpt_5_Tree/Tree"

/*
BST中序遍历 单调非降
*/
type key = int

type value = interface{}

type Entry struct {
	K key
	V value
}

func (T *Entry) init(k key, v value) {
	T.K = k
	T.V = v
}

func (T *Entry) copy(e Entry) {
	T.K = e.K
	T.V = e.V
}

type BST struct {
	Tree.BinTree
	_hot Tree.BinNodePosi
}

func (T *BST) Search(e key, isLeft *bool) Tree.BinNodePosi {
	T._hot = nil
	return T.searchIn(T.Root(), e, isLeft)
}

//递归版
func (T *BST) searchInRec(v Tree.BinNodePosi, e key, isLeft *bool) Tree.BinNodePosi {
	//v 当前(子)树根
	//e 目标关键码
	//hot 记忆热点
	if v == nil || e == v.Data.(int) {
		return v
	}
	T._hot = v
	if e < v.Data.(int) {
		*isLeft = true
		return T.searchInRec(v.LChild, e, isLeft)
	} else {
		*isLeft = false
		return T.searchInRec(v.RChild, e, isLeft)
	}
}

//迭代版
func (T *BST) searchIn(v Tree.BinNodePosi, e key, isLeft *bool) Tree.BinNodePosi {
	//hot 记忆热点
	//e 目标关键码
	//v 当前(子)树根
	for {
		if v == nil || e == v.Data.(int) {
			return v
		}
		T._hot = v
		if e < v.Data.(int) {
			*isLeft = true
			v = v.LChild
		} else {
			*isLeft = false
			v = v.RChild
		}
	}
}

func (T *BST) Insert(e int) Tree.BinNodePosi {
	var isLeft *bool
	x := T.Search(e, isLeft)
	var newNode Tree.BinNode
	if x == nil {
		newNode.Data = e
		newNode.Parent = T._hot
		if *isLeft {
			T._hot.LChild = &newNode
		} else {
			T._hot.RChild = &newNode
		}
		T.SizeAdd(1)
		T.UpdateHeightAbove(&newNode)
	}
	return &newNode
}

func (T *BST) Remove(e int) bool {
	var isLeft *bool
	x := T.Search(e, isLeft)
	if x == nil {
		return false
	}
	T.removeAt(x, isLeft)
	T.SizeAdd(-1)
	T.UpdateHeightAbove(T._hot)
	return true
}

func (T *BST) removeAt(x Tree.BinNodePosi, isLeft *bool) Tree.BinNodePosi {
	//w := x
	var succ Tree.BinNodePosi = nil
	switch {
	case x.LChild == nil && *isLeft:
		succ = x.RChild
		T._hot.LChild = succ
	case x.LChild == nil && !*isLeft:
		succ = x.RChild
		T._hot.RChild = succ
	case x.RChild == nil && *isLeft:
		succ = x.LChild
		T._hot.LChild = succ
	case x.RChild == nil && !*isLeft:
		succ = x.LChild
		T._hot.RChild = succ
	default: //左右子树并存
		w := x.Succ()
		w.Data, x.Data = x.Data, w.Data
		u := w.Parent
		succ = w.RChild
		if u.Data == x.Data {
			u.RChild = succ
		} else {
			u.LChild = succ
		}
		T._hot = u
	}
	if succ != nil {
		succ.Parent = T._hot
	}
	return succ
}

/* ----- AVL ----- */

func (T *BST) Balanced(x Tree.BinNodePosi) bool { //理想平衡
	return x.LChild.Stature() == x.RChild.Stature()
}

func (T *BST) BalFac(x Tree.BinNodePosi) int { //平衡因子
	return x.LChild.Stature() - x.RChild.Stature()
}

func (T *BST) AvlBalanced(x Tree.BinNodePosi) bool { //AVL平衡条件
	return (-2 < T.BalFac(x)) && (T.BalFac(x) < 2)
}

type AVL struct {
	BST
}

//一个节点的插入会引起有可能所有的祖先失衡
//一个节点的删除最多引起该节点的 parent 一个节点的失衡

//func (T *AVL) Search(x int) {}

func (T *AVL) Insert(e int) Tree.BinNodePosi {
	var isLeft *bool

	if x := T.Search(e, isLeft); x != nil {
		return x
	} //找到该 key 无法插入

	//如果没找到 key 则创建并插入
	var x Tree.BinNodePosi
	x.Data = e
	x.Parent = T._hot
	if *isLeft {
		T._hot.LChild = x
	} else {
		T._hot.RChild = x
	}
	T.SizeAdd(1)
	//xx := x

	//以下从 x 的 parent 出发检查祖先 g
	for g := x.Parent; g != nil; g = g.Parent {
		if !T.AvlBalanced(g) {
			// 发现失衡则调节
			//T.FromParentTo(g) = T.rotateAt(T.tallerChild(T.tallerChild(g)))
			*g = *T.rotateAt(T.tallerChild(T.tallerChild(g)))
			break
		} else { //否则代表没失衡
			T.UpdateHeight(g)
		}
	}
	return x
}

func (T *AVL) Remove(e int) bool {
	var isLeft *bool
	x := T.Search(e, isLeft)
	if x == nil {
		return false
	} //如果目标不存在

	//如果找到目标
	T.removeAt(x, isLeft)
	T.SizeAdd(-1)

	//以下 从hot向上检查
	for g := T._hot; g != nil; g = g.Parent {
		if !T.AvlBalanced(g) {
			*g = *T.rotateAt(T.tallerChild(T.tallerChild(g)))
			//T.FromParentTo(g) = T.rotateAt(T.tallerChild(T.tallerChild(g)))
			//g = T.FromParentTo(g)
		}
		T.UpdateHeight(g)
	}
	return true
}

func (T *AVL) tallerChild(x Tree.BinNodePosi) Tree.BinNodePosi {
	switch {
	case x.LChild.Stature() > x.RChild.Stature():
		return x.LChild
	case x.LChild.Stature() < x.RChild.Stature():
		return x.RChild
	default:
		if x == x.Parent.LChild {
			return x.LChild
		} else {
			return x.RChild
		}
	}
}

//R R
func (T *AVL) zagzag(g Tree.BinNodePosi) {
	newg := *g
	p := newg.RChild
	newg.RChild = p.LChild
	p.LChild.Parent = &newg

	p.Parent = newg.Parent

	newg.Parent = p
	p.LChild = &newg

	*g = *p
}

//L L
func (T *AVL) zigzig(g Tree.BinNodePosi) {
	newg := *g
	p := newg.LChild
	newg.LChild = p.RChild
	p.RChild.Parent = &newg

	p.Parent = newg.Parent

	newg.Parent = p
	p.RChild = &newg

	*g = *p
}

//L R
func (T *AVL) zigzag(g Tree.BinNodePosi) {
	newg := *g
	p := newg.RChild
	v := p.LChild
	////zig
	//p.LChild = v.RChild
	//v.RChild.Parent = p
	//
	//v.Parent = p.Parent
	//v.Parent.RChild = v
	//
	//v.RChild = p
	//p.Parent = v!!!!
	//p.Parent = g
	//
	////zag
	//newg.RChild = v.LChild
	//newg.RChild.Parent = &newg
	//
	//v.Parent = newg.Parent
	//
	//v.LChild = &newg
	//newg.Parent = v!!!!
	//newg.Parent = g
	//
	//*g = *v

	//new method
	newg.RChild = v.LChild
	newg.RChild.Parent = &newg

	p.LChild = v.RChild
	p.LChild.Parent = p

	v.LChild = &newg
	//newg.Parent = v
	newg.Parent = g

	v.RChild = p
	//p.Parent = v
	p.Parent = g

	*g = *v
}

//R L
func (T *AVL) zagzig(g Tree.BinNodePosi) {
	newg := *g
	p := newg.RChild
	v := p.LChild

	newg.LChild = v.RChild
	newg.LChild.Parent = &newg

	p.RChild = v.LChild
	p.RChild.Parent = p

	v.RChild = &newg
	newg.Parent = g

	v.LChild = p
	p.Parent = g

	*g = *v
}

func (T *AVL) rotateAt(v Tree.BinNodePosi) Tree.BinNodePosi {
	p := v.Parent
	g := p.Parent
	switch {
	case T.IsLChild(p) && T.IsLChild(v): //zig-zig
		p.Parent = g.Parent
		return T.connect34(v, p, g, v.LChild, v.RChild, p.RChild, g.RChild)
	case T.IsLChild(p) && !T.IsLChild(v): //zig-zag
		v.Parent = g.Parent
		return T.connect34(p, v, g, p.LChild, v.LChild, v.RChild, g.RChild)
	case !T.IsLChild(p) && T.IsLChild(v): //zag-zig
		v.Parent = g.Parent
		return T.connect34(g, v, p, g.LChild, v.LChild, v.RChild, p.RChild)
	//case !T.IsLChild(p) && !T.IsLChild(v): //zag-zag
	default:
		p.Parent = g.Parent
		return T.connect34(g, p, v, g.LChild, p.LChild, v.LChild, v.RChild)
	}
}

func (T *AVL) connect34(a, b, c, T0, T1, T2, T3 Tree.BinNodePosi) Tree.BinNodePosi {
	a.LChild = T0
	if T0 != nil {
		T0.Parent = a
	}
	a.RChild = T1
	if T1 != nil {
		T1.Parent = a
	}
	T.UpdateHeight(a)

	c.LChild = T2
	if T2 != nil {
		T2.Parent = c
	}
	c.RChild = T3
	if T3 != nil {
		T3.Parent = c
	}
	T.UpdateHeight(c)

	b.LChild = a
	a.Parent = b

	b.RChild = c
	c.Parent = b
	T.UpdateHeight(b)
	return b
}
