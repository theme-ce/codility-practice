package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func StoneWall(H []int) int {
	// Implement your solution here
	var stack []int
	blockCount := 0

	for _, height := range H {
		for len(stack) > 0 && stack[len(stack)-1] > height {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 || stack[len(stack)-1] < height {
			stack = append(stack, height)
			blockCount++
		}
	}

	return blockCount
}
