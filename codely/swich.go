package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	for i, n := range numbers {
		switch n % 2 {
		case 0:
			fmt.Printf("the number at index: %d is even and ", i)
			fallthrough
		default:
			fmt.Printf("the number: %d is at index: %d\n", n, i)

		}
	}
}
