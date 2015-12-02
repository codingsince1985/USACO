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

var a []int
var c [101][101]int

func sum() int {
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

func solve(l, r, s int) int {
	if c[l][r] != 0 {
		return c[l][r]
	}
	if l + 1 == r {
		c[l][r] = a[l]
		return c[l][r]
	} else {
		c[l][r] = s - min(solve(l + 1, r, s - a[l]), solve(l, r - 1, s - a[r - 1]))
		return c[l][r]
	}
}

func main() {
	in, _ := os.Open("game1.in")
	defer in.Close()
	out, _ := os.Create("game1.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &a[i])
	}

	s := sum()
	f := solve(0, n, s)
	fmt.Fprintln(out, f, s - f)
}
