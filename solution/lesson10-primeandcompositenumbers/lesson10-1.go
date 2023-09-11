package solution

import "math"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func CountFactors(N int) int {
	// Implement your solution here
	count := 0
	sqrtN := int(math.Sqrt(float64(N)))

	for i := 1; i <= sqrtN; i++ {
		if N%i == 0 {
			if i*i == N {
				count += 1 // i and N/i are the same
			} else {
				count += 2 // i and N/i are distinct factors
			}
		}
	}
	return count
}
