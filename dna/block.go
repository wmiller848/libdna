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
)

const (
	FOUR_BY_THREE BlockSize = 0
)

type BlockSize uint

type BlockConfig struct {
	Size  BlockSize
	Codex gene.Codex
}

type EncodingFunction func(indicies ...gene.Base) (gene.Codon, error)

type Block interface {
	Bases() []gene.Base
	//Encoding() EncodingFunction
	Random() *DNA
	Match(gene.Base) gene.Base
	Decode(...gene.Base) (gene.Codon, error)
}
