package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func FrogRiverOne(X int, A []int) int {
	// Implement your solution here
	neededNumber := make(map[int]bool)
	for i := 0; i < X; i++ {
		neededNumber[i+1] = true
	}

	for i, num := range A {
		if _, exist := neededNumber[num]; exist {
			delete(neededNumber, num)
		}

		if len(neededNumber) == 0 {
			return i
		}
	}
	return -1
}
