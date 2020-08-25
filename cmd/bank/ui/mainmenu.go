package ui

import "github.com/rivo/tview"

func mainMenu(app *tview.Application) tview.Primitive {
	menu := tview.NewList()
	menu.SetTitle("MAIN MENU")
	menu.SetBorder(true)
	menu.SetBorderPadding(1, 1, 2, 2)

	menu.AddItem(
		"Sign-In",
		"Sign-in to a customer account.",
		'i',
		func() {},
	)

	menu.AddItem(
		"Sign-Up",
		"Open a bank account for a new customer.",
		'u',
		func() {},
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
