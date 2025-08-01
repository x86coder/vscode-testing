package main

// How to create array of type: [8]int	
//numbers := [...]int{5, 10, 3, 1, 4, 6, 2, 8}

func bubblesort(numbers []float64) []float64{
	
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				var aux = numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = aux
			}
		}
	}
	
	return numbers
}