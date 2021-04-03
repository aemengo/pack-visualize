package component

import (
	"fmt"
	"github.com/rivo/tview"
)

var keyBindingsText = fmt.Sprintf(`%s edit | %s update | %s rebase | %s build`,
	boldYellow.Sprint(`(e)`), boldYellow.Sprint(`(u)`), boldYellow.Sprint(`(r)`), boldYellow.Sprint(`(b)`))

type KeyBindings struct {}

func NewKeyBindings() *KeyBindings {
	return &KeyBindings{}
}

func (u *KeyBindings) View() *tview.TextView {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText(tview.TranslateANSI(keyBindingsText)).
		SetBackgroundColor(backgroundColor)
	return tv
}