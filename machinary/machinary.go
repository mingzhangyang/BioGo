package machinary

import "fmt"

// PriorityLevel indicates the priority of a task
// The highest priority is 0, with the lowest 9.
type PriorityLevel uint

const (
	I PriorityLevel = iota
	II
	III
	IV
	V
	VI
	VII
	VIII
	IX
	X
)

func (p PriorityLevel) String() string {
	return fmt.Sprintf("Priority Level: %d", p)
}

type Instruction struct {
	arguments string
	priority int
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
