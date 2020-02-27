package main

import "fmt"

func main() {
	data := []int{7, 2, 5, 3, 1, 8, 9, 6, 4}
	var x1, x2 int
	max2(data, 0, 9, &x1, &x2)
	fmt.Println("x1, x2 =", x1, x2)
	fmt.Println("max 2 data:", data[x1], data[x2])
}

func max2(data []int, lo int, hi int, x1 *int, x2 *int) {
	switch {
	case lo+2 == hi:
		if data[lo] > data[lo+1] {
			*x1, *x2 = lo, lo+1
		} else {
			*x1, *x2 = lo+1, lo
		}
		return
	case lo+3 == hi:
		if data[lo] > data[lo+1] {
			*x1, *x2 = lo, lo+1
		} else {
			*x1, *x2 = lo+1, lo
		}
		if data[lo+2] > data[lo] {
			*x1, *x2 = lo+2, lo
		} else if data[lo+2] > data[lo+1] {
			*x2 = lo + 2
		}
		return
	}

	var x1l, x2l, x1r, x2r int
	mi := (lo + hi) >> 1
	max2(data, lo, mi, &x1l, &x2l)
	fmt.Printf("left [%d, %d): x1l=%d, x2l=%d\n", lo, mi, x1l, x2l)
	max2(data, mi, hi, &x1r, &x2r)
	fmt.Printf("right [%d, %d): x1r=%d, x2r=%d\n", mi, hi, x1r, x2r)

	fmt.Printf("x1l = %d, x1r = %d, x2l = %d, x2r = %d, x1 = %d, x2 = %d\n", x1l, x1r, x2l, x2r, *x1, *x2)
	if data[x1l] > data[x1r] {
		*x1 = x1l
		if data[x2l] > data[x1r] {
			*x2 = x2l
		} else {
			*x2 = x1r
		}
	} else {
		*x1 = x1r
		if data[x1l] > data[x2r] {
			*x2 = x1l
		} else {
			*x2 = x2r
		}
	}
	fmt.Println("------ after combine l & r ------")
	fmt.Printf("x1l = %d, x1r = %d, x2l = %d, x2r = %d, x1 = %d, x2 = %d\n", x1l, x1r, x2l, x2r, *x1, *x2)
	fmt.Println("==== end ====")
}
