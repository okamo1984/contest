package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	scanned := make([]string, num)
	var text string
	for i := 0; i < num; i++ {
		scanner.Scan()
		text = scanner.Text()
		if text == "" {
			break
		}
		scanned[num-(i+1)] = text
	}
	fmt.Println(strings.Join(scanned, "\n"))
}
