package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MinAvgTwoSlice(A []int) int {
	// Implement your solution here
	N := len(A)
	minAvg := float64(A[0]+A[1]) / 2.0
	minAvgPos := 0

	for i := 0; i < N-2; i++ {
		avg2 := float64(A[i]+A[i+1]) / 2.0
		if avg2 < minAvg {
			minAvg = avg2
			minAvgPos = i
		}

		avg3 := float64(A[i]+A[i+1]+A[i+2]) / 3.0
		if avg3 < minAvg {
			minAvg = avg3
			minAvgPos = i
		}
	}

	avg2 := float64(A[N-2]+A[N-1]) / 2.0
	if avg2 < minAvg {
		minAvg = avg2
		minAvgPos = N - 2
	}

	return minAvgPos
}
