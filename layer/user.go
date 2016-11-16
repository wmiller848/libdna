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
	"fmt"

	dnaio "github.com/wmiller848/libdna/io"
)

type UserLayerHandler func(dnaio.Stream) dnaio.Stream

type UserLayerConfig struct {
	Handler UserLayerHandler
}

func newUserLayer(config *UserLayerConfig) (*UserLayer, error) {
	return &UserLayer{
		Config: config,
	}, nil
}

type UserLayer struct {
	Config *UserLayerConfig
}

func (l *UserLayer) Pipe(stream dnaio.Stream) dnaio.Stream {
	fmt.Println("USER-PIPE", stream)
	return l.Config.Handler(stream)
}

func (l *UserLayer) Type() string {
	return "UserLayer"
}
