package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Brackets(S string) int {
	// Implement your solution here
	stack := []rune{}

	for _, s := range S {
		switch s {
		case '(', '{', '[':
			stack = append(stack, s)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return 0
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return 0
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
