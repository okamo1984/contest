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
