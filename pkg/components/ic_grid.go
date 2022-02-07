package components

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ICGrid struct {
	app.Compo

	Children []app.UI
}

func (c *ICGrid) Render() app.UI {
	return app.Div().
		Class("pf-l-gallery", "pf-m-gutter").
		Body(
			app.Range(c.Children).Slice(func(i int) app.UI {
				return app.Div().
					Class("pf-l-gallery__item").
					Body(c.Children[i])
			}),
		)
}
