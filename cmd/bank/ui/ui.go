package ui

import (
	"context"
	"sync"

	"github.com/rivo/tview"

	"github.com/dogmatiq/dodeca/logging"
	"github.com/dogmatiq/dogma"
)

// UI encapsulates an interactive command-line interface for the example bank
// application.
type UI struct {
	logger *logger
	app    *tview.Application
	pages  *tview.Pages

	m       sync.RWMutex
	execCtx context.Context
	exec    dogma.CommandExecutor
}

// executor is a function used by pages to execute commands.
type executor func(dogma.Message)

// New returns a new UI.
func New() *UI {
	app := tview.NewApplication()

	// Initialize a logger that writes to a text view.
	logView := tview.NewTextView()
	logView.SetTitle(" LOG ")
	logView.SetScrollable(true)
	logView.SetBorder(true)

	log := &logger{
		App:  app,
		Text: logView,
	}

	ui := &UI{
		logger: log,
		app:    app,
		pages:  tview.NewPages(),
	}

	// Setup the pages that comprise the main content.
	ui.pages.AddPage(
		"main-menu",
		mainMenuPage(ui.app, ui.pages),
		true, // focus
		true, // visible
	)

	ui.pages.AddPage(
		"new-customer",
		newCustomerPage(ui.pages, ui.executeCommand),
		true,  // focus
		false, // visible
	)

	// Setup a flex box to serve as the horizontal division between the main
	// content and the log view.
	f := tview.NewFlex()
	f.SetDirection(tview.FlexRow)
	f.AddItem(ui.pages, 0, 2, true)
	f.AddItem(logView, 0, 1, true)

	ui.app.SetRoot(f, true)

	return ui
}

// Run executes the UI until ctx is canceled or an error occurs.
func (ui *UI) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ui.m.Lock()
	ui.execCtx = ctx
	ui.m.Unlock()

	go func() {
		<-ctx.Done()
		ui.logger.close()
		ui.app.Stop()
	}()

	return ui.app.Run()
}

// SetExecutor provides the UI with the executor used to execute Dogma commands.
func (ui *UI) SetExecutor(e dogma.CommandExecutor) {
	ui.m.Lock()
	ui.exec = e
	ui.m.Unlock()
}

// Logger returns a logger that logs to the log pane of the UI.
func (ui *UI) Logger() logging.Logger {
	return ui.logger
}

// executeCommand executes a Dogma command.
func (ui *UI) executeCommand(m dogma.Message) {
	ui.m.RLock()
	defer ui.m.RUnlock()

	if ui.exec == nil {
		logging.LogString(ui.logger, "cannot execute command: no executor has been provided")
		return
	}

	if err := ui.exec.ExecuteCommand(ui.execCtx, m); err != nil {
		logging.Log(ui.logger, "cannot execute command: %s", err)
	}
}
