package genbank

// GBRecord represent the genebank data
type GBRecord struct {
	Name       string            `json:"name,omitempty"`
	Locus      string            `json:"locus,omitempty"`
	Accession  []string          `json:"accession,omitempty"`
	Definition string            `json:"definition,omitempty"`
	Version    string            `json:"version,omitempty"`
	DbLink     []string          `json:"db_link,omitempty"`
	Keywords   string            `json:"keywords,omitempty"`
	Source     Source            `json:"source,omitempty"`
	Reference  []Reference       `json:"reference,omitempty"`
	Comment    string            `json:"comment,omitempty"`
	Annotation map[string]string `json:"annotaion,omitempty"`
	Features   Features          `json:"features,omitempty"`
	Contig     string            `json:"contig,omitempty"`
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
