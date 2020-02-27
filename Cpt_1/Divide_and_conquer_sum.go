package main

import "fmt"

func main() {
	data := []int{7, 2, 5, 3, 1, 8, 9, 6, 4}
	fmt.Println("origin data:", data)
	res := divide_sum(data, 0, 9)
	fmt.Println("divide sum =", res)
}

func divide_sum(data []int, lo int, hi int) int {
	if lo == hi-1 {
		return data[lo]
	}
	mi := (lo + hi) >> 1
	return divide_sum(data, lo, mi) + divide_sum(data, mi, hi)
}
