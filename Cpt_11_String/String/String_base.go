package String

func match1(P string, T string) int {
	n, m := len(T), len(P)
	i, j := 0, 0
	for j < m && i < n {
		if T[i] == P[j] {
			i++
			j++
		} else {
			i -= j - 1
			j = 0
		}
	}
	return i - j
}

func match2(P, T string) int {
	n, m := len(T), len(P)
	i, j := 0, 0
	for i = 0; i < n-m+1; i++ {
		for j = 0; j < m; j++ {
			if T[i+j] != P[j] {
				break
			}
		}
		if m <= j {
			break
		}
	}
	return i
}

func matchKMP(P, T string) int {
	//var next []int
	next := buildNext(P)
	n, m := len(T), len(P)
	i, j := 0, 0
	for j < m && i < n {
		if j < 0 || T[i] == P[j] {
			//匹配
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return i - j
}

func buildNext(P string) []int {
	m := len(P)
	j := 0
	N := make([]int, m)
	N[0] = -1
	t := -1
	for j < m-1 {
		switch {
		case t < 0:
			j++
			t++
			//以卵击石
			//N[j] = t
			if P[j] != P[t] {
				N[j] = t
			} else {
				N[j] = N[t]
			}
		case P[j] == P[t]:
			j++
			t++
			//以卵击石
			//N[j] = t
			if P[j] != P[t] {
				N[j] = t
			} else {
				N[j] = N[t]
			}
		default:
			t = N[t]
		}
	}
	return N
}

func buildBC(P string) []int {
	bc := make([]int, 256)
	for j := 0; j < 256; j++ {
		bc[j] = -1
	}
	for m, j := len(P), 0; j < m; j++ {
		bc[int(P[j])] = j
	}
	return bc
}
