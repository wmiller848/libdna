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

import "fmt"

type Expression struct {
	genes CodexGigas
}

func (e *Expression) Codexs() CodexGigas {
	return e.genes
}

func (e *Expression) Type() string {
	return "expression"
}

func NewExpressionGene(codex Codex) *Expression {
	genes := CodexGigas{}

	fmt.Println(codex)

	return &Expression{
		genes: genes,
	}
}
