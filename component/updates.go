package component

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

var updatesText1 = fmt.Sprintf(`An update is available for %s. Please rebase.`,
	boldWhite.Sprint("paketo-buildpacks/run:full"))

var updatesText2 = fmt.Sprintf(`An update is available for %s. Please rebase.
The following CVEs have been detected: %s(%s), %s(%s).`,
boldWhite.Sprint("paketo-buildpacks/run:full"),
boldWhite.Sprint("CVE-2020-1946"),
boldYellow.Sprint("medium"),
boldWhite.Sprint("CVE-2021-27365"),
boldRed.Sprint("high"))

type Updates struct{
	tv *tview.TextView
}

func NewUpdates() *Updates {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLightGray).
		SetBorder(true).
		SetTitle(tview.TranslateANSI(bold.Sprint("| updates |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)

	return &Updates{
		tv: tv,
	}
}

func (u *Updates) View() *tview.TextView {
	return u.tv
}

func (u *Updates) Strobe() {
	flag := true
	for i := 0; i < 16; i++ {
		if flag {
			u.tv.SetText(tview.TranslateANSI(updatesText1))
		} else {
			u.tv.SetText(tview.TranslateANSI(updatesText2))
		}

		time.Sleep(750 * time.Millisecond)
		flag = !flag
	}
}
