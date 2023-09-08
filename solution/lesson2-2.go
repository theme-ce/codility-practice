package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson2Solution2(A []int) int {
	// Implement your solution here
	mapPair := map[int]int{}
	for i := range A {
		mapPair[A[i]] += 1
	}

	for key, value := range mapPair {
		if value%2 != 0 {
			return key
		}
	}
	return 0
}
