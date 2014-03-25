package main

import "fmt"

// Summation of primes
// Problem 10

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

func sieve() []int {
        primes := make([]int, 2000000)
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

func euler10() int {
        sum := 0
        for _, i := range sieve() {
                sum += i
        }
        return sum
}

func main() {
        fmt.Println(euler10())
}
