package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"strconv"
	"strings"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func BinaryGap(N int) int {
	// Implement your solution here
	bin := integerToBinary(N)
	res := strings.Split(bin, "1")
	longest := 0
	for i := range res {
		if len(res[i]) > longest {
			longest = len(res[i])
		}
	}
	return longest
}

func integerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
