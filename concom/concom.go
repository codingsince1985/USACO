/*
  ID:
  PROG: concom
  LANG: Go
*/

package main

import (
	"fmt"
	"os"
)

const MAX_N = 101

var N int
var control [MAX_N][MAX_N]int
var c [MAX_N]int
var v [MAX_N]bool
var out *os.File

func prepare() {
	in, _ := os.Open("concom.in")
	fmt.Fscanf(in, "%d", &N)
	var c1, c2, pt int
	for i := 1; i <= N; i++ {
		fmt.Fscanf(in, "%d %d %d", &c1, &c2, &pt)
		control[c1][c2] = pt
	}

	out, _ = os.Create("concom.out")

	//	fmt.Println(N)
	//	for i := 1; i <= N; i++ {
	//		for j := 1; j <= N; j++ {
	//			fmt.Print(control[i][j], "\t")
	//		}
	//		fmt.Println()
	//	}
}

func initialize() {
	for i := 1; i <= N; i++ {
		c[i] = 0
		v[i] = false
	}
}

func process(company int) {
	if v[company] {
		return
	}
	v[company] = true
	for i := 1; i <= N; i++ {
		if company == i {
			continue
		}
		c[i] += control[company][i]
		if c[i] >= 50 {
			process(i)

		}
	}
}

func print(company int) {
	for i := 1; i <= N; i++ {
		if c[i] >= 50 {
			fmt.Fprintf(out, "%d %d\n", company, i)
		}
	}
}

func do() {
	for i := 1; i <= N; i++ {
		initialize()
		process(i)
		print(i)
	}
}

func main() {
	prepare()
	do()
}
