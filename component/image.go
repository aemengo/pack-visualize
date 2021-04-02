package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Image struct {}

func NewImage() *Image {
	return &Image{}
}

func (b *Image) View() *tview.TreeView {
	root := tview.NewTreeNode(tview.TranslateANSI("image:" + boldWhite.Sprint(" anthony/my-sample-golang:0.1") + bold.Sprintf(" - 7 days ago"))).SetColor(tcell.ColorDimGray)
	runImage := tview.NewTreeNode(tview.TranslateANSI(" run:" + boldWhite.Sprint(" paketo-buildpacks/run:full-cnb") + bold.Sprintf(" - 84eff6936799") )).SetColor(tcell.ColorDimGray)

	root.AddChild(runImage)

	tree := tview.NewTreeView()
	tree.SetGraphics(true).
		SetGraphicsColor(tcell.ColorMediumTurquoise).
		SetBorderPadding(1, 1, 2, 2).
		SetBackgroundColor(backgroundColor)
	tree.SetRoot(root)
	return tree
}
