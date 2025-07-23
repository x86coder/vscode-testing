package main

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1 // Return standard definition 0! = 1
	} else {
		return n * factorial(n-1)
	}
}
