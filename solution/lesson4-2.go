package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func PermCheck(A []int) int {
	// Implement your solution here
	neededNumber := make(map[int]bool)
	for i := range A {
		neededNumber[i+1] = true
	}

	for _, num := range A {
		if _, exist := neededNumber[num]; exist {
			delete(neededNumber, num)
		}

		if len(neededNumber) == 0 {
			return 1
		}
	}

	return 0
}
