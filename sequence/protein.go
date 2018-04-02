package sequence

// Protein type based on Sequence
type Protein struct {
	Sequence
	name  string
	attrs map[string]string
}
