///////////////////////////////////////////////////////
//
//    __           __       ____
//   /\ \       __/\ \     /\  _'\
//   \ \ \     /\_\ \ \____\ \ \/\ \    ___      __
//    \ \ \  __\/\ \ \ '__'\\ \ \ \ \ /' _ '\  /'__'\
//     \ \ \L\ \\ \ \ \ \L\ \\ \ \_\ \/\ \/\ \/\ \L\.\_
//      \ \____/ \ \_\ \_,__/ \ \____/\ \_\ \_\ \__/.\_\
//       \/___/   \/_/\/___/   \/___/  \/_/\/_/\/__/\/_/
//
///////////////////////////////////////////////////////

package dna

import (
	"fmt"

	"github.com/wmiller848/libdna/dna/gene"
)

type Sequence struct {
	Codex    gene.Codex
	CodexID  int
	Index    int
	Elements int
}

func (s *Sequence) Node() *SequenceNode {
	return &SequenceNode{
		Sequence: s,
		Child:    nil,
	}
}

type SequenceNode struct {
	*Sequence
	Child *SequenceNode
}

func (s *SequenceNode) String() string {
	str := fmt.Sprintf("%+v", s.Sequence.Codex)
	if s.Child != nil {
		str += " " + s.Child.String()
	}
	return str
}

func (s *SequenceNode) Bytes() []byte {
	bytes := []byte{}
	bytes = append(bytes, s.Sequence.Codex.Bytes()...)
	if s.Child != nil {
		bytes = append(bytes, byte(' '))
		bytes = append(bytes, s.Child.Bytes()...)
	}
	return bytes
}

func (s *SequenceNode) Clone() *SequenceNode {
	seq := &SequenceNode{
		Sequence: &Sequence{
			Codex:    s.Sequence.Codex,
			CodexID:  s.Sequence.CodexID,
			Index:    s.Sequence.Index,
			Elements: s.Sequence.Elements,
		},
	}
	if s.Child != nil {
		seq.Child = s.Child.Clone()
	}
	return seq
}

func (s *SequenceNode) Merge(seq *Sequence) *SequenceNode {
	node := seq.Node()
	if s.Index > seq.Index {
		node.Child = s
		return node
	} else {
		if s.Child == nil {
			s.Child = node
		} else {
			s.Child = s.Child.Merge(seq)
		}
		return s
	}
}
