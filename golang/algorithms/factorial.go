package main

import(
	"fmt"
	"os"
	"strconv"
)
func factorial(n uint64) uint64{
	if n==0 {
		return 1	// Return standard definition 0! = 1
	} else{
		return n * factorial(n-1)
	}
}

func main(){
	fmt.Printf("factorial(n) - Usage: 'go run factorial.go 5'\n")
	param, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
		fmt.Printf(" > Calculating factorial of %s\n", os.Args[1])
		fmt.Printf(" > result = %d\n", factorial(param))
	} else{
		fmt.Println(" ** Error converting first parameter")
	}
}
