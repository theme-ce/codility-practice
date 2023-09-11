package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MaxDoubleSliceSum(A []int) int {
	// Implement your solution here
	l := len(A)
	if l < 3 {
		return 0
	}

	maxStartingHere := make([]int, l)
	maxEndingHere := make([]int, l)

	for i := 1; i < l-1; i++ {
		maxEndingHere[i] = max(0, maxEndingHere[i-1]+A[i])
	}

	for i := l - 2; i > 0; i-- {
		maxStartingHere[i] = max(0, maxStartingHere[i+1]+A[i])
	}

	maxDoubleSliceSum := 0
	for i := 1; i < l-1; i++ {
		maxDoubleSliceSum = max(maxDoubleSliceSum, maxEndingHere[i-1]+maxStartingHere[i+1])
	}

	return maxDoubleSliceSum
}
