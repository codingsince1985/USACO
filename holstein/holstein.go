/*
  ID:
  PROG: holstein
  LANG: Go
*/

// USACO Section 2.1 - Healthy Holsteins

package main

import (
	"fmt"
	"os"
)

var V, G int
var v []int
var g [][]int
var minCount int
var answer []int

func newTotal(total []int, current int) []int {
	newTotal := make([]int, V)
	for i := range total {
		newTotal[i] = total[i] + g[current][i]
	}
	return newTotal
}

func isEnough(total []int) bool {
	for i := range total {
		if total[i] < v[i] {
			return false
		}
	}
	return true
}

func dfs(used, total []int, current int) {
	if (current == G) {
		return
	}

	newUsed := append(used, current)
	if (len(newUsed) >= minCount) {
		return
	}
	newTotal := newTotal(total, current)

	if isEnough(newTotal) {
		if (len(newUsed) < minCount) {
			//			fmt.Println(newUsed)
			//			fmt.Println(newTotal)
			minCount = len(newUsed)
			answer = make([]int, len(newUsed))
			copy(answer, newUsed)
		}
		return
	}

	dfs(newUsed, newTotal, current + 1)

	dfs(used, total, current + 1)
}

func input() {
	in, _ := os.Open("holstein.in")
	fmt.Fscanf(in, "%d", &V)

	v = make([]int, V)
	for i := 0; i < V; i ++ {
		fmt.Fscanf(in, "%d", &v[i])
	}

	fmt.Fscanf(in, "%d", &G)
	minCount = G + 1

	g = make([][]int, G)
	for i := 0; i < G; i ++ {
		g[i] = make([]int, V)
		for j := 0; j < V; j ++ {
			fmt.Fscanf(in, "%d", &g[i][j])
		}
	}
	in.Close()
}

func output() {
	out, _ := os.Create("holstein.out")
	fmt.Fprint(out, len(answer));
	for _, scoop := range answer {
		fmt.Fprintf(out, " %d", scoop + 1);
	}
	fmt.Fprintln(out)
	out.Close()
}

func main() {
	input()
	dfs(make([]int, 0), make([]int, V), 0)
	output()
}
