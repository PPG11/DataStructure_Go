package main

import "fmt"

func main() {
	n := 64
	for gen := 0; gen < n; gen++ {
		f := fib_dp(gen)
		fmt.Println("gen =", gen, ",\tfib =", f)
	}

}

func fib_dp(n int) int {
	f, g := 0, 1
	for ; n > 0; n-- {
		g = g + f
		f = g - f
	}
	return g
}
