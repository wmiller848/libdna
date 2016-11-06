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

type ExpressionNode struct {
	children []*ExpressionNode
	codon    Codon
	value    interface{}
	flavor   string
}

func (n *ExpressionNode) String() string {
	children := ""
	for _, child := range n.children {
		children += child.String()
	}
	switch n.flavor {
	case "operator":
		return n.codon.String() + children
	case "constant", "variable":
		return n.value.(string) + children
	default:
		return children
	}
}

func (n *ExpressionNode) Debug() string {
	children := ""
	for _, child := range n.children {
		children += child.Debug()
	}
	switch n.value.(type) {
	case string:
		return n.flavor + " ( " + n.value.(string) + " " + children + ") "
	default:
		return n.flavor + " ( " + n.codon.String() + " " + children + ") "
	}
}

func (n *ExpressionNode) Type() string {
	children := ""
	for _, child := range n.children {
		children += child.Type()
	}
	sep := " "
	var endSep, op string
	switch n.flavor {
	case "operator":
		op = "∫"
	case "constant":
		op = "π"
	case "variable":
		op = "?"
	case "stream":
		op = "∑"
		sep = "( "
		endSep = ") "
	}
	return op + sep + children + endSep
}

func NewExpressionTree(codex Codex, nodes ...*ExpressionNode) Node {
	mask := false
	cursor := cursor_node_open
	var current, root *ExpressionNode
	i := 0
	if nodes != nil && nodes[i] != nil {
		cursor = cursor_node_inception
		current = nodes[i]
		root = nodes[i]
	}
	var constNode *ExpressionNode
	for i, codon := range codex {
		switch codon.String() {
		case "&", "|", "^", "%":
			if mask {
				continue
			}
			node := &ExpressionNode{
				children: []*ExpressionNode{},
				flavor:   "operator",
			}
			node.codon = codon
			node.value = codon
			if cursor != cursor_node_open {
				current.children = append(current.children, node)
			} else {
				current = node
				root = node
			}
			cursor = cursor_node_operator
		case "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f":
			if mask {
				continue
			}

			if cursor != cursor_node_constant {
				constNode = &ExpressionNode{
					children: []*ExpressionNode{},
					flavor:   "constant",
				}
				constNode.codon = codon
				constNode.value = codon.String()
				current.children = append(current.children, constNode)
			} else {
				str := constNode.value.(string) + codon.String()
				constNode.value = str
			}
			cursor = cursor_node_constant
		case ",":
			if mask {
				continue
			}
			cursor = cursor_node_seperator
		case "$":
			if mask {
				continue
			}

			if cursor != cursor_node_constant {
				constNode = &ExpressionNode{
					children: []*ExpressionNode{},
					flavor:   "variable",
				}
				constNode.codon = codon
				constNode.value = codon.String()
				current.children = append(current.children, constNode)
			}
			cursor = cursor_node_constant
		case "[":
			if mask {
				continue
			}
			mask = true
			node := &ExpressionNode{
				children: []*ExpressionNode{},
				flavor:   "stream",
			}
			n := i + 1
			if n < len(codex) {
				cn := codex.Find(n)
				var cdx Codex
				if cn < 0 {
					cdx = codex[n:]
				} else {
					cdx = codex[n:cn]
				}
				NewExpressionTree(cdx, node)
			}
			current.children = append(current.children, node)
			cursor = cursor_node_braket_start
		case "]":
			if cursor != cursor_node_braket_start {
				break
			}
			mask = false
			cursor = cursor_node_braket_end
		}
	}
	return root
}
