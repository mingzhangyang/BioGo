package signaling

type Messenger struct {
	message string
	source *Emitter
	target *Receptor
}
