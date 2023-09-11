package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func PermMissingElem(A []int) int {
	// Implement your solution here
	n := len(A)
	totalSum := (n + 1) * (n + 2) / 2

	arraySum := 0
	for _, i := range A {
		arraySum += i
	}

	return totalSum - arraySum
}
