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

const (
	cursor_stream_open         = iota
	cursor_stream_seperator    = iota
	cursor_stream_braket_start = iota
	cursor_stream_braket_end   = iota

	flag_stream_off          = iota
	flag_stream_braket_start = iota
	flag_stream_braket_end   = iota

	mode_stream_unknown   = iota
	mode_stream_literal   = iota
	mode_stream_reference = iota
)

type Stream struct {
	genes CodexGigas
}

func (s *Stream) Nodes() []Node {
	return NewStreamTree(s.genes)
}

func (s *Stream) Codexs() CodexGigas {
	return s.genes
}

func (s *Stream) Type() string {
	return "stream"
}

func NewStreamGene(codex Codex) *Stream {
	cursor := cursor_stream_open
	flag := flag_stream_off
	mode := mode_stream_unknown
	sliced := false
	genes := CodexGigas{}
	healed := Codex{}
	for _, codon := range codex {
		switch codon.String() {
		case "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f":
			if mode != mode_stream_unknown {
				healed = append(healed, codon)
				cursor = cursor_stream_open
			}
		case ",":
			if cursor != cursor_stream_braket_start && cursor != cursor_stream_seperator && flag == flag_stream_braket_start && mode != mode_stream_reference {
				mode = mode_stream_literal
				healed = append(healed, codon)
				cursor = cursor_stream_seperator
			}
		case "$":
			if cursor == cursor_stream_open && mode == mode_stream_unknown {
				mode = mode_stream_reference
				healed = append(healed, codon)
			}
		case ":":
			if cursor == cursor_stream_braket_start && mode == mode_stream_reference && !sliced {
				sliced = true
				healed = append(healed, codon)
			}
		case "[":
			if cursor == cursor_stream_open && flag != flag_stream_braket_start {
				cursor = cursor_stream_braket_start
				flag = flag_stream_braket_start
				healed = append(healed, codon)
			}
		case "]":
			if flag == flag_stream_braket_start {
				cursor = cursor_stream_braket_end
				healed = append(healed, codon)
				genes = append(genes, healed)
				healed = Codex{}
				cursor = cursor_stream_open
				flag = flag_stream_braket_end
				mode = mode_stream_unknown
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
	cleanedGenes := CodexGigas{}
	for _, h := range genes {
		if h != nil {
			cleanedGenes = append(cleanedGenes, h)
		}
	}
	//fmt.Println(genes)
	return &Stream{
		genes: cleanedGenes,
	}
}
