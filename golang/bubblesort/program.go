package main

import (
	"fmt"
	"os"
)

func bubblesort(numbers []int){
	fmt.Println(numbers);
}

func main(){
	N := make([]int, 1);
	fmt.Printf("%T", os.Args[1]);
	bubblesort(N)
}