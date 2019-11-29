package genbank

import (
	"strings"
	"strconv"
)

// GBRecord represent the genebank data
type GBRecord struct {
	Name string
	Locus string
	Accession []string
	Definition string
	Version string
	Dblink []string
	Keywords string
	Source Source
	Reference []Reference
	Comment string
	Annotation map[string]string
	Features Features
	contig string
}

// Source of the genome
type Source struct {
	name string
	organism []string
}

func newSource(a []string) Source {
	return Source{
		name: a[0],
		organism: a[1:],
	}
}

// Reference about the genome
type Reference struct {
	description string
	authors []string
	title string
	journal string
	pubmed string
}

func newReference(a []string) Reference {
	ref := Reference{}
	ref.description = a[0]
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
		switch cur.name {
		case "AUTHORS":
			authors := make([]string, 0)
			for _, d := range cur.data {
				authors = append(authors, strings.Split(d, ", ")...)
			}
			m := len(authors)
			authors = append(authors[:m-1], strings.Split(authors[m-1], " and ")...)
			ref.authors = authors
		case "TITLE":
			ref.title = strings.Join(cur.data, " ")
		case "JOURNAL":
			ref.journal = strings.Join(cur.data, " ")
		case "PUBMED":
			ref.pubmed = cur.data[0]
		default:
			println("bypass this line: ", line)
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
		ref.authors = authors
	case "TITLE":
		ref.title = strings.Join(cur.data, " ")
	case "JOURNAL":
		ref.journal = strings.Join(cur.data, " ")
	case "PUBMED":
		ref.pubmed = cur.data[0]
	default:
		println("bypass this line: ", line)
	}
	return ref
}

// Features of the genome
type Features struct {
	Description FeaturesDescription
	Genes []Gene
}

// FeaturesDescription describe the features
type FeaturesDescription struct {
	Range [2]int
	Organism string
	Type string
	Strain string
	DbXref string
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
		cl := lines[i]
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
	Type string
	Start, End int
	Strand string
	Gene string
	LocusTag string
	OldLocusTag string
	Pseudo bool
	EcNumber []string
	Inference string
	Note string
	CodonStart int
	TranslateTable int
	Product string
	ProteinID string
	Translation string
}

func newGene(cur *holder) Gene {
	g := Gene{
		Type: cur.name,
	}
	line := cur.data[0]
	if strings.HasPrefix(line, "complement") {
		g.Strand = "complement"
		vs := strings.Split(line, "complement")[1]
		t := strings.Split(vs[1:len(vs)-1], "..")
		a := strings.TrimLeft(t[0], "<>")
		b := strings.TrimLeft(t[1], "><")
		var err error
		g.Start, err = strconv.Atoi(a)
		if err != nil {
			println("failed to convert to int at ", line)
		}
		g.End, err = strconv.Atoi(b)
		if err != nil {
			println("failed to convert to int at ", line)
		}
	} else {
		t := strings.Split(line, "..")
		a := strings.TrimLeft(t[0], "<>")
		b := strings.TrimLeft(t[1], "><")
		var err error
		g.Start, err = strconv.Atoi(a)
		if err != nil {
			println("failed to convert to int at ", line)
		}
		g.End, err = strconv.Atoi(b)
		if err != nil {
			println("failed to convert to int at ", line)
		}
	}
	line = ""
	for i, n := 1, len(cur.data); i < n; i++ {
		cl := cur.data[i]
		if !strings.HasPrefix(cl, "/") {
			line += " " + cl
			continue
		}
		if line != "" {
			vs := strings.Split(line, "=")
			p := vs[0]
			q := strings.Trim(vs[1], "\"")
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
			default:
				println("bypass line: ", line)
			}
		}
		line = cl
	}
	if line != "" {
		vs := strings.Split(line, "=")
		p := vs[0]
		q := strings.Trim(vs[1], "\"")
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
		default:
			println("bypass line: ", line)
		}
	}
	return g
}

func extracAnnotation(arr []string) map[string]string {
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