package main

import (
        "fmt"
        "math/big"
        "strconv"
        "strings"
)

func factorial(n int64) *big.Int {
        a := new(big.Int)
        return a.MulRange(1, n)
}

func numtoString(n *big.Int) string {
        return n.String()
}

func sumDigits(s string) int {
        strslice := strings.Split(s, "")
        sum := 0
        for _, i := range strslice {
                a, err := strconv.Atoi(i)
                if err != nil {
                        panic(err)
                }
                sum += a
        }
        return sum
}

func euler20() int {
        return sumDigits(numtoString(factorial(100)))
}

func main() {
        fmt.Println(euler20())
}
