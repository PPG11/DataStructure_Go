package main

import "fmt"

func main() {
	str1 := "didactically"
	str2 := "advantage"
	longest := LCS(str1, str2)
	fmt.Println("the length of the longest common subsequence is", longest)
}

func LCS(str1 string, str2 string) int {
	fmt.Println("---")
	fmt.Println("str1 =", str1)
	fmt.Println("str2 =", str2)
	switch {
	case len(str1) == 0 || len(str2) == 0:
		return 0
	case str1[0] == str2[0]:
		return LCS(str1[1:], str2[1:]) + 1
	default:
		return max(LCS(str1[1:], str2), LCS(str1, str2[1:]))
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
