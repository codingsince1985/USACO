/*
  ID:
  PROG: fact4
  LANG: Go
*/

package main

import (
	"fmt"
	"os"
)

func removeTrailingZeros(n int) int {
	for n % 10 == 0 {
		n /= 10
	}
	return n
}

func do(n int) int {
	d := 1
	var next int
	for i := 1; i <= n; i++ {
		d *= i
		d = removeTrailingZeros(d)

		next = i + 1
		if next <= n && ((d % 10 ) * (removeTrailingZeros(next) % 10)) % 10 == 0 {
			d %= 100
		} else {
			d %= 10
		}
	}
	return d
}

func main() {
	in, _ := os.Open("fact4.in")
	out, _ := os.Create("fact4.out")
	var n int
	for {
		_, err := fmt.Fscanf(in, "%d", &n)
		if err != nil {
			break
		}
		fmt.Fprintln(out, do(n));
	}
	in.Close()
	out.Close()
}