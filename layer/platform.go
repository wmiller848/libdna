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

type PlatformLayerHandler func(dnaio.Buffer) dnaio.Stream

type PlatformLayerConfig struct {
	Handler PlatformLayerHandler
}

func newPlatformLayer(config *PlatformLayerConfig) (*PlatformLayer, error) {
	return &PlatformLayer{
		Config: config,
	}, nil
}

type PlatformLayer struct {
	Config *PlatformLayerConfig
}

func (l *PlatformLayer) Pipe(stream dnaio.Stream) dnaio.Stream {
	return l.Config.Handler(stream.Flatten())
}

func (l *PlatformLayer) Type() string {
	return "PlatformLayer"
}
