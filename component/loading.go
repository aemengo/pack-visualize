package component

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

type Loading struct {
	app      *tview.Application
	textView *tview.TextView
	screen   tview.Primitive
}

func NewLoading(app *tview.Application) *Loading {
	var (
		textView = tview.NewTextView()
		screen   = centered(textView)
	)

	textView.SetChangedFunc(func() { app.Draw() })

	return &Loading{
		app:      app,
		textView: textView,
		screen:   screen,
	}
}

func (l *Loading) Run() <-chan bool {
	var (
		doneChan = make(chan bool)
		ticker   = time.NewTicker(200 * time.Millisecond)
		timeout  = time.After(4 * time.Second)
	)

	l.app.SetRoot(l.screen, true)
	go l.run(ticker, timeout, doneChan)
	return doneChan
}

func (l *Loading) run(ticker *time.Ticker, timeout <-chan time.Time, doneChan chan<- bool) {
	var (
		i        = 0
		doneText = "⌛️ Detected!"
		texts    = []string{
			"⌛️ Detecting",
			"⌛️ Detecting.",
			"⌛️ Detecting..",
			"⌛️ Detecting...",
		}
	)

	for {
		select {
		case <-ticker.C:
			l.textView.Clear()
			fmt.Fprintln(l.textView, texts[i])

			i = i + 1
			if i == len(texts) {
				i = 0
			}
		case <-timeout:
			l.textView.Clear()
			fmt.Fprintln(l.textView, doneText)

			doneChan <- true
			return
		}
	}
}

func centered(p tview.Primitive) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, 20, 0).
		SetRows(0, 1, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}
