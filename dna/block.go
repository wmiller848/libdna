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

const (
	FOUR_BY_THREE BlockSize = 0
)

type BlockSize uint

type BlockConfig struct {
	Size BlockSize
}

type EncodingFunction func(indicies ...Base) (Codon, error)

type Block interface {
	Bases() []Base
	Encoding() EncodingFunction
	Random() *DNA
	Match(Base) Base
	Decode(...Base) (Codon, error)
}
