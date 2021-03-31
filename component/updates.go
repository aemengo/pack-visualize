package component

import "github.com/rivo/tview"

var copy = `
A stack update is available for: . Please rebase.
-
The following CVEs have been detected: .
`

type Updates struct {}

func NewUpdates() *Updates {
	return &Updates{}
}

func (u *Updates) View() tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(copy).
		SetBorder(true).
		SetTitle(tview.TranslateANSI(bold.Sprint("| updates |"))).
		SetTitleAlign(tview.AlignLeft)
}