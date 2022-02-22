package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Toggle struct {
	ID    int
	Title string
	On    bool
}

type ToggleGroup struct {
	app.Compo

	ID       string
	Toggles  []Toggle
	OnToggle func(int)
}

func (c *ToggleGroup) Render() app.UI {
	return app.Div().
		ID(c.ID).
		Class("pf-c-toggle-group pf-m-compact").
		Body(
			app.Range(c.Toggles).Slice(func(i int) app.UI {
				buttonClasses := "pf-c-toggle-group__button"
				if c.Toggles[i].On {
					buttonClasses += " pf-m-selected"
				}

				return app.Div().
					Class("pf-c-toggle-group__item").
					Body(
						app.Button().
							Class(buttonClasses).
							Type("button").
							OnClick(func(ctx app.Context, e app.Event) {
								c.OnToggle(i)
							}).
							Body(
								app.Span().
									Class("pf-c-toggle-group__text").
									Body(
										app.Text(c.Toggles[i].Title),
									),
							),
					)
			}),
		)
}
