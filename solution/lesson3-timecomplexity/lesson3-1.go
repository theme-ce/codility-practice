package solution

import (
	"math"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func FrogJmp(X int, Y int, D int) int {
	// Implement your solution here
	distance := Y - X
	jumpCount := float64(distance) / float64(D)
	jumpCount = math.Ceil(jumpCount)
	return int(jumpCount)
}
