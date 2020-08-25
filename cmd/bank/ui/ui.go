package ui

import (
	"context"
	"log"

	"github.com/rivo/tview"

	"github.com/dogmatiq/dodeca/logging"
	"github.com/dogmatiq/dogma"
)

// UI encapsulates an interactive command-line interface for the example bank
// application.
type UI struct {
	Logger   logging.Logger
	Executor dogma.CommandExecutor

	state *state
}

type state struct {
	App   *tview.Application
	Pages *tview.Pages
}

// New returns a new UI.
func New() *UI {
	st := &state{
		App:   tview.NewApplication(),
		Pages: tview.NewPages(),
	}

	// Setup the pages that comprise the main content.
	st.Pages.AddPage("mainmenu", mainMenu(st), true, true)

	// Initialize a text view for displaying log messages.
	l := tview.NewTextView()
	l.SetTitle(" LOG ")
	l.SetScrollable(true)
	l.SetBorder(true)

	// Setup a flex box to serve as the horizontal division between the main
	// content and the log view.
	f := tview.NewFlex()
	f.SetDirection(tview.FlexRow)
	f.AddItem(st.Pages, 0, 2, true)
	f.AddItem(l, 0, 1, true)

	st.App.SetRoot(f, true)

	return &UI{
		Logger: &logging.StandardLogger{
			Target:       log.New(l, "", 0),
			CaptureDebug: true,
		},
		state: st,
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
