package main

import "fmt"

func main() {
	n := 24
	for gen := 0; gen < n; gen++ {
		f := fib_bad(gen)
		fmt.Println("gen =", gen, ",\tfib =", f)
	}

}

func fib_bad(n int) int {
	switch {
	case n < 2:
		return n
	default:
		return fib(n-1) + fib(n-2)
	}
}
