package datastructure

// Stack type
type Stack struct {
	Entry *Node
}

func (s *Stack) isEmpty() bool {
	return s.Entry == nil
}

// Push an element to the stack
func (s *Stack) Push(str string) {
	n := &Node{str, s.Entry}
	s.Entry = n
}

// Pop an element from the stack
func (s *Stack) Pop() (string, bool) {
	if s.Entry == nil {
		return "", false
	}
	n := s.Entry
	s.Entry = n.Next
	return n.Value, true
}

// String method define custom output
func (s Stack) String() string {
	if s.Entry == nil {
		return ""
	}
	n := s.Entry
	str := n.Value
	for n.Next != nil {
		n = n.Next
		str += (" - " + n.Value)
	}
	return str
}
