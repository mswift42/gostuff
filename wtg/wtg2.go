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
        forArray()
        fibarray()
        slby()
        fmt.Println(fibonacciFuncarray(15))
        slice := []float64{2.1, 3.4, 5.123}
        fmt.Printf("sum of slice is %.4f\n", sumArray(slice))
        slice2 := []int{20, 30, 40, 10, -20, 120, -400, 11}
        fmt.Printf("min of slice2 is: %d, and max is: %d\n",
                minSlice(slice2), maxSlice(slice2))
        fmt.Println(magnifySlice(slice2, 3))
        fmt.Println(even(4))
        slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
        fmt.Println(filterSlice(slice3, even))
        fmt.Println(stringSplit("hallo", 2))
        fmt.Println(mapFunction(func(i int) int { return i * 10 }, slice3))

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
                fmt.Printf("Hello World\n")
        }
        fmt.Printf("%T\n", fv)
        fv()
}
func forArray() {
        var result [15]int
        for i := 0; i < len(result); i++ {
                result[i] = i
        }
        fmt.Println(result)
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
func slby() {
        b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
        fmt.Println(b[1:4])
        fmt.Println(b[:2])
}
func fibonacciFuncarray(n int) []int {
        result := make([]int, n)
        a, b := 1, 1
        for i := range result {
                result[i] = a
                a, b = b, a+b
        }
        return result
}
func sumArray(arr []float64) float64 {
        total := 0.0
        for _, val := range arr {
                total += val
        }
        return total
}
func minSlice(slc []int) int {
        min := slc[0]
        for _, val := range slc {
                if val < min {
                        min = val
                }
        }
        return min
}
func maxSlice(slc []int) int {
        max := slc[0]
        for _, val := range slc {
                if val > max {
                        max = val
                }
        }
        return max
}
func magnifySlice(s []int, fac int) []int {
        for i := range s {
                s[i] = s[i] * fac
        }
        return s
}
func filterSlice(s []int, fn func(int) bool) []int {
        res := make([]int, 0)
        for i := range s {
                if fn(i) {
                        res = append(res, i)
                }

        }
        return res
}
func even(n int) bool {
        return n%2 == 0
}

func stringSplit(s string, i int) (a, b string) {
        a = s[:i]
        b = s[i:]
        return a, b
}
func mapFunction(fn func(int) int, s []int) []int {
        res := make([]int, len(s))
        for i, j := range s {
                res[i] = fn(j)
        }
        return res
}
