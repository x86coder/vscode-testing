package main

import (
	"fmt"
)

func bubblesort() {
	numbers := [...]int{5, 10, 3, 1, 4, 6, 2, 8}
	fmt.Println("bubblesort v0.1 - by x86coder")
	fmt.Println("numbers = ", numbers)

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				var aux = numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = aux
			}
		}
	}

	fmt.Println("sorted numbers =", numbers)
}
