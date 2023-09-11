package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MaxSliceSum(A []int) int {
	// Implement your solution here
	maxSum := A[0]
	currentSum := A[0]

	for i := 1; i < len(A); i++ {
		currentSum = max(A[i], currentSum+A[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
