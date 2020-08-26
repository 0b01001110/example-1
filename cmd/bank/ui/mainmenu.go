package ui

import "github.com/rivo/tview"

func mainMenuPage(
	app *tview.Application,
	pages *tview.Pages,
) tview.Primitive {
	menu := tview.NewList()
	menu.SetTitle(" MAIN MENU ")
	menu.SetBorder(true)
	menu.SetBorderPadding(1, 1, 2, 2)

	menu.AddItem(
		"Sign-In",
		"Sign-in as an existing customer.",
		's',
		func() {},
	)

	menu.AddItem(
		"New Customer",
		"Open a bank account for a new customer.",
		'n',
		func() {
			pages.SwitchToPage("new-customer")
		},
	)

	menu.AddItem(
		"Quit",
		"Quit to terminal.",
		'q',
		func() {
			app.Stop()
		},
	)

	return menu
}
