package util

import (
	"math"
)

// EntropyOfString calculate the Shannon entropy of the string
// Probablity is calculate as the frequency of each unique character
func EntropyOfString(str string) float64 {
	m := make(map[rune]int)
	var l int
	for _, c := range str {
		m[c]++
		l++
	}

	var e float64
	var t float64
	for _, v := range m {
		t = (float64(v) / float64(l))
		e += (t * math.Log2(t))
	}
	if e < 0 {
		return -e
	}
	// to avoid -0, when e is not smaller than 0, the only case is e == 0
	return e // this line will only execute when e == 0 because e never be positive
}