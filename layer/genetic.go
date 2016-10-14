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

import (
	"github.com/wmiller848/libdna/dna"
	dnaio "github.com/wmiller848/libdna/io"
)

type GeneticLayerConfig struct {
	Labled     bool
	Population int
}

func newGeneticLayer(config *GeneticLayerConfig) (*GeneticLayer, error) {
	programs := make([]*dna.Program, config.Population)
	for i := 0; i < config.Population; i++ {
		program, err := dna.NewProgram()
		if err != nil {
			return nil, err
		}
		programs[i] = program
	}
	return &GeneticLayer{
		Config:   config,
		programs: programs,
	}, nil
}

type GeneticLayer struct {
	Config   *GeneticLayerConfig
	exemplar *dna.Program
	programs []*dna.Program
}

func (l *GeneticLayer) Pipe(stream dnaio.Stream) dnaio.Stream {
	if l.exemplar == nil {
		if l.Config.Labled {
			// Supervised
			points := make(map[int]int)
			streamLen := len(stream)

			dataStream := stream[:streamLen-1]
			assertStream := dnaio.Stream{stream[streamLen-1]}

			//
			for i, _ := range l.programs {
				points[i] = 0
				outputStream := l.programs[i].Evaluate(dataStream)
				if len(outputStream) == len(assertStream) {
					points[i] += 10
				}
				for j, _ := range dataStream {
					for k, _ := range assertStream {
						dataBuf := dataStream[j]
						assertBuf := assertStream[k]
						for jj, _ := range dataBuf {
							for kk, _ := range assertBuf {
								if dataBuf[jj] == assertBuf[kk] {
									points[i] += 50
								}
							}
						}
					}
				}
			}
			//fmt.Println(points)
			//
			//for i, v := range points {
			//fmt.Println(i, v)
			//}
		} else {
			// Unsupervised

		}
		return dnaio.Stream{}
	}
	return stream
}

func (l *GeneticLayer) Type() string {
	return "GeneticLayer"
}
