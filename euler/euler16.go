package main

// Power digit sum
// Problem 16

// 215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 21000?

import (
        "fmt"
        "math/big"
        "strconv"
        "strings"
)

// return 2 ** 1000 as slice of it digits as strings
func bigstringslice() []string {
        a := big.NewInt(2)
        b := big.NewInt(1000)
        exp := new(big.Int)
        bignum := exp.Exp(a, b, exp)
        return strings.Split(bignum.String(), "")
}

func euler16() int {
        slice := bigstringslice()
        sum := 0
        for _, i := range slice {
                x, err := strconv.Atoi(i)
                if err != nil {
                        panic(err)
                }
                sum += x
        }
        return sum
}

func main() {
        fmt.Println(euler16())
}
