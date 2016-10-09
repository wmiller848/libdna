///////////////////////////////////////////////////////
//
//    __           __       ____
//   /\ \       __/\ \     /\  _'\
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
	"testing"
	"time"
)

func TestFlood(t *testing.T) {
	flood := make(Flood)
	go func() {
		time.Sleep(2 * time.Millisecond)
		flood <- Stream{}
		close(flood)
	}()

	stream, open := <-flood
	if stream == nil {
		t.Error("Stream is nil")
	}

	if !open {
		t.Error("Flood is still closed")
	}

	stream, open = <-flood
	if stream != nil {
		t.Error("Stream is not nil")
	}

	if open {
		t.Error("Flood is still open")
	}

}

func TestIoReader(t *testing.T) {
	data := []byte("some super amazing data")
	rdr := bytes.NewBuffer(data)
	flood := IoReader(rdr)

	stream, open := <-flood
	if stream.String() != string(data) {
		t.Error("Stream " + stream.String() + " does not match data")
	}

	if !open {
		t.Error("Flood is not open")
	}

	stream, open = <-flood
	if stream.String() != "" {
		t.Error("Stream contains data before new line")
	}

	if open {
		t.Error("Flood is open")
	}
}

func TestIoReaderCached(t *testing.T) {
	data := []byte("some super amazing data")
	rdr := bytes.NewBuffer(data)
	signal := make(chan int)
	flood := IoReaderCached(rdr, signal)

	stream, open := <-flood
	if stream.String() != string(data) {
		t.Error("Stream " + stream.String() + " does not match data")
	}

	if !open {
		t.Error("Flood is not open")
	}

	signal <- 1

	stream, open = <-flood
	if stream.String() != string(data) {
		t.Error("Stream " + stream.String() + " does not match data")
	}

	close(signal)

	stream, open = <-flood
	if stream.String() != "" {
		t.Error("Stream contains data before new line")
	}

	if open {
		t.Error("Flood is open")
	}

}
