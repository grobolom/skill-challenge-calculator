package main

import "testing"

func TestFactorial(t *testing.T) {
	total := Factorial(5)
	if total != 120 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 120)
	}
}

func TestCalculateProbability(t *testing.T) {
	result := CalculateProbability(3, 5, 0.5)
	if result != 0.1875 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 0.1875)
	}
}

func TestCalculateTotalProbability(t *testing.T) {
	result := CalculateTotalProbability(3, 3, 0.5)

	if result != 0.5 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 0.5)
	}

	result = CalculateTotalProbability(2, 3, 0.5)

	if result != 0.6875 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 0.6875)
	}
}

func TestGetProbabilityRange(t *testing.T) {
	r := []int{2, 3}

	result := GetProbabilityRange(r, 3, 0.5)

	expected := Probability{0.6875}
	if result[0] != expected {
		t.Errorf("Probabilities were incorrect, got: %v, want: %v.", result[0], 0.6875)
	}

	expected = Probability{0.5}
	if result[1] != expected {
		t.Errorf("Probabilities were incorrect, got: %v, want: %v.", result[1], 0.5)
	}
}
