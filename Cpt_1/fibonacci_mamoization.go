package main

import "fmt"

var fibonacciList = [100]int{}

func main() {
	n := 64
	for gen := 0; gen < n; gen++ {
		f := fib_mamo(gen)
		fmt.Println("gen =", gen, ",\tfib =", f)
	}
}

func fib_mamo(n int) int {
	switch {
	case n < 2:
		fibonacciList[n] = n
		return n
	default:
		fibonacciList[n] = fibonacciList[n-1] + fibonacciList[n-2]
		return fibonacciList[n]
	}
}
