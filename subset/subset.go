/*
  ID:
  PROG: subset
  LANG: Go
*/

package main

import (
	"fmt"
	"os"
)

const MAX = 40

var N int

func input() {
	in, _ := os.Open("subset.in")
	fmt.Fscanf(in, "%d", &N)
}

func calculate() int {
	if (N * (N + 1)) % 4 != 0 {
		return 0;
	}

	a := make([][]int, MAX)
	for i := 0; i < MAX; i ++ {
		a[i] = make([]int, MAX * (MAX + 1) / 4)
	}

	a[1][0] = 1
	a[1][1] = 1
	for i := 2; i <= N; i ++ {
		for j := 0; j <= i * (i + 1) / 2; j ++ {
			a[i][j] = a[i - 1][j]
			if j >= i {
				a[i][j] += a[i - 1][j - i]
			}
			//			fmt.Print(a[i][j], "\t")
		}
		//		fmt.Println()
	}
	return a[N][N * (N + 1) / 4] / 2;
}

func output(count int) {
	out, _ := os.Create("subset.out")
	fmt.Fprintf(
		out, "%d\n", count)
}

func main() {
	input()
	output(calculate())
}