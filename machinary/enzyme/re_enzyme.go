package enzyme

import (
	"errors"
	"regexp"

	"../../sequence"
)

type RE_Site struct {
	p int
	s string
}

func (s RE_Site) seq() string {
	return s.s
}

var (
	BamHI   = RE_Site{1, "GGATCC"}
	HindIII = RE_Site{1, "AAGCTT"}
	KpnI    = RE_Site{5, "GGTACC"}
	NdeI    = RE_Site{2, "CATATG"}
	XhoI    = RE_Site{1, "CTCGAG"}
)

var dict = make(map[string]RE_Site)

func init() {
	dict["BamHI"] = BamHI
	dict["HindIII"] = HindIII
	dict["KpnI"] = KpnI
	dict["NdeI"] = NdeI
	dict["XhoI"] = XhoI
}

type RE_Enzyme struct {
	Name            string
	RecognitionSite *regexp.Regexp
}

// NewRE_EnzymeFromName create an new RE_Enzyme from the enzyme name
func NewRE_EnzymeFromName(name string) (*RE_Enzyme, error) {
	var s RE_Site
	var ok bool
	if s, ok = dict[name]; !ok {
		return nil, errors.New("enzyme not in the pre-built collection, please try NewRE_EnzymeFromPattern instead")
	}
	return &RE_Enzyme{name, regexp.MustCompile(s.seq())}, nil
}

// NewRE_EnzymeFromPattern create an new RE_Enzyme from the recognition site
func NewRE_EnzymeFromPattern(pat string) (*RE_Enzyme, error) {
	re, err := regexp.Compile(pat)
	if err != nil {
		return nil, err
	}
	return &RE_Enzyme{RecognitionSite: re}, nil
}

// SearchDNA return the index of all the matched region
func (enz RE_Enzyme) SearchDNA(dna *sequence.DNA) [][]int {
	return enz.RecognitionSite.FindAllStringIndex(dna.Seq(), -1)
}

// SearchRawSequence search a raw string
func (enz RE_Enzyme) SearchRawSequence(str string) [][]int {
	return enz.RecognitionSite.FindAllStringIndex(str, -1)
}
