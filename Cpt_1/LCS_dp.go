package main

import "fmt"

// longest common subsequence
var LCSMatrix [][]int

func main() {
	str1 := "program"
	str2 := "algorithm"
	n1, n2 := len(str1), len(str2)
	LCSMatrix = make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		LCSMatrix[i] = make([]int, n2+1)
	}
	LCSdp(str1, str2)
	longest := LCSMatrix[n1][n2]
	for i := 0; i <= n1; i++ {
		fmt.Println(LCSMatrix[i])
	}
	fmt.Println("the length of the longest common subsequence is", longest)
}

func LCSdp(str1 string, str2 string) {
	n1, n2 := len(str1), len(str2)
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if str1[i-1] == str2[j-1] {
				LCSMatrix[i][j] = LCSMatrix[i-1][j-1] + 1
			} else {
				LCSMatrix[i][j] = maxDP(LCSMatrix[i][j-1], LCSMatrix[i-1][j])
			}
		}
	}
}

func maxDP(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
