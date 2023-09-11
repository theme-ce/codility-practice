package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func CountNonDivisible(A []int) []int {
	// Implement your solution here
	N := len(A)

	counts := make(map[int]int)
	for _, a := range A {
		counts[a]++
	}

	divisors := make(map[int]int)
	for a := range counts {
		divisorCount := 0
		for d := 1; d*d <= a; d++ {
			if a%d == 0 {
				divisorCount += counts[d]
				if d*d != a {
					divisorCount += counts[a/d]
				}
			}
		}
		divisors[a] = divisorCount
	}

	result := make([]int, N)
	for i, a := range A {
		result[i] = N - divisors[a]
	}

	return result
}
