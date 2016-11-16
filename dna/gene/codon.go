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

package gene

type Base byte

type Codon []byte

func (c Codon) String() string {
	return string(c)
}

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

func (c Codex) Bytes() []byte {
	bytes := []byte{}
	for _, codon := range c {
		bytes = append(bytes, codon...)
	}
	return bytes
}

func (c Codex) Interface() [][]byte {
	// TODO
	// Get []Codon -> [][]byte not copy :(
	bytes := make([][]byte, len(c))
	for i, codon := range c {
		bytes[i] = codon
	}
	return bytes
}

func (c Codex) Find(start int) int {
	cdx := c[start:]
	depth := 0
	for i, codon := range cdx {
		if codon.String() == "[" {
			depth++
		}

		if codon.String() == "]" {
			depth--
			if depth < 0 {
				return i + start
			}
		}
	}
	return -1
}

type CodexGigas []Codex
