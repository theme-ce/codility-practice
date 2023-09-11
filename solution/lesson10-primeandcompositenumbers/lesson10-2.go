package solution

import (
	"math"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MinPerimeterRectangle(N int) int {
	// Implement your solution here
	sqrtN := int(math.Sqrt(float64(N)))

	minPerimeter := math.MaxInt64
	for i := 1; i <= sqrtN; i++ {
		if N%i == 0 {
			if i*i == N {
				if minPerimeter > 2*(i+i) {
					minPerimeter = 2 * (i + i)
				}
			} else {
				if minPerimeter > 2*(i+N/i) {
					minPerimeter = 2 * (i + N/i)
				}
			}
		}
	}
	return minPerimeter
}
