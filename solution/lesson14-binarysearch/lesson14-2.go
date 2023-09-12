package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int, C []int) int {
	// Implement your solution here
	N := len(A)
	M := len(C)
	result := -1

	// We keep track of the rightmost nail for each plank
	rightmostNail := make([]int, 2*M)

	for i := 0; i < 2*M; i++ {
		rightmostNail[i] = -1
	}

	// Populate the rightmostNail array
	for i := 0; i < M; i++ {
		rightmostNail[C[i]] = i
	}

	for i := 1; i < 2*M; i++ {
		rightmostNail[i] = max(rightmostNail[i], rightmostNail[i-1])
	}

	// Iterate through all the planks to find the rightmost nail
	for i := 0; i < N; i++ {
		minNail := rightmostNail[B[i]]
		if minNail == -1 {
			return -1
		}
		for j := A[i]; j < B[i]; j++ {
			minNail = min(minNail, rightmostNail[j])
		}

		result = max(result, minNail)
	}

	return result + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
