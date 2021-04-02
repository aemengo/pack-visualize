package component

import "github.com/rivo/tview"

var updatesText = `An update is available for *paketo-buildpacks/run:full-cnb*. Please rebase.
The following CVEs have been detected: .`

type Updates struct{}

func NewUpdates() *Updates {
	return &Updates{}
}

func (u *Updates) View() *tview.TextView {
	tv := tview.NewTextView()
	tv.SetTextAlign(tview.AlignCenter).
		SetText(updatesText).
		SetBorder(true).
		SetTitle(tview.TranslateANSI(bold.Sprint("| updates |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)
	return tv
}
