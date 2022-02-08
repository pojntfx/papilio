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
				Class("pf-c-page__main", "pf-c-page__main").
				ID("papilio-main").
				TabIndex(-1).
				Body(
					app.Div().
						Class("pf-c-page__main-section", "pf-m-fill", "pf-l-flex", "pf-m-justify-content-center", "pf-m-align-items-center").
						Body(
							&ICGrid{
								Children: []app.UI{
									&ICCard{
										Link:   "/ic/fe11s",
										ICName: "FE 1.1s",
										ICImg:  "/web/img/fe11s.svg",
									},
									&ICCard{
										Link:   "/ic/fe21",
										ICName: "FE 2.1",
										ICImg:  "/web/img/fe21.svg",
									},
									&ICCard{
										Link:   "/ic/sl22s",
										ICName: "SL 2.2s",
										ICImg:  "/web/img/sl22s.svg",
									},
									&ProposeICCard{
										Link: "https://github.com/pojntfx/papilio/issues/new",
									},
								},
							},
						),
				),
		)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
