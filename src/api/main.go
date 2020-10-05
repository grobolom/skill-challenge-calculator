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

// ProbabilityRangeHandler sends back an array of objects that represent the probability
// of success for each number of successes, based on the skill bonus, check DC, and number
// of failures.
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

	// I don't care about handling this for now
	if err != nil {
	}

	// this calculation should be moved out of this handler
	baseProbability := math.Min(float64(21-(DC-skillBonus))/20, 1)

	result := GetProbabilityRange(successes, failures, baseProbability)

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/v2", ProbabilityRangeHandler)
	http.ListenAndServe(":5051", nil)
}
