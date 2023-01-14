package main

import (
	"fmt"

	"github.com/okamo1984/contest/atcoder"
)

func dfs(visited []bool, gather [][]int, visit int) (result int) {
	result++
	visited[visit] = true
	for _, v := range gather[visit] {
		if visited[v] {
			continue
		}
		result += dfs(visited, gather, v)
	}
	visited[visit] = false
	return result
}

func main() {
	s := atcoder.NewAtCoderStdInScanner()
	n, m := atcoder.ScanNumberTupleLine(s.ScanText())
	gather := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := atcoder.ScanNumberTupleLine(s.ScanText())
		u--
		v--
		gather[u] = append(gather[u], v)
		gather[v] = append(gather[v], u)
	}

	var result int
	visited := make([]bool, n)
	result += dfs(visited, gather, 0)
	fmt.Println(result)
}
