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

const SeedBase int = 100
const SeedMax int = 200

type Base byte

type Codon []byte

var CodonStart Codon = Codon("<")
var CodonStop Codon = Codon(">")

type Codex []Codon

func (c Codex) String() string {
	str := ""
	for _, codon := range c {
		str += string(codon)
	}
	return str
}

type CodexGigas []Codex
