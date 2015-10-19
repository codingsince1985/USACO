package main

import (
	"fmt"
	"os"
	"strconv"
)

var N int

type Node struct {
	n, sum, num, sign int
	exp               string

}

var out *os.File

func input() {
	in, _ := os.Open("zerosum.in")
	fmt.Fscanf(in, "%d", &N)
	out, _ = os.Create("zerosum.out")
}

func dfs(node Node) {
	//	fmt.Println(node)
	if node.n == N {
		if node.sum + node.sign * node.num == 0 {
			fmt.Fprintf(out, "%s\n", node.exp)
		}
		return
	}

	dfs(Node{
		n: node.n + 1,
		sum: node.sum,
		num: node.num * 10 + node.n + 1,
		sign: node.sign,
		exp: node.exp + " " + strconv.Itoa(node.n + 1),
	})
	dfs(Node{
		n:node.n + 1,
		sum:node.sum + node.sign * node.num,
		num:node.n + 1,
		sign:1,
		exp:node.exp + "+" + strconv.Itoa(node.n + 1),
	})
	dfs(Node{
		n:node.n + 1,
		sum:node.sum + node.sign * node.num,
		num:node.n + 1,
		sign:-1,
		exp:node.exp + "-" + strconv.Itoa(node.n + 1),
	})
}

func main() {
	input()
	dfs(Node{n:1, sum:0, num:1, sign:1, exp:"1"})
}