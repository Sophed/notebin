package components

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func NavBar() Node {
	return Div(
		Ul(Class("nav"),
			navItem("Home", "/", "left"),
			navItem("Account", "/account", "right"),
		),
	)
}

func navItem(title, page, align string) Node {
	return Li(Class(align),
		A(Text(title), hx.Get(page), hx.Target("."+contentArea), hx.Swap("innerHTML")),
	)
}
