package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task2Scanner struct {
	*bufio.Scanner
	results     []string
	currentCase int
}

func (s *task2Scanner) scanCase() {
	s.Scan()
	num, _ := strconv.Atoi(s.Text())
	s.Scan()
	lineScanner := bufio.NewScanner(strings.NewReader(s.Text()))
	lineScanner.Split(bufio.ScanWords)
	var odd int
	for i := 0; i < num; i++ {
		lineScanner.Scan()
		digit, _ := strconv.Atoi(lineScanner.Text())
		if digit%2 == 1 {
			odd++
		}
	}
	s.results[s.currentCase] = strconv.Itoa(odd)
	s.currentCase++
}

func main() {
	bufioScanner := bufio.NewScanner(os.Stdin)
	bufioScanner.Scan()
	caseNum, _ := strconv.Atoi(bufioScanner.Text())
	scanner := task2Scanner{bufioScanner, make([]string, caseNum), 0}
	for i := 0; i < caseNum; i++ {
		scanner.scanCase()
	}
	fmt.Println(strings.Join(scanner.results, "\n"))
}
