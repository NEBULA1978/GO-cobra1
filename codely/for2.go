package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	for i, n := range numbers {
		if n%2 == 0 {
			fmt.Printf("the number at index: %d is even\n", i)
		} else {
			fmt.Printf("the number at index: %d is odd\n", i)
		}

	}
}
