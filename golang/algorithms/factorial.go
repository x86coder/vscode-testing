package main

import("fmt")
import("os")
import("strconv")

func factorial(n uint64) uint64{
	if n==0 {
		return 1
	} else{
		return n * factorial(n-1)
	}
}

func main(){
	param, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
		fmt.Printf(" > Calculating factorial of %s\n", os.Args[1])
		fmt.Printf(" > result = %d\n", factorial(param))
	} else{
		fmt.Println("Error converting first parameter")
	}
}
