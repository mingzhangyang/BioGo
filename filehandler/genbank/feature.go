package genbank

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Features of the genome
type Features struct {
	Description FeaturesDescription
	Genes       []*Gene
}

// FeaturesDescription describe the features
type FeaturesDescription struct {
	Range    [2]int `json:"range,omitempty"`
	Organism string `json:"organism,omitempty"`
	Type     string `json:"type,omitempty"`
	Strain   string `json:"strain,omitempty"`
	DbXref   string `json:"db_xref,omitempty"`
}

func newFeatureDescription(lines []string) FeaturesDescription {
	desc := &FeaturesDescription{
		Range: [2]int{0, 0},
	}
	line := lines[0]
	vs := strings.Split(line, "..")
	a, err := strconv.Atoi(vs[0])
	if err != nil {
		println("failed to convert to in at: ", line)
	}
	desc.Range[0] = a
	b, err := strconv.Atoi(vs[1])
	if err != nil {
		println("failed to convert to in at: ", line)
	}
	desc.Range[1] = b
	line = ""
	for i, n := 1, len(lines); i < n; i++ {
		cl := strings.TrimLeft(lines[i], " ")
		if !strings.HasPrefix(cl, "/") {
			line += " " + cl
			continue
		}
		if line != "" {
			desc.setAttribute(line)
		}
		line = cl
	}
	if line != "" {
		desc.setAttribute(line)
	}
	return *desc
}

func (desc *FeaturesDescription) setAttribute(line string) {
	vs := strings.Split(line, "=")
	p, q := vs[0], strings.Trim(vs[1], "\"")
	switch p {
	case "/organism":
		desc.Organism = q
	case "/mol_type":
		desc.Type = q
	case "/strain":
		desc.Strain = q
	case "/db_xref":
		desc.DbXref = q
	default:
		println("bypass this line: ", line)
	}
}

func (desc FeaturesDescription) String() string {
	bs, err := json.Marshal(desc)
	if err != nil {
		return "error happens marshaling the object"
	}
	return string(bs)
}
