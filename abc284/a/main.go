package main

import (
	"fmt"
	"github.com/okamo1984/contest/atcoder"
)

func main() {
	s := atcoder.NewAtCoderStdInScanner()
	num := s.ScanNumber()
	var text string
	for i := 0; i < num; i++ {
		text = s.ScanText()
		if text == "" {
			break
		}
		fmt.Println(text)
	}
}
