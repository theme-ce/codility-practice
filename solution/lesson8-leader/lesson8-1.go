package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Dominator(A []int) int {
	// Implement your solution here
	count := make(map[int]int)

	for _, num := range A {
		count[num]++
	}

	dominator := -1

	for key, value := range count {
		if value > len(A)/2 {
			dominator = key
			break
		}
	}

	if dominator == -1 {
		return -1
	}

	for i, num := range A {
		if num == dominator {
			return i
		}
	}

	return -1
}
