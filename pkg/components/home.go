package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Home struct {
	app.Compo
}

func (c *Home) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#papilio-main").
				Body(
					app.Text("Skip to content"),
				),
			&Navbar{},
			app.Main().
				Class("pf-c-page__main").
				ID("papilio-main").
				TabIndex(-1).
				Body(
					app.H1().Text("Hello, world!"),
				),
		)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
