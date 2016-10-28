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
	"errors"

	"github.com/wmiller848/libdna/dna/gene"
)

const Seed int = 1024 * 20

//const SeedMax int = 2048

type DNA struct {
	strand []byte
	block  Block
}

func New(config *BlockConfig) (*DNA, error) {
	var dna *DNA
	switch config.Size {
	case FOUR_BY_THREE:
		block, err := NewBlock4x3(Block4x3Bases, config.Codex)
		if err != nil {
			return nil, err
		}
		dna = block.Random()
	default:
		return nil, errors.New("Unkown dna block size")
	}
	return dna, nil
}

func (d *DNA) Unwind() gene.CodexGigas {
	strand := d.strand
	leng := len(strand)
	codexGigas := gene.CodexGigas{}
	// TODO make not hard coded 3
	for i := 0; i < 3; i++ {
		codex := gene.Codex{}
		for j := 0; j < leng; j += 3 {
			t0 := i + 0 + j
			t1 := i + 1 + j
			t2 := i + 2 + j
			if t0 > leng-1 {
				t0 -= leng
			}
			if t1 > leng-1 {
				t1 -= leng
			}
			if t2 > leng-1 {
				t2 -= leng
			}
			strand_frag := []gene.Base{gene.Base(strand[t0]), gene.Base(strand[t1]), gene.Base(strand[t2])}
			codon, _ := d.block.Decode(strand_frag...)
			codex = append(codex, codon)
		}
		codexGigas = append(codexGigas, codex)
	}
	return codexGigas
}

func (d *DNA) Sequence(codexGigas gene.CodexGigas) chan *Sequence {
	chanSeq := make(chan *Sequence)
	go func() {
		for codexID, codex := range codexGigas {
			i := 0
			index := 0
			elements := 0
			reading := false
			codexDecoded := gene.Codex{}
			for _, codon := range codex {
				if string(codon) == string(gene.CodonStart) {
					reading = true
					index = i
				} else if string(codon) != string(gene.CodonStop) && reading == true {
					codexDecoded = append(codexDecoded, codon)
					elements++
				} else if string(codon) == string(gene.CodonStop) && reading == true {
					if len(codexDecoded) == 0 {
						reading = false
						continue
					}
					seq := &Sequence{
						Codex:    codexDecoded,
						CodexID:  codexID,
						Index:    codexID + index,
						Elements: elements,
					}
					chanSeq <- seq
					elements = 0
					reading = false
					codexDecoded = gene.Codex{}
				}
				i++
			}
		}
		close(chanSeq)
	}()
	return chanSeq
}

func (d *DNA) SpliceSequence(chanSeq chan *Sequence) *SequenceNode {
	var dnaSeq *SequenceNode
	var head0 *SequenceNode
	var head1 *SequenceNode
	var head2 *SequenceNode
	for {
		seq, open := <-chanSeq
		if open == false {
			break
		}
		switch seq.CodexID {
		case 0:
			if head0 == nil {
				head0 = seq.Node()
			} else {
				head0 = head0.Merge(seq)
			}
		case 1:
			if head1 == nil {
				head1 = seq.Node()
			} else {
				head1 = head1.Merge(seq)
			}
		case 2:
			if head2 == nil {
				head2 = seq.Node()
			} else {
				head2 = head2.Merge(seq)
			}
		}
	}
	dnaSeq = head0
	if dnaSeq == nil || (head1 != nil && dnaSeq.Index > head1.Index) {
		dnaSeq = head1
	}
	if dnaSeq == nil || (head2 != nil && dnaSeq.Index > head2.Index) {
		dnaSeq = head2
	}
	return dnaSeq
}

func (d *DNA) MarshalGenes() ([]gene.Gene, error) {
	codexGigas := d.Unwind()
	channel := d.Sequence(codexGigas)
	dnaSeq := d.SpliceSequence(channel)
	seq := dnaSeq
	genes := []gene.Gene{}
	for seq != nil {
		g, err := gene.New(seq.Sequence.Codex)
		if g != nil && err == nil {
			if len(g.Codexs()) > 0 {
				genes = append(genes, g)
			}
		}
		seq = seq.Child
	}

	if len(genes) > 0 {
		return genes, nil
	}

	return nil, errors.New("Unable to sequence genes")
}
