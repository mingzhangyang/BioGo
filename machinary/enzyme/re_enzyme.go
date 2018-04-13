package enzyme

import (
	"errors"
	"regexp"

	"BioGo/sequence"
)

// ReSite type
type ReSite struct {
	p int
	s string
}

func (s ReSite) seq() string {
	return s.s
}

// common enzymes
var (
	BamHI   = ReSite{1, "GGATCC"}
	HindIII = ReSite{1, "AAGCTT"}
	KpnI    = ReSite{5, "GGTACC"}
	NdeI    = ReSite{2, "CATATG"}
	XhoI    = ReSite{1, "CTCGAG"}
)

var dict = make(map[string]ReSite)

func init() {
	dict["BamHI"] = BamHI
	dict["HindIII"] = HindIII
	dict["KpnI"] = KpnI
	dict["NdeI"] = NdeI
	dict["XhoI"] = XhoI
}

// ReEnzyme type
type ReEnzyme struct {
	Name            string
	RecognitionSite *regexp.Regexp
}

// NewReEnzymeFromName create an new ReEnzyme from the enzyme name
func NewReEnzymeFromName(name string) (*ReEnzyme, error) {
	var s ReSite
	var ok bool
	if s, ok = dict[name]; !ok {
		return nil, errors.New("enzyme not in the pre-built collection, please try NewRE_EnzymeFromPattern instead")
	}
	return &ReEnzyme{name, regexp.MustCompile(s.seq())}, nil
}

// NewReEnzymeFromPattern create an new ReEnzyme from the recognition site
func NewReEnzymeFromPattern(pat string) (*ReEnzyme, error) {
	re, err := regexp.Compile(pat)
	if err != nil {
		return nil, err
	}
	return &ReEnzyme{RecognitionSite: re}, nil
}

// SearchDNA return the index of all the matched region
func (enz ReEnzyme) SearchDNA(dna *sequence.DNA) [][]int {
	return enz.RecognitionSite.FindAllStringIndex(dna.Seq(), -1)
}

// SearchRawSequence search a raw string
func (enz ReEnzyme) SearchRawSequence(str string) [][]int {
	return enz.RecognitionSite.FindAllStringIndex(str, -1)
}
