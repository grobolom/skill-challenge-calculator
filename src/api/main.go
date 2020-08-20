package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

// Probability represents the probably response based on the input
type Probability struct {
	ChanceOfSuccess float64 `json:"ChanceOfSuccess"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// ProbabilityHandler just sends a dummy response for now
func ProbabilityHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	successes := 4
	failures := 3
	DC := 5
	skillBonus := 1
	var err error

	formSuccesses := r.FormValue("successes")
	formFailures := r.FormValue("failures")
	formSkillBonus := r.FormValue("skillBonus")
	formDC := r.FormValue("dc")

	if formSuccesses != "" {
		successes, err = strconv.Atoi(formSuccesses)
	}

	if formFailures != "" {
		failures, err = strconv.Atoi(formFailures)
	}

	if formDC != "" {
		DC, err = strconv.Atoi(formDC)
	}

	if formSkillBonus != "" {
		skillBonus, err = strconv.Atoi(formSkillBonus)
	}

	// for now, it's fine to ignore this shit
	if err != nil {
	}

	probability := math.Min(float64(21-(DC-skillBonus))/20, 1)

	result := Probability{CalculateTotalProbability(successes, failures, probability)}

	json.NewEncoder(w).Encode(result)
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

func main() {
	http.HandleFunc("/", ProbabilityHandler)
	http.ListenAndServe(":5051", nil)
}
