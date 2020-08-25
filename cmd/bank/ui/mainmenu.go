package ui

import "github.com/rivo/tview"

func mainMenu() tview.Primitive {
	menu := tview.NewList()

	menu.AddItem(
		"plugins",
		"browse applications by plugin",
		'p',
		func() {},
	)

	menu.AddItem(
		"quit",
		"quit to terminal",
		'q',
		func() {},
	)

	return menu
}
