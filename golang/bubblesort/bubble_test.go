package main

import "testing"

func TestBubble(t *testing.T) {

	/* Simple case: [5, 4, 3, 0, 0] should output [0, 0, 3, 4, 5] */
	s1 := make([]float64, 5)
	s1[0] = 5
	s1[1] = 4
	s1[2] = 3
	u := bubblesort(s1)
	if u[0] == 0 && u[1] == 0 && u[2] == 3 && u[3] == 4 && u[4] == 5{
	
	} else{
		t.Errorf("Slice %v is not sorted!", u);
	}
	
	/* Negative numbers: [4, 15, -0.001, 10, -5, -6.5] should output [-6.5, -5, -0.001, 4, 10, 15] */
	m := make([]float64, 6)
	m[0] = 4
	m[1] = 15
	m[2] = -0.001
	m[3] = 10
	m[4] = -5
	m[5] = -6.5
	p := bubblesort(m)
	if p[0] == -6.5 && p[1] == -5 && p[2] == -0.001 && p[3] == 4 && p[4] == 10 && p[5] == 15{
	
	} else{
		t.Errorf("Slice %v is not sorted!", p);
	}
}
