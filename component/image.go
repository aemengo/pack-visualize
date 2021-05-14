package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Image struct {
	treeView *tview.TreeView
}

func NewImage() *Image {
	tree := tview.NewTreeView()
	tree.SetGraphics(true).
		SetGraphicsColor(tcell.ColorMediumTurquoise).
		SetBorderPadding(1, 1, 2, 2).
		SetBackgroundColor(backgroundColor)
	return &Image{
		treeView: tree,
	}
}

func (b *Image) View() *tview.TreeView {
	root := tview.NewTreeNode(tview.TranslateANSI("image:" + boldWhite.Sprint(" anthony/my-sample-golang:0.1"))).SetColor(tcell.ColorDimGray)
	runImage := tview.NewTreeNode(tview.TranslateANSI(" run:")).SetColor(tcell.ColorDimGray)

	root.AddChild(runImage)
	b.treeView.SetRoot(root)
	return b.treeView
}

func (b *Image) Update() {
	root := tview.NewTreeNode(tview.TranslateANSI("image:" + boldWhite.Sprint(" anthony/my-sample-golang:0.1") + bold.Sprintf(" - 2878ca7ca986 (just now)"))).SetColor(tcell.ColorDimGray)
	runImage := tview.NewTreeNode(tview.TranslateANSI(" run:" + boldWhite.Sprint(" paketo-buildpacks/run:full") + bold.Sprintf(" - 84eff6936799") )).SetColor(tcell.ColorDimGray)

	root.AddChild(runImage)
	b.treeView.SetRoot(root)
}
