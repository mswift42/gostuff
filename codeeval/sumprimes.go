package main

import "fmt"

func sieve() []int {
	primes := make([]int, 500000)
	for i := 2; i < len(primes); i++ {
		primes[i] = i
	}
	for i := 2; i < len(primes); i++ {
		for j := i + i; j < len(primes); j += i {
			primes[j] = 0
		}

	}
	return primes

}

func main() {
	count := 0
	sum := 0
	for _, i := range sieve() {
		if i != 0 {
			count++
			sum += i
			if count == 1000 {
				break
			}
		}
	}
	fmt.Println(sum)
}
