package main

import "fmt"

func euler01() int {
        sum := 0
        for i := 1; i < 1000; i++ {
                if i%3 == 0 || i%5 == 0 {
                        sum += i
                }
        }
        return sum
}

func main() {
        fmt.Println(euler01())
        fmt.Println(Euler02())
}
