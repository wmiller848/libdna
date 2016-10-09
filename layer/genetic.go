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
	dnaio "github.com/wmiller848/libdna/io"
)

type GeneticLayerConfig struct {
	Labled bool
}

func newGeneticLayer(config *GeneticLayerConfig) (*GeneticLayer, error) {
	return &GeneticLayer{}, nil
}

type GeneticLayer struct {
}

func (l *GeneticLayer) Pipe(dnaio.Flood) dnaio.Flood {
	flood := make(dnaio.Flood)
	return flood
}

func (l *GeneticLayer) Type() string {
	return "GeneticLayer"
}
