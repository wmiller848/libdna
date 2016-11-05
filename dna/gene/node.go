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

const (
	cursor_node_open         = iota
	cursor_node_constant     = iota
	cursor_node_operator     = iota
	cursor_node_variable     = iota
	cursor_node_seperator    = iota
	cursor_node_braket_start = iota
	cursor_node_braket_end   = iota
	cursor_node_inception    = iota
)

type Node interface {
	String() string
	Debug() string
	Type() string
}
