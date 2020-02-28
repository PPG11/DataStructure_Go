package main

import "fmt"

// longest common subsequence
func main() {
	str1 := "didactically"
	str2 := "advantage"
	longest := LCS1(str1, str2)
	fmt.Println("the length of the longest common subsequence is", longest)
}

func LCS1(str1 string, str2 string) int {
	//fmt.Println("---")
	//fmt.Println("str1 =", str1)
	//fmt.Println("str2 =", str2)
	switch {
	case len(str1) == 0 || len(str2) == 0:
		return 0
	case str1[0] == str2[0]:
		return LCS1(str1[1:], str2[1:]) + 1
	default:
		return max(LCS1(str1[1:], str2), LCS1(str1, str2[1:]))
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
