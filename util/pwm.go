package util

import (
	"log"
	"strings"
	"math"
)

/*
* A position weight matrix (PWM), also known as a position-specific weight matrix (PSWM)
* or position-specific scoring matrix (PSSM), is a commonly used representation of motifs
* in biological sequences.
*/ 

type counter struct {
	A, T, G, C int
}

// Composition is the statistics of bases
type Composition struct {
	A, T, G, C float64
}

// CalcPwmDNA compute PWM for DNA alignment
// alig should be a collection of DNA strings separated by "\n" 
func CalcPwmDNA (alig string) []Composition {
	ss := strings.Split(alig, "\n")
	xx := make([]string, 0)
	
	// trim space and check the length of each substring in the slice
	var n int
	for i := 0; i < len(ss); i++ {
		s := strings.TrimSpace(ss[i])
		if (len(s) == 0) {
			continue
		}
		if n == 0 {
			n = len(s)
		}
		if len(s) != n {
			log.Fatal("invalid alignment")
		}
		xx = append(xx, s)
	}

	res := make([]Composition, n)
	m := len(xx)
	count := counter{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch xx[j][i] {
			case 'A':
				count.A++
			case 'T':
				count.T++
			case 'G':
				count.G++
			case 'C':
				count.C++
			default:
				log.Fatalf("invalid character %s detected\n", string(xx[j][i]))
			}
		}
		res[i] = Composition{
			A: math.Log2(float64(count.A) / float64(m) * 4),
			T: math.Log2(float64(count.T) / float64(m) * 4),
			G: math.Log2(float64(count.G) / float64(m) * 4),
			C: math.Log2(float64(count.C) / float64(m) * 4),
		}
		count.A = 0
		count.T = 0
		count.G = 0
		count.C = 0
	}
	return res
}