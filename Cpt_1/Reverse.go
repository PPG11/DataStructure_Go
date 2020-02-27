package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("origin data:", data)
	reverse(data, 0, 10)
	fmt.Println("reverse data:", data)
}

func reverse(data []int, lo int, hi int) {
	data[lo], data[hi-1] = data[hi-1], data[lo]
	if lo+2 < hi {
		reverse(data, lo+1, hi-1)
	}
}
