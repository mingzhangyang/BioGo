package datastructure

type Node struct {
	name string
	value string
	parent *Node
	children []*Node
}

func (n *Node) ToJSON() string {
	
}