package main

import("testing")

func TestFactorial(t *testing.T){
	if factorial(0) != 1{
		t.Errorf("Expected returned value to be: 1")
	}
	if factorial(1) != 1{
		t.Errorf("Expected returned value to be: 1")
	}
	if factorial(2) != 2{
		t.Errorf("Expected returned value to be: 2")
	}
	if factorial(3) != 6{
		t.Errorf("Expected returned value to be: 6")
	}
	if factorial(4) != 24{
		t.Errorf("Expected returned value to be: 1")
	}
	if factorial(5) != 120{
		t.Errorf("Expected returned value to be: 120")
	}
}