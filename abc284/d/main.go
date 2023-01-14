package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	primes := make([]int, 1)
	primes[0] = 2
outer:
	for i := 3; i < n; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue outer
			}
			return true
		}
	}
	return false
}

func findAnswer(n int) (int, int) {
	primes := make([]int, 1)
	primes[0] = 2
outer:
	for i := 3; i < n; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
		if n%(i*i) == 0 {
			c := n / (i * i)
			if isPrime(c) {
				return i, c
			}
		}
	}
	return 0, 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	answers := make([]string, num)
	for i := 0; i < num; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		p, q := findAnswer(n)
		answers[i] = fmt.Sprintf("%d %d", p, q)
	}
	fmt.Println(strings.Join(answers, "\n"))
}
