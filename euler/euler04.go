package main

import "fmt"
import "strconv"

func reverseBytes(s string) string {
        r := make([]byte, len(s))
        for i := 0; i < len(s); i++ {
                r[i] = s[len(s)-1-i]
        }
        return string(r)
}

func isPalindrome(i int) bool {
        s := strconv.Itoa(i)
        return s == reverseBytes(s)
}
func euler04() int {
        max := 0
        for i := 100; i < 1000; i++ {
                for j := 100; j < 1000; j++ {
                        if prod := i * j; isPalindrome(prod) {
                                if prod > max {
                                        max = prod
                                }
                        }
                }
        }
        return max
}

func main() {
        fmt.Println(isPalindrome(10201))
        fmt.Println(euler04())

}
