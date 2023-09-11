package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Flags(A []int) int {
	// Implement your solution here
	l := len(A)
	peaks := make([]int, 0)

	for i := 1; i < l-1; i++ {
		if A[i-1] < A[i] && A[i+1] < A[i] {
			peaks = append(peaks, i)
		}
	}

	maxFlags := 0
	for k := 1; k <= len(peaks); k++ {
		count := 0
		lastFlag := -1
		for _, peak := range peaks {
			if lastFlag == -1 || peak-lastFlag >= k {
				count++
				lastFlag = peak
			}
		}

		if count >= k {
			maxFlags = k
		} else {
			break
		}
	}

	return maxFlags
}
