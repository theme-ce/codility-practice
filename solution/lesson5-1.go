package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson5Solution1(A []int) int {
	// Implement your solution here
	startCount := false
	count := 0
	currentPassingCar := 1
	for _, num := range A {
		if num == 0 {
			if startCount {
				currentPassingCar += 1
			} else {
				startCount = true
			}
		}

		if num == 1 && startCount {
			count += 1 * currentPassingCar
		}
	}

	if count > 1000000000 {
		count = -1
	}

	return count
}
