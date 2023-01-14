package main

import (
	"fmt"

	"github.com/okamo1984/contest/atcoder"
)

func solve(n, i int, t string) string {
	j := n - i
	s := t[i : len(t)-j]
	rs := atcoder.Reverse(s)
	fs := fi(n, i, rs, s)
	if fs == t {
		return rs
	}
	return ""
}

func fi(n, i int, s, rs string) string {
	ri := s[:i]
	bi := s[i:]
	return ri + rs + bi
}

func main() {
	s := atcoder.NewAtCoderStdInScanner()
	num := s.ScanNumber()
	text := s.ScanText()
	for i := 0; i < num; i++ {
		answer := solve(num, i, text)
		if answer != "" {
			fmt.Println(answer)
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
