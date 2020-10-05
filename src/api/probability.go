package main

import "math"

// Probability represents the probably response based on the input
type Probability struct {
	ChanceOfSuccess float64 `json:"probability"`
	Successes       int     `json:"successes"`
}

// GetProbabilityRange gets a series of probabilities for different counts of successes
func GetProbabilityRange(successRange []int, failures int, probability float64) []Probability {
	var pr = []Probability{}

	for i := 0; i < len(successRange); i++ {
		pr = append(pr, Probability{CalculateTotalProbability(successRange[i], failures, probability), successRange[i]})
	}

	return pr
}

// Factorial does something
func Factorial(integer int) int {
	sum := 1

	for i := 1; i <= integer; i++ {
		sum = i * sum
	}

	return sum
}

// CalculateProbability does something
func CalculateProbability(successes int, trials int, probability float64) float64 {
	firstTerm := Factorial(trials-1) / (Factorial(successes-1) * Factorial(trials-successes))
	secondTerm := math.Pow(probability, float64(successes))
	thirdTerm := math.Pow(1-probability, float64(trials-successes))

	return float64(firstTerm) * secondTerm * thirdTerm
}

// CalculateTotalProbability does something
func CalculateTotalProbability(successes int, failures int, probability float64) float64 {
	var sum float64

	for i := 0; i < failures; i++ {
		sum += CalculateProbability(successes, successes+i, probability)
	}

	return sum
}
