/*
  ID:
  PROG: stamps
  LANG: Go
*/

// USACO Section 3.1 - Stamps

package main

import (
	"fmt"
	"os"
)

const MAX_STAMP_VALUE = 10000

var K, N int
var v, f []int

func prepare() {
	in, _ := os.Open("stamps.in")
	fmt.Fscanf(in, "%d %d", &K, &N)
	v = make([]int, N)
	f = make([]int, K * MAX_STAMP_VALUE + 1)
	f[0] = 0
	for i := 1; i <= K * MAX_STAMP_VALUE; i++ {
		f[i] = K * MAX_STAMP_VALUE + 1
	}
	for i := 0; i < N; i++ {
		fmt.Fscanf(in, "%d", &v[i])
	}
	in.Close()
}

func main() {
	prepare()

	var i int
	for i = 1; i <= K * MAX_STAMP_VALUE; i ++ {
		for j := 0; j < N; j++ {
			if i >= v[j] {
				if f[i] > f[i - v[j]] + 1 {
					f[i] = f[i - v[j]] + 1
				}
			}
		}
		if f[i] > K {
			break
		}
	}
	out, _ := os.Create("stamps.out")
	fmt.Fprintln(out, i - 1)
	out.Close()
}