package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Lesson2Solution1(A []int, K int) []int {
	// Implement your solution here
	res := A

	if len(res) <= 0 {
		return nil
	}

	if len(res) == K {
		return res
	}

	for i := 0; i < K; i++ {
		last := res[len(res)-1]
		cutLast := res[:len(res)-1]
		newArray := []int{last}
		newArray = append(newArray, cutLast...)
		res = newArray
	}
	return res
}
