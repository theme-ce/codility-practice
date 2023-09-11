package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func NumberOfDiscIntersections(A []int) int {
	// Implement your solution here
	N := len(A)
	left := make([]int, N)
	right := make([]int, N)

	for i := 0; i < N; i++ {
		left[i] = i - A[i]
		right[i] = i + A[i]
	}

	sort.Ints(left)
	sort.Ints(right)

	intersect := 0
	j := 0

	for i := 0; i < N; i++ {
		for j < N && right[i] >= left[j] {
			j++
		}
		intersect += j - i - 1

		if intersect > 10000000 {
			return -1
		}
	}

	return intersect
}
