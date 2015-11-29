/*
  ID:
  PROG: game1
  LANG: Go
*/

// USACO Section 3.3 - A Game

package main

import (
	"fmt"
	"os"
)

func sum(a []int) int {
	t := 0
	for _, v := range a {
		t += v
	}
	return t
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func solve(a []int) int {
	l := len(a)
	if l == 1 {
		return a[0]
	} else {
		return sum(a) - min(solve(a[1:]), solve(a[:l - 1]))
	}
}

func main() {
	in, _ := os.Open("game1.in")
	defer in.Close()
	out, _ := os.Create("game1.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &a[i])
	}
	f := solve(a)
	fmt.Fprintln(out, f, sum(a) - f)
}
