package main

import "fmt"

func main() {
	data := []int{7, 2, 5, 3, 1, 8, 9, 6, 4}
	fmt.Println("origin data:", data)
	bubbleSort(data)
	fmt.Println("after sort,  data:", data)
}

func bubbleSort(data []int) {
	n := len(data)
	var isSorted bool
	for !isSorted {
		isSorted = true
		for i := 1; i < n; i++ {
			if data[i] < data[i-1] {
				swap(&data[i], &data[i-1])
				isSorted = false
				fmt.Println(data)
				fmt.Println("---")
			}
		}
		n--
	}
	//fmt.Println(data)
}

func swap(i *int, j *int) {
	fmt.Println("swap", *i, *j)
	*i, *j = *j, *i
}
