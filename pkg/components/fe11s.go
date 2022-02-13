package components

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
	"github.com/pojntfx/papilio/pkg/generators/fe11s"
)

const (
	hexUint16Pattern = `[a-fA-F\d]{4}`
)

type FE11sModal struct {
	app.Compo

	OnSubmit func(
		idVendor uint16,
		idProduct uint16,
		bcdDevice uint16,
		numberOfDownstreamPorts uint8,
	)
	OnCancel func()

	idVendor                string
	idProduct               string
	bcdDevice               string
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

					vendorID, err := strconv.ParseUint(c.idVendor, 16, 16)
					if err != nil {
						log.Println("Could not parse vendor ID:", err)

						return
					}

					productID, err := strconv.ParseUint(c.idProduct, 16, 16)
					if err != nil {
						log.Println("Could not parse product ID:", err)

						return
					}

					bcdDevice, err := strconv.ParseUint(c.bcdDevice, 16, 16)
					if err != nil {
						log.Println("Could not parse device release number:", err)

						return
					}

					numberOfDownstreamPorts, err := strconv.ParseUint(c.numberOfDownstreamPorts, 10, 8)
					if err != nil {
						log.Println("Could not parse product ID:", err)

						return
					}

					c.OnSubmit(
						uint16(vendorID),
						uint16(productID),
						uint16(bcdDevice),
						uint8(numberOfDownstreamPorts),
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
												Text("Vendor ID (HEX)"),
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
											Pattern(hexUint16Pattern).
											Placeholder("046d").
											ID("vendor-id-input").
											OnInput(func(ctx app.Context, e app.Event) {
												c.idVendor = ctx.JSSrc().Get("value").String()
											}).
											Value(c.idVendor),
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
														Text("Product ID (HEX)"),
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
												Pattern(hexUint16Pattern).
												Placeholder("082d").
												ID("product-id-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.idProduct = ctx.JSSrc().Get("value").String()
												}).
												Value(c.idProduct),
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
														Text("Device Release Number (HEX)"),
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
												Pattern(hexUint16Pattern).
												Placeholder("0001").
												ID("device-release-number-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.bcdDevice = ctx.JSSrc().Get("value").String()
												}).
												Value(c.bcdDevice),
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
												Type("number").
												Placeholder("4").
												Min(0).
												Max(127).
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

func (c *FE11sModal) OnMount(ctx app.Context) {
	c.idVendor = fmt.Sprintf("%04x", fe11s.DefaultIdVendor)
	c.idProduct = fmt.Sprintf("%04x", fe11s.DefaultIdProduct)
	c.bcdDevice = fmt.Sprintf("%04x", fe11s.DefaultBcdDevice)
	c.numberOfDownstreamPorts = fmt.Sprintf("%v", fe11s.DefaultNumberOfDownstreamPorts)
}

func (c *FE11sModal) cancel() {
	c.idVendor = ""
	c.idProduct = ""
	c.bcdDevice = ""
	c.numberOfDownstreamPorts = ""

	c.OnCancel()
}
