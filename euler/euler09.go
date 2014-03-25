package main

import "fmt"

// Special Pythagorean triplet
// Problem 9

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a2 + b2 = c2

// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func euler09() int {
        for c := 1; c < 1000; c++ {
                for b := 1; b < c; b++ {
                        for a := 1; a < b; a++ {
                                if a*a+b*b == c*c && a+b+c == 1000 {
                                        return a * b * c
                                }
                        }
                }
        }
        return 0
}
func main() {
        fmt.Println(euler09())
}
