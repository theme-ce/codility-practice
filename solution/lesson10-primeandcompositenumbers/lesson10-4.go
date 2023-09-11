package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Peaks(A []int) int {
	// Implement your solution here
	N := len(A)

	peaks := []int{}
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	numPeaks := len(peaks)

	maxBlocks := 0
	for k := 1; k <= numPeaks; k++ {
		if N%k != 0 {
			continue
		}

		blockSize := N / k
		foundPeaks := 0
		blockCount := 0

		for _, peak := range peaks {
			if peak/blockSize == blockCount {
				foundPeaks++
				blockCount++
			}

			if blockCount == k {
				break
			}
		}

		if blockCount == k {
			maxBlocks = k
		}
	}

	return maxBlocks
}
