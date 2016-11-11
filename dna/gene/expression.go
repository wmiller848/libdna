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
	cursor_expression_open      = iota
	cursor_expression_operator  = iota
	cursor_expression_variable  = iota
	cursor_expression_seperator = iota
	cursor_expression_constant  = iota
	cursor_expression_braket    = iota

	flag_expression_off          = iota
	flag_expression_braket_start = iota
	flag_expression_braket_end   = iota

	mode_expression_unknown = iota
	mode_expression_valid   = iota
)

type Expression struct {
	gene Codex
}

func (e *Expression) Node() Node {
	return NewExpressionTree(e.gene)
}

func (e *Expression) Codex() Codex {
	return e.gene
}

func (e *Expression) Type() string {
	return "expression"
}

func NewExpressionGene(codex Codex) *Expression {
	cursor := cursor_expression_open
	flag := flag_expression_off
	flagCount := 0
	mode := mode_expression_unknown
	healed := Codex{}

	for _, codon := range codex {
		switch codon.String() {
		case "&", "|", "^", "%":
			if mode == mode_expression_unknown {
				mode = mode_expression_valid
			}
			if cursor != cursor_expression_operator &&
				cursor != cursor_expression_seperator &&
				cursor != cursor_expression_braket &&
				cursor != cursor_expression_variable {
				healed = append(healed, codon)
			}
			cursor = cursor_expression_operator
		case "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f":
			if mode == mode_expression_valid {
				healed = append(healed, codon)
				if cursor != cursor_expression_variable {
					cursor = cursor_expression_constant
				}
			}
		case ",":
			if mode == mode_expression_valid {
				if cursor != cursor_expression_seperator && cursor != cursor_expression_operator {
					healed = append(healed, codon)
					cursor = cursor_expression_seperator
				}
			}
		case "$":
			if mode == mode_expression_valid && flag != flag_expression_braket_start {
				if cursor != cursor_expression_variable && cursor != cursor_expression_constant && cursor != cursor_expression_braket {
					healed = append(healed, codon)
					cursor = cursor_expression_variable
				}
			}
		case "[":
			if mode == mode_expression_valid {
				if flagCount < max_depth {
					healed = append(healed, codon)
					cursor = cursor_expression_braket
					flag = flag_expression_braket_start
					flagCount++
				}
			}
		case "]":
			if mode == mode_expression_valid {
				if flagCount > 0 {
					healed = append(healed, codon)
					cursor = cursor_expression_braket
					flag = flag_expression_braket_end
					flagCount--
				}
			}
		}
	}
	for i := 0; i < flagCount; i++ {
		healed = append(healed, Codon("]"))
	}

	return &Expression{
		gene: healed,
	}
}
