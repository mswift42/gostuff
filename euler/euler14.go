package main

import "fmt"

// Longest Collatz sequence Problem 14

// The following iterative sequence is defined for the set of positive
// integers:

// n → n/2 (n is even) n → 3n + 1 (n is odd)

// Using the rule above and starting with 13, we generate the following
// sequence: 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

// It can be seen that this sequence (starting at 13 and finishing at 1)
// contains 10 terms. Although it has not been proved yet (Collatz
// Problem), it is thought that all starting numbers finish at 1.

// Which starting number, under one million, produces the longest chain?

// NOTE: Once the chain starts the terms are allowed to go above one
// million.

func main() {
        fmt.Println(collatz(13))
        fmt.Println(collatzlength(13))
        fmt.Println(euler14())
}

func collatz(num int) int {
        if num%2 == 0 {
                return num / 2
        }
        return num*3 + 1
}

func collatzlength(num int) int {
        count := 1
        n := num
        for n != 1 {
                n = collatz(n)
                count++
        }
        return count
}

func euler14() int {
        current, max := 0, 0
        for i := 10; i < 1000000; i++ {
                if maxlength := collatzlength(i); maxlength > max {
                        max, current = maxlength, i
                }
        }
        return current
}
