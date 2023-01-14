package main

import (
	"fmt"
	"github.com/okamo1984/contest/atcoder"
)

type task2Scanner struct {
	*atcoder.AtCoderStdInScanner
}

func (s *task2Scanner) scanCase() (odd int) {
	num := s.ScanNumber()
	ls := atcoder.NewAtCoderLineScanner(s.ScanText())
	for i := 0; i < num; i++ {
		digit := ls.ScanNumber()
		if digit%2 == 1 {
			odd++
		}
	}
	return
}

func main() {
	s := atcoder.NewAtCoderStdInScanner()
	caseNum := s.ScanNumber()
	t2s := task2Scanner{s}
	for i := 0; i < caseNum; i++ {
		fmt.Println(t2s.scanCase())
	}
}
