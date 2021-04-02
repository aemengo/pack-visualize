package component

import "github.com/rivo/tview"

type Logs struct {}

func NewLogs() *Logs {
	return &Logs{}
}

func (l *Logs) View() *tview.TextView {
	tv := tview.NewTextView()
	tv.SetTextAlign(tview.AlignLeft).
		SetText(" ... ").
		SetBorder(true).
		SetTitle(tview.TranslateANSI(bold.Sprint("| logs |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)
	return tv
}