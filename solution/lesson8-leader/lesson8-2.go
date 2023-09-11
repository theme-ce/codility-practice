package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func EquiLeader(A []int) int {
	// Implement your solution here
	count := make(map[int]int)
	l := len(A)

	for _, num := range A {
		count[num]++
	}

	var leader int
	leaderCount := 0

	for key, value := range count {
		if value > l/2 {
			leader = key
			leaderCount = value
		}
	}

	equiLeaders := 0
	leftCount := 0

	for i := 0; i < l; i++ {
		if A[i] == leader {
			leftCount++
		}

		if leftCount > (i+1)/2 && leaderCount-leftCount > (l-i-1)/2 {
			equiLeaders++
		}
	}

	return equiLeaders
}
