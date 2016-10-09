package main

import (
	"fmt"
	"os"

	"github.com/wmiller848/libdna"
	"github.com/wmiller848/libdna/layer"
)

func main() {
	model := libdna.New()
	config := &layer.GeneticLayerConfig{
		Labled: true,
	}
	classifyLayer, err := layer.New(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	model.AddLayer(classifyLayer)
	model.Run(os.Stdin)
}
