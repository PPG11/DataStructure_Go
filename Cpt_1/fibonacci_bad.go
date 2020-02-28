package main

import "fmt"

func main() {
	n := 24
	for gen := 0; gen < n; gen++ {
		f := fib(gen)
		fmt.Println("gen =", gen, ",\tfib =", f)
	}

}

func fib(n int) int {
	switch {
	case n < 2:
		return n
	default:
		return fib(n-1) + fib(n-2)
	}
}
