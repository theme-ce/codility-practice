package solution

import "math"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson3Solution3(A []int) int {
	// Implement your solution here
	totalSum := 0
	for _, num := range A {
		totalSum += num
	}

	leftSum := 0
	rightSum := totalSum
	minDiff := math.MaxInt64

	for p := 0; p < len(A)-1; p++ {
		leftSum += A[p]
		rightSum -= A[p]
		diff := int(math.Abs(float64(leftSum - rightSum)))

		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
