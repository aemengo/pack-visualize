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
		// TODO: flash the updates
		updatesTextView     = NewUpdates().View()
		keybindingsTextView = NewKeyBindings().View()
		logsTextView        = NewLogs().View()
		builderView         = NewBuilder().View()
		imageView           = NewImage().View()
		buildPlan           = NewBuildPlan().View()
	)

	screen := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(updatesTextView, 4, 0, false).
			AddItem(imageView, 4, 0, false).
			AddItem(builderView, 7, 0, false).
			AddItem(buildPlan, 0, 1, true).
			AddItem(keybindingsTextView, 1, 0, false), 0, 2, true).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(logsTextView, 0, 1, false).
			AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 1, 0, false), 0, 1, false)

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
