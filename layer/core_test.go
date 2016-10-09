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

package layer

import "testing"

func TestNewGeneticLayer(t *testing.T) {
	config := &GeneticLayerConfig{}
	layer, err := New(config)
	if err != nil {
		t.Error(err.Error())
	}
	switch layer.(type) {
	case *GeneticLayer:
	default:
		t.Error("Layer is not GeneticLayer")
	}
}

func TestNewUserLayer(t *testing.T) {
	config := &UserLayerConfig{}
	layer, err := New(config)
	if err != nil {
		t.Error(err.Error())
	}
	switch layer.(type) {
	case *UserLayer:
	default:
		t.Error("Layer is not UserLayer")
	}
}

func TestNewPlatformLayer(t *testing.T) {
	config := &PlatformLayerConfig{}
	layer, err := New(config)
	if err != nil {
		t.Error(err.Error())
	}
	switch layer.(type) {
	case *PlatformLayer:
	default:
		t.Error("Layer is not PlatformLayer")
	}
}

func TestNewUnknownLayer(t *testing.T) {
	config := "not a real layer config"
	layer, err := New(config)
	if err != UNKOWN_LAYER {
		t.Error(err.Error())
	}
	switch layer.(type) {
	case *GeneticLayer:
		t.Error("Layer is PlatformLayer")
	case *UserLayer:
		t.Error("Layer is PlatformLayer")
	case *PlatformLayer:
		t.Error("Layer is PlatformLayer")
	default:
	}
}
