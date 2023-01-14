package main

import (
	"fmt"
	"github.com/okamo1984/contest/atcoder"
)

func dfs(visited []bool, gather [][]int, visit int) {
	visited[visit] = true
	for _, v := range gather[visit] {
		if visited[v] {
			continue
		}
		dfs(visited, gather, v)
	}
}

func main() {
	s := atcoder.NewAtCoderStdInScanner()
	var line string
	line = s.ScanText()
	n, m := atcoder.ScanNumberTupleLine(line)
	gather := make([][]int, n)
	for i := 0; i < m; i++ {
		s.Scan()
		line = s.Text()
		u, v := atcoder.ScanNumberTupleLine(line)
		u--
		v--
		gather[u] = append(gather[u], v)
		gather[v] = append(gather[v], u)
	}

	visited := make([]bool, n)
	var result int
	for i := 0; i < n; i++ {
		if !visited[i] {
			result++
			dfs(visited, gather, i)
		}
	}
	fmt.Println(result)
}
