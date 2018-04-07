package datastructure

type Node struct {
  Value string
  Next *Node
}

type Stack struct {
  Entry *Node
}

func (s *Stack) isEmpty() bool {
  return s.Entry == nil
}

func (s *Stack) Push(str string) {
  n := &Node{str, s.Entry}
  s.Entry = n
}

func (s *Stack) Pop() (string, bool) {
  if s.Entry == nil {
    return "", false
  }
  n := s.Entry
  s.Entry = n.Next
  return n.Value, true
}

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
