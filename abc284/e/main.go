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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, m := scanLine(scanner.Text())
	gather := make([][]int, n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		u, v := scanLine(scanner.Text())
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
