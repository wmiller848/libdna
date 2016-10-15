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
	"testing"

	"github.com/wmiller848/libdna/dna/gene"
)

func TestSequence(t *testing.T) {
	seq := &Sequence{
		Codex: gene.Codex{
			gene.Codon("a"),
			gene.Codon("b"),
		},
	}

	node := seq.Node()
	if node.Sequence != seq {
		t.Error("Node sequence does not match original sequence")
	}
}

func TestSequenceNodeString(t *testing.T) {
	seq := &Sequence{
		Codex: gene.Codex{
			gene.Codon("a"),
			gene.Codon("b"),
		},
	}

	node := seq.Node()
	AssertStr(t, node.String(), "ab")
}

func TestSequenceNodeBytes(t *testing.T) {
	seq := &Sequence{
		Codex: gene.Codex{
			gene.Codon("a"),
			gene.Codon("b"),
		},
	}

	node := seq.Node()
	AssertStr(t, string(node.Bytes()), "ab")
}

func TestSequenceNodeClone(t *testing.T) {
	seq := &Sequence{
		Codex: gene.Codex{
			gene.Codon("a"),
			gene.Codon("b"),
			gene.Codon("c"),
			gene.Codon("d"),
		},
	}

	node := seq.Node()
	clone := node.Clone()
	AssertStr(t, node.String(), clone.String())
}

func TestSequenceNodeMerge(t *testing.T) {
	seq1 := &Sequence{
		Index: 1,
		Codex: gene.Codex{
			gene.Codon("a"),
			gene.Codon("b"),
		},
	}

	seq2 := &Sequence{
		Index: 0,
		Codex: gene.Codex{
			gene.Codon("c"),
			gene.Codon("d"),
		},
	}

	node := seq1.Node()
	merge := node.Merge(seq2)
	AssertStr(t, merge.String(), "cd ab")

	node = seq2.Node()
	merge = node.Merge(seq1)
	AssertStr(t, merge.String(), "cd ab")
}
