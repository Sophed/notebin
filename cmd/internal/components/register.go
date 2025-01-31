package components

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func Register() Node {
	return Div(Class("account-form"),
		P(Text("Create an account to continue")),
		Input(ID("inp-username"), Placeholder("Username..."), MaxLength("128")), Br(),
		Input(ID("inp-password"), Placeholder("Password..."), MaxLength("128"), Type("password")),
		Div(Class("form-actions"),
			Button(Text("Create account"), Attr("onclick", "register()")),
			Button(Text("Login instead"), hx.Get("/login"), hx.Target("."+contentArea), hx.Swap("innerHTML")),
		),
	)
}
