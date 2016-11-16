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
	"github.com/wmiller848/libdna/dna/gene"
	dnaio "github.com/wmiller848/libdna/io"
)

func NewProgram() (*Program, error) {
	codex := gene.Codex{
		gene.Codon("&"), gene.Codon("|"), gene.Codon("^"), gene.Codon("%"),
		gene.Codon("0"), gene.Codon("1"), gene.Codon("2"), gene.Codon("3"),
		gene.Codon("4"), gene.Codon("5"), gene.Codon("6"), gene.Codon("7"),
		gene.Codon("8"), gene.Codon("9"), gene.Codon("a"), gene.Codon("b"),
		gene.Codon("c"), gene.Codon("d"), gene.Codon("e"), gene.Codon("f"),
		gene.Codon("{"), gene.Codon("}"), gene.Codon("["), gene.Codon("]"),
		gene.Codon(","), gene.Codon("."), gene.Codon(":"), gene.Codon("!"),
		gene.Codon("$"), gene.Codon("∫"), gene.Codon("ƒ"), gene.Codon("»"),
	}
	blockConfig := &BlockConfig{
		Size:  FOUR_BY_THREE,
		Codex: codex,
	}
	dna, err := New(blockConfig)
	if err != nil {
		return nil, err
	}
	return &Program{
		dna: dna,
	}, nil
}

type Program struct {
	dna *DNA
}

func (p *Program) Evaluate(stream dnaio.Stream) dnaio.Stream {
	genes, err := p.dna.MarshalGenes()
	if err == nil {
		context := gene.NewContext(genes)
		return context.Evaluate(stream)
	}

	return nil
}
