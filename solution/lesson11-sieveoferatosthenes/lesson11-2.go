package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func CountSemiprimes(N int, P []int, Q []int) []int {
	// Implement your solution here
	primes := make([]bool, N+1)
	for i := range primes {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	for i := 2; i*i <= N; i++ {
		if primes[i] {
			for j := i * i; j <= N; j += i {
				primes[j] = false
			}
		}
	}

	semiprimes := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if !primes[i] {
			continue
		}
		for j := i; j*i <= N; j++ {
			if primes[j] {
				semiprimes[i*j] = true
			}
		}
	}

	prefixSum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		prefixSum[i] = prefixSum[i-1]
		if semiprimes[i] {
			prefixSum[i]++
		}
	}

	M := len(P)
	result := make([]int, M)
	for i := 0; i < M; i++ {
		result[i] = prefixSum[Q[i]] - prefixSum[P[i]-1]
	}

	return result
}
