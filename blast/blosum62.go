package blast

import (
	"strings"
	"strconv"
	"fmt"
)

const aa = "ARNDCQEGHILKMFPSTWYV*"

const s = `   A  R  N  D  C  Q  E  G  H  I  L  K  M  F  P  S  T  W  Y  V  B  Z  X  *
A  4 -1 -2 -2  0 -1 -1  0 -2 -1 -1 -1 -1 -2 -1  1  0 -3 -2  0 -2 -1  0 -4 
R -1  5  0 -2 -3  1  0 -2  0 -3 -2  2 -1 -3 -2 -1 -1 -3 -2 -3 -1  0 -1 -4 
N -2  0  6  1 -3  0  0  0  1 -3 -3  0 -2 -3 -2  1  0 -4 -2 -3  3  0 -1 -4 
D -2 -2  1  6 -3  0  2 -1 -1 -3 -4 -1 -3 -3 -1  0 -1 -4 -3 -3  4  1 -1 -4 
C  0 -3 -3 -3  9 -3 -4 -3 -3 -1 -1 -3 -1 -2 -3 -1 -1 -2 -2 -1 -3 -3 -2 -4 
Q -1  1  0  0 -3  5  2 -2  0 -3 -2  1  0 -3 -1  0 -1 -2 -1 -2  0  3 -1 -4 
E -1  0  0  2 -4  2  5 -2  0 -3 -3  1 -2 -3 -1  0 -1 -3 -2 -2  1  4 -1 -4 
G  0 -2  0 -1 -3 -2 -2  6 -2 -4 -4 -2 -3 -3 -2  0 -2 -2 -3 -3 -1 -2 -1 -4 
H -2  0  1 -1 -3  0  0 -2  8 -3 -3 -1 -2 -1 -2 -1 -2 -2  2 -3  0  0 -1 -4 
I -1 -3 -3 -3 -1 -3 -3 -4 -3  4  2 -3  1  0 -3 -2 -1 -3 -1  3 -3 -3 -1 -4 
L -1 -2 -3 -4 -1 -2 -3 -4 -3  2  4 -2  2  0 -3 -2 -1 -2 -1  1 -4 -3 -1 -4 
K -1  2  0 -1 -3  1  1 -2 -1 -3 -2  5 -1 -3 -1  0 -1 -3 -2 -2  0  1 -1 -4 
M -1 -1 -2 -3 -1  0 -2 -3 -2  1  2 -1  5  0 -2 -1 -1 -1 -1  1 -3 -1 -1 -4 
F -2 -3 -3 -3 -2 -3 -3 -3 -1  0  0 -3  0  6 -4 -2 -2  1  3 -1 -3 -3 -1 -4 
P -1 -2 -2 -1 -3 -1 -1 -2 -2 -3 -3 -1 -2 -4  7 -1 -1 -4 -3 -2 -2 -1 -2 -4 
S  1 -1  1  0 -1  0  0  0 -1 -2 -2  0 -1 -2 -1  4  1 -3 -2 -2  0  0  0 -4 
T  0 -1  0 -1 -1 -1 -1 -2 -2 -1 -1 -1 -1 -2 -1  1  5 -2 -2  0 -1 -1  0 -4 
W -3 -3 -4 -4 -2 -2 -3 -2 -2 -3 -2 -3 -1  1 -4 -3 -2 11  2 -3 -4 -3 -2 -4 
Y -2 -2 -2 -3 -2 -1 -2 -3  2 -1 -1 -2 -1  3 -3 -2 -2  2  7 -1 -3 -2 -1 -4 
V  0 -3 -3 -3 -1 -2 -2 -3 -3  3  1 -2  1 -1 -2 -2  0 -3 -1  4 -3 -2 -1 -4 
B -2 -1  3  4 -3  0  1 -1  0 -3 -4  0 -3 -3 -2  0 -1 -4 -3 -3  4  1 -1 -4 
Z -1  0  0  1 -3  3  4 -2  0 -3 -3  1 -1 -3 -1  0 -1 -3 -2 -2  1  4 -1 -4 
X  0 -1 -1 -1 -2 -1 -1 -1 -1 -1 -1 -1 -1 -1 -2  0  0 -2 -1 -1 -1 -1 -1 -4 
* -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4 -4  1 `

func filter(arr []string, target string) []string {
	res := make([]string, 0)
	for _, str := range arr {
		if str != target {
			res = append(res, str)
		}
	}
	return res
}

func prep(s string) map[string]int {
	lines := strings.Split(s, "\n")
	cols := filter(strings.Split(lines[0], " "), "")
	dict := make(map[string]int)
	for i, max := 1, len(lines); i < max; i++ {
		row := filter(strings.Split(lines[i], " "), "")
		for j, maximum := 1, len(row); j < maximum; j++ {
			key := row[0] + cols[j-1]
			value, err := strconv.Atoi(row[j])
			if err != nil {
				panic(fmt.Sprintf("string to integer conversion failed at row #%d, column #%d\n", i, j))
			}
			dict[key] = value
		}
	}
	return dict
}

// BLOSUM62 exported
var BLOSUM62 = prep(s)

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