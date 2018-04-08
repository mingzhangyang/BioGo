package misc

// Combination returns all the combinations of n numbers ranged from 1 to m
// source: http://rosettacode.org/wiki/Combinations#Go
func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
