package sequence

import (
	"errors"
	"sort"
)

// Sequence define the base type of DNA, Protein, etc.
type Sequence string

// Length method to get the length
func (s *Sequence) Length() int {
	return len(*s)
}

// Seq method return the sequence as string
func (s *Sequence) Seq() string {
	return string(*s)
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

// Uniq return the uniq characters in the sequence
func (s *Sequence) Uniq() []string {
	m := make(map[byte]int)
	res := make([]string, 0)
	b := []byte(*s)
	c := -1
	for i := 0; i < len(b); i++ {
		if _, ok := m[b[i]]; !ok {
			c += 1
			m[b[i]] = c
			res = append(res, string(b[i]))
		}
	}
	return res
}

// Range can select a fragment of the sequence, inclusive on the start and exclusive on the end index
// if the end index is omitted, fragment will end till the last character
func (s *Sequence) Range(index ...int) (string, error) {
	var start = index[0]
	var end int = 0
	if len(index) > 1 {
		end = index[1]
	}
	h := len(*s)
	switch {
	case start > h:
		return "", errors.New("start index out of range: too large")
	case start < -h:
		return "", errors.New("start index out of range: too small")
	case start < 0:
		start += h
	}
	switch {
	case end > h:
		return "", errors.New("end index out of range: too large")
	case end < -h:
		return "", errors.New("end index out of range: too small")
	case end <= 0:
		end += h
	}
	switch {
	case end < start:
		return "", errors.New("start index is larger than end index")
	default:
		return (string(*s))[start:end], nil
	}
}
