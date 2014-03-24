package main

import "fmt"
import "math"

// Largest prime factor
// Problem 3

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

var limit = int(math.Sqrt(600851475143))

func isprime(i int) bool {
        if i == 2 {
                return false
        }
        for a := 3; a < int(math.Sqrt(float64(i))); a = a + 2 {
                if i%a == 0 {
                        return false
                }
        }
        return true
}

func euler03() int {
        for i := limit; i > 0; i-- {
                if 600851475143%i == 0 && isprime(i) {
                        return i
                }
        }
        return 0
}

func main() {
        fmt.Println(euler03())
}
