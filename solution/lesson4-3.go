package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson4Solution3(N int, A []int) []int {
	// Implement your solution here
	arr := make([]int, N)
	currentMax := 0
	lastMax := 0
	for _, num := range A {
		if num < N+1 && num > 0 {
			if arr[num-1] < lastMax {
				arr[num-1] = lastMax
			}

			arr[num-1]++
			if arr[num-1] > currentMax {
				currentMax = arr[num-1]
			}
		} else if num == N+1 {
			lastMax = currentMax
		}
	}

	for i := range arr {
		if arr[i] < lastMax {
			arr[i] = lastMax
		}
	}

	return arr
}
