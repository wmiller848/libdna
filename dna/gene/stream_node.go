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
}

func (n *StreamNode) String() string {
	children := ""
	for _, child := range n.children {
		children += child.String()
	}
	return n.codon.String() + children
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

func NewStreamTree(codexGigas CodexGigas) []Node {
	cursor := cursor_node_open
	trees := []Node{}
	var constNode *StreamNode
	for _, codex := range codexGigas {
		var current *StreamNode
		for _, codon := range codex {
			switch codon.String() {
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
			case "$":
				node := &StreamNode{
					children: []*StreamNode{},
					codon:    codon,
					value:    codon.String(),
				}
				if cursor != cursor_node_open {
					current.children = append(current.children, node)
				} else {
					cursor = cursor_node_variable
					current = node
				}
			}
		}
		trees = append(trees, current)
	}
	return trees
}
