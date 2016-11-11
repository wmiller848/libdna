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

type StreamNode struct {
	children []*StreamNode
	codon    Codon
	value    interface{}
	flavor   string
}

func (n *StreamNode) String() string {
	children := ""
	for _, child := range n.children {
		children += child.String()
	}
	switch n.flavor {
	case "inception", "reference", "literal":
		return n.flavor + " ( " + n.value.(string) + " " + children + ") "
	default:
		return n.value.(string) + " " + children
	}
}

func (n *StreamNode) Debug() string {
	children := ""
	for _, child := range n.children {
		children += child.Debug()
	}
	return n.codon.String() + children
}

func (n *StreamNode) Type() string {
	return "stream"
}

func NewStreamTree(codex Codex, nodes ...*StreamNode) Node {
	cursor := cursor_node_open
	mode := mode_stream_unknown
	var current, root, constNode *StreamNode
	if nodes != nil && nodes[0] != nil {
		mode = mode_stream_inception
		current = nodes[0]
		root = nodes[0]
	}
	for i := 0; i < len(codex); i++ {
		codon := codex[i]
		switch codon.String() {
		case "$":
			constNode = &StreamNode{
				children: []*StreamNode{},
				codon:    codon,
				value:    codon.String(),
			}
			if mode == mode_stream_unknown {
				current = constNode
				root = constNode
				mode = mode_stream_reference
			}
			cursor = cursor_node_constant
		case ",":
			cursor = cursor_node_seperator
		case "[":
			node := &StreamNode{
				children: []*StreamNode{},
				value:    "",
			}
			n := i + 1
			if n < len(codex) {
				cn := codex.Find(n)
				var cdx Codex
				if cn <= 0 {
					cdx = codex[n:]
					cn = len(codex) - 1
				} else {
					cdx = codex[n:cn]
				}
				i = cn - 1
				NewStreamTree(cdx, node)
				if mode == mode_stream_unknown {
					mode = mode_stream_literal
					current = node
					root = node
				} else {
					current.children = append(current.children, node)
				}
				cursor = cursor_node_braket_start
			}
		case "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f":
			if cursor != cursor_node_constant {
				constNode = &StreamNode{
					children: []*StreamNode{},
					codon:    codon,
					value:    codon.String(),
				}
				current.children = append(current.children, constNode)
			} else {
				str := constNode.value.(string) + codon.String()
				constNode.value = str
			}
			cursor = cursor_node_constant
		case ":":
			node := &StreamNode{
				children: []*StreamNode{},
				codon:    codon,
				value:    codon.String(),
			}
			current.children = append(current.children, node)
			cursor = cursor_node_seperator
		case "]":
			cursor = cursor_node_braket_end
		}
	}

	if mode == mode_stream_reference {
		root.flavor = "reference"
	} else if mode == mode_stream_literal {
		root.flavor = "literal"
	} else if mode == mode_stream_inception {
		root.flavor = "inception"
	}
	return root
}
