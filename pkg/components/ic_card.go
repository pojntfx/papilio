package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ICCard struct {
	app.Compo

	Open   func()
	ICName string
	ICImg  string
}

func (c *ICCard) Render() app.UI {
	return app.Div().
		Role("button").
		TabIndex(0).
		OnKeyDown(func(ctx app.Context, e app.Event) {
			if key := e.Get("key").String(); key == " " || key == "Enter" {
				c.open(ctx, e)
			}
		}).
		OnClick(c.open).
		Class("pf-u-color-100 pf-x-m-no-decoration pf-x-u-cursor-pointer").
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

func (c *ICCard) open(ctx app.Context, e app.Event) {
	e.PreventDefault()

	c.Open()
}
