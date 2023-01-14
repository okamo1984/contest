package atcoder

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func (s *AtCoderStdInScanner) ScanText() string {
	s.Scan()
	return s.Text()
}

type AtCoderLineScanner struct {
	*bufio.Scanner
}

func NewAtCoderLineScanner(line string) *AtCoderLineScanner {
	s := bufio.NewScanner(strings.NewReader(line))
	s.Split(bufio.ScanWords)
	return &AtCoderLineScanner{s}
}

func (s *AtCoderLineScanner) ScanNumber() (num int) {
	s.Scan()
	num, _ = strconv.Atoi(s.Text())
	return
}
