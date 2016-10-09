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

import (
	"bytes"
	"io"
)

const (
	BUFFER_SIZE int = 512
)

type Flood chan Stream

func IoReader(rdr io.Reader) Flood {
	flood := make(Flood)
	buf := make(Buffer, BUFFER_SIZE)
	var lastLn Buffer
	go func() {
		for {
			n, err := rdr.Read(buf)
			if err == io.EOF {
				close(flood)
				return
			}

			if bytes.Count(buf[:n], []byte("\n")) == 0 {
				lastLn = append(lastLn, buf[:n]...)
				continue
			}
			ln := bytes.Split(buf[:n], []byte("\n"))
			for i := 0; i < len(ln); i++ {
				if len(ln[i]) > 0 {
					flood <- Stream{append(lastLn, Buffer(ln[i])...)}
				}
			}
			lastLn = ln[len(ln)-1]
		}
	}()
	return flood
}
