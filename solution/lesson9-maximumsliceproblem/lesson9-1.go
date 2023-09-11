package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func MaxProfit(A []int) int {
	// Implement your solution here
	if len(A) < 2 {
		return 0
	}

	minPrice := A[0]
	maxProfit := 0

	for _, price := range A {
		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
