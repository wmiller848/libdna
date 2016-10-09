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

	dnaio "github.com/wmiller848/libdna/io"
	"github.com/wmiller848/libdna/layer"
)

func New() *Model {
	return &Model{}
}

type Model struct {
	currentLayer int
	Layers       []layer.Layer
}

func (m *Model) AddLayer(l layer.Layer) *Model {
	m.Layers = append(m.Layers, l)
	return m
}

func (m *Model) Run(stdin io.Reader) {
	flood := dnaio.IoReader(stdin)
	m.currentLayer = 0
	out := m.pipe(flood, m.nextLayer())
	for {
		stream, open := <-out
		if !open {
			return
		}
		fmt.Println(stream.String(), open)
	}
}

func (m *Model) pipe(flood dnaio.Flood, l layer.Layer) dnaio.Flood {
	if l == nil {
		return flood
	}
	downstream := make(dnaio.Flood)
	go func() {
		for {
			stream, open := <-flood
			if len(stream) > 0 {
				downstream <- l.Pipe(stream)
			}
			if !open {
				close(downstream)
				return
			}
		}
	}()
	return m.pipe(downstream, m.nextLayer())
}

func (m *Model) nextLayer() layer.Layer {
	defer m.tickLayer()
	if m.currentLayer < len(m.Layers) {
		return m.Layers[m.currentLayer]
	} else {
		return nil
	}
}

func (m *Model) tickLayer() {
	m.currentLayer++
}
