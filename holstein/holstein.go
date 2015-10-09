/*
  ID:
  PROG: holstein
  LANG: Go
*/

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

func valid(total []int) bool {
	for i := range total {
		if total[i] < v[i] {
			return false
		}
	}
	return true
}

func dfs(used, total []int, depth int) {
	if (depth == G) {
		return
	}

	keepUsed := make([]int, len(used));
	copy(keepUsed, used);
	keepTotal := make([]int, V, V)
	copy(keepTotal, total)

	used = append(used, depth)
	for i := range total {
		total[i] += g[depth][i]
	}

	if valid(total) {
		if (minCount > len(used)) {
			//			fmt.Println(used)
			//			fmt.Println(total)
			minCount = len(used)
			answer = used
		}
		return
	}

	dfs(used, total, depth + 1)

	used = make([]int, len(keepUsed));
	copy(used, keepUsed);
	copy(total, keepTotal)

	dfs(used, total, depth + 1)
}

func input() {
	in, _ := os.Open("holstein.in")
	fmt.Fscanf(in, "%d", &V)

	v = make([]int, V, V)
	for i := 0; i < V; i ++ {
		fmt.Fscanf(in, "%d", &v[i])
	}

	fmt.Fscanf(in, "%d", &G)
	minCount = G + 1

	g = make([][]int, G, G)
	for i := 0; i < G; i ++ {
		g[i] = make([]int, V, V)
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
	used := make([]int, 0)
	total := make([]int, V, V)
	dfs(used, total, 0)
	output()
}
