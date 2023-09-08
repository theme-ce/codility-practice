package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson4Solution4(A []int) int {
	// Implement your solution here
	uniqueNum := make(map[int]bool)

	for _, num := range A {
		if num > 0 {
			uniqueNum[num] = true
		}
	}

	for i := 1; i <= len(A)+1; i++ {
		if !uniqueNum[i] {
			return i
		}
	}

	return 1
}
