package blast

import (

)

// Split a string into an array of words of n characters
// Only work for strings containing ASCII characters solely
// The function takes for granted that the input string are valid
func Split(s string, n int) map[string][]int {
	dict := make(map[string][]int)
	for i, max := 0, len(s)-n+1; i < max; i++ {
		sub := s[i:i+n]
		if v, ok := dict[sub]; ok {
			dict[sub] = append(v, i)
		} else {
			dict[sub] = []int{i}
		}
	}
	return dict
}