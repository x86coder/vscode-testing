package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	// Make slice of type "[]float64" -- slice of floats
	//  with 1 element.
	N := make([]float64, 1);

	fmt.Println("bubblesort v0.2 -- by x86coder");
	
	if len(os.Args) < 2{
		fmt.Println("Usage: bubblesort 5 2 1 -1 0 ...");
	}else{
	
		for i := 0; i < len(os.Args) -1; i++{
			//fmt.Println(os.Args[i+1])
			k, err := strconv.ParseFloat(os.Args[i+1], 64)
			if err == nil{
				N = append(N, k)
			}
		}

		fmt.Println(" > N=", N);
		sortedNumbers := bubblesort(N) // Create new variable, infer data type automatically
		fmt.Println(" > Sorted numbers =", sortedNumbers)
	}
}