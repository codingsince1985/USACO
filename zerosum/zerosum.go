/*
  ID:
  PROG: zerosum
  LANG: Go
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

var N int

type Node struct {
	n, sum, nextNum, sign int
	exp                   string

}

var out *os.File

func prepare() {
	in, _ := os.Open("zerosum.in")
	fmt.Fscanf(in, "%d", &N)
	out, _ = os.Create("zerosum.out")
}

func dfs(node Node) {
	//	fmt.Println(node)
	if node.n == N {
		if node.sum + node.sign * node.nextNum == 0 {
			fmt.Fprintf(out, "%s\n", node.exp)
		}
		return
	}

	dfs(Node{
		n: node.n + 1,
		sum: node.sum,
		nextNum: node.nextNum * 10 + node.n + 1,
		sign: node.sign,
		exp: node.exp + " " + strconv.Itoa(node.n + 1),
	})
	dfs(Node{
		n: node.n + 1,
		sum: node.sum + node.sign * node.nextNum,
		nextNum: node.n + 1,
		sign: 1,
		exp: node.exp + "+" + strconv.Itoa(node.n + 1),
	})
	dfs(Node{
		n: node.n + 1,
		sum: node.sum + node.sign * node.nextNum,
		nextNum: node.n + 1,
		sign: -1,
		exp: node.exp + "-" + strconv.Itoa(node.n + 1),
	})
}

func main() {
	prepare()
	dfs(Node{n:1, sum:0, nextNum:1, sign:1, exp:"1"})
}