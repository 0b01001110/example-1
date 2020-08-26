package ui

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rivo/tview"
)

// maxLogLength is the maximum number of bytes of the log to be retained.
const maxLogLength = 2000000

// logger is an implementation of Dodeca's logging.Logger that writes log
// messages to a TextView.
type logger struct {
	App  *tview.Application
	Text *tview.TextView

	m    sync.RWMutex
	done bool
}

func (l *logger) Log(f string, v ...interface{}) {
	l.LogString(fmt.Sprintf(f, v...))
}

func (l *logger) LogString(s string) {
	l.m.RLock()
	defer l.m.RUnlock()

	if !l.done {
		l.App.QueueUpdateDraw(func() {
			l.append(s)
		})
	}
}

func (l *logger) Debug(f string, v ...interface{}) {
	l.DebugString(fmt.Sprintf(f, v...))
}

func (l *logger) DebugString(s string) {
	l.LogString(s)
}

func (l *logger) IsDebug() bool {
	return true
}

// flush appends any pending log lines to the text view.
func (l *logger) append(s string) {
	// Append the new string to the existing log text. GetText() always returns
	// a LF terminated string, so we do not need to add our own.
	t := l.Text.GetText(false) + s

	// Trim the oldest lines until the length of the text falls back below the
	// maximum length.
	for len(t) > maxLogLength {
		i := strings.IndexByte(t, '\n')

		if i == -1 {
			break
		}

		t = t[i+1:]
	}

	// Update the text view with the new text.
	l.Text.SetText(t)
}

// close marks the logger as done so that it does not attempt to enqueue any
// more update events with the application.
//
// It must be called before the application is stopped.
func (l *logger) close() {
	l.m.Lock()
	defer l.m.Unlock()

	l.done = true
}
