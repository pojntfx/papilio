package components

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ProposeICCard struct {
	app.Compo

	Link string
}

func (c *ProposeICCard) Render() app.UI {
	return app.A().
		Href(c.Link).
		Target("_blank").
		Class("pf-u-color-100 pf-x-m-no-decoration").
		Body(
			app.Div().
				Class("pf-c-card pf-m-hoverable-raised pf-u-h-100 pf-l-flex pf-m-justify-content-center pf-m-align-items-center").
				Body(
					app.Div().
						Class("pf-c-card__header pf-l-flex pf-m-justify-content-center").
						Body(
							app.Div().
								Class("pf-c-card__header-main").
								Body(
									app.I().
										Class("pf-icon fas fa-plus"),
								),
						),
					app.Div().
						Class("pf-c-card__title pf-u-text-align-center").
						Body(
							app.Text("Suggest another IC"),
						),
				),
		)
}
