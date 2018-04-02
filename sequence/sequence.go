package sequence

import "sort"

// Sequence define the base type of DNA, Protein, etc.
type Sequence string

// Length method to get the length
func (s *Sequence) Length() int {
	return len(*s)
}

// Reverse method return the reversed sequence
func (s *Sequence) Reverse() string {
	b := []byte((*s))
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Unit for counting the frequency of every single character
type Unit struct {
	char  string
	count int
}

// Composition returns the composition stat of the sequence
func (s *Sequence) Composition() []Unit {
	m := make(map[byte]int)
	b := []byte(*s)
	for i := 0; i < len(b); i++ {
		m[b[i]]++
	}
	res := make([]Unit, 0)
	for k, v := range m {
		res = append(res, Unit{string(k), v})
	}
	sort.SliceStable(res, func(i, j int) bool { return res[j].count < res[i].count })
	return res
}
