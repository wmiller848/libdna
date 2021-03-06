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

import (
	dnaio "github.com/wmiller848/libdna/io"
)

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
	mode_stream_inception = iota

	gene_type_stream = "stream"
)

func NewStreamGene(codex Codex) *Stream {
	cursor := cursor_stream_open
	flag := flag_stream_off
	flagCount := 0
	mode := mode_stream_unknown
	sliced := false
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
			if cursor != cursor_stream_seperator && flag == flag_stream_braket_start && mode != mode_stream_reference {
				healed = append(healed, codon)
				cursor = cursor_stream_seperator
			}
		case "$":
			if cursor == cursor_stream_open && mode == mode_stream_unknown {
				mode = mode_stream_reference
				healed = append(healed, codon)
			}
		case ":":
			if flag == flag_stream_braket_start && !sliced && mode == mode_stream_reference {
				sliced = true
				healed = append(healed, codon)
				cursor = cursor_stream_seperator
			}
		case "[":
			if mode == mode_stream_unknown {
				mode = mode_stream_literal
			}
			cursor = cursor_stream_braket_start
			flag = flag_stream_braket_start
			if mode == mode_stream_literal && flagCount < max_depth {
				flagCount++
				healed = append(healed, codon)
			} else if mode == mode_stream_reference && flagCount == 0 {
				flagCount++
				healed = append(healed, codon)
			}
		case "]":
			if flagCount > 0 {
				cursor = cursor_stream_braket_end
				healed = append(healed, codon)
				flag = flag_stream_braket_end
				flagCount--
				if flagCount == 0 {
					healed = Codex{}
					cursor = cursor_stream_open
					mode = mode_stream_unknown
				}
			}
		}
	}
	for i := 0; i < flagCount; i++ {
		healed = append(healed, Codon("]"))
	}
	if healed.String() == "[]" ||
		healed.String() == "$[]" ||
		healed.String() == "[:]" ||
		healed.String() == "$[:]" {
		healed = nil
	}
	return &Stream{
		gene: healed,
	}
}

type Stream struct {
	gene Codex
}

func (s *Stream) Node() Node {
	return NewStreamTree(s.gene)
}

func (s *Stream) Codex() Codex {
	return s.gene
}

func (s *Stream) Type() string {
	return gene_type_stream
}

func (s *Stream) Evaluate(runtime *Runtime) dnaio.Stream {
	node := s.Node()
	codex := node.Evaluate(runtime)
	return codex.Interface()
}
