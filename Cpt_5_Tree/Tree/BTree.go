package Tree

import (
	"datastructure/DataStructure_Go/Cpt_2_Vector/Vector"
)

type BTNodePosi = *BTNode

type BTNode struct {
	parent *BTNode
	key    Vector.Vector
	child  Vector.Vector
}

func (T *BTNode) Init() {
	T.parent = nil
	T.child.Insert(0, nil)
}

func (T *BTNode) InitWithParam(e int, lc BTNodePosi, rc BTNodePosi) {
	T.parent = nil
	T.key.Insert(0, e)
	T.child.Insert(0, lc)
	T.child.Insert(1, rc)
	if lc != nil {
		lc.parent = T
	}
	if rc != nil {
		rc.parent = T
	}
}

type BTree struct {
	_size  int
	_order int
	_root  BTNodePosi
	_hot   BTNodePosi
}

func (T *BTree) Search(e int) BTNodePosi {
	v := T._root
	T._hot = nil
	for v != nil {
		r := v.key.Search(e)
		if 0 <= r && e == v.key.Get(r).(int) { //找到
			return v
		}
		T._hot = v
		v = v.child.Get(r + 1).(BTNodePosi)
	}
	return nil
}

func (T *BTree) Insert(e int) bool {
	v := T.Search(e)
	if v != nil {
		return false
	}
	r := T._hot.key.Search(e) //确认插入的位置
	T._hot.key.Insert(r+1, e)
	T._hot.child.Insert(r+2, nil) //创建新的空子树
	T._size++
	T.solveOverflow(T._hot) //如果上溢则分裂
	return true
}

func (T *BTree) solveOverflow(v BTNodePosi) {
	if T._order >= v.child.Size() {
		return
	}
	s := T._order / 2
	var u BTNode
	for j := 0; j < T._order-s-1; j++ {
		u.child.Insert(j, v.child.Remove(s+1))
		u.key.Insert(j, v.key.Remove(s+1))
	}
	u.child.Put(T._order-s-1, v.child.Remove(s+1))
	if u.child[0] != nil {
		for j := 0; j < T._order-s; j++ {
			u.child[j].(BTNodePosi).parent = &u
		}
		p := v.parent
		if p == nil {
			var p BTNode
			T._root = &p
			p.child.Put(0, &v)
			v.parent = &p
		}
		r := 1 + p.key.Search(v.key.Get(0)) //p中指向u指针的大小
		p.key.Insert(r, v.key.Remove(s))    //轴点关键码上升
		p.child.Insert(r+1, u)              //新节点和父节点链接1
		u.parent = p                        //新节点和父节点链接2
		T.solveOverflow(p)
	}
}

func (T *BTree) Remove(e int) bool {
	v := T.Search(e)
	if v == nil {
		return false
	}
	r := v.key.Search(e)
	if v.child.Get(0) != nil {
		u := v.child.Get(r + 1).(BTNodePosi)
		for u.child.Get(0) != nil {
			u = u.child.Get(0).(BTNodePosi)
		}
		v.key.Put(r, u.key.Get(0))
		v = u
		r = 0
	}
	v.key.Remove(r)
	v.child.Remove(r + 1)
	T._size--
	T.solveUnderflow(v)
	return true
}

func (T *BTree) solveUnderflow(v BTNodePosi) {
	if (T._order+1)/2 <= v.child.Size() {
		return
	} //递归基 没有下溢
	p := v.parent
	if p == nil { //递归基 到达根节点 没有孩子下限
		if v.key.Size() == 0 && v.child.Get(0) != nil {
			//删除树根后没关键码 但是却有非空孩子
			T._root = v.child.Get(0).(BTNodePosi)
			v.child.Put(0, nil)
		} //树高度下降1
		return
	}
	r := 0
	if p.child.Get(r).(BTNodePosi) != v {
		r++
	}
	//先确定v是p的第r个孩子--此时v可能不含关键码, 所以不能通过关键码查找
	//在实现孩子指针判断后 可以直接调用 Vector.Find()定位
	//-----情况1: 向左兄弟借关键码
	if 0 < r { //如果v不是p的第一个孩子则
		ls := p.child.Get(r - 1).(BTNodePosi) //左侧兄弟必然存在
		if (T._order+1)/2 < ls.child.Size() { //左侧兄弟够胖
			v.key.Insert(0, p.key.Get(r-1).(int))                //p借一个关键码给v 作为最小关键码
			p.key.Put(r-1, ls.key.Remove(ls.key.Size()-1).(int)) //ls的最大关键码转入p
			v.child.Insert(0, ls.child.Remove(ls.child.Size()-1).(BTNodePosi))
			//同时ls右侧孩子过继给v
			if v.child.Get(0) != nil {
				v.child.Get(0).(BTNodePosi).parent = v
			}
			return
		}
	} // 至此, 左兄弟要么空, 要么太瘦
	//-----情况2: 向右兄弟借关键码
	if p.child.Size()-1 > r { //若v不是p的最后一个孩子
		rs := p.child.Get(r + 1).(BTNodePosi) //右侧兄弟必然存在
		if (T._order+1)/2 < rs.child.Size() { //右侧兄弟够胖
			v.key.Insert(v.key.Size(), p.key.Get(r).(BTNodePosi)) //p借一个关键码给v 作为最大关键码
			p.key.Put(r, rs.key.Remove(0).(int))                  //rs的最小关键码转入p
			v.child.Insert(v.child.Size(), rs.child.Remove(0).(BTNodePosi))
			//同时rs最左侧孩子过继给v
			if v.child.Get(v.child.Size()-1) != nil {
				v.child.Get(v.child.Size() - 1).(BTNodePosi).parent = v
			}
			return //至此, 通过左旋完成当前层及所有层下溢处理
		}
	} //至此 右兄弟要么空 要么太瘦
	//-----情况3: 左右兄弟要么空(但不可能同时), 要么都太瘦---合并
	if 0 < r { //与左兄弟合并
		ls := p.child.Get(r - 1).(BTNodePosi) //左兄弟必存在
		ls.key.Insert(ls.key.Size(), p.key.Remove(r-1).(int))
		p.child.Remove(r)
		//p的第 r-1 个关键码转入 ls, v不再是p的第一个孩子
		ls.child.Insert(ls.child.Size(), v.child.Remove(0).(BTNodePosi))
		if ls.child.Get(ls.child.Size()-1) != nil { //v最左侧过继给ls作右孩子
			ls.child.Get(ls.child.Size() - 1).(BTNodePosi).parent = ls
		}
		for !v.key.Empty() { //其余关键码和孩子转入ls
			ls.key.Insert(ls.key.Size(), v.key.Remove(0).(int))
			ls.child.Insert(ls.child.Size(), v.child.Remove(0).(BTNodePosi))
			if ls.child.Get(ls.child.Size()-1) != nil {
				ls.child.Get(ls.child.Size() - 1).(BTNodePosi).parent = ls
			}
		}
	} else { //与右兄弟合并
		rs := p.child.Get(r + 1).(BTNodePosi) //右兄弟必存在
		rs.key.Insert(0, p.key.Remove(r).(int))
		p.child.Remove(r)
		//p的第 r 个关键码转入 rs, v不再是p的第r个孩子
		rs.child.Insert(0, v.child.Remove(v.child.Size()-1).(BTNodePosi))
		if rs.child.Get(0) != nil { //v最左侧过继给rs作左孩子
			rs.child.Get(0).(BTNodePosi).parent = rs
		}
		for !v.key.Empty() { //其余关键码和孩子转入rs
			rs.key.Insert(0, v.key.Remove(v.key.Size()-1).(int))
			rs.child.Insert(0, v.child.Remove(v.child.Size()-1).(BTNodePosi))
			if rs.child.Get(0) != nil {
				rs.child.Get(0).(BTNodePosi).parent = rs
			}
		}
	}
	T.solveUnderflow(p) //上升一层 如有必要继续分裂 最多O(logn)
	return
}
