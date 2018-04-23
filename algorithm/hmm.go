package algorithm

// This file try to implement HMM using golang

// Probability type describing the probability of an event
type Probability struct {
	event       string
	probability float64
}

// HMM type
type HMM struct {
	States                []string
	Observations          []string
	TransistionProbablities map[string][]Probability
	EmissionProbabilities   map[string][]Probability
	InitialProbabilities []Probability
}

// Monitor method to calculate the probability of a state given observation
// the arguments are a list of CONSECUTIVE observations
func (h *HMM) Monitor(obs ...string) {

}

// MostLikelyPath method calculate the most likely path through which
// the hidden states go
func (h *HMM) MostLikelyPath(seq []string) {

}
