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
	"io"

	dnaio "github.com/wmiller848/libdna/io"
	"github.com/wmiller848/libdna/layer"
)

type Iterator func(dnaio.Buffer) bool

func New() *Model {
	return &Model{}
}

type Model struct {
	currentLayer int
	Layers       []layer.Layer
	Trained      bool
}

func (m *Model) AddLayer(l layer.Layer) *Model {
	m.Layers = append(m.Layers, l)
	return m
}

func (m *Model) Run(stdin io.Reader, iterator Iterator) chan dnaio.Buffer {
	streams := make(chan dnaio.Buffer)
	go func() {
		signal := make(chan int)
		generation := 0
		//flood := dnaio.IoReader(stdin)
		flood := dnaio.IoReaderCached(stdin, signal)
		m.reset()
		out := m.pipe(flood, m.nextLayer())
		for {
			go func() {
				for {
					stream, open := <-out
					if len(stream) > 0 {
						//term := iterator(stream.Flatten())
						streams <- stream.Flatten()
					}
					if !open {
						close(streams)
						return
					}
				}
			}()
			generation++
			if generation == 2 {
				close(signal)
				return
			}
			signal <- generation
		}
	}()
	return streams
}

func (m *Model) pipe(flood dnaio.Flood, l layer.Layer) dnaio.Flood {
	if l == nil {
		return flood
	}
	downstream := make(dnaio.Flood)
	go func() {
		for {
			stream, open := <-flood
			if !open {
				// Flush the stream
				downstream <- l.Pipe(stream)
				close(downstream)
				return
			}
			if len(stream) > 0 {
				downstream <- l.Pipe(stream)
			}
		}
	}()
	nextLayer := m.nextLayer()
	return m.pipe(downstream, nextLayer)
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

func (m *Model) reset() {
	m.currentLayer = 0
}
