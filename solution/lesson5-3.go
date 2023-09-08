package solution

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	N := len(S)
	A := make([]int, N+1)
	C := make([]int, N+1)
	G := make([]int, N+1)

	for i := 0; i < N; i++ {
		A[i+1] = A[i]
		C[i+1] = C[i]
		G[i+1] = G[i]
		switch S[i] {
		case 'A':
			A[i+1]++
		case 'C':
			C[i+1]++
		case 'G':
			G[i+1]++
		}

		fmt.Printf("A: ")
		for a := range A {
			fmt.Printf("%d ", A[a])
		}
		fmt.Printf("\n")
		fmt.Printf("C: ")
		for c := range C {
			fmt.Printf("%d ", C[c])
		}
		fmt.Printf("\n")
		fmt.Printf("G: ")
		for g := range G {
			fmt.Printf("%d ", G[g])
		}
		fmt.Printf("\n")
	}

	M := len(P)
	result := make([]int, M)

	for i := 0; i < M; i++ {
		p, q := P[i], Q[i]
		if A[q+1]-A[p] > 0 {
			result[i] = 1
		} else if C[q+1]-C[p] > 0 {
			result[i] = 2
		} else if G[q+1]-G[p] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	return result
}
