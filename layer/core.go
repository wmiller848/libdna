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
	"errors"

	dnaio "github.com/wmiller848/libdna/io"
)

type Layer interface {
	Pipe(dnaio.Stream) dnaio.Stream
	Type() string
}

func New(config interface{}) (Layer, error) {
	switch config.(type) {
	case *GeneticLayerConfig:
		return newGeneticLayer(config.(*GeneticLayerConfig))
	case *UserLayerConfig:
		return newUserLayer(config.(*UserLayerConfig))
	case *PlatformLayerConfig:
		return newPlatformLayer(config.(*PlatformLayerConfig))
	default:
		return nil, errors.New("Unkown layer config type")
	}
}
