package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type FE11sModal struct {
	app.Compo

	OnSubmit func(
		vendorID string,
		productID string,
		deviceReleaseNumber string,
		numberOfDownstreamPorts string,
	)
	OnCancel func()

	vendorID                string
	productID               string
	deviceReleaseNumber     string
	numberOfDownstreamPorts string
}

func (c *FE11sModal) Render() app.UI {
	return &components.Modal{
		ID:           "fe11s-modal",
		Title:        "Configure your FE 1.1s",
		DisableFocus: true,
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("fe11s-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					c.OnSubmit(
						c.vendorID,
						c.productID,
						c.deviceReleaseNumber,
						c.numberOfDownstreamPorts,
					)

					c.cancel()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("vendor-id-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Vendor ID"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									&components.Autofocused{
										Component: app.Input().
											Class("pf-c-form-control").
											Required(true).
											Type("text").
											Placeholder("0x1a0x40").
											ID("vendor-id-input").
											OnInput(func(ctx app.Context, e app.Event) {
												c.vendorID = ctx.JSSrc().Get("value").String()
											}).
											Value(c.vendorID),
									},
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("product-id-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Text("Product ID"),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Text("*"),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												Placeholder("0x010x01").
												ID("product-id-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.productID = ctx.JSSrc().Get("value").String()
												}).
												Value(c.productID),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("device-release-number-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Text("Device Release Number"),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Text("*"),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												Placeholder("0x30x83").
												ID("device-release-number-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.deviceReleaseNumber = ctx.JSSrc().Get("value").String()
												}).
												Value(c.deviceReleaseNumber),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("number-of-downstream-ports-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Text("Port Number"),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Text("*"),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												Placeholder("4").
												ID("number-of-downstream-ports-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.numberOfDownstreamPorts = ctx.JSSrc().Get("value").String()
												}).
												Value(c.numberOfDownstreamPorts),
										),
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Text("Generate and download").
				Form("fe11s-form"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					c.cancel()
				}),
		},
		OnClose: func() {
			c.cancel()
		},
	}

}

func (c *FE11sModal) cancel() {
	c.vendorID = ""
	c.productID = ""
	c.deviceReleaseNumber = ""
	c.numberOfDownstreamPorts = ""

	c.OnCancel()
}
