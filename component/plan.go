package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Plan struct{}

type Info struct {
	Name        string
	Version     string
	HasUpdate   bool
	HasSucceeded   bool
	Description string
}

func NewBuildPlan() *Plan {
	return &Plan{}
}

func (p *Plan) View() *tview.List {
	builderList := tview.NewList()
	builderList.
		SetMainTextColor(tcell.ColorMediumTurquoise).
		SetSelectedTextColor(tcell.ColorMediumTurquoise).
		SetSelectedBackgroundColor(tcell.ColorDarkSlateGray).
		SetSecondaryTextColor(tcell.ColorDimGray).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1).
		SetTitle(tview.TranslateANSI(bold.Sprint("| plan |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)

	for _, info := range p.buildpackInfo() {
		var roone rune = 0
		if info.HasSucceeded {
			roone = '✔'
		}

		title := info.Name+"@"+info.Version
		//if info.HasUpdate {
		//	title = title + " (update available)"
		//}

		builderList.AddItem(
			title,
			info.Description,
			roone,
			func() {},
		)
	}

	return builderList
}

func (p *Plan) buildpackInfo() []Info {
	return []Info{
		{
			Name:        "paketo-buildpacks/ca-certificates",
			Version:     "2.1.0",
			HasUpdate:   false,
			HasSucceeded: true,
			Description: "Adds custom CA Certificates",
		},
		{
			Name:        "paketo-buildpacks/go-dist",
			Version:     "0.3.1",
			HasUpdate:   true,
			HasSucceeded: true,
			Description: "Provides Go Binary",
		},
		{
			Name:        "paketo-buildpacks/go-build",
			Version:     "0.3.0",
			HasUpdate:   false,
			HasSucceeded: true,
			Description: "Executes Go Build",
		},
		{
			Name:        "paketo-buildpacks/procfile",
			Version:     "4.1.0",
			HasUpdate:   false,
			HasSucceeded: true,
			Description: "Parses any Procfiles",
		},
		{
			Name:        "paketo-buildpacks/environment-variables",
			Version:     "3.0.0",
			HasUpdate:   false,
			HasSucceeded: true,
			Description: "Embeds environment variables",
		},
		{
			Name:        "paketo-buildpacks/image-labels",
			Version:     "3.0.0",
			HasUpdate:   false,
			HasSucceeded: false,
			Description: "Configures label metadata",
		},
		{
			Name:        "<custom>",
			Version:     "_",
			HasUpdate:   false,
			HasSucceeded: false,
			Description: "Does a specific thing that I want for this application",
		},
	}
}
