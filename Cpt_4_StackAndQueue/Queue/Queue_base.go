package Queue

import "datastructure/DataStructure_Go/Cpt_3_List/List"

type Queue struct {
	List.List
}

func (Q *Queue) Enqueue(e interface{}) {
	Q.InsertAsLast(e)
}

func (Q *Queue) Dequeue() interface{} {
	return Q.Remove(Q.First())
}

func (Q *Queue) Front() interface{} {
	return Q.First().Data
}

func (Q *Queue) Rear() interface{} {
	return Q.Last().Data
}
