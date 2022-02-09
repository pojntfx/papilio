package components

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo

	showFE11sModalOpen bool
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
										Open: func() {
											c.showFE11sModalOpen = true

											c.Update()
										},
										ICName: "FE 1.1s",
										ICImg:  "/web/img/fe11s.svg",
									},
									&ICCard{
										Open: func() {
											log.Println("Opening FE 2.1")
										},
										ICName: "FE 2.1",
										ICImg:  "/web/img/fe21.svg",
									},
									&ICCard{
										Open: func() {
											log.Println("Opening SL 2.2s")
										},
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
			app.If(
				c.showFE11sModalOpen,
				&FE11sModal{
					OnSubmit: func(vendorID, productID, deviceReleaseNumber, numberOfDownstreamPorts string) {
						log.Println("Generating EEPROM config for", vendorID, productID, deviceReleaseNumber, numberOfDownstreamPorts)
					},
					OnCancel: func() {
						c.showFE11sModalOpen = false

						c.Update()
					},
				},
			),
		)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
