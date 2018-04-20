package filehandler

import (
	"fmt"
	"math"
	"os"
	"log"
	"bufio"

	"BioGo/util"
)

// Fastq type is a handler to deal with fastq files
type Fastq struct {
	Path string
	Name string
}

// NewFastq return a Fastq pointer given a string
func NewFastq(path string) *Fastq{
	return &Fastq{Path: path}
}

// Hist method read the fastq file line by line and return
// a slice of the mean score of each position across all the reads
func (f *Fastq) Hist() []float64 {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal(err)
	}
	report := make(map[int][]uint8)
	scanner := bufio.NewScanner(file)
	var i uint8 = 1
	for scanner.Scan() {
		switch i {
		case 1: 
			i++
		case 2:
			i++
		case 3:
			i++
		default:
			scores := scanner.Bytes()
			for idx, score := range scores {
				report[idx] = append(report[idx], (score - 33))
			}
			i = 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	var res = make([]float64, 0)
	for _, v := range report {
		res = append(res, util.MeanUint8(v))
	}
	return res
}

// QCounts counts the number of each Q score
func (f *Fastq) QCounts() map[byte]int {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal(err)
	}
	report := make(map[uint8]int)
	scanner := bufio.NewScanner(file)
	var i uint8 = 1
	for scanner.Scan() {
		switch i {
		case 1: 
			i++
		case 2:
			i++
		case 3:
			i++
		default:
			scores := scanner.Bytes()
			for _, score := range scores {
				report[score - 33]++
			}
			i = 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	
	return report
}

func p2Q(p float64) uint8 {
	return uint8(math.Round(-10 * math.Log10(p)))
}

func q2P(q uint8) float64 {
	t := float64(q) / float64(-10)
	return math.Pow(10, t)
}

// QtoPhred33 return the ASCII code in string
func QtoPhred33(q uint8) string {
	return string(q + 33)
}

// Phred33toQ return the Q value given char
func Phred33toQ(c byte) uint8 {
	return c - 33
}