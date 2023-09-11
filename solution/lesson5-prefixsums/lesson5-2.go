package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func CountDiv(A int, B int, K int) int {
	// Implement your solution here
	if A%K == 0 {
		return (B / K) - (A / K) + 1
	}
	return (B / K) - (A / K)
}
