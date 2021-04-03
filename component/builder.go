package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Builder struct {}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) View() *tview.TreeView {
	root := tview.NewTreeNode(tview.TranslateANSI("builder:" + boldWhite.Sprint(" paketo-buildpacks/builder:full-cf"))).SetColor(tcell.ColorDimGray)
	lifecycle := tview.NewTreeNode(tview.TranslateANSI(" lifecycle:" + boldWhite.Sprint(" 0.11.1") + boldGreen.Sprintf(" (update available)") )).SetColor(tcell.ColorDimGray)
	buildImage := tview.NewTreeNode(tview.TranslateANSI(" build:" + boldWhite.Sprint(" paketo-buildpacks/build:full") + bold.Sprintf(" - d0e8bfb686fd") )).SetColor(tcell.ColorDimGray)
	runImage := tview.NewTreeNode(tview.TranslateANSI(" run:" + boldWhite.Sprint(" paketo-buildpacks/run:full") + bold.Sprintf(" - 84eff6936799")  + boldGreen.Sprintf(" (update available)") )).SetColor(tcell.ColorDimGray)

	root.AddChild(lifecycle)
	root.AddChild(buildImage)
	root.AddChild(runImage)

	tree := tview.NewTreeView()
	tree.SetGraphics(true).
		SetGraphicsColor(tcell.ColorMediumTurquoise).
		SetBorderPadding(1, 1, 2, 2).
		SetBackgroundColor(backgroundColor)
	tree.SetRoot(root)
	return tree
}
