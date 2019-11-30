package genbank

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Gene is a generic for gene, cds, tRNA, rRNA
type Gene struct {
	Type              string   `json:"id,omitempty"`
	Range             [][2]int `json:"range,omitempty"`
	Strand            string   `json:"strand,omitempty"`
	Gene              string   `json:"gene,omitempty"`
	LocusTag          string   `json:"locus_tag,omitempty"`
	OldLocusTag       string   `json:"old_locus_tag,omitempty"`
	Pseudo            bool     `json:"pseudo,omitempty"`
	EcNumber          []string `json:"EC_number,omitempty"`
	Inference         string   `json:"inference,omitempty"`
	Note              string   `json:"note,omitempty"`
	CodonStart        int      `json:"codon_start,omitempty"`
	TranslateTable    int      `json:"transl_table,omitempty"`
	Product           string   `json:"product,omitempty"`
	ProteinID         string   `json:"protein_id,omitempty"`
	Translation       string   `json:"translation,omitempty"`
	DbXref            string   `json:"db_xref,omitempty"`
	RibosomalSlippage bool     `json:"ribosomal_slippage,omitempty"`
	GeneSynonym       string   `json:"gene_synonym,omitempty"`
	AntiCodon         string   `json:"anti_codon,omitempty"`
	Function          string   `json:"function,omitempty"`
	NcRNAClass        string   `json:"ncRNA_class,omitempty"`
	TranslateExcept   string   `json:"translate_except,omitempty"`
	RegulatoryClass   string   `json:"regulatory_class,omitempty"`
	BoundMoiety       string   `json:"bound_moiety,omitempty"`
}

func newGene(cur *holder) *Gene {
	g := &Gene{
		Type: cur.name,
	}
	line := cur.data[0]
	if strings.HasPrefix(line, "complement") {
		g.Strand = "complement"
		line = strings.Split(line, "complement")[1]
		line = strings.Trim(line, "()")
	}
	if strings.HasPrefix(line, "join") {
		line = strings.Split(line, "join")[1]
		line = strings.Trim(line, "()")
	}
	ts := strings.Split(line, ",")
	g.Range = make([][2]int, 0, 2)
	for _, pa := range ts {
		t := strings.Split(pa, "..")
		a := strings.TrimLeft(t[0], "<>")
		b := strings.TrimLeft(t[1], "><")
		var err error
		var r [2]int
		r[0], err = strconv.Atoi(a)
		if err != nil {
			println("failed to convert to int at ", line)
		}
		r[1], err = strconv.Atoi(b)
		if err != nil {
			println("failed to convert to int at ", line)
		}
		g.Range = append(g.Range, r)
	}
	line = ""
	for i, n := 1, len(cur.data); i < n; i++ {
		cl := cur.data[i]
		if !strings.HasPrefix(cl, "/") {
			line += " " + cl
			continue
		}
		if line != "" {
			g.SetAttribute(line)
		}
		line = cl
	}
	if line != "" {
		g.SetAttribute(line)
	}
	return g
}

// SetAttribute does what it says
func (g *Gene) SetAttribute(line string) {
	vs := strings.Split(line, "=")
	var p, q string
	p = vs[0]
	if len(vs) > 1 {
		q = strings.Trim(vs[1], "\"")
	}

	switch p {
	case "/locus_tag":
		g.LocusTag = q
	case "/gene":
		g.Gene = q
	case "/old_locus_tag":
		g.OldLocusTag = q
	case "/note":
		g.Note = q
	case "/codon_start":
		v, err := strconv.Atoi(q)
		if err != nil {
			println("failed to convert to int at line: ", line)
		}
		g.CodonStart = v
	case "/transl_table":
		v, err := strconv.Atoi(q)
		if err != nil {
			println("failed to convert to int at line: ", line)
		}
		g.TranslateTable = v
	case "/pseudo":
		g.Pseudo = true
	case "/inference":
		g.Inference = q
	case "/product":
		g.Product = q
	case "/protein_id":
		g.ProteinID = q
	case "/translation":
		g.Translation = q
	case "/EC_number":
		if g.EcNumber == nil {
			g.EcNumber = []string{}
		}
		g.EcNumber = append(g.EcNumber, q)
	case "/db_xref":
		g.DbXref = q
	case "/gene_synonym":
		g.GeneSynonym = q
	case "/anticodon":
		g.AntiCodon = q
	case "/ribosomal_slippage":
		g.RibosomalSlippage = true
	case "/function":
		g.Function = q
	case "/ncRNA_class":
		g.NcRNAClass = q
	case "/transl_except":
		g.TranslateExcept = q
	case "/regulatory_class":
		g.RegulatoryClass = q
	case "/bound_moiety":
		g.BoundMoiety = q
	default:
		println("bypass line: ", line)
	}
}

func (g *Gene) String() string {
	bs, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(bs)
}

func extractAnnotation(arr []string) map[string]string {
	res := make(map[string]string)
	tmp := [2]string{"", ""}
	for _, line := range arr {
		vs := strings.Split(line, "::")
		if len(vs) == 2 {
			if tmp[0] != "" {
				res[tmp[0]] = tmp[1]
			}
			tmp[0], tmp[1] = strings.Trim(vs[0], " "), strings.Trim(vs[1], " ")
			continue
		}
		tmp[1] += strings.Trim(line, " ")
	}
	res[tmp[0]] = tmp[1]
	return res
}
