package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MinMaxDivision(K int, M int, A []int) int {
	// Implement your solution here
	low := 0
	high := 0

	for _, num := range A {
		high += num
		if num > low {
			low = num
		}
	}

	result := high
	for low <= high {
		mid := (high + low) / 2
		if isPossible(A, mid, K) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

func isPossible(arr []int, barrier int, k int) bool {
	block := 1
	sum := 0

	for _, num := range arr {
		if num > barrier {
			return false
		}

		if sum+num > barrier {
			sum = num
			block += 1
			if block > k {
				return false
			}
		} else {
			sum += num
		}
	}

	return true
}
