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

func (n *ExpressionNode) Evaluate(runtime *Runtime) Codex {
	//fmt.Println("Codex", n.Type(), n.String())
	return Codex{}
}

func NewExpressionTree(codex Codex, nodes ...*ExpressionNode) Node {
	cursor := cursor_node_open
	var current, root, constNode *ExpressionNode
	if nodes != nil && nodes[0] != nil {
		current = nodes[0]
		root = nodes[0]
	}
	for i := 0; i < len(codex); i++ {
		codon := codex[i]
		switch codon.String() {
		case "&", "|", "^", "%":
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
			if cursor != cursor_node_constant {
				constNode = &ExpressionNode{
					children: []*ExpressionNode{},
					codon:    codon,
					value:    codon.String(),
					flavor:   "constant",
				}
				current.children = append(current.children, constNode)
			} else {
				str := constNode.value.(string) + codon.String()
				constNode.value = str
			}
			cursor = cursor_node_constant
		case ",":
			cursor = cursor_node_seperator
		case "$":
			if cursor != cursor_node_constant {
				constNode = &ExpressionNode{
					children: []*ExpressionNode{},
					codon:    codon,
					value:    codon.String(),
					flavor:   "variable",
				}
				current.children = append(current.children, constNode)
			}
			cursor = cursor_node_constant
		case "[":
			node := &ExpressionNode{
				children: []*ExpressionNode{},
				flavor:   "stream",
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
				NewExpressionTree(cdx, node)
				current.children = append(current.children, node)
				cursor = cursor_node_braket_start
			}
		case "]":
			cursor = cursor_node_braket_end
		}
	}
	return root
}
