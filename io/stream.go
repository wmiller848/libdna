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

package io

type Buffer []byte

type Stream [][]byte

func (s Stream) Flatten() Buffer {
	bigbuf := Buffer{}
	for _, buf := range s {
		bigbuf = append(bigbuf, buf...)
	}
	return bigbuf
}

func (s Stream) String() string {
	return string(s.Flatten())
}
