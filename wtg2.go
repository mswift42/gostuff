package main

import "fmt"
import "math"
import "errors"

func main() {
        sum, product, difference := multReturn(20, 40)
        fmt.Printf("sum: %d, product: %d, difference: %d\n", sum, product, difference)
        sum, product, difference = multReturnNamed(33, 722)
        fmt.Printf("sum: %d, product: %d, difference: %d\n", sum, product,
                difference)
        varArgs(2, 3, 4, 5, 6, 7)
        lambda()
        fibarray()

}

func multReturn(a, b int) (int, int, int) {
        sum := a + b
        product := a * b
        difference := a - b
        return sum, product, difference
}
func multReturnNamed(a, b int) (sum, product, difference int) {
        sum = a + b
        product = a * b
        difference = a - b
        return
}
func errorReturnVal(x float64) (error, float64) {
        if x < 0 {
                return errors.New("negative number"), x
        }
        return nil, math.Sqrt(x)
}
func varArgs(args ...int) {
        for _, i := range args {
                fmt.Println(i)
        }
}
func lambda() {
        fv := func() {
                fmt.Printf("Hello World")
        }
        fmt.Printf("%T\n", fv)
        fv()
}
func fibarray() {
        var result [50]int
        a, b := 1, 1
        for i := range result {
                result[i] = a
                a, b = b, a+b
        }
        fmt.Println(result)
}
