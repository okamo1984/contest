package atcoder

import (
	"bufio"
	"os"
	"strconv"
)

type AtCoderStdInScanner struct {
	*bufio.Scanner
}

func NewAtCoderStdInScanner() *AtCoderStdInScanner {
	return &AtCoderStdInScanner{bufio.NewScanner(os.Stdin)}
}

func (s *AtCoderStdInScanner) ScanNumber() (num int) {
	s.Scan()
	num, _ = strconv.Atoi(s.Text())
	return
}
