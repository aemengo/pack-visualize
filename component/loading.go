package component

import (
	"github.com/gdamore/tcell/v2"
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

	textView.SetBackgroundColor(backgroundColor)
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
	)

	l.app.SetRoot(l.screen, true)
	go func() {
		l.pull(ticker, time.After(4 * time.Second), doneChan)
		l.run(ticker, time.After(4 * time.Second), doneChan)
	}()
	return doneChan
}

func (l *Loading) pull(ticker *time.Ticker, timeout <-chan time.Time, doneChan chan<- bool) {
	var (
		i        = 0
		texts    = []string{
			"⏳️ Pulling",
			"⏳️ Pulling.",
			"⏳️ Pulling..",
			"⏳️ Pulling...",
		}
	)

	for {
		select {
		case <-ticker.C:
			l.textView.SetText(texts[i])

			i = i + 1
			if i == len(texts) {
				i = 0
			}
		case <-timeout:
			return
		}
	}
}

func (l *Loading) run(ticker *time.Ticker, timeout <-chan time.Time, doneChan chan<- bool) {
	var (
		i        = 0
		doneText = "⌛️ Detected!"
		texts    = []string{
			"⏳️ Detecting",
			"⏳️ Detecting.",
			"⏳️ Detecting..",
			"⏳️ Detecting...",
		}
	)

	for {
		select {
		case <-ticker.C:
			l.textView.SetText(texts[i])

			i = i + 1
			if i == len(texts) {
				i = 0
			}
		case <-timeout:
			l.textView.SetText(doneText)

			doneChan <- true
			return
		}
	}
}

func centered(p tview.Primitive) tview.Primitive {
	grid := tview.NewGrid().
		SetColumns(0, 20, 0).
		SetRows(0, 1, 0).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 0, 0, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 0, 1, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 0, 2, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 1, 0, 1, 1, 0, 0, true).
		AddItem(p, 1, 1, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 1, 2, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 2, 0, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 2, 1, 1, 1, 0, 0, true).
		AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 2, 2, 1, 1, 0, 0, true)

	logs := tview.NewTextView().
		SetBackgroundColor(tcell.NewRGBColor(10, 35, 45))

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
			AddItem(logs, 0, 3, false).
			AddItem(grid, 0, 1, false)
}
