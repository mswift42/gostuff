package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindrome(i int) bool {
	str := strconv.Itoa(i)
	return str == reverseBytes(str)
}

func reverseBytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func isPrime(i int) bool {
	if i == 2 {
		return false
	}
	for a := 3; a < int(math.Sqrt(float64(i))); a += 2 {
		if i%a == 0 {
			return false
		}
	}
	return true
}
func main() {

	for i := 999; i > 1; i-- {
		if isPrime(i) && isPalindrome(i) {
			fmt.Printf("%d", i)
			break
		}
	}
	return
}
