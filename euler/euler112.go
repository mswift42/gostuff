// Bouncy numbers Problem 112

// Working from left-to-right if no digit is exceeded by the digit to its
// left it is called an increasing number; for example, 134468.

// Similarly if no digit is exceeded by the digit to its right it is
// called a decreasing number; for example, 66420.

// We shall call a positive integer that is neither increasing nor
// decreasing a "bouncy" number; for example, 155349.

// Clearly there cannot be any bouncy numbers below one-hundred, but just
// over half of the numbers below one-thousand (525) are bouncy. In fact,
// the least number for which the proportion of bouncy numbers first
// reaches 50% is 538.

// Surprisingly, bouncy numbers become more and more common and by the
// time we reach 21780 the proportion of bouncy numbers is equal to 90%.

// Find the least number for which the proportion of bouncy numbers is
// exactly 99%.

package main

import "fmt"
import "strings"
import "strconv"

func numberToSlice(i int) []int {
        strslice := strings.Split(strconv.Itoa(i), "")
        intslice := make([]int, len(strslice))
        for i, j := range strslice {
                num, err := strconv.Atoi(j)
                if err != nil {
                        panic(err)
                }
                intslice[i] = num
        }
        return intslice
}

func isIncreasing(sl []int) bool {
        if len(sl) < 2 {
                return true
        }
        for i := 0; i < len(sl)-1; i++ {
                if sl[i+1] < sl[i] {
                        return false
                }
        }
        return true
}
func isDecreasing(sl []int) bool {
        if len(sl) < 2 {
                return true
        }
        for i, j := 0, 1; j < len(sl); i, j = i+1, j+1 {
                if sl[j] > sl[i] {
                        return false
                }
        }
        return true
}

func isBouncy(sl []int) bool {
        return !isDecreasing(sl) && !isIncreasing(sl)
}

func euler112() int {
        sumbouncy := 0
        count := 1
        for {
                count++
                if isBouncy(numberToSlice(count)) {
                        sumbouncy++
                }

                if (sumbouncy * 100 / count) >= 99 {
                        return count
                }
        }
}

func main() {
        fmt.Println(euler112())
}
