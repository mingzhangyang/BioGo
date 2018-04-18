package util

// CountSubsequentOneChar method count the numbers of every unique char after any char in the string
func CountSubsequentOneChar(str string, dict map[string]map[string]int) map[string]map[string]int {
	if dict == nil {
		dict = make(map[string]map[string]int)
	}
	for i := 1; i < len(str); i++ {
		a, b := string(str[i-1]), string(str[i])
		if dict[a] == nil {
			dict[a] = make(map[string]int)
		}
		dict[a][b]++
	}
	return dict
}

// CountPrecedingOneChar method count the numbers of every unique char before any char in the string
func CountPrecedingOneChar(str string, dict map[string]map[string]int) map[string]map[string]int {
	if dict == nil {
		dict = make(map[string]map[string]int)
	}
	for i := 1; i < len(str); i++ {
		a, b := string(str[i-1]), string(str[i])
		if dict[b] == nil {
			dict[b] = make(map[string]int)
		}
		dict[b][a]++
	}
	return dict
}