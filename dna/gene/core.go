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

import "errors"

type Gene interface {
	Codexs() CodexGigas
	Type() string
}

func New(codex Codex) (Gene, error) {
	if len(codex) == 0 {
		return nil, errors.New("Codex contains no genes")
	}
	switch codex[0].String() {
	case "$":
		//fmt.Println("Stream")
		return NewStreamGene(codex[1:]), nil
	case "∫":
		//fmt.Println("Expression")
		return NewExpressionGene(codex[1:]), nil
		return nil, nil
	case "ƒ":
		//fmt.Println("Function")
		return nil, nil
	case "»":
		//fmt.Println("FlowControl")
		return nil, nil
	default:
		return nil, errors.New("Codex is invalid")
	}
}
