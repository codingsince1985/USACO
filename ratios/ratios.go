/*
  ID:
  PROG: ratios
  LANG: Go
*/

// USACO Section 3.2 - Feed Ratios

package main

import (
	"fmt"
	"os"
)

var r [4][3]int
var found bool

func okay(i, j, k int) int {
	var t [3]int
	t[0] = i * r[1][0] + j * r[2][0] + k * r[3][0]
	t[1] = i * r[1][1] + j * r[2][1] + k * r[3][1]
	t[2] = i * r[1][2] + j * r[2][2] + k * r[3][2]

	if t[0] * r[0][1] == t[1] * r[0][0] && t[0] * r[0][2] == t[2] * r[0][0] {
		return t[0] / r[0][0]
	}

	return 0;
}

func prepare() {
	in, _ := os.Open("ratios.in")
	for i := 0; i < 4; i++ {
		fmt.Fscanf(in, "%d %d %d", &r[i][0], &r[i][1], &r[i][2])
	}
	in.Close()
}

func main() {
	prepare()

	var num int
	out, _ := os.Create("ratios.out")
	for t := 1; t <= 300 && !found; t++ {
		for i := 0; i <= 100 && !found; i++ {
			if i > t {
				break
			}
			for j := 0; j <= 100; j++ {
				if i + j > t {
					break;
				}
				k := t - i - j
				num = okay(i, j, k)
				if num != 0 {
					fmt.Fprintln(out, i, j, k, num)
					found = true
					break
				}
			}
		}
	}

	if !found {
		fmt.Fprintln(out, "NONE")
	}
	out.Close()
}
