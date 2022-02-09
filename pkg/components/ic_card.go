package components

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ICCard struct {
	app.Compo

	Open   func()
	ICName string
	ICImg  string
}

func (c *ICCard) Render() app.UI {
	return app.A().
		OnClick(func(ctx app.Context, e app.Event) {
			c.Open()
		}).
		Class("pf-u-color-100 pf-x-m-no-decoration").
		Body(
			app.Div().
				Class("pf-c-card pf-m-hoverable-raised").
				Body(
					app.Div().
						Class("pf-c-card__header pf-l-flex pf-m-justify-content-center").
						Body(
							app.Div().
								Class("pf-c-card__header-main").
								Body(
									app.Img().
										Class("pf-x-u-height-10").
										Src(c.ICImg).
										Alt("Preview image of "+c.ICImg),
								),
						),
					app.Div().
						Class("pf-c-card__title pf-u-text-align-center").
						Body(
							app.Text(c.ICName),
						),
				),
		)
}
