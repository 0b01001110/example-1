package ui

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/dogmatiq/example/messages/commands"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func newCustomerPage(
	pages *tview.Pages,
	exec executor,
) tview.Primitive {
	errors := tview.NewTextView()
	errors.SetTextColor(tcell.ColorRed)

	form := tview.NewForm()

	customerName := tview.NewInputField().
		SetLabel("Customer Name").
		SetFieldWidth(20)
	form.AddFormItem(customerName)

	accountName := tview.NewInputField().
		SetLabel("Account Name").
		SetFieldWidth(20)
	form.AddFormItem(accountName)

	reset := func() {
		errors.SetText("")
		customerName.SetText("")
		accountName.SetText("Savings")
	}
	reset()

	form.AddButton("Submit", func() {
		errors.SetText("")
		ok := true

		if customerName.GetText() == "" {
			io.WriteString(errors, "Customer name must not be empty.\n")
			ok = false
		}

		if accountName.GetText() == "" {
			io.WriteString(errors, "Account name must not be empty.\n")
			ok = false
		}

		if !ok {
			return
		}

		defer reset()

		exec(commands.OpenAccountForNewCustomer{
			CustomerID:   generateCustomerID(),
			CustomerName: customerName.GetText(),
			AccountID:    generateAccountID(),
			AccountName:  accountName.GetText(),
		})

		pages.SwitchToPage("main-menu")
	})

	cancel := func() {
		reset()
		pages.SwitchToPage("main-menu")
	}

	form.AddButton("Cancel", cancel)
	form.SetCancelFunc(cancel)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(form, 7, 0, true)
	flex.AddItem(errors, 0, 1, false)

	frame := tview.NewFrame(flex)
	frame.SetTitle(" NEW CUSTOMER ")
	frame.SetBorder(true)
	frame.SetBorderPadding(0, 1, 2, 2)

	return frame
}

// generateCustomerID returns a new unique customer ID.
func generateCustomerID() string {
	return fmt.Sprintf(
		"%03d-%04d",
		100+rand.Intn(899),
		1000+rand.Intn(8999),
	)
}
