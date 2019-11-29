package genbank

import (
	"fmt"
	"strconv"
	"strings"
)

// GBRecord represent the genebank data
type GBRecord struct {
	Name       string
	Locus      string
	Accession  []string
	Definition string
	Version    string
	DbLink     []string
	Keywords   string
	Source     Source
	Reference  []Reference
	Comment    string
	Annotation map[string]string
	Features   Features
	Contig     string
}

// Source of the genome
type Source struct {
	name     string
	organism []string
}

func newSource(a []string) Source {
	return Source{
		name:     a[0],
		organism: a[1:],
	}
}

// Reference about the genome
type Reference struct {
	Description string
	Authors     []string
	Title       string
	Journal     string
	PubMed      string
}

func (ref Reference) String() string {
	res := strings.Builder{}
	res.WriteString(fmt.Sprintln("Description: ", ref.Description))
	res.WriteString(fmt.Sprintln("Authors: ", strings.Join(ref.Authors, " ")))
	res.WriteString(fmt.Sprintln("Title: ", ref.Title))
	res.WriteString(fmt.Sprintln("Journal: ", ref.Journal))
	res.WriteString(fmt.Sprintln("PubMed ID: ", ref.PubMed))
	return res.String()
}

func newReference(a []string) Reference {
	ref := Reference{}
	ref.Description = a[0]
	var line string
	cur := holder{
		data: make([]string, 0, 1024),
	}
	for i, n := 1, len(a); i < n; i++ {
		line = a[i]
		head, body := strings.Trim(line[:12], " "), line[12:]
		if head == "" {
			cur.data = append(cur.data, body)
			continue
		}
		if cur.name != "" {
			switch cur.name {
			case "AUTHORS":
				authors := make([]string, 0)
				for _, d := range cur.data {
					authors = append(authors, strings.Split(d, ", ")...)
				}
				m := len(authors)
				authors = append(authors[:m-1], strings.Split(authors[m-1], " and ")...)
				ref.Authors = authors
			case "TITLE":
				ref.Title = strings.Join(cur.data, " ")
			case "JOURNAL":
				ref.Journal = strings.Join(cur.data, " ")
			case "PUBMED":
				ref.PubMed = cur.data[0]
			default:
				println("bypass this line: ", line)
			}
		}

		cur.name = head
		cur.data = cur.data[:1]
		cur.data[0] = body
	}

	switch cur.name {
	case "AUTHORS":
		authors := make([]string, 0)
		for _, d := range cur.data {
			authors = append(authors, strings.Split(d, ", ")...)
		}
		m := len(authors)
		authors = append(authors[:m-1], strings.Split(authors[m-1], " and ")...)
		ref.Authors = authors
	case "TITLE":
		ref.Title = strings.Join(cur.data, " ")
	case "JOURNAL":
		ref.Journal = strings.Join(cur.data, " ")
	case "PUBMED":
		ref.PubMed = cur.data[0]
	default:
		println("bypass this line: ", line)
	}
	return ref
}

// Features of the genome
type Features struct {
	Description FeaturesDescription
	Genes       []*Gene
}

// FeaturesDescription describe the features
type FeaturesDescription struct {
	Range    [2]int
	Organism string
	Type     string
	Strain   string
	DbXref   string
}

func newFeatureDescription(lines []string) FeaturesDescription {
	desc := FeaturesDescription{
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
		line = cl
	}
	if line != "" {
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
	return desc
}

// Gene is a generic for gene, cds, tRNA, rRNA
type Gene struct {
	Type              string
	Range             [][2]int
	Strand            string
	Gene              string
	LocusTag          string
	OldLocusTag       string
	Pseudo            bool
	EcNumber          []string
	Inference         string
	Note              string
	CodonStart        int
	TranslateTable    int
	Product           string
	ProteinID         string
	Translation       string
	DbXref            string
	RibosomalSlippage bool
	GeneSynonym       string
	AntiCodon         string
	Function          string
	NcRNAClass        string
	TranslateExcept   string
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
		cl := strings.TrimLeft(cur.data[i], " ")
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
	default:
		println("bypass line: ", line, p, q)
	}
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
