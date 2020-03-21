package Graph

import (
	"datastructure/DataStructure_Go/Cpt_4_StackAndQueue/Queue"
	"math"
)

/* --------- Vertex --------- */
type VStatus uint8

const (
	UNDISCOVERED VStatus = 0
	DISCOVERED   VStatus = 1
	VISITED      VStatus = 2
)

type Vertex struct {
	data                interface{}
	inDegree, outDegree int     //入出度
	status              VStatus //遍历时候的状态
	dTime, fTime        int     //时间标签
	parent              int     //遍历树中的父节点
	priority            int     //遍历树中的优先级
}

func (T *Vertex) init(d interface{}) {
	T.data = d
	T.inDegree = 0
	T.outDegree = 0
	T.status = UNDISCOVERED
	T.dTime = -1
	T.fTime = -1
	T.parent = -1
	T.priority = math.MaxInt32
}

/* --------- Edge --------- */
type EStatus uint8

const (
	UNDETERMINED EStatus = 0
	TREE         EStatus = 1
	CROSS        EStatus = 2
	FORWARD      EStatus = 3
	BACKWARD     EStatus = 4
)

type Edge struct {
	data   interface{}
	weight int
	status EStatus
}

func (T *Edge) init(d interface{}, w int) {
	T.data = d
	T.weight = w
	T.status = UNDETERMINED
}

/* --------- Graph --------- */
type GMatrix struct {
	V []Vertex //点集 内部元素是 Vertex 结构
	E [][]Edge //边集 即邻接矩阵
	n int      // 顶点数
	e int      // 边数量
}

/* --------- 基本的访问操作: 点 --------- */
func (T *GMatrix) Vertex(i int) interface{} {
	return T.V[i].data
}

func (T *GMatrix) InDegree(i int) int {
	return T.V[i].inDegree
}

func (T *GMatrix) OutDegree(i int) int {
	return T.V[i].outDegree
}

func (T *GMatrix) VStatus(i int) VStatus {
	return T.V[i].status
}

func (T *GMatrix) DTime(i int) int {
	return T.V[i].dTime
}

func (T *GMatrix) FTime(i int) int {
	return T.V[i].fTime
}

func (T *GMatrix) Parent(i int) int {
	return T.V[i].parent
}

func (T *GMatrix) Priority(i int) int {
	return T.V[i].priority
}

/* --------- 基本的访问操作: 边 --------- */
func (T *GMatrix) Edge(i, j int) interface{} {
	return T.E[i][j].data
}

func (T *GMatrix) EStatus(i, j int) EStatus {
	return T.E[i][j].status
}

func (T *GMatrix) Weight(i, j int) int {
	return T.E[i][j].weight
}

/* --------- 静态操作 --------- */
//首个邻居
func (T *GMatrix) firstNbr(i int) int {
	return T.nextNbr(i, T.n)
}

//枚举邻接顶点(neighbor)
func (T *GMatrix) nextNbr(i int, j int) int {
	j--
	for (-1 < j) && !T.exist(i, j) {
		j--
	}
	return j
}

//判断边(i, j)是否存在
func (T *GMatrix) exist(i int, j int) bool {
	if T.E[i][j].data != nil {
		return true
	}
	return false
}

/* --------- 动态操作 --------- */
//插入边(i, j, w)
func (T *GMatrix) InsertEdge(edge interface{}, w int, i int, j int) {
	//如果已有边则忽略操作
	if T.exist(i, j) {
		return
	}
	var newEdge Edge
	newEdge.init(edge, w)
	T.E[i][j] = newEdge
	T.e++              //更新边计数
	T.V[i].outDegree++ //更新i出度
	T.V[j].inDegree++  //更新j入度
}

//删除边(i, j)返回被删除边的data
func (T *GMatrix) RemoveEdge(i, j int) interface{} {
	eBak := T.Edge(i, j)
	T.E[i][j] = *new(Edge)
	T.e--
	T.V[i].outDegree--
	T.V[j].inDegree--
	return eBak
}

//插入顶点n
func (T *GMatrix) InsertVertex(vertex interface{}) int {
	var newVertex Vertex
	newVertex.init(vertex)
	T.V = append(T.V, newVertex)
	for j := 0; j < T.n; j++ {
		T.E[j] = append(T.E[j], *new(Edge))
	}
	T.n++
	RowN := make([]Edge, T.n)
	T.E = append(T.E, RowN)
	return T.n - 1
}

//删除顶点 i 返回被删除顶点的data
func (T *GMatrix) RemoveVertex(i int) interface{} {
	for j := 0; j < T.n; j++ {
		if T.exist(i, j) {
			//T.E[i][j] = *new(Edge)
			T.V[j].inDegree--
		}
	}
	T.E = append(T.E[:i], T.E[i+1:]...)
	for j := 0; j < T.n; j++ {
		if T.exist(j, i) {
			//T.E[j][j] = *new(Edge)
			T.E[j] = append(T.E[j][:i], T.E[j][i+1:]...)
			T.V[j].outDegree--
		}
	}
	vBak := T.Vertex(i)
	T.V = append(T.V[:i], T.V[i+1:]...)
	return vBak
}

//一组bfs的连通域
func (T *GMatrix) bfsAtom(v int, clock *int) {
	var Q Queue.Queue
	//入队时变为DISCOVERED
	T.V[v].status = DISCOVERED
	Q.Enqueue(v)
	for !Q.Empty() {
		vv := Q.Dequeue().(int) //取出队首节点vv
		*clock++
		T.V[vv].dTime = *clock                                  //加注时间标签
		for u := T.firstNbr(vv); -1 < u; u = T.nextNbr(vv, u) { //考察vv的每个邻居u
			if T.VStatus(u) == UNDISCOVERED { //如果u没发现
				T.V[u].status = DISCOVERED //则发现新顶点u
				Q.Enqueue(u)               //新顶点u入队
				T.E[vv][u].status = TREE   //vv到u的连线作为查找树的树边
				T.V[u].parent = vv         //u在查找方向上的父节点改为vv
			} else {
				T.E[vv][u].status = CROSS //边(vv, u)归类为跨边
			}
		}
		//全部考察完毕后变为VISITED
		//*clock++
		//T.V[vv].fTime = *clock
		T.V[vv].status = VISITED
	}
}

func (T *GMatrix) BFS(s int) {
	T.reset()
	clock := 0
	T.bfsAtom(s, &clock)
	for v := (s + 1) % T.n; s != v; v = (v + 1) % T.n {
		if T.VStatus(v) == UNDISCOVERED {
			T.bfsAtom(v, &clock)
		}
	}
}

func (T *GMatrix) reset() {
	for i := 0; i < T.n; i++ {
		T.V[i].status = UNDISCOVERED
		for j := 0; j < T.n; j++ {
			T.E[i][j].status = UNDETERMINED
		}
	}
}

//DFS
func (T *GMatrix) dfsAtom(v int, clock *int) {
	*clock++
	T.V[v].dTime = *clock
	T.V[v].status = DISCOVERED
	for u := T.firstNbr(v); -1 < u; u = T.nextNbr(v, u) {
		//根据u状态分别处理
		switch T.VStatus(u) {
		case UNDISCOVERED: //支撑树可进一步拓展
			T.E[v][u].status = TREE
			T.V[u].parent = v
			T.dfsAtom(u, clock)
		case DISCOVERED: //已被发现但是未访问完毕，即本case为后代指向祖先
			T.E[v][u].status = BACKWARD
		default: //u已访问完毕(VISITED 有向图) 根据继承关系选择前向边或者跨边 祖先指向后代
			if T.DTime(v) < T.DTime(u) {
				T.E[v][u].status = FORWARD
			} else {
				T.E[v][u].status = CROSS
			}
		}
	}
	T.V[v].status = VISITED
	*clock++
	T.V[v].fTime = *clock
}
