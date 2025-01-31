package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func App() Node {
	return Div(
		Img(
			Src("https://firefox-settings-attachments.cdn.mozilla.net/main-workspace/newtab-wallpapers-v2/6b8eb3cf-f232-4c7b-a179-afd174555134.avif"),
			Width("200px"),
		),
	)
}
