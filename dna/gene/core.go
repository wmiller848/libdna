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
	"errors"
	"fmt"
)

type Gene interface {
	Type() string
}

func New(codex Codex) (Gene, error) {
	if len(codex) == 0 {
		return nil, errors.New("Codex contains no genes")
	}
	switch string(codex[0]) {
	case Codon("$").String():
		fmt.Println("Stream")
		return NewStreamGene(codex), nil
	case Codon("∫").String():
		fmt.Println("Expression")
		return nil, nil
	case Codon("ƒ").String():
		fmt.Println("Function")
		return nil, nil
	case Codon("»").String():
		fmt.Println("FlowControl")
		return nil, nil
	default:
		return nil, errors.New("Codex is invalid")
	}
}
