package main

import (
        "fmt"
)

func sumquares() int {
        result := 0
        for i := 1; i <= 100; i++ {
                result += i * i
        }
        return result
}

func squaresum() int {
        result := 0
        for i := 1; i <= 100; i++ {
                result += i
        }
        return result * result
}

func main() {
        fmt.Println(squaresum() - sumquares())
}
