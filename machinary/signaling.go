package machinary

type Messenger struct {
	message string
	source
	target *Receptor
}
