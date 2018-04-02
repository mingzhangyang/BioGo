package sequence

import (
	"errors"
	"strings"
)

const baseA byte = 65
const baseC byte = 67
const baseG byte = 71
const baseT byte = 84

// DNA type based on Sequence type
type DNA struct {
	Sequence
	name  string
	attrs map[string]string
}

// NewDNA create a new DNA object
func NewDNA(seq string) (*DNA, error) {
	seq = strings.ToUpper(seq)
	b := []byte(seq)
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case baseA, baseC, baseG, baseT:
		default:
			return nil, errors.New("Illegal character found")
		}
	}
	return &DNA{Sequence(seq), "", make(map[string]string)}, nil
}

// GetName get the name field
func (d *DNA) GetName() string {
	return d.name
}

// SetName set the name field
func (d *DNA) SetName(name string) {
	d.name = name
}

// SetAttr set attr
func (d *DNA) SetAttr(k, v string) {
	d.attrs[k] = v
}

// GCContent get the gc content
func (d *DNA) GCContent() float32 {
	b := []byte(d.Sequence)
	n := 0
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case baseC, baseG:
			n++
		}
	}
	return float32(n) / float32(len(b))
}

// Complement get the complementary sequence
func (d *DNA) Complement() string {
	b := []byte(d.Sequence)
	nb := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case baseA:
			nb[i] = baseT
		case baseT:
			nb[i] = baseA
		case baseC:
			nb[i] = baseG
		case baseG:
			nb[i] = baseC
		}
	}
	return string(nb)
}

// ReverseComplement get the reversed complementray seq
func (d *DNA) ReverseComplement() string {
	b := []byte(d.Sequence)
	nb := make([]byte, len(b))
	for i, j := 0, len(b)-1; i < len(b); i, j = i+1, j-1 {
		switch b[i] {
		case baseA:
			nb[j] = baseT
		case baseT:
			nb[j] = baseA
		case baseC:
			nb[j] = baseG
		case baseG:
			nb[j] = baseC
		}
	}
	return string(nb)
}
