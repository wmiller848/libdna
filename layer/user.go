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

type UserLayerConfig struct {
}

func newUserLayer(config *UserLayerConfig) (*UserLayer, error) {
	return &UserLayer{}, nil
}

type UserLayer struct {
}

func (l *UserLayer) Pipe(dnaio.Flood) dnaio.Flood {
	flood := make(dnaio.Flood)
	return flood
}

func (l *UserLayer) Type() string {
	return "UserLayer"
}
