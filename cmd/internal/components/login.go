package components

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func Login() Node {
	return Div(Class("account-form"),
		P(Text("Log in to continue")),
		Input(ID("inp-username"), Placeholder("Username..."), MaxLength("128")), Br(),
		Input(ID("inp-password"), Placeholder("Password..."), MaxLength("128"), Type("password")),
		Div(Class("form-actions"),
			Button(Text("Login"), Attr("onclick", "login()")),
			Button(Text("Register instead"), hx.Get("/register"), hx.Target("."+contentArea), hx.Swap("innerHTML")),
		),
	)
}
