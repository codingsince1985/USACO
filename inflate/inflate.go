/*
  ID:
  PROG: inflate
  LANG: Go
*/

package main

import (
	"fmt"
	"math"
	"os"
)

var M, N int
var c [][]int
var point, minute []int

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func knapsack(n, w int) int {
	if n == 0 {
		return 0
	}

	if w < 0 {
		return math.MinInt32
	}

	if c[n][w] != -1 {
		return c[n][w]
	}

	c[n][w] = max(knapsack(n - 1, w), knapsack(n, w - minute[n]) + point[n])
	return c[n][w]
}

func prepare() {
	in, _ := os.Open("inflate.in")
	fmt.Fscanf(in, "%d %d", &M, &N)

	point = make([]int, N + 1)
	minute = make([]int, N + 1)

	for i := 1; i <= N; i++ {
		fmt.Fscanf(in, "%d %d", &point[i], &minute[i])
	}
	in.Close()

	c = make([][]int, N + 1)
	for i := 0; i <= N; i ++ {
		c[i] = make([]int, M + 1)
		for j := 0; j <= M; j++ {
			c[i][j] = -1
		}
	}
}

func main() {
	prepare()

	out, _ := os.Create("inflate.out")
	fmt.Fprintln(out, knapsack(N, M))
	out.Close()
}
