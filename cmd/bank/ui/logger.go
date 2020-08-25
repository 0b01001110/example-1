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
	Text *tview.TextView

	m       sync.RWMutex
	pending strings.Builder
}

func (l *logger) Log(f string, v ...interface{}) {
	l.LogString(fmt.Sprintf(f, v...))
}

func (l *logger) LogString(s string) {
	l.m.Lock()
	defer l.m.Unlock()

	l.pending.WriteString(s + "\n")
}

func (l *logger) Debug(f string, v ...interface{}) {
	l.Log(f, v...)
}

func (l *logger) DebugString(s string) {
	l.LogString(s)
}

func (l *logger) IsDebug() bool {
	return true
}

// flush appends any pending log lines to the text view.
func (l *logger) flush() {
	l.m.Lock()
	defer l.m.Unlock()

	// Append the new string to the existing log text. GetText() always returns
	// a LF terminated string, so we do not need to add our own.
	t := l.Text.GetText(false) + l.pending.String()
	l.pending.Reset()

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
