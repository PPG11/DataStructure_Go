package main

import "fmt"

func main() {
	data := []int{7, 2, 5, 3, 1, 8, 9, 6, 4, 10}
	fmt.Println("origin data:", data)
	res := decrease_sum(data, 0, 10)
	fmt.Println("decrease sum =", res)
}

func decrease_sum(data []int, lo int, hi int) int {
	if lo == hi-1 {
		return data[lo]
	}
	return decrease_sum(data, lo+1, hi) + data[lo]
}
