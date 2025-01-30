package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Titles() Node {
	return Div(Class("titles"),
		H2(Style("float: left"), Text("Notebin")),
		H2(Style("float: right"), Text("⸜( ´ ꒳ ` )⸝")),
	)
}
