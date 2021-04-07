package component

import (
	"github.com/rivo/tview"
)

type DashBoard struct {
	app     *tview.Application
	screen  tview.Primitive
	updates *Updates
	logs    *Logs
	plan    *Plan
}

func NewDashboard(app *tview.Application) *DashBoard {
	var (
		updates         = NewUpdates()
		updatesTextView = updates.View()

		logs         = NewLogs()
		logsTextView = logs.View()

		plan     = NewBuildPlan()
		planView = plan.View()

		keybindingsTextView = NewKeyBindings().View()
		builderView         = NewBuilder().View()
		imageView           = NewImage().View()
	)

	screen := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(updatesTextView, 4, 0, false).
			AddItem(imageView, 4, 0, false).
			AddItem(builderView, 7, 0, false).
			AddItem(planView, 0, 1, true).
			AddItem(keybindingsTextView, 1, 0, false), 0, 1, true).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(logsTextView, 0, 1, false).
			AddItem(tview.NewBox().SetBackgroundColor(backgroundColor), 1, 0, false), 0, 1, false)

	l := &DashBoard{
		app:     app,
		screen:  screen,
		updates: updates,
		logs:    logs,
		plan:    plan,
	}

	updatesTextView.SetChangedFunc(func() { app.Draw() })
	keybindingsTextView.SetChangedFunc(func() { app.Draw() })
	logsTextView.SetChangedFunc(func() { app.Draw() })

	plan.SetChangedFunc(func() { app.Draw() })
	plan.SetItemSelectedFunc(func() {
		go l.logs.Strobe()
		go l.plan.Strobe()
	})

	return l
}

func (l *DashBoard) Run() <-chan bool {
	var (
		doneChan = make(chan bool)
	)

	l.app.SetRoot(l.screen, true).SetFocus(l.screen)
	go l.updates.Strobe()
	return doneChan
}
