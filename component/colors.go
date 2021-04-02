package component

import (
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

var (
	bold            = color.New(color.Bold)
	boldWhite       = color.New(color.Bold, color.FgHiWhite)
	boldGreen       = color.New(color.Bold, color.FgGreen)
	boldMagenta     = color.New(color.Bold, color.FgMagenta)
	backgroundColor = tcell.NewRGBColor(5, 30, 40)
)
