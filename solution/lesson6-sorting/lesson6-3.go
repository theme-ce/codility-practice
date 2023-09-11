package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Triangle(A []int) int {
	// Implement your solution here
	l := len(A)

	if l < 3 {
		return 0
	}

	sort.Ints(A)

	for i := 0; i < l-2; i++ {
		P := A[i]
		Q := A[i+1]
		R := A[i+2]

		if P+Q > R {
			return 1
		}
	}

	return 0
}
