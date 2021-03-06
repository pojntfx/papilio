package components

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
	"github.com/pojntfx/papilio/pkg/generators/fe21"
)

type FE21Modal struct {
	app.Compo

	OnSubmit func(
		idVendor uint16,
		idProduct uint16,
		bcdDevice uint16,
		numberOfDownstreamPorts uint8,

		serial string,
		portsWithRemovableDevices [fe21.MaxNumberOfDownstreamPorts]bool,
		portIndicatorSupport bool,
		compoundDevice bool,
		maximumCurrent500mA bool,
	)
	OnCancel func()

	idVendor                string
	idProduct               string
	bcdDevice               string
	numberOfDownstreamPorts uint64

	serial                    string
	portsWithRemovableDevices [fe21.MaxNumberOfDownstreamPorts]bool
	portIndicatorSupport      bool
	compoundDevice            bool
	maximumCurrent500mA       bool
}

func (c *FE21Modal) Render() app.UI {
	return &components.Modal{
		ID:           "fe21-modal",
		Title:        "Configure your FE 2.1",
		DisableFocus: true,
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("fe21-form").
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

					c.OnSubmit(
						uint16(vendorID),
						uint16(productID),
						uint16(bcdDevice),
						uint8(c.numberOfDownstreamPorts),

						c.serial,
						c.portsWithRemovableDevices,
						c.portIndicatorSupport,
						c.compoundDevice,
						c.maximumCurrent500mA,
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
												Placeholder(strconv.Itoa(int(fe21.MaxNumberOfDownstreamPorts))).
												Min(1).
												Max(fe21.MaxNumberOfDownstreamPorts).
												ID("number-of-downstream-ports-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													numberOfDownstreamPorts, err := strconv.ParseUint(ctx.JSSrc().Get("value").String(), 10, 8)
													if err != nil {
														log.Println("Could not parse number of downstream ports:", err)

														return
													}

													if numberOfDownstreamPorts > uint64(fe21.MaxNumberOfDownstreamPorts) {
														log.Println("Maximum amount of downstream ports exceeded, ignoring")

														return
													}

													c.numberOfDownstreamPorts = numberOfDownstreamPorts
												}).
												Value(c.numberOfDownstreamPorts),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("serial-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Serial Number"),
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
										Required(true).
										Type("text").
										Max(15).
										Placeholder("abcdefghijk").
										ID("serial-input").
										OnInput(func(ctx app.Context, e app.Event) {
											c.serial = ctx.JSSrc().Get("value").String()
										}).
										Value(c.serial),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("ports-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Ports with Removable Devices"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control pf-l-flex pf-m-justify-content-center").
								Body(
									&ToggleGroup{
										ID: "ports-input",
										Toggles: func() []Toggle {
											toggles := []Toggle{}

											for i, port := range c.portsWithRemovableDevices[:c.numberOfDownstreamPorts] {
												toggles = append(toggles, Toggle{
													ID:    i,
													Title: fmt.Sprintf("Port %v", i+1),
													On:    port,
												})
											}

											return toggles
										}(),
										OnToggle: func(id int) {
											if len(c.portsWithRemovableDevices) <= id {
												log.Println("Could not find port with ID", id)

												return
											}

											c.portsWithRemovableDevices[id] = !c.portsWithRemovableDevices[id]
										},
									},
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-check").
										Body(
											&components.Controlled{
												Component: app.Input().
													Class("pf-c-check__input").
													Type("checkbox").
													ID("port-indicator-support").
													OnInput(func(ctx app.Context, e app.Event) {
														c.portIndicatorSupport = !c.portIndicatorSupport
													}),
												Properties: map[string]interface{}{
													"checked": !c.portIndicatorSupport,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("port-indicator-support").
												Body(
													app.Text("Port Indicator Support"),
												),
											app.Span().
												Class("pf-c-check__description").
												Text("Whether port indicators are supported on the downstream facing ports and PORT_INDICATOR requests control the indicators."),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-check").
										Body(
											&components.Controlled{
												Component: app.Input().
													Class("pf-c-check__input").
													Type("checkbox").
													ID("compound-device").
													OnInput(func(ctx app.Context, e app.Event) {
														c.compoundDevice = !c.compoundDevice
													}),
												Properties: map[string]interface{}{
													"checked": !c.compoundDevice,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("compound-device").
												Body(
													app.Text("Compound Device"),
												),
											app.Span().
												Class("pf-c-check__description").
												Text("Whether the hub identifies a compound device."),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-check").
										Body(
											&components.Controlled{
												Component: app.Input().
													Class("pf-c-check__input").
													Type("checkbox").
													ID("maximum-current").
													OnInput(func(ctx app.Context, e app.Event) {
														c.maximumCurrent500mA = !c.maximumCurrent500mA
													}),
												Properties: map[string]interface{}{
													"checked": !c.maximumCurrent500mA,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("maximum-current").
												Body(
													app.Text("Allow 500mA"),
												),
											app.Span().
												Class("pf-c-check__description").
												Text("Whether the maximum current requirements of the hub controller electronics should be allowed to reach 500mA."),
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
				Form("fe21-form"),
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

func (c *FE21Modal) OnMount(ctx app.Context) {
	c.idVendor = fmt.Sprintf("%04x", fe21.DefaultIdVendor)
	c.idProduct = fmt.Sprintf("%04x", fe21.DefaultIdProduct)
	c.bcdDevice = fmt.Sprintf("%04x", fe21.DefaultBcdDevice)
	c.numberOfDownstreamPorts = uint64(fe21.MaxNumberOfDownstreamPorts)

	c.serial = ""
	c.portsWithRemovableDevices = fe21.DefaultPortsWithRemovableDevices
	c.compoundDevice = fe21.DefaultCompoundDevice
	c.maximumCurrent500mA = fe21.DefaultMaximumCurrent500mA
}

func (c *FE21Modal) cancel() {
	c.idVendor = ""
	c.idProduct = ""
	c.bcdDevice = ""
	c.numberOfDownstreamPorts = uint64(fe21.MaxNumberOfDownstreamPorts)

	c.serial = ""
	c.portsWithRemovableDevices = [fe21.MaxNumberOfDownstreamPorts]bool{false, false, false, false, false, false, false}
	c.compoundDevice = false
	c.maximumCurrent500mA = false

	c.OnCancel()
}
