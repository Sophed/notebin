package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Account() Node {
	return Div(
		P(Text("waow account pagee.,.,....")),
		Button(Text("Log out"), Attr("onclick", "logOut()")),
	)
}
