package util

import (
	"fmt"
	"bytes"
)

// WordCounter type/class
type WordCounter struct {
	seq string
	wordLength uint
	legalChars string
	words map[string][]int
}

// NewWordCounter create a WordCounter object
func NewWordCounter(str string, length uint, chars string) *WordCounter {
	return &WordCounter{str, length, chars, make(map[string][]int)}
}

// Count the sequence
func (wc *WordCounter) Count() {
	// m is the return value
	m := make(map[string][]int)
	b := []byte(wc.legalChars)

	s := make([]rune, 0)
	for _, ch := range wc.seq {
		if !bytes.ContainsRune(b, ch) {
			// panic(fmt.Sprintf("Illegal chatacter found: %s", string(s[i])))
			fmt.Printf("Illegal chatacter found: %q, ignored...\n", ch)
			continue
		}
		s = append(s, ch)
	}
	// number of characters in the word
	n := int(wc.wordLength)
	
	for i, j := 0, len(s); i < j-n+1; i++ {
		w := string(s[i:i+n])
		if m[w] == nil {
			m[w] = []int{i}
		} else {
			m[w] = append(m[w], i)
		}
	}
	wc.words = m
}

// Words return the words
func (wc *WordCounter) Words() map[string][]int {
	return wc.words
}

// Ratio calculate the ratio of words found / all possible words
func (wc *WordCounter) Ratio() float64 {
	var c int
	c = len(wc.words)
	if c == 0 {
		wc.Count()
		c = len(wc.words)
	}
	n := len([]rune(wc.legalChars))
	a := 1
	for i, j := 0, int(wc.wordLength); i < j; i++ {
		a *= n
	}
	fmt.Printf("Number of words found: %d\n", c)
	fmt.Printf("Number of total possible words: %d\n", a)
	return float64(c) / float64(a)
}

// SetInput change the internal sequence of the object
func (wc *WordCounter) SetInput(s string) {
	wc.seq = s
	if len(wc.words) != 0 {
		wc.words = make(map[string][]int)
	}
}

// SetWordLength change the internal wordLength of the object
func (wc *WordCounter) SetWordLength(i uint) {
	wc.wordLength = i
	if len(wc.words) != 0 {
		wc.words = make(map[string][]int)
	}
}

// SetLegalChars change the internal legalChars of the object
func (wc *WordCounter) SetLegalChars(s string) {
	wc.legalChars = s
	if len(wc.words) != 0 {
		wc.words = make(map[string][]int)
	}
}

// AddLegalChars add a character to the legalChars
func (wc *WordCounter) AddLegalChars(cs string) {
	wc.legalChars += cs
	if len(wc.words) != 0 {
		wc.words = make(map[string][]int)
	}
}