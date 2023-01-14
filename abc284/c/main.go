package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanLine(line string) (n, m int) {
	ls := bufio.NewScanner(strings.NewReader(line))
	ls.Split(bufio.ScanWords)
	ls.Scan()
	n, _ = strconv.Atoi(ls.Text())
	ls.Scan()
	m, _ = strconv.Atoi(ls.Text())
	return
}

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
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	var line string
	line = s.Text()
	n, m := scanLine(line)
	gather := make([][]int, n)
	for i := 0; i < m; i++ {
		s.Scan()
		line = s.Text()
		u, v := scanLine(line)
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
