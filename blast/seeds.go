package blast

const aa = "ARNDCQEGHILKMFPSTWYV*"

// SimilarWords produce a list of similar words with score higher than threshold
// This function works only for words with length = 3
func SimilarWords(s string, t float64) []string {
	res := make([]string, 0)
	word := make([]byte, 3)
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			for k := 0; k < 21; k++ {
				word[0] = aa[i]
				word[1] = aa[j]
				word[2] = aa[k]
				v1 := BLOSUM62[string([]byte{s[0], aa[i]})]
				v2 := BLOSUM62[string([]byte{s[1], aa[j]})]
				v3 := BLOSUM62[string([]byte{s[2], aa[k]})]
				if float64(v1) + float64(v2) + float64(v3) >= t {
					res = append(res, string(word))
				}
			}
		}
	}
	return res
}