package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Fish(A []int, B []int) int {
	// Implement your solution here
	stack := []int{}
	aliveCount := 0

	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			stack = append(stack, A[i])
		} else {
			for len(stack) > 0 {
				topFish := stack[len(stack)-1]
				if topFish > A[i] {
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}

			if len(stack) == 0 {
				aliveCount++
			}
		}
	}

	return aliveCount + len(stack)
}
