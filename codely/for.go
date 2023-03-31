package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Println("for...range example:")

	for i, n := range numbers {
		fmt.Printf("the number: %d is at index: %d\n", n, i)
	}
	// No declaramos el valor de i
	// for _, n := range numbers {
	// 	fmt.Printf("the number: %d is at index: %d\n", n, i)
	// }

	fmt.Println("simple for example:")

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("the number: %d is at index: %d\n", numbers[i], i)
	}
}
