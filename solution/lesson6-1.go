package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// Implement your solution here
	distinct := make(map[int]bool)

	for _, num := range A {
		distinct[num] = true
	}

	return len(distinct)
}
