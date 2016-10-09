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

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/wmiller848/libdna"
	dnaio "github.com/wmiller848/libdna/io"
	"github.com/wmiller848/libdna/layer"
)

func main() {
	model := libdna.New()

	// Build a platform helper layer to ingest csv data
	pConfig := &layer.PlatformLayerConfig{
		Handler: func(buf dnaio.Buffer) dnaio.Stream {
			stream := bytes.Split(buf, []byte(","))
			return stream
		},
	}
	csvLayer, err := layer.New(pConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Build a genetic layer to learn from labled data
	gConfig := &layer.GeneticLayerConfig{
		Labled: true,
	}
	classifyLayer, err := layer.New(gConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	model.
		AddLayer(csvLayer).
		AddLayer(classifyLayer)

	model.Run(os.Stdin)
}
