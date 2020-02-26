package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	data := [51]int{55, 34, 31, 27, 21, 21, 20, 17, 15, 15, 15, 13, 12, 11, 11, 11, 11, 10, 10, 10, 10,
		9, 9, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3}
	sum := 0
	for _, n := range data {
		sum += n
	}
	// fmt.Println("sum =", sum)
	var half int = sum / 2

	// filter use binary to choose province
	numProvince := len(data)
	maxFilter := int64(math.Pow(2, float64(numProvince))) - 1
	fmt.Printf("maxFilter = %d\n", maxFilter)
	//fmt.Printf("maxFilter = %051b\n", maxFilter)

	var filter int64
	find := false
	for filter = 1; filter < maxFilter; filter++ {
		if filter%(10000000) == 0 {
			rate := float64(filter) / float64(maxFilter)
			fmt.Printf("%.4f%%\n", rate*100)
		}
		bFilter := ItoB(filter, numProvince)
		numVote, isOverflow := calculate(data, bFilter, half)
		if isOverflow {
			continue
		}
		if numVote == half {
			find = true
			break
		}
	}
	if find {
		fmt.Println(filter)
	} else {
		fmt.Println("Do not find the result")
	}
}

func calculate(data [51]int, filter string, threshold int) (numVote int, isOverflow bool) {
	isOverflow = false
	for idx, item := range filter {
		if item == 1 {
			numVote += data[idx]
			if numVote > threshold {
				return 0, true
			}
		}
	}
	return
}

func ItoB(i int64, length int) (s string) {
	bin := strconv.FormatInt(i, 2)
	zeroLength := length - len(bin)
	for k := 0; k < zeroLength; k++ {
		s += "0"
	}
	s += bin
	return
}
