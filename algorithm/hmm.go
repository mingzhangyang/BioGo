package algorithm

import "log"

// This file try to implement HMM using golang

// // Probability type describing the probability of an event
// type Probability struct {
// 	event       string
// 	probability float64
// }

/*
	TransistionProbablities
	(s0: START state; s1-N: states; sE: END state)

		s0	s1	s2	s3	...		sN	sE
	s0	[]	[]	[]	[]	...		[]	[]
	s1	[]	[]	[]	[]	...		[]	[]
	s2	[]	[]	[]	[]	...		[]	[]
	s3	[]	[]	[]	[]	...		[]	[]
	.	.	.	.	.	...		.	.
	.	.	.	.	.	...		.	.
	.	.	.	.	.	...		.	.
	sN	[]	[]	[]	[]	...		[]	[]
	SE	.0	.0	.0	.0	...		.0	.0


	EmissionProbablities
	(s0: START state; s1-N: states; o0-M: observations)

		o0	o1	o2	o3	...		oM
	s0	.0	.0	.0	.0	...		.0
	s1	[]	[]	[]	[]	...		[]
	s2	[]	[]	[]	[]	...		[]
	s3	[]	[]	[]	[]	...		[]
	.	.	.	.	.	...		.
	.	.	.	.	.	...		.
	.	.	.	.	.	...		.
	sN	[]	[]	[]	[]	...		[]
	SE	.0	.0	.0	.0	...		.0
*/

// HMM type
// states: array of states
// observations: array of observations
// stateIndexMap: find the index of a state in TP / EP tables
// observationIndexMap: find the index of an observation in EP table
// TPTable: transition probablity table
// EPTable: emission probablity table
type HMM struct {
	states              []string
	observations        []string
	stateIndexMap       map[string]int
	observationIndexMap map[string]int
	TPTable             [][]float64
	EPTable             [][]float64
}

// NewHMM create a HMM model
func NewHMM(states, obs []string) *HMM {
	// statesIndex starts from 1, START state as 0
	ms := make(map[string]int)
	for i, s := range states {
		ms[s] = (i + 1)
	}
	ms["START"] = 0
	ms["END"] = (len(states) + 1)

	// observationIndex starts from 0
	os := make(map[string]int)
	for i, s := range obs {
		os[s] = i
	}

	// len(states) + 2 = START + states + END
	tp := make([][]float64, (len(states) + 2))
	ep := make([][]float64, (len(states) + 2))
	for i := 0; i < len(tp); i++ {
		tp[i] = make([]float64, len(states)+2)
	}
	for i := 0; i < len(ep); i++ {
		ep[i] = make([]float64, len(obs))
	}
	return &HMM{
		states:              states,
		observations:        obs,
		stateIndexMap:       ms,
		observationIndexMap: os,
		TPTable:             tp,
		EPTable:             ep,
	}
}

// SetTransitionProbablity set transition probablity
// stateA: current state; stateB: next state; p: probablity
func (h *HMM) SetTransitionProbablity(stateA, stateB string, p float64) {
	var i, j int
	ok := true
	if i, ok = h.stateIndexMap[stateA]; !ok {
		log.Panic(stateA + " not found")
	}
	if j, ok = h.stateIndexMap[stateB]; !ok {
		log.Panic(stateB + " not found")
	}
	h.TPTable[i][j] = p
}

// SetEmissionProbablity set emission probablity
// state: state; obs: observation; p: probablity
func (h *HMM) SetEmissionProbablity(state, obs string, p float64) {
	var i, j int
	ok := true
	if i, ok = h.stateIndexMap[state]; !ok {
		log.Panic(state + " not found")
	}
	if j, ok = h.observationIndexMap[obs]; !ok {
		log.Panic(obs + " not found")
	}
	h.EPTable[i][j] = p
}

// Monitor method to calculate the probability of a state given observation
// the arguments are a list of CONSECUTIVE observations
// Implement the Forward algorithm
func (h *HMM) Monitor(obs []string) float64 {
	var res float64
	// len(h.states) + 2 = START + states + END
	var dp = make([]float64, len(h.states)+2)

	// to avoid allocate memory space repeatedly
	var tmp = make([]float64, len(dp))
	var zeros = make([]float64, len(dp))

	// init
	dp = h.TPTable[0]
	// oi: observation index
	oi, ok := h.observationIndexMap[obs[0]]
	if !ok {
		log.Panic(obs[0] + " not found")
	}
	for j := 1; j < len(dp)-1; j++ {
		dp[j] *= h.EPTable[j][oi]
	}
	// foward
	for i := 1; i < len(obs); i++ {
		oi, ok := h.observationIndexMap[obs[i]]
		if !ok {
			log.Panic(obs[0] + " not found")
		}
		// reset tmp to all zeros
		// for idx := range tmp {
		// 	tmp[idx] = 0
		// }
		copy(tmp, zeros)

		// begin to update tmp
		for j := 1; j < len(dp)-1; j++ {
			for k := 1; k < len(dp)-1; k++ {
				tmp[k] += dp[j] * h.TPTable[j][k] * h.EPTable[k][oi]
			}
		}

		// update dp
		copy(dp, tmp)
	}
	for _, v := range dp {
		res += v
	}
	return res
}

// MostLikelyPath method calculate the most likely path through which
// the hidden states go
func (h *HMM) MostLikelyPath(seq []string) {

}
