package blast

import (
	"strings"
	"fmt"
)

// BlosumScore compute the score between two strings
// The length of s1 and s2 should be equal, and space should be represented as *
func BlosumScore(s1, s2 string) int {
	if len(s1) != len(s2) {
		panic("Two strings with the same length are expected.")
	}
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)
	var res int
	for i, max := 0, len(s1); i < max; i++ {
		key := string([]byte{s1[i], s2[i]})
		if v, ok := BLOSUM62[key]; ok {
			res += v
		} else {
			panic(fmt.Sprintf("Invalid combination of %s and %s\n", string(s1[i]), string(s2[i])))
		}
	}
	return res
}