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

type Stream struct {
	genes CodexGigas
}

func (s *Stream) Codexs() CodexGigas {
	return s.genes
}

func (s *Stream) Type() string {
	return "stream"
}

func NewStreamGene(codex Codex) *Stream {
	cursor := cursor_open
	mode := mode_unknown
	sliced := false
	genes := CodexGigas{}
	healed := Codex{}
	for _, codon := range codex {
		switch codon.String() {
		case "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9":
			if mode != mode_unknown {
				healed = append(healed, codon)
			}
		case "a", "b", "c", "d", "e", "f":
			if mode != mode_unknown {
				healed = append(healed, codon)
			}
		case ",":
			if cursor == cursor_braket_start && mode != mode_reference {
				mode = mode_literal
				if healed[len(healed)-1].String() != "[" {
					healed = append(healed, codon)
				}
			}
		case "$":
			if cursor == cursor_open && mode == mode_unknown {
				mode = mode_reference
				healed = append(healed, codon)
			}
		case ":":
			if cursor == cursor_braket_start && mode == mode_reference && !sliced {
				sliced = true
				healed = append(healed, codon)
			}
		case "[":
			if cursor == cursor_open {
				cursor = cursor_braket_start
				healed = append(healed, codon)
			}
		case "]":
			if cursor == cursor_braket_start {
				cursor = cursor_braket_end
				healed = append(healed, codon)
				genes = append(genes, healed)
				healed = Codex{}
				cursor = cursor_open
				mode = mode_unknown
			}
		}
	}
	for i, h := range genes {
		hl := len(h)
		if hl > 0 && h[hl-1].String() != "]" {
			h = append(h, Codon("]"))
		}
		if h.String() == "[]" || h.String() == "$[]" {
			genes[i] = nil
		}
	}
	//fmt.Println(genes)
	return &Stream{
		genes: genes,
	}
}
