package component

import (
	"github.com/rivo/tview"
)

type DashBoard struct {
	app    *tview.Application
	screen tview.Primitive
}

func NewDashboard(app *tview.Application) *DashBoard {
	var (
		updates = NewUpdates()
		screen    = tview.NewGrid().
			SetRows(4, 0).
			AddItem(updates.View(), 0, 0, 1, 2, 0, 0, false)
	)

	return &DashBoard{
		app:    app,
		screen: screen,
	}
}

func (l *DashBoard) Run() <-chan bool {
	var (
		doneChan = make(chan bool)
	)

	l.app.SetRoot(l.screen, true).SetFocus(l.screen)
	return doneChan
}