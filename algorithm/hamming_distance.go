package algorithm

// HammingDistance count the number of positions
// at which the correspondingg symbols are different
// if the two strings of distinct length, return -1
func HammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	n := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			n++
		}
	}
	return n
}
