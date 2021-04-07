package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type Plan struct {
	list               *tview.List
	changed            func()
}

type Info struct {
	Name         string
	Version      string
	HasUpdate    bool
	HasSucceeded bool
	Description  string
}

func NewBuildPlan() *Plan {
	list := tview.NewList()
	list.SetMainTextColor(tcell.ColorMediumTurquoise).
		SetSelectedTextColor(tcell.ColorMediumTurquoise).
		SetSelectedBackgroundColor(tcell.ColorDarkSlateGray).
		SetSecondaryTextColor(tcell.ColorDimGray).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1).
		SetTitle(tview.TranslateANSI(bold.Sprint("| plan |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)

	for _, info := range buildpackInfo() {
		title := info.Name + "@" + info.Version
		if info.HasUpdate {
			title = title + boldGreen.Sprint(" (update available)")
		}

		list.AddItem(
			tview.TranslateANSI(title),
			info.Description,
			' ',
			func() {},
		)
	}

	return &Plan{
		list: list,
	}
}

func (p *Plan) View() *tview.List {
	return p.list
}

func (p *Plan) SetChangedFunc(handler func()) {
	p.changed = handler
}

func (p *Plan) SetItemSelectedFunc(handler func()) {
	p.list.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		handler()
	})
}

func (p *Plan) Strobe() {
	time.Sleep(5 * time.Second)

	for index, info := range buildpackInfo() {
		title := info.Name + "@" + info.Version
		if info.HasUpdate {
			title = title + boldGreen.Sprint(" (update available)")
		}

		var roone rune = ' '
		if info.HasSucceeded {
			roone = '✔'
		}

		prevIndex := p.list.GetCurrentItem()
		p.list.RemoveItem(index)
		p.list.InsertItem(
			index,
			tview.TranslateANSI(title),
			info.Description,
			roone,
			func() {},
		)
		p.list.SetCurrentItem(prevIndex)

		time.Sleep(750 * time.Millisecond)

		if p.changed != nil {
			p.changed()
		}
	}
}

func buildpackInfo() []Info {
	return []Info{
		{
			Name:         "paketo-buildpacks/ca-certificates",
			Version:      "2.1.0",
			HasUpdate:    false,
			HasSucceeded: true,
			Description:  "Adds custom CA Certificates",
		},
		{
			Name:         "paketo-buildpacks/go-dist",
			Version:      "0.3.1",
			HasUpdate:    true,
			HasSucceeded: true,
			Description:  "Provides Go Binary",
		},
		{
			Name:         "paketo-buildpacks/go-build",
			Version:      "0.3.0",
			HasUpdate:    false,
			HasSucceeded: true,
			Description:  "Executes Go Build",
		},
		{
			Name:         "paketo-buildpacks/procfile",
			Version:      "4.1.0",
			HasUpdate:    false,
			HasSucceeded: true,
			Description:  "Parses any Procfiles",
		},
		{
			Name:         "paketo-buildpacks/environment-variables",
			Version:      "3.0.0",
			HasUpdate:    false,
			HasSucceeded: true,
			Description:  "Embeds environment variables",
		},
		{
			Name:         "paketo-buildpacks/image-labels",
			Version:      "3.0.0",
			HasUpdate:    false,
			HasSucceeded: true,
			Description:  "Configures label metadata",
		},
		{
			Name:         "<custom>",
			Version:      "_",
			HasUpdate:    false,
			HasSucceeded: false,
			Description:  "Does a specific thing that I want for this application",
		},
	}
}
