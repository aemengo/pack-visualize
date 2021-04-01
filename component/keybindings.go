package component

import (
	"fmt"
	"github.com/rivo/tview"
)

var keyBindingsText = fmt.Sprintf(`%s edit | %s update | %s rebase | %s build`,
	boldMagenta.Sprint(`(e)`), boldMagenta.Sprint(`(u)`), boldMagenta.Sprint(`(r)`), boldMagenta.Sprint(`(b)`))

type KeyBindings struct {}

func NewKeyBindings() *KeyBindings {
	return &KeyBindings{}
}

func (u *KeyBindings) View() *tview.TextView {
	tv := tview.NewTextView()
	tv.SetTextAlign(tview.AlignLeft).
		SetText(keyBindingsText)
	return tv
}