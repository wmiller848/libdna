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

import dnaio "github.com/wmiller848/libdna/io"

func NewProgram() (*Program, error) {
	codex := Codex{
		Codon("&"), Codon("|"), Codon("^"),
		Codon("+"), Codon("-"), Codon("*"), Codon("/"),
		Codon("0"), Codon("1"), Codon("2"), Codon("3"),
		Codon("4"), Codon("5"), Codon("6"), Codon("7"),
		Codon("8"), Codon("9"), Codon("a"), Codon("b"),
		Codon("c"), Codon("d"), Codon("e"), Codon("f"),
		Codon("{"), Codon("}"), Codon("["), Codon("]"),
		Codon(","), Codon("."), Codon(":"), Codon("!"),
		Codon("$"), Codon("∫"), Codon("ƒ"), Codon("»"),
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
	return stream
}
