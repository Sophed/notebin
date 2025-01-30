package components

import (
	"bytes"
	"notebin/config"

	"github.com/sophed/lg"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

const contentArea = "content-area"

func View(title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("icon"), Type("image/x-icon"), Href("/static/favicon.png")),
			Link(Rel("stylesheet"), Href(config.Get().StaticDir+"/css/global.css")),
			Link(Rel("stylesheet"), Href(config.Get().StaticDir+"/css/nav.css")),
			Script(Src(config.Get().StaticDir+"/htmx.min.js"), Defer()),
		},
		Body: []Node{
			Div(Class("container"),
				Titles(),
				NavBar(),
				Div(Class(contentArea),
					Group(elements),
				),
			),
		},
	})
}

func Render(view Node) string {
	buf := new(bytes.Buffer)
	err := view.Render(buf)
	if err != nil {
		lg.Fatl(err)
	}
	return buf.String()
}
