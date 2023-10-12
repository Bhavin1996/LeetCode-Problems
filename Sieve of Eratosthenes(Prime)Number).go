package leetcode

import (
	"fmt"
)

func sieveOfEratosthenes(limit int) []int {
	// Create a boolean slice to represent numbers up to 'limit'.
	isPrime := make([]bool, limit+1)

	// Initialize all numbers as prime.
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	// Start marking non-prime numbers.
	for p := 2; p*p <= limit; p++ {
		if isPrime[p] == true {
			// Mark all multiples of p as non-prime.
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}

	// Collect prime numbers into a slice.
	primes := make([]int, 0)
	for p := 2; p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func main_func() {
	limit := 100 // Change this limit as needed
	primes := sieveOfEratosthenes(limit)
	fmt.Println("Prime numbers up to", limit, "are:", primes)
}
