package calc

import "fmt"
import "github.com/fatih/color"

// returns sum of two integers
func Add(numbers ...int) int {
    sum := 0

    for _, num := range numbers {
        sum = sum + num
    }

    fmt.Println(color.RedString("test color"))

    return sum
}
