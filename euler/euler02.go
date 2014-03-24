package main

import "fmt"

func euler02() int {
        a, b := 1, 1
        sum := 0
        for a < 4000000 {
                a, b = b, a+b
                if a%2 == 0 {
                        sum += a
                }

        }
        return sum
}

func main() {
        fmt.Println(euler02())
}
