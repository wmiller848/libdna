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
	"fmt"
)

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
		fmt.Println(dna.Unwind())
	default:
		return nil, errors.New("Unkown dna block size")
	}
	return dna, nil
}

func (d *DNA) Unwind() CodexGigas {
	strand := d.strand
	leng := len(strand)
	codexGigas := CodexGigas{}
	for i := 0; i < 3; i++ {
		codex := Codex{}
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
			strand_frag := []Base{Base(strand[t0]), Base(strand[t1]), Base(strand[t2])}
			codon, _ := d.block.Decode(strand_frag...)
			codex = append(codex, codon)
		}
		codexGigas = append(codexGigas, codex)
	}
	return codexGigas
}
