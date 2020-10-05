package main

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN"))
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

	baseProbability := math.Min(float64(21-(DC-skillBonus))/20, 1)

	result := Probability{CalculateTotalProbability(successes, failures, baseProbability), successes}

	json.NewEncoder(w).Encode(result)
}

func ProbabilityRangeHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	successes := []int{3, 4, 5, 6, 7, 8, 9, 10, 11}
	failures := 3
	DC := 14
	skillBonus := 3
	var err error

	formFailures := r.FormValue("failures")
	formSkillBonus := r.FormValue("skillBonus")
	formDC := r.FormValue("dc")

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

	baseProbability := math.Min(float64(21-(DC-skillBonus))/20, 1)

	result := GetProbabilityRange(successes, failures, baseProbability)

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", ProbabilityHandler)
	http.HandleFunc("/v2", ProbabilityRangeHandler)
	http.ListenAndServe(":5051", nil)
}
