package filehandler


// GBRecord represent the genebank data
type GBRecord struct {
	name string
	locus string
	reference []ref
	features []feature
}

type ref struct {
	authors []string
	title string
	journal string
}

type feature struct {
	source string
	gene []string
	mRNA []string
	exon []string
	cds string
	sts []string
	misc []string
}
