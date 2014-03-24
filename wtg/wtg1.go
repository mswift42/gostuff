// exercises for the way to go.
package main

import "fmt"

func main() {
        forLoop()
        forCharacter()
        forCharacterConc()
        bitwiseComplement()
        fizzBuzz()
}

func forLoop() {
        for i := 0; i < 15; i++ {
                fmt.Printf("the int is %d\n", i)
        }
}
func forCharacter() {
        for i := 0; i < 26; i++ {
                fmt.Printf("\n")
                for j := 0; j < i; j++ {
                        fmt.Printf("G")
                }
        }
}
func forCharacterConc() {
        for i, j := 0, ""; i < 25; i, j = i+1, j+"G" {
                fmt.Println(j)
        }
}
func bitwiseComplement() {
        for i := 0; i <= 10; i++ {
                fmt.Printf("bitwise complement: %b\n", i)
        }
}
func fizzBuzz() {
        for i := 1; i <= 100; i++ {
                switch {
                case i%3 == 0 && i%5 == 0:
                        fmt.Printf("fizzbuzz\n")
                case i%5 == 0:
                        fmt.Printf("buzz\n")
                case i%3 == 0:
                        fmt.Printf("fizz\n")
                default:
                        fmt.Printf("%d\n", i)
                }
        }
}
