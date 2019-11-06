package misc

// Perm gives all possible combinations of a given list of characters with a given length
// e.g. given []byte{'A', 'B', 'C'} and a length of 2, the below list will be generated
// {"AA", "AB", "AC", "BA", "BB", "BC", "CA", "CB", "CC"}
// The chars argument should contains unique characters
func Perm(chars []byte, length int) []string {
	res := make([]string, 0)
	if length == 1 {
		for _, c := range chars {
			res = append(res, string(c))
		}
		return res
	}
	tmp := make(map[int][]string)
	tmp[1] = make([]string, 0)
	for _, c := range chars {
		tmp[1] = append(tmp[1], string(c))
	}

	var i = 2

	for i <= length {
		tmp[i] = make([]string, 0)
		for _, c := range chars {
			for _, s := range tmp[i-1] {
				tmp[i] = append(tmp[i], string(c) + s)
			}
		}
		i++
	}
	return tmp[length]
}