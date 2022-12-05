package util

// https://blog.devgenius.io/golang-data-structures-stack-and-queue-2-5a6822baeb0d
type Node struct {
	Next  *Node
	Value string
}

type Stack struct {
	Top, Bottom *Node
	Length      int
}

func (s *Stack) Push(value string) {
	n := &Node{
		Value: value,
	}

	if s.Length == 0 {
		s.Bottom = n
	} else {
		n.Next = s.Top
	}

	s.Top = n
	s.Length++
}

func (s *Stack) Pop() (n *Node) {
	if s.Length == 0 {
		return
	}

	n = s.Top
	if s.Length == 1 {
		s.Top = nil
		s.Bottom = nil
	} else {
		s.Top = n.Next
		n.Next = nil
	}
	s.Length--

	return
}
