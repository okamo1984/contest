package atcoder

import (
	"bufio"
	"strconv"
	"strings"
)

func ScanNumberTupleLine(line string) (n, m int) {
	ls := bufio.NewScanner(strings.NewReader(line))
	ls.Split(bufio.ScanWords)
	ls.Scan()
	n, _ = strconv.Atoi(ls.Text())
	ls.Scan()
	m, _ = strconv.Atoi(ls.Text())
	return
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
