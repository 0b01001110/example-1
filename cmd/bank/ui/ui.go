package ui

import (
	"context"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/dogmatiq/dodeca/logging"
	"github.com/dogmatiq/dogma"
)

// UI encapsulates an interactive command-line interface for the example bank
// application.
type UI struct {
	Logger   logging.Logger
	Executor dogma.CommandExecutor

	state state
}

type state struct {
	App   *tview.Application
	Pages *tview.Pages
}

// New returns a new UI.
func New() *UI {
	// Setup the pages that comprise the main content.
	p := tview.NewPages()
	// p.AddPage("mainmenu", mainMenu(), true, true)

	// Initialize a text view for displaying log messages.
	l := tview.NewTextView()
	l.SetTitle("LOG")
	l.SetScrollable(true)
	l.SetBorder(true)

	// Setup a flex box to serve as the horizontal division between the main
	// content and the log view.
	f := tview.NewFlex()
	f.SetDirection(tview.FlexRow)
	f.AddItem(p, 0, 2, true)
	f.AddItem(l, 0, 1, true)

	app := tview.NewApplication()
	app.SetRoot(f, true)

	// Initialize the logger and wire it to the app such that pending log
	// messages are flushed to the text view before the app is drawn.
	log := &logger{
		Text: l,
	}
	app.SetBeforeDrawFunc(func(tcell.Screen) bool {
		log.flush()
		return false
	})

	return &UI{
		Logger: log,
		state: state{
			App:   app,
			Pages: p,
		},
	}
}

// Run executes the UI until ctx is canceled or an error occurs.
func (ui *UI) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-ctx.Done()
		ui.state.App.Stop()
	}()

	return ui.state.App.Run()
}
