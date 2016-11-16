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

package gene

import dnaio "github.com/wmiller848/libdna/io"

func NewContext(genes []Gene) *Context {
	var expressionGenes, streamGenes []Gene = []Gene{}, []Gene{}
	for _, g := range genes {
		switch g.Type() {
		case gene_type_stream:
			streamGenes = append(streamGenes, g)
		case gene_type_expression:
			expressionGenes = append(expressionGenes, g)
		}
	}
	return &Context{
		streamGenes:     streamGenes,
		expressionGenes: expressionGenes,
	}
}

type Runtime struct {
	stream     dnaio.Stream
	references []Node
}

type Context struct {
	streamGenes     []Gene
	expressionGenes []Gene
}

func (c *Context) Evaluate(stream dnaio.Stream) dnaio.Stream {
	in := stream
	nodes := []Node{}
	for _, g := range c.streamGenes {
		nodes = append(nodes, g.Node())
	}
	//fmt.Printf("%p\n", c)
	//fmt.Println("Stream Nodes", len(nodes), nodes)
	for _, g := range c.expressionGenes {
		runtime := &Runtime{
			stream:     in,
			references: nodes,
		}
		in = g.Evaluate(runtime)
	}
	return dnaio.Stream{}
}
