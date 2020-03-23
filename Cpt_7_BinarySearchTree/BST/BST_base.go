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

func (T *BST) Search(e key) Tree.BinNodePosi {
	T._hot = nil
	return T.searchInRec(T.Root(), e)
}

//递归版
func (T *BST) searchInRec(v Tree.BinNodePosi, e key) Tree.BinNodePosi {
	//v 当前(子)树根
	//e 目标关键码
	//hot 记忆热点
	if v == nil || e == v.Data.(int) {
		return v
	}
	T._hot = v
	if e < v.Data.(int) {
		return T.searchInRec(v.LChild, e)
	} else {
		return T.searchInRec(v.RChild, e)
	}
}
