package Tree

import "datastructure/DataStructure_Go/Cpt_7_BinarySearchTree/BST"

type RedBlack struct {
	BST.BST
}

func (T *RedBlack) Insert(e int) BinNodePosi {
	var isLeft *bool

	if x := T.Search(e, isLeft); x != nil {
		return x
	}
	x := &BinNode{Parent: T.Hot(), Data: e, Height: -1}
	T._size++
	//如果有必要, 需要双红修正
	T.solveDoubleRed(x)
	if x != nil {
		return x
	}
	return T.Hot().Parent
}

/* --------- RBTreeNode 方法 --------- */
func (T *BinNode) IsBlack() bool {
	if T == nil || T.Color == RBBLACK {
		return true
	}
	return false
}

func (T *BinNode) IsRed() bool {
	return !T.IsBlack()
}

func (T *RedBlack) BlackHeightUpdate(x BinNodePosi) bool {
	if x.LChild.Stature() == x.RChild.Stature() {
		switch {
		case x.IsRed() && x.Height == x.LChild.Stature():
			return true
		case x.IsBlack() && x.Height == (x.LChild.Stature()+1):
			return true
		}
	}
	return false
}

func (T *RedBlack) updateHeight(x BinNodePosi) int {
	x.Height = max(x.LChild.Stature(), x.RChild.Stature())
	if x.IsBlack() {
		x.Height++
	}
	return x.Height
}

func (T *RedBlack) solveDoubleRed(x *BinNode) {
	if T.IsRoot(x) { //到树根转黑, 调整黑树高度
		T.Root().Color = RBBLACK
		T.Root().Height++
		return
	} //否则父亲必然存在
	p := x.Parent
	if p.IsBlack() {
		return
	} //递归基 父亲黑则无需双红修正
	g := p.Parent
	u := T.Uncle(x)
	if u.IsBlack() {
		if T.IsLChild(x) == T.IsLChild(p) {
			p.Color = RBBLACK
		} else {
			x.Color = RBBLACK
		}
		x.Color = RBRED

		gg := g.Parent
		*g = *T.RotateAt(x)
		r := g
		r.Parent = gg
	} else {
		p.Color = RBBLACK
		p.Height++
		u.Color = RBBLACK
		u.Height++
		if !T.IsRoot(g) {
			g.Color = RBRED
		}
		T.solveDoubleRed(g)
	}
}

func (T *RedBlack) Remove(e int) bool {
	var isLeft *bool
	x := T.Search(e, isLeft)
	if x == nil {
		return false
	}
	r := T.RemoveAt(x, isLeft)
	T._size--
	if T._size == 0 {
		return true
	}
	if T.Hot() == nil { //删除的是根节点
		T._root.Color = RBBLACK
		T.UpdateHeight(T._root)
		return true
	}
	if T.BlackHeightUpdate(T.Hot()) {
		return true
	}
	T.solveDoubleBlack(r)
	return true
}

/*
* 双黑调整算法
* 三大类共四种情况
* BB-1 :2次颜色反转, 2次黑高度更新, 1~2次旋转, 不递归
* BB-2R:2次颜色反转, 2次黑高度更新,  0 次旋转, 不递归
* BB-2B:1次颜色反转, 1次黑高度更新,  0 次旋转, 要递归
* BB-3 :2次颜色反转, 2次黑高度更新,  1 次旋转, 转为BB-1或BB-2R
 */
func (T *RedBlack) solveDoubleBlack(r BinNodePosi) {
	var p, s BinNodePosi
	if r != nil {
		p = r.Parent
	} else {
		p = T.Hot()
	}
	if p == nil {
		return
	}
	if r == p.LChild {
		s = p.RChild
	} else {
		s = p.LChild
	}
	if s.IsBlack() {
		var t BinNodePosi
		if s.RChild.IsRed() {
			t = s.RChild
		}
		if s.LChild.IsRed() {
			t = s.LChild
		}
		if t != nil { //黑s有红孩子 BB-1
			oldColor := p.Color
			*p = *T.RotateAt(t)
			b := p
			if b.LChild != nil {
				b.LChild.Color = RBBLACK
				T.UpdateHeight(b.LChild)
			}
			if b.RChild != nil {
				b.RChild.Color = RBBLACK
				T.UpdateHeight(b.RChild)
			}
			b.Color = oldColor
			T.UpdateHeight(b)
		} else { //黑s无红孩子
			s.Color = RBRED
			s.Height--
			if p.IsRed() {
				p.Color = RBBLACK
			} else {
				p.Height--
				T.solveDoubleBlack(p)
			}
		}
	} else { //兄弟 s 为红: BB-3
		s.Color = RBBLACK
		p.Color = RBRED
		var t BinNodePosi
		if s.IsRed() {
			t = s.LChild
		} else {
			t = s.RChild
		}
		T.SetHot(p)
		*p = *T.RotateAt(t)
		T.solveDoubleBlack(r)
	}
}
