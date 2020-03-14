package Stack

import "datastructure/DataStructure_Go/Cpt_2_Vector/Vector"

type Stack struct {
	Vector.Vector
}

func (T *Stack) Push(e interface{}) {
	T.Insert(T.Size(), e)
}

func (T *Stack) Pop() interface{} {
	return T.Remove(T.Size() - 1)
}

func (T *Stack) Top() interface{} {
	return T.Get(T.Size() - 1)
}
