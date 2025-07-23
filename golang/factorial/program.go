package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("factorial(n) - Usage: 'go run factorial.go 5'\n")
	param, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
		fmt.Printf(" > Calculating factorial of %s\n", os.Args[1])
		fmt.Printf(" > result = %d\n", factorial(param))
	} else {
		fmt.Println(" ** Error converting first parameter")
	}
}
