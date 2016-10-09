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

	// Build a helper layer to ingest csv data
	uConfig := &layer.UserLayerConfig{
		Handler: func(buf dnaio.Buffer) dnaio.Stream {
			stream := bytes.Split(buf, []byte(","))
			return stream
		},
	}
	csvLayer, err := layer.New(uConfig)
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
