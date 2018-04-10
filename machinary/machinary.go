package machinary

type Instruction struct {
	command   func()
	arguments string
	immediate bool
}

type Messenger interface {
	Transduce() Instruction
	Pair(r Receptor) bool
}

type Receptor interface {
	Receive(m Messenger)
}

type Emitter interface {
	Emit() Messenger
}

type Worker interface {
	Work()
}
