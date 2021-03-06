/*
  ID:
  PROG: money
  LANG: Go
*/

// USACO Section 2.3 - Money Systems

package main

import (
	"fmt"
	"os"
)

var V, N int
var coins []int
var c []uint64

func input() {
	in, _ := os.Open("money.in")
	fmt.Fscanf(in, "%d %d", &V, &N)
	coins = make([]int, V + 1)
	for i := 1; i <= V; i++ {
		fmt.Fscanf(in, "%d", &coins[i])
	}
	in.Close()
}

func compute() {
	c = make([]uint64, N + 1)
	c[0] = 1
	for i := 1; i <= V; i ++ {
		for j := coins[i]; j <= N; j++ {
			c[j] += c[j - coins[i]]
		}
		//		fmt.Println(c)
	}
}

func main() {
	input()
	//	fmt.Println(V, N, coins)
	compute()
	out, _ := os.Create("money.out")
	fmt.Fprintln(out, c[N])
	out.Close()
}
