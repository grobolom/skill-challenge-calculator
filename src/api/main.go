package main

import (
	"encoding/json"
	"net/http"
)

// Probability represents the probably response based on the input
type Probability struct {
	ChanceOfSuccess int `json:"ChanceOfSuccess"`
}

// ProbabilityHandler just sends a dummy response for now
func ProbabilityHandler(w http.ResponseWriter, r *http.Request) {
	probability := Probability{82}

	json.NewEncoder(w).Encode(probability)
}

func main() {
	http.HandleFunc("/", ProbabilityHandler)
	http.ListenAndServe(":5051", nil)
}
