package main

import (
	"fmt"

	"github.com/okamo1984/contest/atcoder"
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
	s := atcoder.NewAtCoderStdInScanner()
	num := s.ScanNumber()
	for i := 0; i < num; i++ {
		n := s.ScanNumber()
		p, q := findAnswer(n)
		fmt.Println(fmt.Sprintf("%d %d", p, q))
	}
}
