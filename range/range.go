/*
  ID:
  PROG: range
  LANG: Go
*/

// USACO Section 3.3 - Home on the Range

package main

import (
	"fmt"
	"os"
)

const MAX = 250

var (
	l [][]byte
	s [][] int
	o [MAX + 1]int
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func solve(n int) {
	var u, lft, ul int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			u, lft, ul = 0, 0, 0
			if l[i][j] == '1' {
				if i != 0 {
					u = s[i - 1][j]
				}

				if j != 0 {
					lft = s[i][j - 1]
				}

				if i != 0 && j != 0 {
					ul = s[i - 1][j - 1]
				}

				s[i][j] = min(min(u, lft), ul) + 1;
				for k := s[i][j]; k >= 2; k-- {
					o[k] ++
				}
			}
		}
	}
}

func initialise(n int) {
	l = make([][]byte, n)
	for i := 0; i < n; i++ {
		l[i] = make([]byte, n + 1) // 1 more to store \n
	}

	s = make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
	}
}

func output(out *os.File) {
	for i := 0; i <= MAX; i++ {
		if o[i] != 0 {
			fmt.Fprintln(out, i, o[i])
		}
	}
}

func main() {
	in, _ := os.Open("range.in")
	defer in.Close()
	out, _ := os.Create("range.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	initialise(n)

	for i := 0; i < n; i++ {
		in.Read(l[i])
	}

	solve(n)
	output(out)
}
