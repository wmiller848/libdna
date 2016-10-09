///////////////////////////////////////////////////////
//
//    __           __       ____
//	 /\ \       __/\ \     /\  _'\
//   \ \ \     /\_\ \ \____\ \ \/\ \    ___      __
//    \ \ \  __\/\ \ \ '__'\\ \ \ \ \ /' _ '\  /'__'\
//     \ \ \L\ \\ \ \ \ \L\ \\ \ \_\ \/\ \/\ \/\ \L\.\_
//      \ \____/ \ \_\ \_,__/ \ \____/\ \_\ \_\ \__/.\_\
//       \/___/   \/_/\/___/   \/___/  \/_/\/_/\/__/\/_/
//
///////////////////////////////////////////////////////

package libdna

import (
	"fmt"
	"io"

	"github.com/wmiller848/libdna/layer"
)

func New() *Model {
	return &Model{}
}

type Model struct {
	Layers []layer.Layer
}

func (m *Model) AddLayer(l layer.Layer) *Model {
	m.Layers = append(m.Layers, l)
	return m
}

func (m *Model) Run(stdin io.Reader) {
	fmt.Println("test 123")
}
