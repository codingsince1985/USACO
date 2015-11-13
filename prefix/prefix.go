/*
  ID:
  PROG: prefix
  LANG: Go
*/

// USACO Section 2.3 - The Longest Prefix

package main

import (
	"fmt"
	"os"
)

var tokens []string
var S string
var d []int

func input() {
	in, _ := os.Open("prefix.in")
	tokens = make([]string, 0)
	var tmp string
	for ; tmp != "."; {
		fmt.Fscanf(in, "%s", &tmp)
		if (tmp != ".") {
			tokens = append(tokens, tmp)
		}
	}
	//	fmt.Println(tokens)

	for {
		count, _ := fmt.Fscanf(in, "%s", &tmp)
		if (count > 0) {
			S += tmp
		} else {
			break;
		}
	}
	//	fmt.Println(S)
	in.Close()
}

func compute() {
	length := len(S)
	d = make([]int, length)
	tokenNum := len(tokens)
	for i := length - 1; i >= 0; i -- {
		subString := S[i:]
		//		fmt.Println(subString)
		for j := 0; j < tokenNum; j ++ {
			//			fmt.Println(" ", tokens[j])
			if len(subString) < len(tokens[j]) {
				continue
			}

			if subString[:len(tokens[j])] == tokens[j] {
				newLength := len(tokens[j]) + d[i + len(tokens[j])]
				if newLength > d[i] {
					d[i] = newLength
				}
			}
		}
	}
}

func main() {
	input()
	compute()
	out, _ := os.Create("prefix.out")
	fmt.Fprintln(out, d[0])
	out.Close()
}
