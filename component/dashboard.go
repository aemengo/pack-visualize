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
		//TODO: flash the updates
		updatesTextView = NewUpdates().View()
		keybindingsTextView = NewKeyBindings().View()
		logsTextView = NewLogs().View()

		middle = tview.NewGrid().
			SetColumns(0, 0).
			AddItem(tview.NewBox(), 0, 0, 1, 1, 0, 0, false).
			AddItem(logsTextView, 0, 1, 1, 1, 0, 0, false)

		screen = tview.NewGrid().
				SetRows(4, 0, 1).
				AddItem(updatesTextView, 0, 0, 1, 3, 0, 0, false).
				AddItem(middle, 1, 0, 1, 3, 0, 0, false).
				AddItem(keybindingsTextView, 2, 0, 1, 3, 0, 0, false)
	)

	updatesTextView.SetChangedFunc(func() { app.Draw() })
	keybindingsTextView.SetChangedFunc(func() { app.Draw() })
	logsTextView.SetChangedFunc(func() { app.Draw() })

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
