package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MaxProductOfThree(A []int) int {
	// Implement your solution here
	sort.Ints(A)
	l := len(A)

	maxLastTriplet := A[l-3] * A[l-2] * A[l-1]
	maxFirstTriplet := A[0] * A[1] * A[l-1]

	if maxLastTriplet >= maxFirstTriplet {
		return maxLastTriplet
	} else {
		return maxFirstTriplet
	}
}
