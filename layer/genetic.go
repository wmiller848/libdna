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

package layer

import (
	"github.com/wmiller848/libdna/dna"
	dnaio "github.com/wmiller848/libdna/io"
)

type GeneticLayerConfig struct {
	Labled bool
}

func newGeneticLayer(config *GeneticLayerConfig) (*GeneticLayer, error) {
	return &GeneticLayer{
		Config: config,
	}, nil
}

type GeneticLayer struct {
	Config   *GeneticLayerConfig
	Actor    *dna.Program
	Programs []*dna.Program
}

func (l *GeneticLayer) Pipe(stream dnaio.Stream) dnaio.Stream {
	if l.Actor == nil {

		// TODO evolve a program that meets the requirments

		return dnaio.Stream{}
	}
	return stream
}

func (l *GeneticLayer) Type() string {
	return "GeneticLayer"
}
